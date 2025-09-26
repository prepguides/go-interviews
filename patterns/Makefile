# Makefile for Go Patterns and Examples

.PHONY: help
help: ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Development

.PHONY: build
build: ## Build all examples
	go build -o bin/worker-pool examples/concurrency/worker_pool.go
	go build -o bin/pipeline examples/concurrency/pipeline.go
	go build -o bin/cli cmd/cli/main.go

.PHONY: run-examples
run-examples: ## Run all examples
	@echo "Running worker pool example..."
	go run examples/concurrency/worker_pool.go
	@echo "\nRunning pipeline example..."
	go run examples/concurrency/pipeline.go

.PHONY: run-cli
run-cli: ## Run CLI examples
	@echo "Running CLI validation example..."
	go run cmd/cli/main.go validate -input "hello world"
	@echo "\nRunning CLI retry example..."
	go run cmd/cli/main.go retry -max-attempts 3

.PHONY: test
test: ## Run tests
	go test ./...

.PHONY: test-verbose
test-verbose: ## Run tests with verbose output
	go test -v ./...

.PHONY: benchmark
benchmark: ## Run benchmarks
	go test -bench=. ./...

.PHONY: clean
clean: ## Clean build artifacts
	rm -rf bin/

.PHONY: fmt
fmt: ## Format code
	go fmt ./...

.PHONY: vet
vet: ## Run go vet
	go vet ./...

.PHONY: mod-tidy
mod-tidy: ## Tidy go modules
	go mod tidy

.PHONY: all
all: fmt vet test build ## Run all checks and build
