PROTOWRAP=hack/bin/protowrap
PROTOC_GEN_GO=hack/bin/protoc-gen-go
GOLANGCI_LINT=hack/bin/golangci-lint
export GO111MODULE=on
GOLIST=go list -f "{{ .Dir }}" -m

all:

vendor:
	go mod vendor

$(PROTOC_GEN_GO):
	cd ./hack; \
	go build -v \
		-o ./bin/protoc-gen-go \
		github.com/golang/protobuf/protoc-gen-go

$(PROTOWRAP):
	cd ./hack; \
	go build -v \
		-o ./bin/protowrap \
		github.com/square/goprotowrap/cmd/protowrap

$(GOLANGCI_LINT):
	cd ./hack; \
	go build -v \
		-o ./bin/golangci-lint \
		github.com/golangci/golangci-lint/cmd/golangci-lint

gengo: $(PROTOWRAP) $(PROTOC_GEN_GO) vendor
	export PATH=$$(pwd)/hack/bin:$${PATH}; \
	cd ./generator && ./update_protos.bash; \
	cd ../apigen && ./apigen.bash

lint: $(GOLANGCI_LINT)
	$(GOLANGCI_LINT) run ./...


