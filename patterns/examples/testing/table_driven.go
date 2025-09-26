package testing

import (
	"context"
	"fmt"
	"testing"
	"time"
)

// TableDrivenTests demonstrates table-driven testing in Go
// This is a common testing pattern asked about in Go interviews

// Calculator demonstrates a simple service for testing
type Calculator struct {
	Logger Logger
}

func NewCalculator(logger Logger) *Calculator {
	return &Calculator{
		Logger: logger,
	}
}

func (c *Calculator) Add(a, b int) int {
	result := a + b
	c.Logger.Info("Addition performed", "a", a, "b", b, "result", result)
	return result
}

func (c *Calculator) Subtract(a, b int) int {
	result := a - b
	c.Logger.Info("Subtraction performed", "a", a, "b", b, "result", result)
	return result
}

func (c *Calculator) Multiply(a, b int) int {
	result := a * b
	c.Logger.Info("Multiplication performed", "a", a, "b", b, "result", result)
	return result
}

func (c *Calculator) Divide(a, b int) (int, error) {
	if b == 0 {
		c.Logger.Error(nil, "Division by zero attempted", "a", a, "b", b)
		return 0, fmt.Errorf("division by zero")
	}
	result := a / b
	c.Logger.Info("Division performed", "a", a, "b", b, "result", result)
	return result, nil
}

// TestCalculatorAdd demonstrates table-driven testing for the Add method
func TestCalculatorAdd(t *testing.T) {
	tests := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{
			name:     "positive numbers",
			a:        2,
			b:        3,
			expected: 5,
		},
		{
			name:     "negative numbers",
			a:        -2,
			b:        -3,
			expected: -5,
		},
		{
			name:     "mixed signs",
			a:        5,
			b:        -3,
			expected: 2,
		},
		{
			name:     "zero values",
			a:        0,
			b:        5,
			expected: 5,
		},
		{
			name:     "large numbers",
			a:        1000000,
			b:        2000000,
			expected: 3000000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockLogger := NewMockLogger()
			calc := NewCalculator(mockLogger)
			
			result := calc.Add(tt.a, tt.b)
			
			if result != tt.expected {
				t.Errorf("Add(%d, %d) = %d, expected %d", tt.a, tt.b, result, tt.expected)
			}
			
			// Verify logging
			if !mockLogger.AssertLogContains("info", "Addition performed") {
				t.Error("Expected addition to be logged")
			}
		})
	}
}

// TestCalculatorDivide demonstrates table-driven testing with error cases
func TestCalculatorDivide(t *testing.T) {
	tests := []struct {
		name        string
		a           int
		b           int
		expected    int
		expectError bool
		errorMsg    string
	}{
		{
			name:        "normal division",
			a:           10,
			b:           2,
			expected:    5,
			expectError: false,
		},
		{
			name:        "division by zero",
			a:           10,
			b:           0,
			expected:    0,
			expectError: true,
			errorMsg:    "division by zero",
		},
		{
			name:        "negative result",
			a:           -10,
			b:           2,
			expected:    -5,
			expectError: false,
		},
		{
			name:        "fractional result (integer division)",
			a:           7,
			b:           3,
			expected:    2,
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockLogger := NewMockLogger()
			calc := NewCalculator(mockLogger)
			
			result, err := calc.Divide(tt.a, tt.b)
			
			if tt.expectError {
				if err == nil {
					t.Errorf("Divide(%d, %d) expected error, got nil", tt.a, tt.b)
				}
				if err != nil && err.Error() != tt.errorMsg {
					t.Errorf("Divide(%d, %d) expected error '%s', got '%s'", tt.a, tt.b, tt.errorMsg, err.Error())
				}
				// Verify error logging
				if !mockLogger.AssertLogContains("error", "Division by zero attempted") {
					t.Error("Expected division by zero to be logged as error")
				}
			} else {
				if err != nil {
					t.Errorf("Divide(%d, %d) unexpected error: %v", tt.a, tt.b, err)
				}
				if result != tt.expected {
					t.Errorf("Divide(%d, %d) = %d, expected %d", tt.a, tt.b, result, tt.expected)
				}
				// Verify success logging
				if !mockLogger.AssertLogContains("info", "Division performed") {
					t.Error("Expected division to be logged")
				}
			}
		})
	}
}

// TestCalculatorConcurrent demonstrates testing concurrent operations
func TestCalculatorConcurrent(t *testing.T) {
	mockLogger := NewMockLogger()
	calc := NewCalculator(mockLogger)
	
	// Test concurrent additions
	results := make(chan int, 10)
	
	for i := 0; i < 10; i++ {
		go func(i int) {
			result := calc.Add(i, i+1)
			results <- result
		}(i)
	}
	
	// Collect results
	collectedResults := make([]int, 0, 10)
	for i := 0; i < 10; i++ {
		select {
		case result := <-results:
			collectedResults = append(collectedResults, result)
		case <-time.After(5 * time.Second):
			t.Fatal("Timeout waiting for concurrent results")
		}
	}
	
	// Verify we got 10 results
	if len(collectedResults) != 10 {
		t.Errorf("Expected 10 results, got %d", len(collectedResults))
	}
	
	// Verify all results are positive (since we're adding i + (i+1))
	for i, result := range collectedResults {
		if result <= 0 {
			t.Errorf("Result %d should be positive, got %d", i, result)
		}
	}
}

// BenchmarkCalculatorAdd demonstrates benchmarking in Go
func BenchmarkCalculatorAdd(b *testing.B) {
	mockLogger := NewMockLogger()
	calc := NewCalculator(mockLogger)
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		calc.Add(i, i+1)
	}
}

// BenchmarkCalculatorAddParallel demonstrates parallel benchmarking
func BenchmarkCalculatorAddParallel(b *testing.B) {
	mockLogger := NewMockLogger()
	calc := NewCalculator(mockLogger)
	
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			calc.Add(i, i+1)
			i++
		}
	})
}
