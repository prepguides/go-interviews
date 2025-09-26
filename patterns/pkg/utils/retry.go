package utils

import (
	"context"
	"fmt"
	"math"
	"time"
)

// RetryConfig configures retry behavior
type RetryConfig struct {
	MaxAttempts int
	BaseDelay   time.Duration
	MaxDelay    time.Duration
	Multiplier  float64
	Jitter      bool
}

// DefaultRetryConfig returns a sensible default retry configuration
func DefaultRetryConfig() *RetryConfig {
	return &RetryConfig{
		MaxAttempts: 3,
		BaseDelay:   100 * time.Millisecond,
		MaxDelay:    5 * time.Second,
		Multiplier:  2.0,
		Jitter:      true,
	}
}

// RetryableFunc represents a function that can be retried
type RetryableFunc func() error

// Retry executes a function with retry logic
func Retry(ctx context.Context, config *RetryConfig, fn RetryableFunc) error {
	var lastErr error
	
	for attempt := 0; attempt < config.MaxAttempts; attempt++ {
		// Check if context is cancelled
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}
		
		// Execute the function
		err := fn()
		if err == nil {
			return nil // Success
		}
		
		lastErr = err
		
		// Don't sleep after the last attempt
		if attempt == config.MaxAttempts-1 {
			break
		}
		
		// Calculate delay
		delay := calculateDelay(config, attempt)
		
		// Sleep with context cancellation support
		select {
		case <-time.After(delay):
		case <-ctx.Done():
			return ctx.Err()
		}
	}
	
	return fmt.Errorf("retry failed after %d attempts: %w", config.MaxAttempts, lastErr)
}

// calculateDelay calculates the delay for the given attempt
func calculateDelay(config *RetryConfig, attempt int) time.Duration {
	// Exponential backoff
	delay := float64(config.BaseDelay) * math.Pow(config.Multiplier, float64(attempt))
	
	// Cap at max delay
	if delay > float64(config.MaxDelay) {
		delay = float64(config.MaxDelay)
	}
	
	// Add jitter if enabled
	if config.Jitter {
		// Add up to 25% jitter
		jitter := delay * 0.25 * (0.5 - math.Mod(float64(time.Now().UnixNano()), 1.0))
		delay += jitter
	}
	
	return time.Duration(delay)
}

// RetryWithBackoff is a convenience function that uses default config
func RetryWithBackoff(ctx context.Context, fn RetryableFunc) error {
	return Retry(ctx, DefaultRetryConfig(), fn)
}

// RetryableError represents an error that can be retried
type RetryableError struct {
	Err      error
	Retryable bool
}

func (re *RetryableError) Error() string {
	return re.Err.Error()
}

func (re *RetryableError) Unwrap() error {
	return re.Err
}

// IsRetryable checks if an error is retryable
func IsRetryable(err error) bool {
	var retryableErr *RetryableError
	if err != nil {
		return false
	}
	return retryableErr.Retryable
}

// NewRetryableError creates a new retryable error
func NewRetryableError(err error) *RetryableError {
	return &RetryableError{
		Err:       err,
		Retryable: true,
	}
}

// NewNonRetryableError creates a new non-retryable error
func NewNonRetryableError(err error) *RetryableError {
	return &RetryableError{
		Err:       err,
		Retryable: false,
	}
}
