package concurrency

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Pipeline demonstrates the pipeline pattern in Go
// This is a common concurrency pattern asked about in interviews

// Stage represents a stage in the pipeline
type Stage[T, U any] interface {
	Process(ctx context.Context, input <-chan T) <-chan U
}

// TransformStage transforms data from type T to type U
type TransformStage[T, U any] struct {
	Transform func(T) U
	Name      string
}

// Process processes input data and returns transformed data
func (s *TransformStage[T, U]) Process(ctx context.Context, input <-chan T) <-chan U {
	output := make(chan U)
	
	go func() {
		defer close(output)
		for {
			select {
			case data, ok := <-input:
				if !ok {
					return
				}
				transformed := s.Transform(data)
				select {
				case output <- transformed:
				case <-ctx.Done():
					return
				}
			case <-ctx.Done():
				return
			}
		}
	}()
	
	return output
}

// FilterStage filters data based on a predicate
type FilterStage[T any] struct {
	Predicate func(T) bool
	Name      string
}

// Process processes input data and returns filtered data
func (s *FilterStage[T]) Process(ctx context.Context, input <-chan T) <-chan T {
	output := make(chan T)
	
	go func() {
		defer close(output)
		for {
			select {
			case data, ok := <-input:
				if !ok {
					return
				}
				if s.Predicate(data) {
					select {
					case output <- data:
					case <-ctx.Done():
						return
					}
				}
			case <-ctx.Done():
				return
			}
		}
	}()
	
	return output
}

// Pipeline represents a data processing pipeline
type Pipeline[T any] struct {
	stages []Stage[any, any]
}

// NewPipeline creates a new pipeline
func NewPipeline[T any]() *Pipeline[T] {
	return &Pipeline[T]{
		stages: make([]Stage[any, any], 0),
	}
}

// AddStage adds a stage to the pipeline
func (p *Pipeline[T]) AddStage(stage Stage[any, any]) *Pipeline[T] {
	p.stages = append(p.stages, stage)
	return p
}

// Process processes data through the pipeline
func (p *Pipeline[T]) Process(ctx context.Context, input <-chan T) <-chan any {
	current := make(chan any)
	
	// Convert input to any type
	go func() {
		defer close(current)
		for {
			select {
			case data, ok := <-input:
				if !ok {
					return
				}
				select {
				case current <- any(data):
				case <-ctx.Done():
					return
				}
			case <-ctx.Done():
				return
			}
		}
	}()
	
	// Process through each stage
	for _, stage := range p.stages {
		current = stage.Process(ctx, current)
	}
	
	return current
}

// FanOutStage distributes data to multiple channels
type FanOutStage[T any] struct {
	NumWorkers int
	Name       string
}

// Process processes input data and distributes it to multiple workers
func (s *FanOutStage[T]) Process(ctx context.Context, input <-chan T) <-chan T {
	output := make(chan T)
	
	go func() {
		defer close(output)
		
		var wg sync.WaitGroup
		
		for i := 0; i < s.NumWorkers; i++ {
			wg.Add(1)
			go func(workerID int) {
				defer wg.Done()
				for {
					select {
					case data, ok := <-input:
						if !ok {
							return
						}
						select {
						case output <- data:
						case <-ctx.Done():
							return
						}
					case <-ctx.Done():
						return
					}
				}
			}(i)
		}
		
		wg.Wait()
	}()
	
	return output
}

// Example usage function
func ExamplePipeline() {
	// Create input data
	input := make(chan int)
	go func() {
		defer close(input)
		for i := 1; i <= 10; i++ {
			input <- i
		}
	}()
	
	// Create pipeline
	pipeline := NewPipeline[int]()
	
	// Add stages
	pipeline.AddStage(&TransformStage[any, any]{
		Transform: func(x any) any {
			return x.(int) * 2 // Double the number
		},
		Name: "double",
	})
	
	pipeline.AddStage(&FilterStage[any]{
		Predicate: func(x any) bool {
			return x.(int) > 5 // Filter numbers greater than 5
		},
		Name: "filter",
	})
	
	pipeline.AddStage(&TransformStage[any, any]{
		Transform: func(x any) any {
			return fmt.Sprintf("processed: %d", x.(int)) // Convert to string
		},
		Name: "stringify",
	})
	
	// Process data
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	output := pipeline.Process(ctx, input)
	
	// Collect results
	for result := range output {
		fmt.Printf("Result: %v\n", result)
	}
}
