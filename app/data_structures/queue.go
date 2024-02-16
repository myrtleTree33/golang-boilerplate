package datastructures

type Queue[T any] struct {
	elements []T
}

// NewQueue creates a new queue
func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}

// Enqueue adds an element to the queue
func (q *Queue[T]) Enqueue(element T) *Queue[T] {
	q.elements = append(q.elements, element)
	return q
}

// Dequeue removes an element from the queue
func (q *Queue[T]) Dequeue() *T {
	if len(q.elements) == 0 {
		return nil
	}
	element := q.elements[0]
	q.elements = q.elements[1:]
	return &element
}

// Peek returns the front element of the queue
func (q *Queue[T]) Peek() *T {
	if len(q.elements) == 0 {
		return nil
	}
	element := q.elements[0]
	return &element
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.elements) == 0
}
