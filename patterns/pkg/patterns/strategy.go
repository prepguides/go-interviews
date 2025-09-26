package patterns

import (
	"context"
	"fmt"
)

// Strategy pattern implementation - another common Go interview topic
// This demonstrates interface-based design and polymorphism

// ProcessingStrategy defines the interface for different processing strategies
type ProcessingStrategy interface {
	Process(ctx context.Context, data interface{}) (interface{}, error)
	GetName() string
}

// DataProcessor uses the strategy pattern to process data
type DataProcessor struct {
	strategy ProcessingStrategy
}

// NewDataProcessor creates a new data processor with a strategy
func NewDataProcessor(strategy ProcessingStrategy) *DataProcessor {
	return &DataProcessor{
		strategy: strategy,
	}
}

// SetStrategy allows changing the processing strategy at runtime
func (dp *DataProcessor) SetStrategy(strategy ProcessingStrategy) {
	dp.strategy = strategy
}

// Process processes data using the current strategy
func (dp *DataProcessor) Process(ctx context.Context, data interface{}) (interface{}, error) {
	if dp.strategy == nil {
		return nil, fmt.Errorf("no strategy set")
	}
	return dp.strategy.Process(ctx, data)
}

// GetCurrentStrategy returns the current strategy name
func (dp *DataProcessor) GetCurrentStrategy() string {
	if dp.strategy == nil {
		return "none"
	}
	return dp.strategy.GetName()
}

// Concrete strategies

// JSONProcessingStrategy processes data as JSON
type JSONProcessingStrategy struct{}

func (j *JSONProcessingStrategy) Process(ctx context.Context, data interface{}) (interface{}, error) {
	// Simulate JSON processing
	return fmt.Sprintf("JSON processed: %v", data), nil
}

func (j *JSONProcessingStrategy) GetName() string {
	return "JSON"
}

// XMLProcessingStrategy processes data as XML
type XMLProcessingStrategy struct{}

func (x *XMLProcessingStrategy) Process(ctx context.Context, data interface{}) (interface{}, error) {
	// Simulate XML processing
	return fmt.Sprintf("XML processed: %v", data), nil
}

func (x *XMLProcessingStrategy) GetName() string {
	return "XML"
}

// BinaryProcessingStrategy processes data as binary
type BinaryProcessingStrategy struct{}

func (b *BinaryProcessingStrategy) Process(ctx context.Context, data interface{}) (interface{}, error) {
	// Simulate binary processing
	return fmt.Sprintf("Binary processed: %v", data), nil
}

func (b *BinaryProcessingStrategy) GetName() string {
	return "Binary"
}
