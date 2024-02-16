package datastructures

type Stack[T any] struct {
	elements []T
}

// NewStack creates a new stack
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

// Push adds an element to the stack
func (s *Stack[T]) Push(element T) *Stack[T] {
	s.elements = append(s.elements, element)
	return s
}

// Pop removes an element from the stack
func (s *Stack[T]) Pop() *T {
	if len(s.elements) == 0 {
		return nil
	}
	element := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return &element
}

// Peek returns the top element of the stack
func (s *Stack[T]) Peek() *T {
	if len(s.elements) == 0 {
		return nil
	}
	element := s.elements[len(s.elements)-1]
	return &element
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.elements) == 0
}
