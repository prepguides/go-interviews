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
│   ├── concurrency/        # Concurrency patterns
│   │   ├── worker_pool.go  # Worker pool implementation
│   │   └── pipeline.go     # Pipeline pattern
│   └── testing/            # Testing examples
│       ├── mocks.go        # Mock implementations
│       └── table_driven.go # Table-driven tests
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

### **2. Design Patterns**
- **Location**: `pkg/patterns/`
- **Key Files**: `observer.go`, `strategy.go`, `builder.go`
- **Concepts**: Observer pattern, Strategy pattern, Builder pattern

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

### **3. Concurrency Patterns**
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

### **4. Error Handling**
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

### **5. Testing**
- **Location**: `examples/testing/`
- **Key Files**: `mocks.go`, `table_driven.go`
- **Concepts**: Table-driven tests, mock implementations, benchmarking

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

### **6. CLI Development**
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
# Worker pool example
go run examples/concurrency/worker_pool.go

# Pipeline example
go run examples/concurrency/pipeline.go

# CLI example
go run cmd/cli/main.go validate -input "test"
```

### **Run Tests**
```bash
# Run all tests
go test ./...

# Run specific test
go test ./examples/testing -v

# Run benchmarks
go test -bench=. ./examples/testing
```

### **Build CLI**
```bash
# Build CLI tool
go build -o bin/cli cmd/cli/main.go

# Use the CLI
./bin/cli validate -input "hello world"
```

## 🎤 Interview Scenarios

### **Scenario 1: Go Fundamentals**
**Focus**: Core Go concepts and language features
1. Show interface design (`pkg/interfaces/`)
2. Demonstrate error handling (`pkg/utils/`)
3. Explain concurrency patterns (`examples/concurrency/`)
4. Walk through testing examples (`examples/testing/`)

### **Scenario 2: Design Patterns**
**Focus**: Software design and architecture
1. Explain Observer pattern (`pkg/patterns/observer.go`)
2. Show Strategy pattern (`pkg/patterns/strategy.go`)
3. Demonstrate Builder pattern (`pkg/patterns/builder.go`)
4. Discuss when to use each pattern

### **Scenario 3: Concurrency**
**Focus**: Concurrent programming and performance
1. Show worker pool implementation (`examples/concurrency/worker_pool.go`)
2. Explain pipeline pattern (`examples/concurrency/pipeline.go`)
3. Discuss channel communication and synchronization
4. Demonstrate context usage

### **Scenario 4: Testing and Quality**
**Focus**: Testing strategies and code quality
1. Show table-driven tests (`examples/testing/table_driven.go`)
2. Demonstrate mock implementations (`examples/testing/mocks.go`)
3. Explain benchmarking (`examples/testing/table_driven.go`)
4. Discuss test organization and best practices

## 🛠️ Common Interview Questions

### **Interfaces**
- "How do you design interfaces in Go?"
- "What's the difference between interfaces and concrete types?"
- "How do you mock dependencies for testing?"

### **Concurrency**
- "How do you implement a worker pool in Go?"
- "Explain the difference between channels and mutexes."
- "How do you handle context cancellation?"

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
3. **Error Handling**: Comprehensive error handling with custom error types
4. **Testing**: Table-driven tests and proper mocking strategies
5. **Concurrency**: Safe concurrent programming with proper synchronization
6. **Documentation**: Comprehensive documentation and examples
7. **Real-World Examples**: Practical implementations of common patterns

This collection demonstrates strong Go language skills and software engineering best practices, making it perfect for technical interviews focused on Go development.
