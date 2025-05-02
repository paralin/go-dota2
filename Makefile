# https://github.com/aperturerobotics/template

SHELL:=bash
PROTOC=protoc
PROTOWRAP=tools/bin/protowrap
PROTOC_GEN_GO=tools/bin/protoc-gen-go
GOIMPORTS=tools/bin/goimports
GOFUMPT=tools/bin/gofumpt
GOLANGCI_LINT=tools/bin/golangci-lint
GO_MOD_OUTDATED=tools/bin/go-mod-outdated
GOLIST=go list -f "{{ .Dir }}" -m

export GO111MODULE=on
undefine GOARCH
undefine GOOS

all:

vendor:
	go mod vendor

$(PROTOC_GEN_GO):
	cd ./tools; \
	go build -v \
		-o ./bin/protoc-gen-go \
		google.golang.org/protobuf/cmd/protoc-gen-go

$(GOIMPORTS):
	cd ./tools; \
	go build -v \
		-o ./bin/goimports \
		golang.org/x/tools/cmd/goimports

$(GOFUMPT):
	cd ./tools; \
	go build -v \
		-o ./bin/gofumpt \
		mvdan.cc/gofumpt

$(PROTOWRAP):
	cd ./tools; \
	go build -v \
		-o ./bin/protowrap \
		github.com/aperturerobotics/goprotowrap/cmd/protowrap

$(GOLANGCI_LINT):
	cd ./tools; \
	go build -v \
		-o ./bin/golangci-lint \
		github.com/golangci/golangci-lint/v2/cmd/golangci-lint

$(GO_MOD_OUTDATED):
	cd ./tools; \
	go build -v \
		-o ./bin/go-mod-outdated \
		github.com/psampaz/go-mod-outdated

.PHONY: genproto
genproto: vendor $(GOIMPORTS) $(PROTOWRAP) $(PROTOC_GEN_GO)
	export PATH=$$(pwd)/tools/bin:$${PATH}; \
	cd ./protocol; \
	GO_OPTS="paths=source_relative"; \
	for proto in *.proto; do \
		GO_OPTS="$$GO_OPTS,M$$proto=./;protocol"; \
	done; \
	$(PROTOC) \
		-I ./ \
		--go_out=./ \
		--go_opt="$$GO_OPTS" \
		./*.proto
	$(GOIMPORTS) -w ./

.PHONY: updateproto
updateproto:
	bash generator/update_protos.bash

.PHONY: apigen
apigen:
	cd ./apigen && bash apigen.bash

.PHONY: gen
gen: updateproto genproto apigen

.PHONY: outdated
outdated: $(GO_MOD_OUTDATED)
	go list -mod=mod -u -m -json all | $(GO_MOD_OUTDATED) -update -direct

.PHONY: list
list: $(GO_MOD_OUTDATED)
	go list -mod=mod -u -m -json all | $(GO_MOD_OUTDATED)

.PHONY: lint
lint: $(GOLANGCI_LINT)
	$(GOLANGCI_LINT) run

.PHONY: fix
fix: $(GOLANGCI_LINT)
	$(GOLANGCI_LINT) run --fix

.PHONY: test
test:
	go test -v ./...

.PHONY: format
format: $(GOFUMPT) $(GOIMPORTS)
	$(GOIMPORTS) -w ./
	$(GOFUMPT) -w ./
