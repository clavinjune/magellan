include scripts/make/buf.mk
include scripts/make/build.mk
include scripts/make/docker.mk
include scripts/make/test.mk
include scripts/make/tools.mk

REVISION:=$(shell git rev-parse --short HEAD)
APP_NAME:=biko

.PHONY: generate
generate: buf/generate wire

.PHONY: lint
lint:
	@go run $(BUF) lint
	@go vet ./...
	@go run $(GOIMPORTS) -w .
	@gofmt -w -s .
	@go mod tidy
	@go run $(GOVULNCHECK) ./...

.PHONY: wire
wire:
	@find . -name "*wire_gen.go" -delete
	@go run $(WIRE) ./...
