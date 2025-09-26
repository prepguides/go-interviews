package patterns

import (
	"fmt"
	"sync"
)

// Observer interface defines the update method
type Observer interface {
	Update(data interface{})
}

// Subject interface defines methods for managing observers
type Subject interface {
	Attach(observer Observer)
	Detach(observer Observer)
	Notify(data interface{})
}

// ConcreteSubject implements the Subject interface
type ConcreteSubject struct {
	observers []Observer
	mutex     sync.RWMutex
}

// NewConcreteSubject creates a new concrete subject
func NewConcreteSubject() *ConcreteSubject {
	return &ConcreteSubject{
		observers: make([]Observer, 0),
	}
}

// Attach adds an observer to the subject
func (s *ConcreteSubject) Attach(observer Observer) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.observers = append(s.observers, observer)
}

// Detach removes an observer from the subject
func (s *ConcreteSubject) Detach(observer Observer) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	for i, obs := range s.observers {
		if obs == observer {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			break
		}
	}
}

// Notify notifies all observers
func (s *ConcreteSubject) Notify(data interface{}) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	for _, observer := range s.observers {
		observer.Update(data)
	}
}

// ConcreteObserver implements the Observer interface
type ConcreteObserver struct {
	name string
}

// NewConcreteObserver creates a new concrete observer
func NewConcreteObserver(name string) *ConcreteObserver {
	return &ConcreteObserver{
		name: name,
	}
}

// Update handles the update notification
func (o *ConcreteObserver) Update(data interface{}) {
	fmt.Printf("Observer %s received update: %v\n", o.name, data)
}

// Example usage
func ExampleObserver() {
	subject := NewConcreteSubject()

	observer1 := NewConcreteObserver("Observer1")
	observer2 := NewConcreteObserver("Observer2")

	subject.Attach(observer1)
	subject.Attach(observer2)

	subject.Notify("Hello, Observers!")

	subject.Detach(observer1)
	subject.Notify("Observer1 has been removed")
}
