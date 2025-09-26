package tests

import (
	"testing"

	datastructures "github.com/kubermatic/go-interviews/patterns/examples/data-structures"
)

func TestStack(t *testing.T) {
	stack := datastructures.NewStack()

	// Test initial state
	if !stack.IsEmpty() {
		t.Error("New stack should be empty")
	}

	if stack.Size() != 0 {
		t.Error("New stack should have size 0")
	}

	// Test push
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	if stack.IsEmpty() {
		t.Error("Stack should not be empty after pushing elements")
	}

	if stack.Size() != 3 {
		t.Errorf("Expected size 3, got %d", stack.Size())
	}

	// Test peek
	if stack.Peek() != 3 {
		t.Errorf("Expected peek to return 3, got %v", stack.Peek())
	}

	// Test pop
	popped := stack.Pop()
	if popped != 3 {
		t.Errorf("Expected pop to return 3, got %v", popped)
	}

	if stack.Size() != 2 {
		t.Errorf("Expected size 2 after pop, got %d", stack.Size())
	}

	// Test multiple pops
	stack.Pop()
	stack.Pop()

	if !stack.IsEmpty() {
		t.Error("Stack should be empty after popping all elements")
	}

	// Test pop on empty stack
	popped = stack.Pop()
	if popped != nil {
		t.Errorf("Expected pop on empty stack to return nil, got %v", popped)
	}
}
