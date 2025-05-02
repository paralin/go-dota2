//go:build deps_only
// +build deps_only

package tools

import (
	// _ imports protowrap
	_ "github.com/aperturerobotics/goprotowrap/cmd/protowrap"
	// _ imports protoc-gen-go-lite
	_ "github.com/aperturerobotics/protobuf-go-lite/cmd/protoc-gen-go-lite"
	// _ imports golangci-lint
	_ "github.com/golangci/golangci-lint/v2/cmd/golangci-lint"
	// _ imports golangci-lint commands
	_ "github.com/golangci/golangci-lint/v2/pkg/commands"
	// _ imports go-mod-outdated
	_ "github.com/psampaz/go-mod-outdated"
	// _ imports goimports
	_ "golang.org/x/tools/cmd/goimports"
	// _ imports gofumpt
	_ "mvdan.cc/gofumpt"
)
