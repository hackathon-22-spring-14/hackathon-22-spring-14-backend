BINARY := ./main

.PHONY: all
all: clean mod build

.PHONY: clean
clean:
	@go clean $(BINARY)

.PHONY: mod
mod:
	@go mod download

.PHONY: build
build:
	@go build -o main ./...

# run with live reload
.PHONY: dev
dev:
	@go run github.com/cosmtrek/air@latest

.PHONY: lint
lint:
	@go run github.com/golangci/golangci-lint/cmd/golangci-lint@latest run --fix ./...
