APP_NAME=serverpulse

.PHONY: help
help:
	@echo "Available targets:"
	@echo "  make test        - run unit tests"
	@echo "  make race        - run race tests"
	@echo "  make fmt         - format code using golangci-lint"
	@echo "  make lint        - run golangci-lint"
	@echo "  make build       - build binary"
	@echo "  make ci          - run all CI checks"

.PHONY: test
test:
	go test ./...

.PHONY: race
race:
	go test -race ./...

.PHONY: fmt
fmt:
	golangci-lint fmt

.PHONY: lint
lint:
	golangci-lint run

.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
	go build -o bin/$(APP_NAME) ./cmd

.PHONY: ci
ci: lint test race build
