package tests

import (
	"testing"

	"github.com/kubermatic/go-interviews/patterns/examples/patterns"
)

func TestSingleton(t *testing.T) {
	// Get first instance
	instance1 := patterns.GetInstance()
	if instance1 == nil {
		t.Error("GetInstance() should not return nil")
	}

	// Get second instance
	instance2 := patterns.GetInstance()
	if instance2 == nil {
		t.Error("GetInstance() should not return nil")
	}

	// Check if both instances are the same
	if instance1 != instance2 {
		t.Error("GetInstance() should return the same instance")
	}

	// Test data manipulation
	instance1.SetData("test data")
	if instance2.GetData() != "test data" {
		t.Error("Both instances should share the same data")
	}

	// Test initial data
	instance3 := patterns.GetInstance()
	if instance3.GetData() != "test data" {
		t.Error("New instance should have the same data as previous instances")
	}
}
