// Command generate updates Dota protobuf sources, regenerates Go protobufs, and
// refreshes the generated API wrappers.
package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const aptreModule = "github.com/aperturerobotics/common/cmd/aptre@v0.34.1"

const protocolImportPrefix = "github.com/paralin/go-dota2/protocol/"

var sourceProtoGlobs = []string{
	"dota_gcmessages_*.proto",
	"dota_client_enums.proto",
	"networkbasetypes.proto",
	"network_connection.proto",
	"base_gcmessages.proto",
	"gcsdk_gcmessages.proto",
	"econ_*.proto",
	"dota_match_metadata.proto",
	"dota_shared_enums.proto",
	"events.proto",
	"steammessages.proto",
	"steammessages_unified_base.steamworkssdk.proto",
	"steammessages_steamlearn.steamworkssdk.proto",
	"valveextensions.proto",
	"gcsystemmsgs.proto",
}

func main() {
	verbose := flag.Bool("verbose", false, "verbose output")
	skipSubmodule := flag.Bool("skip-submodule", false, "skip git submodule update")
	skipApigen := flag.Bool("skip-apigen", false, "skip the apigen step")
	flag.Parse()

	if err := run(context.Background(), *verbose, *skipSubmodule, *skipApigen); err != nil {
		log.Fatal(err)
	}
}

func run(ctx context.Context, verbose, skipSubmodule, skipApigen bool) error {
	wd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("getwd: %w", err)
	}
	repoRoot, err := findRepoRoot(wd)
	if err != nil {
		return fmt.Errorf("finding repo root: %w", err)
	}
	if verbose {
		log.Printf("repo root: %s", repoRoot)
	}

	protocolDir := filepath.Join(repoRoot, "protocol")
	submoduleDir := filepath.Join(repoRoot, "generator", "Protobufs")

	if !skipSubmodule {
		log.Println("==> Updating submodule...")
		if err := runCmd(ctx, repoRoot, verbose, "git", "submodule", "update", "--init", "generator/Protobufs"); err != nil {
			return fmt.Errorf("submodule update: %w", err)
		}
	}

	log.Println("==> Preprocessing protos...")
	if err := preprocessProtos(submoduleDir, protocolDir, verbose); err != nil {
		return fmt.Errorf("preprocess protos: %w", err)
	}

	log.Println("==> Vendoring protobuf dependencies...")
	if err := runCmd(ctx, repoRoot, verbose, "go", "mod", "vendor"); err != nil {
		return fmt.Errorf("go mod vendor: %w", err)
	}
	if err := vendorProtobufIncludes(ctx, repoRoot); err != nil {
		return fmt.Errorf("vendor protobuf includes: %w", err)
	}

	log.Println("==> Generating protobufs...")
	if err := generateProtos(ctx, repoRoot, verbose); err != nil {
		return fmt.Errorf("generate protobufs: %w", err)
	}

	if !skipApigen {
		log.Println("==> Running apigen...")
		if err := runApigen(ctx, repoRoot, verbose); err != nil {
			return fmt.Errorf("apigen: %w", err)
		}
	}

	log.Println("==> Running goimports...")
	if err := runAptre(ctx, repoRoot, verbose, "goimports", "--project-dir", repoRoot); err != nil {
		return fmt.Errorf("goimports: %w", err)
	}

	log.Println("==> Done.")
	return nil
}

func findRepoRoot(dir string) (string, error) {
	for {
		modPath := filepath.Join(dir, "go.mod")
		data, err := os.ReadFile(modPath)
		if err == nil && strings.Contains(string(data), "module github.com/paralin/go-dota2\n") {
			return dir, nil
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return "", fmt.Errorf("could not find repo root")
		}
		dir = parent
	}
}

func preprocessProtos(submoduleDir, protocolDir string, verbose bool) error {
	srcDir := filepath.Join(submoduleDir, "dota2")

	var srcFiles []string
	for _, glob := range sourceProtoGlobs {
		matches, err := filepath.Glob(filepath.Join(srcDir, glob))
		if err != nil {
			return fmt.Errorf("glob %q: %w", glob, err)
		}
		srcFiles = append(srcFiles, matches...)
	}
	if len(srcFiles) == 0 {
		return fmt.Errorf("no proto files found in %s", srcDir)
	}

	seen := make(map[string]bool)
	uniqueFiles := make([]string, 0, len(srcFiles))
	for _, file := range srcFiles {
		if seen[file] {
			continue
		}
		seen[file] = true
		uniqueFiles = append(uniqueFiles, file)
	}
	if verbose {
		log.Printf("found %d source proto files", len(uniqueFiles))
	}

	if err := os.MkdirAll(protocolDir, 0o755); err != nil {
		return fmt.Errorf("create protocol dir: %w", err)
	}
	writtenProtos := make(map[string]bool, len(uniqueFiles))
	for _, src := range uniqueFiles {
		data, err := os.ReadFile(src)
		if err != nil {
			return fmt.Errorf("read %s: %w", src, err)
		}

		name := filepath.Base(src)
		content := normalizeProto(name, string(data))
		dst := filepath.Join(protocolDir, name)
		if err := os.WriteFile(dst, []byte(content), 0o644); err != nil {
			return fmt.Errorf("write %s: %w", dst, err)
		}
		writtenProtos[name] = true
	}

	entries, err := os.ReadDir(protocolDir)
	if err != nil {
		return fmt.Errorf("read protocol dir: %w", err)
	}
	for _, entry := range entries {
		if !strings.HasSuffix(entry.Name(), ".proto") || writtenProtos[entry.Name()] {
			continue
		}
		path := filepath.Join(protocolDir, entry.Name())
		if verbose {
			log.Printf("removing old proto: %s", entry.Name())
		}
		if err := os.Remove(path); err != nil {
			return fmt.Errorf("remove %s: %w", path, err)
		}
	}

	if err := os.WriteFile(filepath.Join(protocolDir, "protocol.go"), []byte("package protocol\n"), 0o644); err != nil {
		return fmt.Errorf("write protocol.go: %w", err)
	}
	return nil
}

func vendorProtobufIncludes(ctx context.Context, repoRoot string) error {
	cmd := exec.CommandContext(ctx, "go", "list", "-m", "-f", "{{.Dir}}", "github.com/aperturerobotics/protobuf")
	cmd.Dir = repoRoot
	out, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("go list protobuf module: %w", err)
	}
	src := filepath.Join(strings.TrimSpace(string(out)), "src", "google", "protobuf")
	dst := filepath.Join(repoRoot, "vendor", "github.com", "aperturerobotics", "protobuf", "src", "google", "protobuf")
	return copyDir(src, dst)
}

func copyDir(src, dst string) error {
	return filepath.WalkDir(src, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		rel, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}
		target := filepath.Join(dst, rel)
		if d.IsDir() {
			return os.MkdirAll(target, 0o755)
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		return os.WriteFile(target, data, 0o644)
	})
}

func normalizeProto(name, body string) string {
	content := "syntax = \"proto2\";\npackage protocol;\n\n" + body
	content = strings.ReplaceAll(content, "optional .", "optional ")
	content = strings.ReplaceAll(content, "required .", "required ")
	content = strings.ReplaceAll(content, "repeated .", "repeated ")
	content = strings.ReplaceAll(content, "\t.", "\t")
	content = strings.ReplaceAll(content, "google/protobuf/valve_extensions.proto", "valveextensions.proto")
	content = rewriteProtoImports(content)

	if name == "steammessages_steamlearn.steamworkssdk.proto" {
		content = strings.ReplaceAll(content, "(.CMsgSteamLearn", "(CMsgSteamLearn")
	}
	if name == "dota_gcmessages_common.proto" {
		content = removeMessageBlock(content, "CDotaMsgStructuredTooltipProperties")
	}
	if name == "networkbasetypes.proto" {
		content = strings.ReplaceAll(content, "maximum_size_bytes = 50000", "maximum_size_bytes = 60002")
	}
	if name == "steammessages_unified_base.steamworkssdk.proto" {
		content = strings.ReplaceAll(content, "service_description = 50000", "service_description = 60011")
		content = strings.ReplaceAll(content, "service_execution_site = 50008", "service_execution_site = 60018")
		content = strings.ReplaceAll(content, "method_description = 50000", "method_description = 60012")
		content = strings.ReplaceAll(content, "enum_description = 50000", "enum_description = 60013")
		content = strings.ReplaceAll(content, "enum_value_description = 50000", "enum_value_description = 60014")
		content = strings.ReplaceAll(content, "description = 50000", "description = 60010")
	}
	return content
}

func rewriteProtoImports(content string) string {
	lines := strings.Split(content, "\n")
	for i, line := range lines {
		trimmed := strings.TrimSpace(line)
		if !strings.HasPrefix(trimmed, "import ") || !strings.HasSuffix(trimmed, "\";") {
			continue
		}
		firstQuote := strings.Index(trimmed, "\"")
		if firstQuote < 0 {
			continue
		}
		path := strings.TrimSuffix(trimmed[firstQuote+1:], "\";")
		if strings.HasPrefix(path, "google/") || strings.Contains(path, "/") {
			continue
		}
		lines[i] = strings.Replace(line, path, protocolImportPrefix+path, 1)
	}
	return strings.Join(lines, "\n")
}

func removeMessageBlock(content, messageName string) string {
	lines := strings.Split(content, "\n")
	result := make([]string, 0, len(lines))
	inBlock := false
	depth := 0
	for _, line := range lines {
		if !inBlock {
			if strings.HasPrefix(line, "message "+messageName+" {") {
				inBlock = true
				depth = 1
				continue
			}
			result = append(result, line)
			continue
		}

		depth += strings.Count(line, "{") - strings.Count(line, "}")
		if depth <= 0 {
			inBlock = false
		}
	}
	return strings.Join(result, "\n")
}

func generateProtos(ctx context.Context, repoRoot string, verbose bool) error {
	args := []string{
		"generate",
		"--project-dir", repoRoot,
		"--language", "go",
		"--rpc", "none",
		"--targets", "protocol/*.proto",
		"--force",
	}
	if verbose {
		args = append(args, "--verbose")
	}
	return runAptre(ctx, repoRoot, verbose, args...)
}

func runApigen(ctx context.Context, repoRoot string, verbose bool) error {
	if err := runCmd(ctx, repoRoot, verbose, "go", "mod", "vendor"); err != nil {
		return fmt.Errorf("go mod vendor: %w", err)
	}

	apigenDir := filepath.Join(repoRoot, "apigen")
	apigenBin := filepath.Join(apigenDir, "apigen")
	if err := runCmd(ctx, apigenDir, verbose, "go", "build", "-o", "apigen", "./"); err != nil {
		return fmt.Errorf("build apigen: %w", err)
	}

	cmds := []struct {
		args     []string
		snapshot string
	}{
		{[]string{apigenBin, "print-type-list"}, "snapshot-type-list.txt"},
		{[]string{apigenBin, "generate-api"}, "snapshot-apigen.txt"},
		{[]string{apigenBin, "print-messages"}, "snapshot-messages.txt"},
	}

	for _, cmdSpec := range cmds {
		label := cmdSpec.args[len(cmdSpec.args)-1]
		log.Printf("  apigen %s...", label)
		cmd := exec.CommandContext(ctx, cmdSpec.args[0], cmdSpec.args[1:]...)
		cmd.Dir = apigenDir
		cmd.Env = append(os.Environ(), "GO111MODULE=on")

		var buf bytes.Buffer
		cmd.Stdout = &buf
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("apigen %s: %w", label, err)
		}

		snapshotPath := filepath.Join(apigenDir, cmdSpec.snapshot)
		if err := os.WriteFile(snapshotPath, buf.Bytes(), 0o644); err != nil {
			return fmt.Errorf("write %s: %w", cmdSpec.snapshot, err)
		}
		if verbose {
			fmt.Print(buf.String())
		}
	}
	return nil
}

func runAptre(ctx context.Context, repoRoot string, verbose bool, args ...string) error {
	goArgs := append([]string{"run", aptreModule}, args...)
	return runCmd(ctx, repoRoot, verbose, "go", goArgs...)
}

func runCmd(ctx context.Context, dir string, verbose bool, name string, args ...string) error {
	cmd := exec.CommandContext(ctx, name, args...)
	cmd.Dir = dir
	if verbose {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	} else {
		var stderr bytes.Buffer
		cmd.Stderr = &stderr
		defer func() {
			if stderr.Len() > 0 {
				fmt.Fprint(os.Stderr, stderr.String())
			}
		}()
	}
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("%s %s: %w", name, strings.Join(args, " "), err)
	}
	return nil
}
