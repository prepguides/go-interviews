# Go Interviews - Comprehensive Project Collection

This repository contains a comprehensive collection of Go projects designed to demonstrate various Go concepts and Kubernetes operator development skills for technical interviews.

## 📁 Project Structure

```
go-interviews/
├── operator/              # Kubernetes Operator Project
│   ├── api/              # CRD definitions
│   ├── controllers/      # Operator controller logic
│   ├── config/           # Kubernetes manifests
│   ├── main.go           # Operator entry point
│   ├── Makefile          # Build and deployment commands
│   └── README.md         # Operator-specific documentation
├── patterns/             # Go Patterns and Examples
│   ├── pkg/              # Reusable Go packages
│   ├── examples/         # Go concept examples
│   ├── cmd/              # Command-line applications
│   └── README.md         # Patterns-specific documentation
├── INTERVIEW_GUIDE.md    # Comprehensive interview guide
├── DEMO_SUMMARY.md       # Project summary and features
└── README.md             # This file
```

## 🎯 What This Repository Contains

### **1. Operator Project** (`operator/`)
A complete Kubernetes operator that demonstrates:
- **Custom Resource Definitions (CRDs)**: API design and validation
- **Controller Logic**: Reconciliation loops and resource management
- **RBAC**: Proper security and permissions
- **Status Management**: Conditions and health monitoring
- **End-to-End Demo**: Working operator with accessible web interface

**Perfect for**: Kubernetes/DevOps engineer interviews, operator development roles

### **2. Go Patterns Project** (`patterns/`)
A comprehensive collection of Go concepts including:
- **Interfaces & Polymorphism**: Clean interface design with mock implementations
- **Design Patterns**: Observer, Strategy, Builder patterns
- **Concurrency**: Worker pools, pipelines, channel communication
- **Error Handling**: Custom error types, retry mechanisms, validation
- **Testing**: Table-driven tests, mocking, benchmarking
- **CLI Development**: Command-line argument parsing and subcommands

**Perfect for**: Generic Go developer interviews, backend development roles

## 🚀 Quick Start

### **For Operator Development Interviews**
```bash
cd operator/
./test-demo.sh
# Access the web server at http://localhost:8080
```

### **For Go Patterns Interviews**
```bash
cd patterns/
go run examples/concurrency/worker_pool.go
go run cmd/cli/main.go validate -input "hello world"
```

### **For Full-Stack Go Interviews**
```bash
# Start with patterns
cd patterns/
go run examples/concurrency/pipeline.go

# Then show operator
cd ../operator/
make build && make install && make deploy
```

## 🎤 Interview Scenarios

### **Scenario 1: Generic Go Developer**
- **Focus**: Core Go concepts, patterns, and best practices
- **Location**: `patterns/` directory
- **Key Topics**: Interfaces, concurrency, testing, error handling

### **Scenario 2: Kubernetes/DevOps Engineer**
- **Focus**: Operator development and Kubernetes integration
- **Location**: `operator/` directory
- **Key Topics**: CRDs, controllers, RBAC, reconciliation

### **Scenario 3: Full-Stack Go Developer**
- **Focus**: Both Go fundamentals and Kubernetes knowledge
- **Location**: Both directories
- **Key Topics**: Complete Go ecosystem knowledge

## 📚 Documentation

- **[INTERVIEW_GUIDE.md](INTERVIEW_GUIDE.md)**: Comprehensive interview guide with scenarios, questions, and exercises
- **[DEMO_SUMMARY.md](DEMO_SUMMARY.md)**: Project summary and key features
- **[operator/README.md](operator/README.md)**: Operator-specific documentation
- **[patterns/README.md](patterns/README.md)**: Go patterns documentation

## 🛠️ Prerequisites

- **Go 1.19+**: For all Go development
- **Kubernetes cluster**: For operator demonstrations
- **kubectl**: For Kubernetes interactions
- **Docker**: For container builds (optional)

## 🎯 Key Strengths

1. **Comprehensive Coverage**: Both generic Go concepts and Kubernetes-specific knowledge
2. **Real-World Examples**: Working code that demonstrates actual use cases
3. **Interview-Ready**: Designed specifically for technical interviews
4. **Well-Documented**: Extensive documentation and examples
5. **Modular Design**: Can focus on specific areas based on interview type
6. **Production-Ready**: Follows best practices and industry standards

## 🏆 Perfect For

- **Go Developer Interviews**: Backend, full-stack, or systems programming roles
- **Kubernetes Engineer Interviews**: DevOps, platform engineering, or SRE roles
- **Technical Lead Interviews**: Architecture and system design discussions
- **Code Review Sessions**: Demonstrating code quality and best practices
- **Learning and Practice**: Self-study and skill development

This repository provides everything needed to demonstrate strong Go skills and Kubernetes knowledge in technical interviews, with the flexibility to focus on specific areas based on the role and interview requirements.