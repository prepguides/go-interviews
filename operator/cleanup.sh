#!/bin/bash

# Cleanup script for the Webserver Operator Demo
set -e

echo "🧹 Cleaning up Webserver Operator Demo"
echo "====================================="

# Delete sample instance
echo "🗑️  Deleting sample instance..."
make delete-sample 2>/dev/null || echo "Sample instance not found or already deleted"

# Undeploy operator
echo "🗑️  Undeploying operator..."
make undeploy 2>/dev/null || echo "Operator not deployed or already undeployed"

# Uninstall CRDs
echo "🗑️  Uninstalling CRDs..."
make uninstall 2>/dev/null || echo "CRDs not installed or already uninstalled"

echo "✅ Cleanup complete!"
