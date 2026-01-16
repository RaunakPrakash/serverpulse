.DEFAULT_GOAL := help

# ------------------------
# Directories & variables
# ------------------------
BUILD_DIR := build
CMD_DIR := cmd
VERSION ?= dev

COMMANDS := $(notdir $(wildcard $(CMD_DIR)/*))

# ------------------------
# Help
# ------------------------
.PHONY: help
help:
	@echo "Available targets:"
	@echo "  make test              - run unit tests"
	@echo "  make race              - run race tests"
	@echo "  make coverage          - run tests with coverage report"
	@echo "  make fmt               - format code (gofmt + goimports)"
	@echo "  make lint              - run golangci-lint"
	@echo "  make build             - build all binaries"
	@echo "  make mod-update        - update Go dependencies"
	@echo "  make mod-tidy          - tidy go.mod/go.sum"
	@echo "  make clean             - clean build & cache files"
	@echo "  make git-latest-release - show latest git tag"
	@echo "  make ci                - run all CI checks"

# ------------------------
# Testing
# ------------------------
.PHONY: test
test:
	go test ./...

.PHONY: race
race:
	go test -race ./...

.PHONY: coverage
coverage:
	go test ./... -coverprofile=coverage.out
	go tool cover -func=coverage.out

# ------------------------
# Formatting
# ------------------------
.PHONY: fmt
fmt:
	gofmt -w .
	goimports -w .

# ------------------------
# Linting
# ------------------------
.PHONY: lint
lint:
	golangci-lint run

# ------------------------
# Build (dynamic, per cmd/*)
# ------------------------
.PHONY: build
build: $(COMMANDS:%=build-%)

.PHONY: $(COMMANDS:%=build-%)
$(COMMANDS:%=build-%):
	@cmd=$(@:build-%=%); \
	echo "🔨 Building $$cmd"; \
	mkdir -p $(BUILD_DIR); \
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
	go build \
		-ldflags="-X main.version=$(VERSION)" \
		-o $(BUILD_DIR)/$$cmd \
		./$(CMD_DIR)/$$cmd

# ------------------------
# Go modules
# ------------------------
.PHONY: mod-update
mod-update:
	go get -u ./...
	$(MAKE) mod-tidy

.PHONY: mod-tidy
mod-tidy:
	rm -f go.sum
	go mod tidy -v

# ------------------------
# Git helpers
# ------------------------
.PHONY: git-latest-release
git-latest-release:
	@git tag --list --sort=v:refname \
		--format="%(refname:short) => %(creatordate:short)" \
		| tail -n 1

# ------------------------
# Cleanup
# ------------------------
.PHONY: golangci-lint-cache-clean
golangci-lint-cache-clean:
	golangci-lint cache clean

.PHONY: clean
clean:
	git clean -fdX
	go clean -cache -testcache
	$(MAKE) golangci-lint-cache-clean

# ------------------------
# CI
# ------------------------
.PHONY: ci
ci: lint test race build
