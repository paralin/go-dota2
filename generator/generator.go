/*
Adapted from go-steam's generator/generator.go to generate dota 2 proto files.

https://github.com/Philipp15b/go-steam/blob/master/generator/generator.go

-------------------------------------------------------------------------------

Copyright (c) 2014 The go-steam Authors. All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are
met:

   * Redistributions of source code must retain the above copyright
notice, this list of conditions and the following disclaimer.
   * Redistributions in binary form must reproduce the above
copyright notice, this list of conditions and the following disclaimer
in the documentation and/or other materials provided with the
distribution.
   * The names of its contributors may not be used to endorse or promote
products derived from this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
"AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

*/

/*
This program generates the protobuf and SteamLanguage files from the SteamKit data.
*/
package main

import (
	"bytes"
	"go/parser"
	"go/token"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

var printCommands = false

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
	"base_gcmessages.proto":        "base.pb.go",
	"gcsdk_gcmessages.proto":       "gcsdk.pb.go",
	"dota_gcmessages_client.proto": "dota_client.pb.go",
	"dota_gcmessages_common.proto": "dota_common.pb.go",
	"gcsystemmsgs.proto":           "system.pb.go",
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

var pkgRegex = regexp.MustCompile(`(package \w+)`)
var pkgCommentRegex = regexp.MustCompile(`(?s)(\/\*.*?\*\/\n)package`)

func fixProto(path string) {
	// goprotobuf is really bad at dependencies, so we must fix them manually...
	// It tries to load each dependency of a file as a seperate package (but in a very, very wrong way).
	// Because we want some files in the same package, we'll remove those imports to local files.

	file, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, path, file, parser.ImportsOnly)
	if err != nil {
		panic("Error parsing " + path + ": " + err.Error())
	}

	importsToRemove := make([]string, 0)
	for _, i := range f.Imports {
		// We remove all imports that include ".pb". This assumes unified and protobuf packages don't share anything.
		if i.Name.Name != "google_protobuf" && strings.Contains(i.Path.Value, ".pb") {
			importsToRemove = append(importsToRemove, i.Name.Name)
		}
	}

	for _, itr := range importsToRemove {
		// remove the package name from all types
		file = bytes.Replace(file, []byte(itr+"."), []byte{}, -1)
		// and remove the import itself
		file = bytes.Replace(file, []byte("import "+itr+" \"pb\"\n"), []byte{}, -1)
	}

	// remove the package comment because it just includes a list of all messages and
	// creates collisions between the others.
	file = cutAllSubmatch(pkgCommentRegex, file, 1)

	// fix the package name
	file = pkgRegex.ReplaceAll(file, []byte("package "+inferPackageName(path)))

	// fix the google dependency;
	// we just reuse the one from protoc-gen-go
	file = bytes.Replace(file, []byte("google/protobuf/descriptor.pb"), []byte("code.google.com/p/goprotobuf/protoc-gen-go/descriptor"), -1)

	err = ioutil.WriteFile(path, file, os.ModePerm)
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
