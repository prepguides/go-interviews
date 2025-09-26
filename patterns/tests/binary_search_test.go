package tests

import (
	"testing"

	"github.com/kubermatic/go-interviews/patterns/examples/algorithms"
)

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		target   int
		expected int
	}{
		{
			name:     "Element found at beginning",
			arr:      []int{1, 2, 3, 4, 5},
			target:   1,
			expected: 0,
		},
		{
			name:     "Element found at end",
			arr:      []int{1, 2, 3, 4, 5},
			target:   5,
			expected: 4,
		},
		{
			name:     "Element found in middle",
			arr:      []int{1, 2, 3, 4, 5},
			target:   3,
			expected: 2,
		},
		{
			name:     "Element not found",
			arr:      []int{1, 2, 3, 4, 5},
			target:   6,
			expected: -1,
		},
		{
			name:     "Empty array",
			arr:      []int{},
			target:   1,
			expected: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := algorithms.BinarySearch(tt.arr, tt.target)
			if result != tt.expected {
				t.Errorf("BinarySearch() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestBinarySearchRecursive(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		target   int
		expected int
	}{
		{
			name:     "Element found",
			arr:      []int{1, 2, 3, 4, 5},
			target:   3,
			expected: 2,
		},
		{
			name:     "Element not found",
			arr:      []int{1, 2, 3, 4, 5},
			target:   6,
			expected: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := algorithms.BinarySearchRecursive(tt.arr, tt.target, 0, len(tt.arr)-1)
			if result != tt.expected {
				t.Errorf("BinarySearchRecursive() = %v, want %v", result, tt.expected)
			}
		})
	}
}
