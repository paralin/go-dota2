BIN = apigen
GO_FILES = $(shell find . -type f -name '*.go')

.PHONY: go-generate
go-generate:
	go generate .

.PHONY: build
build: $(BIN)

$(BIN): $(GO_FILES)
	go build -o $@ .

.PHONY: generate-api
generate-api: go-generate build
	./apigen generate-api
