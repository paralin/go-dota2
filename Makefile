PROTOWRAP=hack/bin/protowrap
PROTOC_GEN_GO=hack/bin/protoc-gen-go
GOLANGCI_LINT=hack/bin/golangci-lint
export GO111MODULE=on
GOLIST=go list -f "{{ .Dir }}" -m

all: deps
	cd ./generator && ./update_protos.sh
	cd ./apigen && ./apigen.sh

deps:
	go install github.com/square/goprotowrap/cmd/protowrap@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

lint:
	golangci-lint run ./...
