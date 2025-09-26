package interfaces

import (
	"context"
	"time"
)

// Reconciler defines the interface for reconciling resources
// This is a common pattern in Kubernetes operators and Go applications
type Reconciler interface {
	Reconcile(ctx context.Context, req Request) (Result, error)
}

// Request represents a reconciliation request
type Request struct {
	Namespace string
	Name      string
}

// Result represents the result of a reconciliation
type Result struct {
	Requeue      bool
	RequeueAfter time.Duration
}

// Manager defines the interface for managing reconcilers
type Manager interface {
	AddReconciler(name string, reconciler Reconciler) error
	Start(ctx context.Context) error
	Stop() error
}

// ResourceManager defines the interface for managing Kubernetes resources
type ResourceManager interface {
	Get(ctx context.Context, key string) (interface{}, error)
	Create(ctx context.Context, obj interface{}) error
	Update(ctx context.Context, obj interface{}) error
	Delete(ctx context.Context, key string) error
	List(ctx context.Context, selector string) ([]interface{}, error)
}
