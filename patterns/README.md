# Go Patterns and Examples

This directory contains comprehensive examples of Go language concepts, design patterns, and best practices. Perfect for demonstrating Go skills in technical interviews.

## 📁 Structure

```
patterns/
├── pkg/                    # Reusable Go packages
│   ├── interfaces/         # Interface definitions
│   │   ├── reconciler.go   # Reconciler interface
│   │   └── logger.go       # Logger interface
│   ├── patterns/           # Design patterns
│   │   ├── observer.go     # Observer pattern
│   │   ├── strategy.go     # Strategy pattern
│   │   └── builder.go      # Builder pattern
│   └── utils/              # Utility functions
│       ├── retry.go        # Retry mechanism
│       └── validation.go   # Validation utilities
├── examples/               # Go concept examples
│   ├── algorithms/         # Algorithm implementations
│   │   ├── binary_search.go # Binary search algorithm
│   │   └── quick_sort.go   # Quick sort algorithm
│   ├── concurrency/        # Concurrency patterns
│   │   ├── worker_pool.go  # Worker pool implementation
│   │   └── pipeline.go     # Pipeline pattern
│   ├── data-structures/    # Data structure implementations
│   │   ├── stack.go        # Stack implementation
│   │   └── queue.go        # Queue implementation
│   ├── patterns/           # Design pattern examples
│   │   ├── singleton.go    # Singleton pattern
│   │   └── observer.go     # Observer pattern
│   └── testing/            # Testing examples
│       ├── mocks.go        # Mock implementations
│       └── table_driven.go # Table-driven tests
├── tests/                  # Test files
│   ├── binary_search_test.go # Binary search tests
│   ├── singleton_test.go   # Singleton pattern tests
│   └── stack_test.go       # Stack tests
├── cmd/                    # Command-line applications
│   └── cli/                # CLI demonstration
│       └── main.go         # CLI with subcommands
└── README.md               # This file
```

## 🎯 Go Concepts Demonstrated

### **1. Interfaces and Polymorphism**
- **Location**: `pkg/interfaces/`
- **Key Files**: `reconciler.go`, `logger.go`
- **Concepts**: Interface design, dependency injection, mock implementations

**Example Usage**:
```go
// Define a logger interface
type Logger interface {
    Info(msg string, keysAndValues ...interface{})
    Error(err error, msg string, keysAndValues ...interface{})
}

// Use dependency injection
func NewService(logger Logger) *Service {
    return &Service{logger: logger}
}
```

### **2. Algorithms**
- **Location**: `examples/algorithms/`
- **Key Files**: `binary_search.go`, `quick_sort.go`
- **Concepts**: Binary search, Quick sort, Recursive algorithms, Time complexity

**Example Usage**:
```go
// Binary search
arr := []int{1, 2, 3, 4, 5}
index := algorithms.BinarySearch(arr, 3) // Returns 2

// Quick sort
unsorted := []int{64, 34, 25, 12, 22, 11, 90}
algorithms.QuickSort(unsorted)
```

### **3. Data Structures**
- **Location**: `examples/data-structures/`
- **Key Files**: `stack.go`, `queue.go`
- **Concepts**: Stack (LIFO), Queue (FIFO), Generic implementations

**Example Usage**:
```go
// Stack operations
stack := datastructures.NewStack()
stack.Push(1)
stack.Push(2)
top := stack.Peek() // Returns 2
popped := stack.Pop() // Returns 2

// Queue operations
queue := datastructures.NewQueue()
queue.Enqueue("first")
queue.Enqueue("second")
front := queue.Dequeue() // Returns "first"
```

### **4. Design Patterns**
- **Location**: `pkg/patterns/` and `examples/patterns/`
- **Key Files**: `observer.go`, `strategy.go`, `builder.go`, `singleton.go`
- **Concepts**: Observer pattern, Strategy pattern, Builder pattern, Singleton pattern

**Example Usage**:
```go
// Observer pattern
eventBus := NewEventBus()
eventBus.Subscribe(observer)
eventBus.NotifyObservers(ctx, event)

// Strategy pattern
processor := NewDataProcessor(jsonStrategy)
result, err := processor.Process(ctx, data)

// Builder pattern
config := NewWebServerConfigBuilder().
    Host("localhost").
    Port(8080).
    WithTLS("cert.pem", "key.pem").
    Build()
```

### **5. Concurrency Patterns**
- **Location**: `examples/concurrency/`
- **Key Files**: `worker_pool.go`, `pipeline.go`
- **Concepts**: Worker pools, pipelines, channel communication, context usage

**Example Usage**:
```go
// Worker pool
pool := NewWorkerPool(3, 10)
pool.Start()
defer pool.Stop()

// Pipeline
pipeline := NewPipeline[int]()
pipeline.AddStage(doubleStage)
pipeline.AddStage(filterStage)
output := pipeline.Process(ctx, input)
```

### **6. Error Handling**
- **Location**: `pkg/utils/`
- **Key Files**: `retry.go`, `validation.go`
- **Concepts**: Custom error types, error wrapping, retry mechanisms, validation

**Example Usage**:
```go
// Retry mechanism
err := RetryWithBackoff(ctx, func() error {
    return someOperation()
})

// Validation
validator := &StringValidator{
    Field: "name",
    Value: "John",
    MinLen: 1,
    MaxLen: 50,
    Required: true,
}
err := validator.Validate()
```

### **7. Testing**
- **Location**: `examples/testing/` and `tests/`
- **Key Files**: `mocks.go`, `table_driven.go`, `*_test.go`
- **Concepts**: Table-driven tests, mock implementations, benchmarking, unit testing

**Example Usage**:
```go
// Table-driven test
tests := []struct {
    name     string
    input    int
    expected int
}{
    {"positive", 5, 25},
    {"negative", -3, 9},
    {"zero", 0, 0},
}

for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
        result := Square(tt.input)
        if result != tt.expected {
            t.Errorf("Square(%d) = %d, expected %d", tt.input, result, tt.expected)
        }
    })
}
```

### **8. CLI Development**
- **Location**: `cmd/cli/`
- **Key Files**: `main.go`
- **Concepts**: Flag parsing, subcommands, context usage

**Example Usage**:
```bash
# Validate input
go run cmd/cli/main.go validate -input "hello world"

# Retry with custom config
go run cmd/cli/main.go retry -max-attempts 5 -base-delay 200ms

# Start server
go run cmd/cli/main.go server -host 0.0.0.0 -port 9090
```

## 🚀 Quick Start

### **Run Examples**
```bash
# CLI example (only runnable example)
go run cmd/cli/main.go validate -input "test"

# Note: Most examples are packages, not runnable programs
# Use 'make run-examples' to see available runnable examples
```

### **Run Tests**
```bash
# Run all tests
make test

# Run specific test categories
make test-algorithms
make test-data-structures
make test-patterns

# Run tests with verbose output
make test-verbose

# Run benchmarks
make benchmark
```

### **Build Examples**
```bash
# Build all examples
make build

# Build specific examples
go build -o bin/cli cmd/cli/main.go
go build -o bin/worker-pool examples/concurrency/worker_pool.go

# Use the CLI
./bin/cli validate -input "hello world"
```

## 🛠️ Makefile Commands

The project includes a comprehensive Makefile with the following targets:

### **Development Commands**
```bash
make help              # Display all available commands
make build             # Build all examples
make clean             # Clean build artifacts
make fmt               # Format code
make vet               # Run go vet
make mod-tidy          # Tidy go modules
make all               # Run all checks and build
```

### **Testing Commands**
```bash
make test              # Run all tests
make test-verbose      # Run tests with verbose output
make test-algorithms   # Run algorithm tests
make test-data-structures # Run data structure tests
make test-patterns     # Run design pattern tests
make test-examples     # Run example tests
make benchmark         # Run benchmarks
```

### **Example Commands**
```bash
make run-examples      # Run all examples (CLI only)
make run-algorithms    # Show algorithm examples info
make run-data-structures # Show data structure examples info
make run-patterns      # Show design pattern examples info
make run-concurrency   # Show concurrency examples info
make run-cli           # Run CLI examples
```

## 🎤 Interview Scenarios

### **Scenario 1: Go Fundamentals**
**Focus**: Core Go concepts and language features
1. Show interface design (`pkg/interfaces/`)
2. Demonstrate error handling (`pkg/utils/`)
3. Explain concurrency patterns (`examples/concurrency/`)
4. Walk through testing examples (`examples/testing/` and `tests/`)
5. Show algorithm implementations (`examples/algorithms/`)
6. Demonstrate data structures (`examples/data-structures/`)

### **Scenario 2: Design Patterns**
**Focus**: Software design and architecture
1. Explain Observer pattern (`pkg/patterns/observer.go`)
2. Show Strategy pattern (`pkg/patterns/strategy.go`)
3. Demonstrate Builder pattern (`pkg/patterns/builder.go`)
4. Show Singleton pattern (`examples/patterns/singleton.go`)
5. Discuss when to use each pattern

### **Scenario 3: Algorithms and Data Structures**
**Focus**: Computer science fundamentals and problem-solving
1. Show binary search implementation (`examples/algorithms/binary_search.go`)
2. Explain quick sort algorithm (`examples/algorithms/quick_sort.go`)
3. Demonstrate stack operations (`examples/data-structures/stack.go`)
4. Show queue implementation (`examples/data-structures/queue.go`)
5. Discuss time and space complexity
6. Walk through test cases (`tests/`)

### **Scenario 4: Concurrency**
**Focus**: Concurrent programming and performance
1. Show worker pool implementation (`examples/concurrency/worker_pool.go`)
2. Explain pipeline pattern (`examples/concurrency/pipeline.go`)
3. Discuss channel communication and synchronization
4. Demonstrate context usage

### **Scenario 5: Testing and Quality**
**Focus**: Testing strategies and code quality
1. Show table-driven tests (`examples/testing/table_driven.go`)
2. Demonstrate mock implementations (`examples/testing/mocks.go`)
3. Explain benchmarking (`examples/testing/table_driven.go`)
4. Discuss test organization and best practices
5. Show unit tests for algorithms and data structures (`tests/`)

## 🛠️ Common Interview Questions

### **Interfaces**
- "How do you design interfaces in Go?"
- "What's the difference between interfaces and concrete types?"
- "How do you mock dependencies for testing?"

### **Concurrency**
- "How do you implement a worker pool in Go?"
- "Explain the difference between channels and mutexes."
- "How do you handle context cancellation?"

### **Algorithms**
- "How do you implement binary search in Go?"
- "What's the time complexity of quick sort?"
- "How do you handle edge cases in recursive algorithms?"
- "Explain the difference between iterative and recursive approaches"

### **Data Structures**
- "How do you implement a stack in Go?"
- "What's the difference between a stack and a queue?"
- "How do you handle memory management in data structures?"
- "When would you use a stack vs a queue?"

### **Error Handling**
- "How do you handle errors in Go applications?"
- "What's the difference between panic and error?"
- "How do you implement retry logic?"

### **Testing**
- "How do you write table-driven tests?"
- "How do you mock external dependencies?"
- "What's the difference between unit tests and integration tests?"

## 🏆 Key Strengths

1. **Clean Architecture**: Well-organized code with clear separation of concerns
2. **Interface Design**: Proper use of interfaces for testability and flexibility
3. **Algorithm Implementation**: Efficient implementations of common algorithms
4. **Data Structures**: Well-tested implementations of fundamental data structures
5. **Error Handling**: Comprehensive error handling with custom error types
6. **Testing**: Table-driven tests, unit tests, and proper mocking strategies
7. **Concurrency**: Safe concurrent programming with proper synchronization
8. **Documentation**: Comprehensive documentation and examples
9. **Real-World Examples**: Practical implementations of common patterns
10. **Makefile Integration**: Easy-to-use build and test automation

This collection demonstrates strong Go language skills and software engineering best practices, making it perfect for technical interviews focused on Go development.
