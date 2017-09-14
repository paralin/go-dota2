package main

import (
	"bytes"
	gofmt "go/format"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

var printCommands = false
var pkgImportPath = "github.com/paralin/go-dota2/protocol"

func main() {
	args := strings.Join(os.Args[1:], " ")

	found := false
	if strings.Contains(args, "clean") {
		clean()
		found = true
	}
	if strings.Contains(args, "proto") {
		buildProto()
		found = true
	}

	if !found {
		os.Stderr.WriteString("Invalid target!\nAvailable targets: clean, proto\n")
		os.Exit(1)
	}
}

func clean() {
	print("# Cleaning")
	cleanGlob("../internal/**/*.pb.go")
}

func cleanGlob(pattern string) {
	protos, _ := filepath.Glob(pattern)
	for _, proto := range protos {
		err := os.Remove(proto)
		if err != nil {
			panic(err)
		}
	}
}

func buildProto() {
	print("# Building Protobufs")

	buildProtoMap(dota2ProtoFiles, "../protocol")
}

func buildProtoMap(files map[string]string, outDir string) {
	os.MkdirAll(outDir, os.ModePerm)
	for proto, out := range files {
		full := filepath.Join(outDir, out)
		compileProto("Protobufs", proto, full)
		fixProto(full)
	}
}

// Maps the proto files to their target files.
// See `SteamKit/Resources/Protobufs/steamclient/generate-base.bat` for reference.
var dota2ProtoFiles = map[string]string{
	"base_gcmessages.proto":                         "base.pb.go",
	"gcsdk_gcmessages.proto":                        "gcsdk.pb.go",
	"dota_gcmessages_client.proto":                  "dota_client.pb.go",
	"dota_gcmessages_common.proto":                  "dota_common.pb.go",
	"dota_gcmessages_common_match_management.proto": "dota_common_match_management.pb.go",
	"dota_shared_enums.proto":                       "dota_shared_enums.pb.go",
	"steammessages.proto":                           "steammessages.pb.go",
	"gcsystemmsgs.proto":                            "system.pb.go",
}

func compileProto(srcBase, proto, target string) {
	outDir, _ := filepath.Split(target)
	err := os.MkdirAll(outDir, os.ModePerm)
	if err != nil {
		panic(err)
	}
	execute("protoc", "--go_out="+outDir, "-I="+srcBase, filepath.Join(srcBase, proto))
	out := strings.Replace(filepath.Join(outDir, proto), ".proto", ".pb.go", 1)
	err = forceRename(out, target)
	if err != nil {
		panic(err)
	}
}

func forceRename(from, to string) error {
	if from != to {
		os.Remove(to)
	}
	return os.Rename(from, to)
}

var pkgRegex = regexp.MustCompile(`(package )(\w+)`)
var pkgCommentRegex = regexp.MustCompile(`(?s)(\/\*.*?\*\/\n)package`)
var localImportRegex = regexp.MustCompile(`(import )(\w+)( ".")`)

func fixProto(path string) {
	// goprotobuf is really bad at dependencies, so we must fix them manually...
	// It tries to load each dependency of a file as a seperate package (but in a very, very wrong way).
	// Because we want some files in the same package, we'll remove those imports to local files.
	file, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	fileDir := filepath.Dir(path)

	// remove the package comment because it just includes a list of all messages and
	// creates collisions between the others.
	file = cutAllSubmatch(pkgCommentRegex, file, 1)

	// fix the package name
	// find the package name
	pkgNameParts := pkgRegex.FindSubmatch(file)
	if len(pkgNameParts) == 0 {
		panic("Error determining package name in " + path)
	}
	packageName := string(pkgNameParts[2])

	// fix the google dependency;
	// we just reuse the one from protoc-gen-go
	file = bytes.Replace(file, []byte("google/protobuf/descriptor.pb"), []byte("code.google.com/p/goprotobuf/protoc-gen-go/descriptor"), -1)
	file = bytes.Replace(file, []byte("import _ \".\"\n"), []byte{}, 1)
	file = bytes.Replace(file, []byte("import google_protobuf \"google/protobuf\""), []byte("import google_protobuf \"github.com/golang/protobuf/protoc-gen-go/descriptor\""), 1)

	matches := localImportRegex.FindAllSubmatch(file, -1)
	for _, match := range matches {
		completeMatch := match[0]
		pkgName := string(match[2])
		replaceWith := "import " + pkgName + " \"" + pkgImportPath + "/" + pkgName + "\""
		if pkgName == "_" {
			replaceWith = ""
		}
		print("Replacing " + string(completeMatch) + " with " + replaceWith)
		file = bytes.Replace(file, completeMatch, []byte(replaceWith), 1)
	}

	pkgSubdir := filepath.Join(fileDir, packageName)
	finalPath := filepath.Join(pkgSubdir, packageName+".go")
	if err := os.MkdirAll(pkgSubdir, 0755); err != nil {
		panic(err)
	}

	fmted, err := gofmt.Source(file)
	if err != nil {
		panic(err)
	}
	if err := os.Remove(path); err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(finalPath, fmted, 0755)
	if err != nil {
		panic(err)
	}
}

func inferPackageName(path string) string {
	pieces := strings.Split(path, string(filepath.Separator))
	return pieces[len(pieces)-2]
}

func cutAllSubmatch(r *regexp.Regexp, b []byte, n int) []byte {
	i := r.FindSubmatchIndex(b)
	return bytesCut(b, i[2*n], i[2*n+1])
}

// Removes the given section from the byte array
func bytesCut(b []byte, from, to int) []byte {
	buf := new(bytes.Buffer)
	buf.Write(b[:from])
	buf.Write(b[to:])
	return buf.Bytes()
}

func print(text string) { os.Stdout.WriteString(text + "\n") }

func printerr(text string) { os.Stderr.WriteString(text + "\n") }

// This writer appends a "> " after every newline so that the outpout appears quoted.
type QuotedWriter struct {
	w       io.Writer
	started bool
}

func NewQuotedWriter(w io.Writer) *QuotedWriter {
	return &QuotedWriter{w, false}
}

func (w *QuotedWriter) Write(p []byte) (n int, err error) {
	if !w.started {
		_, err = w.w.Write([]byte("> "))
		if err != nil {
			return n, err
		}
		w.started = true
	}

	for i, c := range p {
		if c == '\n' {
			nw, err := w.w.Write(p[n : i+1])
			n += nw
			if err != nil {
				return n, err
			}

			_, err = w.w.Write([]byte("> "))
			if err != nil {
				return n, err
			}
		}
	}
	if n != len(p) {
		nw, err := w.w.Write(p[n:len(p)])
		n += nw
		return n, err
	}
	return
}

func execute(command string, args ...string) {
	if printCommands {
		print(command + " " + strings.Join(args, " "))
	}
	cmd := exec.Command(command, args...)
	cmd.Stdout = NewQuotedWriter(os.Stdout)
	cmd.Stderr = NewQuotedWriter(os.Stderr)
	err := cmd.Run()
	if err != nil {
		printerr(err.Error())
		os.Exit(1)
	}
}
