package concurrency

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// WorkerPool demonstrates common Go concurrency patterns
// This is frequently asked about in Go interviews

// Job represents a unit of work
type Job struct {
	ID     int
	Data   interface{}
	Result chan interface{}
	Error  chan error
}

// Worker represents a worker in the pool
type Worker struct {
	ID       int
	JobQueue chan Job
	Quit     chan bool
	Wg       *sync.WaitGroup
}

// NewWorker creates a new worker
func NewWorker(id int, jobQueue chan Job, wg *sync.WaitGroup) *Worker {
	return &Worker{
		ID:       id,
		JobQueue: jobQueue,
		Quit:     make(chan bool),
		Wg:       wg,
	}
}

// Start starts the worker
func (w *Worker) Start() {
	w.Wg.Add(1)
	go func() {
		defer w.Wg.Done()
		for {
			select {
			case job := <-w.JobQueue:
				// Process the job
				result, err := w.processJob(job)
				if err != nil {
					job.Error <- err
				} else {
					job.Result <- result
				}
			case <-w.Quit:
				fmt.Printf("Worker %d stopping\n", w.ID)
				return
			}
		}
	}()
}

// Stop stops the worker
func (w *Worker) Stop() {
	w.Quit <- true
}

// processJob simulates job processing
func (w *Worker) processJob(job Job) (interface{}, error) {
	// Simulate some work
	time.Sleep(100 * time.Millisecond)
	return fmt.Sprintf("Worker %d processed job %d with data: %v", w.ID, job.ID, job.Data), nil
}

// WorkerPool manages a pool of workers
type WorkerPool struct {
	Workers    []*Worker
	JobQueue   chan Job
	NumWorkers int
	Wg         sync.WaitGroup
}

// NewWorkerPool creates a new worker pool
func NewWorkerPool(numWorkers int, jobQueueSize int) *WorkerPool {
	jobQueue := make(chan Job, jobQueueSize)
	workers := make([]*Worker, numWorkers)

	pool := &WorkerPool{
		JobQueue:   jobQueue,
		NumWorkers: numWorkers,
	}

	// Create workers
	for i := 0; i < numWorkers; i++ {
		workers[i] = NewWorker(i, jobQueue, &pool.Wg)
	}
	pool.Workers = workers

	return pool
}

// Start starts all workers in the pool
func (wp *WorkerPool) Start() {
	for _, worker := range wp.Workers {
		worker.Start()
	}
}

// Stop stops all workers in the pool
func (wp *WorkerPool) Stop() {
	for _, worker := range wp.Workers {
		worker.Stop()
	}
	wp.Wg.Wait()
}

// SubmitJob submits a job to the pool
func (wp *WorkerPool) SubmitJob(job Job) {
	wp.JobQueue <- job
}

// ProcessJobs processes a batch of jobs concurrently
func (wp *WorkerPool) ProcessJobs(ctx context.Context, jobs []Job) ([]interface{}, []error) {
	results := make([]interface{}, len(jobs))
	errors := make([]error, len(jobs))

	// Submit all jobs
	for i, job := range jobs {
		job.Result = make(chan interface{}, 1)
		job.Error = make(chan error, 1)
		wp.SubmitJob(job)

		// Wait for result
		select {
		case result := <-job.Result:
			results[i] = result
		case err := <-job.Error:
			errors[i] = err
		case <-ctx.Done():
			errors[i] = ctx.Err()
		}
	}

	return results, errors
}

// Example usage function
func ExampleWorkerPool() {
	// Create a worker pool with 3 workers
	pool := NewWorkerPool(3, 10)
	pool.Start()
	defer pool.Stop()

	// Create some jobs
	jobs := make([]Job, 5)
	for i := 0; i < 5; i++ {
		jobs[i] = Job{
			ID:   i,
			Data: fmt.Sprintf("data-%d", i),
		}
	}

	// Process jobs
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	results, errors := pool.ProcessJobs(ctx, jobs)

	// Print results
	for i, result := range results {
		if errors[i] != nil {
			fmt.Printf("Job %d failed: %v\n", i, errors[i])
		} else {
			fmt.Printf("Job %d result: %v\n", i, result)
		}
	}
}
