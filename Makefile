.DEFAULT_GOAL := help

BIN_DIR := bin
CMD_DIR := cmd

COMMANDS := $(notdir $(wildcard $(CMD_DIR)/*))

.PHONY: help
help:
	@echo "Available targets:"
	@echo "  make test        - run unit tests"
	@echo "  make race        - run race tests"
	@echo "  make fmt         - format code (gofmt + goimports)"
	@echo "  make lint        - run golangci-lint"
	@echo "  make build       - build all binaries"
	@echo "  make ci          - run all CI checks"

.PHONY: test
test:
	go test ./...

.PHONY: race
race:
	go test -race ./...

.PHONY: fmt
fmt:
	gofmt -w .
	goimports -w .

.PHONY: lint
lint:
	golangci-lint run

.PHONY: build
build: $(COMMANDS:%=build-%)

# ------------------------
# Dynamic per-command build
# ------------------------
.PHONY: $(COMMANDS:%=build-%)
$(COMMANDS:%=build-%):
	@cmd=$(@:build-%=%); \
	echo "🔨 Building $$cmd"; \
	mkdir -p $(BIN_DIR); \
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
	go build -o $(BIN_DIR)/$$cmd ./$(CMD_DIR)/$$cmd

.PHONY: ci
ci: lint test race build
