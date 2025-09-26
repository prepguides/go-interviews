package patterns

import (
	"sync"
)

// Singleton represents a singleton instance
type Singleton struct {
	data string
}

var (
	instance *Singleton
	once     sync.Once
)

// GetInstance returns the singleton instance
func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{
			data: "Singleton instance created",
		}
	})
	return instance
}

// GetData returns the data from the singleton
func (s *Singleton) GetData() string {
	return s.data
}

// SetData sets the data in the singleton
func (s *Singleton) SetData(data string) {
	s.data = data
}
