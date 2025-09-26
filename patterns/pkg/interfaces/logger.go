package interfaces

import (
	"context"
)

// Logger defines a common logging interface
// This demonstrates interface design and dependency injection
type Logger interface {
	Debug(msg string, keysAndValues ...interface{})
	Info(msg string, keysAndValues ...interface{})
	Warn(msg string, keysAndValues ...interface{})
	Error(err error, msg string, keysAndValues ...interface{})
	WithValues(keysAndValues ...interface{}) Logger
	WithName(name string) Logger
}

// ContextLogger extends Logger with context support
type ContextLogger interface {
	Logger
	FromContext(ctx context.Context) Logger
}

// MetricsCollector defines the interface for collecting metrics
type MetricsCollector interface {
	IncrementCounter(name string, labels map[string]string)
	RecordHistogram(name string, value float64, labels map[string]string)
	RecordGauge(name string, value float64, labels map[string]string)
}

// HealthChecker defines the interface for health checks
type HealthChecker interface {
	CheckHealth(ctx context.Context) error
	IsHealthy() bool
	GetHealthStatus() HealthStatus
}

// HealthStatus represents the health status of a component
type HealthStatus struct {
	Healthy bool
	Message string
	Details map[string]interface{}
}
