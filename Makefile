# Build variables
VERSION ?= 0.1.0
BUILD_TIME = $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
LDFLAGS = -ldflags "-X main.Version=$(VERSION) -X main.BuildTime=$(BUILD_TIME)"

.PHONY: build clean run test lint help

build:
	@echo "Building application..."
	@go build $(LDFLAGS) -o bin/dynamic ./cmd/...

clean:
	@echo "Cleaning build artifacts..."
	@rm -rf bin/
	@go clean

run: build
	@echo "Running application..."
	@./bin/dynamic

test:
	@echo "Running tests..."
	@go test -v ./...

lint:
	@echo "Linting code..."
	@if command -v golangci-lint &> /dev/null; then \
		golangci-lint run ./...; \
	else \
		echo "golangci-lint not installed. Run: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest"; \
		exit 1; \
	fi

help:
	@echo "Available commands:"
	@echo "  make build       - Build the application"
	@echo "  make clean       - Remove build artifacts"
	@echo "  make run         - Build and run the application"
	@echo "  make test        - Run tests"
	@echo "  make lint        - Run linter"
	@echo "  make help        - Show this help message"

# Default target
.DEFAULT_GOAL := help