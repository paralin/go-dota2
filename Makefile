# https://github.com/aperturerobotics/template

SHELL:=bash
APTRE=go run github.com/aperturerobotics/common/cmd/aptre@v0.34.1

export GO111MODULE=on
undefine GOARCH
undefine GOOS

all:

vendor:
	go mod vendor

.PHONY: genproto
genproto: vendor
	$(APTRE) generate --language go --rpc none --targets 'protocol/*.proto' --force

.PHONY: updateproto
updateproto:
	go run ./cmd/generate -skip-apigen

.PHONY: apigen
apigen:
	cd ./apigen && go run . generate-api

.PHONY: gen
gen:
	go run ./cmd/generate

.PHONY: outdated
outdated:
	$(APTRE) outdated

.PHONY: list
list:
	$(APTRE) outdated

.PHONY: lint
lint:
	$(APTRE) lint

.PHONY: fix
fix:
	$(APTRE) fix

.PHONY: test
test:
	go test -v ./...

.PHONY: format
format:
	$(APTRE) goimports
	$(APTRE) format
