# Makefile for Go Patterns and Examples

.PHONY: help
help: ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Development

.PHONY: build
build: ## Build all examples
	go build -o bin/worker-pool examples/concurrency/worker_pool.go
	go build -o bin/pipeline examples/concurrency/pipeline.go
	go build -o bin/binary-search examples/algorithms/binary_search.go
	go build -o bin/quick-sort examples/algorithms/quick_sort.go
	go build -o bin/stack examples/data-structures/stack.go
	go build -o bin/queue examples/data-structures/queue.go
	go build -o bin/singleton examples/patterns/singleton.go
	go build -o bin/observer examples/patterns/observer.go
	go build -o bin/cli cmd/cli/main.go

.PHONY: run-examples
run-examples: ## Run all examples (only CLI has main function)
	@echo "Note: Most examples are packages, not runnable programs"
	@echo "Running CLI example..."
	go run cmd/cli/main.go validate -input "hello world"

.PHONY: run-algorithms
run-algorithms: ## Show algorithm examples (packages, not runnable)
	@echo "Algorithm examples are packages, not runnable programs"
	@echo "Use 'make test-algorithms' to test them"

.PHONY: run-data-structures
run-data-structures: ## Show data structure examples (packages, not runnable)
	@echo "Data structure examples are packages, not runnable programs"
	@echo "Use 'make test-data-structures' to test them"

.PHONY: run-patterns
run-patterns: ## Show design pattern examples (packages, not runnable)
	@echo "Design pattern examples are packages, not runnable programs"
	@echo "Use 'make test-patterns' to test them"

.PHONY: run-concurrency
run-concurrency: ## Show concurrency examples (packages, not runnable)
	@echo "Concurrency examples are packages, not runnable programs"
	@echo "Use 'make test' to test them"

.PHONY: run-cli
run-cli: ## Run CLI examples
	@echo "Running CLI validation example..."
	go run cmd/cli/main.go validate -input "hello world"
	@echo "\nRunning CLI retry example..."
	go run cmd/cli/main.go retry -max-attempts 3

.PHONY: test
test: ## Run all tests
	go test ./...

.PHONY: test-verbose
test-verbose: ## Run tests with verbose output
	go test -v ./...

.PHONY: test-algorithms
test-algorithms: ## Run algorithm tests
	go test -v ./tests -run TestBinarySearch

.PHONY: test-data-structures
test-data-structures: ## Run data structure tests
	go test -v ./tests -run TestStack

.PHONY: test-patterns
test-patterns: ## Run design pattern tests
	go test -v ./tests -run TestSingleton

.PHONY: test-examples
test-examples: ## Run example tests
	go test -v ./examples/...

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
