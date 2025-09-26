#!/bin/bash

# Test script for the Webserver Operator Demo
set -e

echo "🚀 Starting Webserver Operator Demo Test"
echo "========================================"

# Check if kubectl is available
if ! command -v kubectl &> /dev/null; then
    echo "❌ kubectl is not installed or not in PATH"
    exit 1
fi

# Check if we can connect to a cluster
if ! kubectl cluster-info &> /dev/null; then
    echo "❌ Cannot connect to Kubernetes cluster"
    echo "Please ensure you have a running cluster and kubectl is configured"
    exit 1
fi

echo "✅ Kubernetes cluster connection verified"

# Build the operator
echo "🔨 Building the operator..."
make build
echo "✅ Operator built successfully"

# Install CRDs
echo "📦 Installing CRDs..."
make install
echo "✅ CRDs installed"

# Deploy the operator
echo "🚀 Deploying the operator..."
make deploy
echo "✅ Operator deployed"

# Wait for operator to be ready
echo "⏳ Waiting for operator to be ready..."
kubectl wait --for=condition=available --timeout=300s deployment/controller-manager -n system
echo "✅ Operator is ready"

# Create sample instance
echo "📝 Creating sample Webserver instance..."
make create-sample
echo "✅ Sample instance created"

# Wait for deployment to be ready
echo "⏳ Waiting for web server deployment to be ready..."
kubectl wait --for=condition=available --timeout=300s deployment/webserver-sample-deployment
echo "✅ Web server deployment is ready"

# Check status
echo "📊 Checking status..."
make status

# Get service info
echo "🌐 Service information:"
kubectl get service webserver-sample-service

echo ""
echo "🎉 Demo setup complete!"
echo ""
echo "To access the web server:"
echo "1. Run: make port-forward"
echo "2. Open: http://localhost:8080"
echo ""
echo "Or use: make get-url (if LoadBalancer is available)"
echo ""
echo "To clean up:"
echo "make delete-sample && make undeploy && make uninstall"
