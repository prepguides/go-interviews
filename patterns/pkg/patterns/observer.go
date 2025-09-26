package patterns

import (
	"context"
	"sync"
)

// Observer pattern implementation - commonly asked in Go interviews
// This demonstrates interfaces, goroutines, and synchronization

// Event represents an event that can be observed
type Event struct {
	Type      string
	Data      interface{}
	Timestamp int64
}

// Observer defines the interface for event observers
type Observer interface {
	Notify(ctx context.Context, event Event) error
	GetID() string
}

// Subject defines the interface for event subjects
type Subject interface {
	Subscribe(observer Observer) error
	Unsubscribe(observerID string) error
	NotifyObservers(ctx context.Context, event Event) error
}

// EventBus implements the Observer pattern
type EventBus struct {
	observers map[string]Observer
	mutex     sync.RWMutex
}

// NewEventBus creates a new event bus
func NewEventBus() *EventBus {
	return &EventBus{
		observers: make(map[string]Observer),
	}
}

// Subscribe adds an observer to the event bus
func (eb *EventBus) Subscribe(observer Observer) error {
	eb.mutex.Lock()
	defer eb.mutex.Unlock()

	eb.observers[observer.GetID()] = observer
	return nil
}

// Unsubscribe removes an observer from the event bus
func (eb *EventBus) Unsubscribe(observerID string) error {
	eb.mutex.Lock()
	defer eb.mutex.Unlock()

	delete(eb.observers, observerID)
	return nil
}

// NotifyObservers notifies all observers of an event
func (eb *EventBus) NotifyObservers(ctx context.Context, event Event) error {
	eb.mutex.RLock()
	observers := make([]Observer, 0, len(eb.observers))
	for _, observer := range eb.observers {
		observers = append(observers, observer)
	}
	eb.mutex.RUnlock()

	// Notify all observers concurrently
	var wg sync.WaitGroup
	for _, observer := range observers {
		wg.Add(1)
		go func(obs Observer) {
			defer wg.Done()
			obs.Notify(ctx, event)
		}(observer)
	}

	wg.Wait()
	return nil
}

// GetObserverCount returns the number of observers
func (eb *EventBus) GetObserverCount() int {
	eb.mutex.RLock()
	defer eb.mutex.RUnlock()
	return len(eb.observers)
}
