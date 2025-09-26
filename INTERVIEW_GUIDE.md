# Go Interview Guide - Comprehensive Project Collection

This repository contains a comprehensive collection of Go projects designed to demonstrate various Go concepts and Kubernetes operator development skills for technical interviews.

## 🎯 Project Structure Overview

```
go-interviews/
├── operator/              # Kubernetes Operator Project
│   ├── api/              # CRD definitions
│   ├── controllers/      # Operator controller logic
│   ├── config/           # Kubernetes manifests
│   └── main.go           # Operator entry point
├── patterns/             # Go Patterns and Examples
│   ├── pkg/              # Reusable Go packages
│   ├── examples/         # Go concept examples
│   └── cmd/              # Command-line applications
└── INTERVIEW_GUIDE.md    # This file
```

## 📚 Go Concepts Demonstrated

### 1. **Interfaces and Polymorphism**
- **Location**: `patterns/pkg/interfaces/`
- **Key Files**: `reconciler.go`, `logger.go`
- **Interview Topics**:
  - Interface design principles
  - Dependency injection
  - Mock implementations for testing

**Example Questions**:
- "How would you design an interface for a logger that can be easily mocked?"
- "Explain the difference between concrete types and interfaces in Go."

### 2. **Design Patterns**
- **Location**: `patterns/pkg/patterns/`
- **Key Files**: `observer.go`, `strategy.go`, `builder.go`
- **Interview Topics**:
  - Observer pattern implementation
  - Strategy pattern for different processing types
  - Builder pattern for complex object construction

**Example Questions**:
- "Implement the Observer pattern in Go."
- "How would you use the Strategy pattern to handle different data formats?"

### 3. **Concurrency Patterns**
- **Location**: `patterns/examples/concurrency/`
- **Key Files**: `worker_pool.go`, `pipeline.go`
- **Interview Topics**:
  - Worker pools
  - Pipeline processing
  - Channel communication
  - Context usage

**Example Questions**:
- "How would you implement a worker pool in Go?"
- "Explain the difference between buffered and unbuffered channels."

### 4. **Error Handling**
- **Location**: `patterns/pkg/utils/retry.go`, `patterns/pkg/utils/validation.go`
- **Interview Topics**:
  - Custom error types
  - Error wrapping
  - Retry mechanisms
  - Validation patterns

**Example Questions**:
- "How do you handle errors in Go applications?"
- "Implement a retry mechanism with exponential backoff."

### 5. **Testing**
- **Location**: `patterns/examples/testing/`
- **Key Files**: `mocks.go`, `table_driven.go`
- **Interview Topics**:
  - Table-driven tests
  - Mock implementations
  - Benchmarking
  - Test helpers

**Example Questions**:
- "Write table-driven tests for a calculator function."
- "How do you mock dependencies in Go tests?"

### 6. **Command-Line Applications**
- **Location**: `patterns/cmd/cli/`
- **Key Files**: `main.go`
- **Interview Topics**:
  - Flag parsing
  - Subcommands
  - Context usage
  - CLI design

**Example Questions**:
- "How do you parse command-line arguments in Go?"
- "Design a CLI with subcommands."

## 🚀 Kubernetes Operator Concepts

### 1. **Custom Resource Definitions (CRDs)**
- **Location**: `operator/api/v1alpha1/`
- **Key Files**: `webserver_types.go`, `groupversion_info.go`
- **Interview Topics**:
  - CRD design
  - API versioning
  - Validation rules
  - Status subresources

### 2. **Controller Logic**
- **Location**: `operator/controllers/`
- **Key Files**: `webserver_controller.go`
- **Interview Topics**:
  - Reconciliation loops
  - Resource management
  - Owner references
  - Status updates

### 3. **Kubernetes Integration**
- **Location**: `operator/config/`
- **Key Files**: RBAC, CRDs, deployments
- **Interview Topics**:
  - RBAC configuration
  - Resource manifests
  - Operator deployment

## 🎤 Interview Scenarios

### **Scenario 1: Generic Go Developer**
**Focus**: Core Go concepts, patterns, and best practices

**Demo Flow**:
1. Show interface design (`patterns/pkg/interfaces/`)
2. Demonstrate design patterns (`patterns/pkg/patterns/`)
3. Walk through concurrency examples (`patterns/examples/concurrency/`)
4. Explain testing strategies (`patterns/examples/testing/`)
5. Show utility functions (`patterns/pkg/utils/`)

**Key Talking Points**:
- Interface design and dependency injection
- Concurrency patterns and channel usage
- Error handling strategies
- Testing methodologies

### **Scenario 2: Kubernetes/DevOps Engineer**
**Focus**: Operator development and Kubernetes integration

**Demo Flow**:
1. Explain CRD design (`operator/api/v1alpha1/`)
2. Walk through controller logic (`operator/controllers/`)
3. Show Kubernetes manifests (`operator/config/`)
4. Demonstrate end-to-end operator functionality
5. Access the web server to show it working

**Key Talking Points**:
- Operator pattern and reconciliation
- Resource management and garbage collection
- RBAC and security considerations
- Status management and conditions

### **Scenario 3: Full-Stack Go Developer**
**Focus**: Both Go fundamentals and Kubernetes knowledge

**Demo Flow**:
1. Start with Go patterns and interfaces
2. Show concurrency and testing examples
3. Transition to operator development
4. Demonstrate end-to-end functionality
5. Show CLI tool (`patterns/cmd/cli/`)

**Key Talking Points**:
- Go language features and best practices
- System design and architecture
- Kubernetes operator development
- Production readiness considerations

## 🛠️ Hands-On Exercises

### **Exercise 1: Add a New Design Pattern**
**Task**: Implement the Factory pattern in `patterns/pkg/patterns/`
**Skills**: Interface design, pattern implementation
**Time**: 15-20 minutes

### **Exercise 2: Extend the Controller**
**Task**: Add a new field to the CRD and handle it in the controller
**Skills**: CRD development, controller logic
**Time**: 20-25 minutes

### **Exercise 3: Add Concurrency**
**Task**: Implement a concurrent processing feature
**Skills**: Goroutines, channels, synchronization
**Time**: 15-20 minutes

### **Exercise 4: Write Tests**
**Task**: Add comprehensive tests for a utility function
**Skills**: Testing, mocking, table-driven tests
**Time**: 10-15 minutes

## 📋 Common Interview Questions

### **Go Fundamentals**
1. "Explain the difference between `make` and `new` in Go."
2. "How do you handle panics in Go applications?"
3. "What's the difference between a slice and an array?"
4. "Explain Go's garbage collector and how it affects performance."

### **Concurrency**
1. "How do you prevent race conditions in Go?"
2. "Explain the difference between `sync.Mutex` and `sync.RWMutex`."
3. "When would you use `select` vs `switch` in Go?"
4. "How do you implement a timeout for a goroutine?"

### **Kubernetes/Operators**
1. "What is the operator pattern and why is it useful?"
2. "How do you handle resource conflicts in a controller?"
3. "Explain the difference between `CreateOrUpdate` and `Create` in controller-runtime."
4. "How do you implement proper error handling in a Kubernetes controller?"

### **System Design**
1. "How would you design a scalable web service in Go?"
2. "Explain how you would implement circuit breakers in Go."
3. "How do you handle configuration management in Go applications?"
4. "What strategies would you use for graceful shutdown?"

## 🎯 Quick Demo Scripts

### **5-Minute Go Concepts Demo**
```bash
# Show interfaces
cd patterns/pkg/interfaces && go doc Logger

# Show patterns
cd ../patterns && go run observer.go

# Show concurrency
cd ../../examples/concurrency && go run worker_pool.go
```

### **5-Minute Operator Demo**
```bash
# Build and deploy
cd operator/
make build
make install
make deploy

# Create instance and show it working
make create-sample
make port-forward
# Open http://localhost:8080
```

### **10-Minute Full Demo**
```bash
# Run the complete demo
cd operator/
./test-demo.sh

# Show the web server
make port-forward
# Open http://localhost:8080

# Clean up
./cleanup.sh
```

## 🏆 Key Strengths to Highlight

1. **Clean Architecture**: Well-organized code with clear separation of concerns
2. **Interface Design**: Proper use of interfaces for testability and flexibility
3. **Error Handling**: Comprehensive error handling with custom error types
4. **Testing**: Table-driven tests and proper mocking strategies
5. **Concurrency**: Safe concurrent programming with proper synchronization
6. **Production Ready**: Proper logging, metrics, and health checks
7. **Documentation**: Comprehensive documentation and examples
8. **Real-World Application**: Actual working operator with tangible results

This project demonstrates both deep Go knowledge and practical Kubernetes operator development skills, making it perfect for a wide range of Go interview scenarios.
