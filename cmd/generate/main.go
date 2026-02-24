// Command generate orchestrates the full go-dota2 protobuf generation pipeline
// using WASI-based protoc (no system protoc install required).
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

	protoc "github.com/aperturerobotics/go-protoc-wasi"
	"github.com/tetratelabs/wazero"
)

// sourceProtoGlobs defines which proto files to copy from the upstream submodule.
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

	if err := run(*verbose, *skipSubmodule, *skipApigen); err != nil {
		log.Fatal(err)
	}
}

func run(verbose, skipSubmodule, skipApigen bool) error {
	ctx := context.Background()

	// Determine repo root (two levels up from cmd/generate).
	exDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("getwd: %w", err)
	}
	// If run via "cd cmd/generate && go run .", cwd is cmd/generate.
	// Walk up to find the repo root by looking for go.mod at root.
	repoRoot, err := findRepoRoot(exDir)
	if err != nil {
		return fmt.Errorf("finding repo root: %w", err)
	}
	if verbose {
		log.Printf("repo root: %s", repoRoot)
	}

	protocolDir := filepath.Join(repoRoot, "protocol")
	generatorDir := filepath.Join(repoRoot, "generator")
	submoduleDir := filepath.Join(generatorDir, "Protobufs")
	toolsDir := filepath.Join(repoRoot, "tools")
	toolsBinDir := filepath.Join(toolsDir, "bin")

	// Step 1: Update submodule
	if !skipSubmodule {
		log.Println("==> Updating submodule...")
		if err := runCmd(repoRoot, verbose, "git", "submodule", "update", "--init", "generator/Protobufs"); err != nil {
			return fmt.Errorf("submodule update: %w", err)
		}
	}

	// Step 2: Preprocess protos
	log.Println("==> Preprocessing protos...")
	if err := preprocessProtos(submoduleDir, protocolDir, verbose); err != nil {
		return fmt.Errorf("preprocess protos: %w", err)
	}

	// Step 3: Build tools
	log.Println("==> Building tools...")
	if err := buildTools(toolsDir, toolsBinDir, verbose); err != nil {
		return fmt.Errorf("build tools: %w", err)
	}

	// Step 4: Run WASI protoc
	log.Println("==> Running protoc (WASI)...")
	if err := runProtoc(ctx, repoRoot, protocolDir, submoduleDir, toolsBinDir, verbose); err != nil {
		return fmt.Errorf("protoc: %w", err)
	}

	// Step 5: Run apigen
	if !skipApigen {
		log.Println("==> Running apigen...")
		if err := runApigen(repoRoot, verbose); err != nil {
			return fmt.Errorf("apigen: %w", err)
		}
	}

	// Step 6: Run goimports (after apigen so generated files are formatted)
	log.Println("==> Running goimports...")
	goimportsPath := filepath.Join(toolsBinDir, "goimports")
	if err := runCmd(repoRoot, verbose, goimportsPath, "-w", "./"); err != nil {
		return fmt.Errorf("goimports: %w", err)
	}

	log.Println("==> Done.")
	return nil
}

// findRepoRoot walks up from dir looking for the repository root.
// It identifies the root by the presence of go.mod with the expected module path.
func findRepoRoot(dir string) (string, error) {
	for {
		modPath := filepath.Join(dir, "go.mod")
		data, err := os.ReadFile(modPath)
		if err == nil && strings.Contains(string(data), "module github.com/paralin/go-dota2\n") {
			return dir, nil
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return "", fmt.Errorf("could not find repo root (go.mod with github.com/paralin/go-dota2)")
		}
		dir = parent
	}
}

// preprocessProtos replicates generator/update_protos.bash in pure Go.
func preprocessProtos(submoduleDir, protocolDir string, verbose bool) error {
	srcDir := filepath.Join(submoduleDir, "dota2")

	// Collect source proto files by matching globs.
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

	// Deduplicate (overlapping globs).
	seen := make(map[string]bool)
	var uniqueFiles []string
	for _, f := range srcFiles {
		if !seen[f] {
			seen[f] = true
			uniqueFiles = append(uniqueFiles, f)
		}
	}
	srcFiles = uniqueFiles

	if verbose {
		log.Printf("found %d source proto files", len(srcFiles))
	}

	// Track which proto files we're writing (for cleanup).
	writtenProtos := make(map[string]bool)

	// Process each proto file.
	for _, src := range srcFiles {
		data, err := os.ReadFile(src)
		if err != nil {
			return fmt.Errorf("read %s: %w", src, err)
		}

		content := string(data)

		// Prepend syntax and package.
		content = "syntax = \"proto2\";\npackage protocol;\n\n" + content

		// Apply string replacements.
		content = strings.ReplaceAll(content, "optional .", "optional ")
		content = strings.ReplaceAll(content, "required .", "required ")
		content = strings.ReplaceAll(content, "repeated .", "repeated ")
		content = strings.ReplaceAll(content, "\t.", "\t")
		content = strings.ReplaceAll(content, "google/protobuf/valve_extensions.proto", "valveextensions.proto")

		fname := filepath.Base(src)

		// Special fix for steammessages_steamlearn.
		if fname == "steammessages_steamlearn.steamworkssdk.proto" {
			content = strings.ReplaceAll(content, "(.CMsgSteamLearn", "(CMsgSteamLearn")
		}

		// Special fix: remove CDotaMsgStructuredTooltipProperties message block.
		if fname == "dota_gcmessages_common.proto" {
			content = removeMessageBlock(content, "CDotaMsgStructuredTooltipProperties")
		}

		dst := filepath.Join(protocolDir, fname)
		if err := os.WriteFile(dst, []byte(content), 0o644); err != nil {
			return fmt.Errorf("write %s: %w", dst, err)
		}
		writtenProtos[fname] = true
	}

	// Clean up old .proto files not in the new set (rsync --delete behavior).
	entries, err := os.ReadDir(protocolDir)
	if err != nil {
		return fmt.Errorf("read protocol dir: %w", err)
	}
	for _, e := range entries {
		if strings.HasSuffix(e.Name(), ".proto") && !writtenProtos[e.Name()] {
			old := filepath.Join(protocolDir, e.Name())
			if verbose {
				log.Printf("removing old proto: %s", e.Name())
			}
			if err := os.Remove(old); err != nil {
				return fmt.Errorf("remove %s: %w", old, err)
			}
		}
	}

	// Write protocol.go.
	protoGo := filepath.Join(protocolDir, "protocol.go")
	if err := os.WriteFile(protoGo, []byte("package protocol\n"), 0o644); err != nil {
		return fmt.Errorf("write protocol.go: %w", err)
	}

	return nil
}

// removeMessageBlock removes a top-level message block by name from proto content.
func removeMessageBlock(content, messageName string) string {
	var result []string
	lines := strings.Split(content, "\n")
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
		} else {
			// Count braces to handle nested messages.
			depth += strings.Count(line, "{") - strings.Count(line, "}")
			if depth <= 0 {
				inBlock = false
			}
		}
	}
	return strings.Join(result, "\n")
}

// buildTools builds protoc-gen-go and goimports if they don't already exist.
func buildTools(toolsDir, toolsBinDir string, verbose bool) error {
	if err := os.MkdirAll(toolsBinDir, 0o755); err != nil {
		return err
	}

	tools := []struct {
		binary string
		pkg    string
	}{
		{"protoc-gen-go", "google.golang.org/protobuf/cmd/protoc-gen-go"},
		{"goimports", "golang.org/x/tools/cmd/goimports"},
	}

	for _, t := range tools {
		binPath := filepath.Join(toolsBinDir, t.binary)
		if _, err := os.Stat(binPath); err == nil {
			if verbose {
				log.Printf("  %s already exists, skipping build", t.binary)
			}
			continue
		}
		log.Printf("  building %s...", t.binary)
		if err := runCmd(toolsDir, verbose, "go", "build", "-v", "-o", binPath, t.pkg); err != nil {
			return fmt.Errorf("build %s: %w", t.binary, err)
		}
	}
	return nil
}

// runProtoc runs protoc via WASI to generate Go code from proto files.
func runProtoc(ctx context.Context, repoRoot, protocolDir, submoduleDir, toolsBinDir string, verbose bool) error {
	// Collect proto files.
	protoFiles, err := filepath.Glob(filepath.Join(protocolDir, "*.proto"))
	if err != nil {
		return fmt.Errorf("glob protos: %w", err)
	}
	if len(protoFiles) == 0 {
		return fmt.Errorf("no proto files in %s", protocolDir)
	}

	// Build --go_opt with M mappings.
	goOpts := "paths=source_relative"
	for _, f := range protoFiles {
		base := filepath.Base(f)
		goOpts += ",M" + base + "=./;protocol"
	}

	// Copy google/protobuf well-known types to a temp dir for WASI protoc.
	wktTmpDir, err := os.MkdirTemp("", "go-dota2-wkt-*")
	if err != nil {
		return fmt.Errorf("create wkt temp dir: %w", err)
	}
	defer os.RemoveAll(wktTmpDir)

	wktSrcDir := filepath.Join(submoduleDir, "google", "protobuf")
	wktDstDir := filepath.Join(wktTmpDir, "google", "protobuf")
	if err := os.MkdirAll(wktDstDir, 0o755); err != nil {
		return err
	}
	if err := copyDir(wktSrcDir, wktDstDir); err != nil {
		return fmt.Errorf("copy well-known types: %w", err)
	}

	// Set up WASI protoc.
	runtime := wazero.NewRuntime(ctx)
	defer runtime.Close(ctx)

	var stdout, stderr bytes.Buffer
	fsConfig := wazero.NewFSConfig().
		WithDirMount(protocolDir, protocolDir).
		WithDirMount(wktTmpDir, wktTmpDir)

	protocGenGoPath := filepath.Join(toolsBinDir, "protoc-gen-go")
	cfg := &protoc.Config{
		Stdout:        &stdout,
		Stderr:        &stderr,
		FSConfig:      fsConfig,
		PluginHandler: &pluginHandler{protocGenGoPath: protocGenGoPath},
	}

	p, err := protoc.NewProtoc(ctx, runtime, cfg)
	if err != nil {
		return fmt.Errorf("create protoc: %w", err)
	}
	defer p.Close(ctx)

	if err := p.Init(ctx); err != nil {
		return fmt.Errorf("init protoc: %w", err)
	}

	// Build args.
	args := []string{
		"protoc",
		"-I", protocolDir,
		"-I", wktTmpDir,
		"--go_out=" + protocolDir,
		"--go_opt=" + goOpts,
	}
	for _, f := range protoFiles {
		args = append(args, f)
	}

	if verbose {
		log.Printf("  protoc args: %v", args)
	}

	exitCode, err := p.Run(ctx, args)
	if err != nil {
		return fmt.Errorf("run protoc: %w", err)
	}

	if verbose || exitCode != 0 {
		if stdout.Len() > 0 {
			fmt.Print(stdout.String())
		}
		if stderr.Len() > 0 {
			fmt.Fprint(os.Stderr, stderr.String())
		}
	}

	if exitCode != 0 {
		return fmt.Errorf("protoc exited with code %d", exitCode)
	}

	return nil
}

// pluginHandler implements protoc.PluginHandler by executing protoc-gen-go natively.
type pluginHandler struct {
	protocGenGoPath string
}

func (h *pluginHandler) Communicate(ctx context.Context, program string, searchPath bool, input []byte) ([]byte, error) {
	// Always use our built protoc-gen-go regardless of program name.
	cmd := exec.CommandContext(ctx, h.protocGenGoPath)
	cmd.Stdin = bytes.NewReader(input)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		if stderr.Len() > 0 {
			return nil, fmt.Errorf("%s: %w: %s", program, err, stderr.String())
		}
		return nil, fmt.Errorf("%s: %w", program, err)
	}
	return stdout.Bytes(), nil
}

// runApigen replicates apigen/apigen.bash.
func runApigen(repoRoot string, verbose bool) error {
	// go mod vendor (needed by apigen for reflection).
	if err := runCmd(repoRoot, verbose, "go", "mod", "vendor"); err != nil {
		return fmt.Errorf("go mod vendor: %w", err)
	}

	apigenDir := filepath.Join(repoRoot, "apigen")
	apigenBin := filepath.Join(apigenDir, "apigen")

	// Build apigen.
	if err := runCmd(apigenDir, verbose, "go", "build", "-o", "apigen", "./"); err != nil {
		return fmt.Errorf("build apigen: %w", err)
	}

	// Run apigen commands.
	cmds := []struct {
		args     []string
		snapshot string
	}{
		{[]string{apigenBin, "print-type-list"}, "snapshot-type-list.txt"},
		{[]string{apigenBin, "generate-api"}, "snapshot-apigen.txt"},
		{[]string{apigenBin, "print-messages"}, "snapshot-messages.txt"},
	}

	for _, c := range cmds {
		label := c.args[len(c.args)-1]
		log.Printf("  apigen %s...", label)

		cmd := exec.Command(c.args[0], c.args[1:]...)
		cmd.Dir = apigenDir
		cmd.Env = append(os.Environ(), "GO111MODULE=on")

		var buf bytes.Buffer
		cmd.Stdout = &buf
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			return fmt.Errorf("apigen %s: %w", label, err)
		}

		// Write snapshot file.
		snapshotPath := filepath.Join(apigenDir, c.snapshot)
		if err := os.WriteFile(snapshotPath, buf.Bytes(), 0o644); err != nil {
			return fmt.Errorf("write %s: %w", c.snapshot, err)
		}

		if verbose {
			fmt.Print(buf.String())
		}
	}

	return nil
}

// copyDir copies all files from src to dst (non-recursive, files only).
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

// runCmd runs an external command, optionally printing output.
func runCmd(dir string, verbose bool, name string, args ...string) error {
	cmd := exec.Command(name, args...)
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
	return cmd.Run()
}
