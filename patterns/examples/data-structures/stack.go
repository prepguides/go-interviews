package datastructures

// Stack represents a LIFO data structure
type Stack struct {
	items []interface{}
}

// NewStack creates a new stack
func NewStack() *Stack {
	return &Stack{
		items: make([]interface{}, 0),
	}
}

// Push adds an item to the top of the stack
func (s *Stack) Push(item interface{}) {
	s.items = append(s.items, item)
}

// Pop removes and returns the top item from the stack
func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}

	index := len(s.items) - 1
	item := s.items[index]
	s.items = s.items[:index]
	return item
}

// Peek returns the top item without removing it
func (s *Stack) Peek() interface{} {
	if s.IsEmpty() {
		return nil
	}

	return s.items[len(s.items)-1]
}

// IsEmpty returns true if the stack is empty
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

// Size returns the number of items in the stack
func (s *Stack) Size() int {
	return len(s.items)
}
