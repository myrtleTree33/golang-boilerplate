package datastructures

type LinkedListNode[T comparable] struct {
	Value T
	Next  *LinkedListNode[T]
	Prev  *LinkedListNode[T]
}

type DoublyLinkedList[T comparable] struct {
	Head *LinkedListNode[T]
	Tail *LinkedListNode[T]
}

func NewDoublyLinkedList[T comparable]() *DoublyLinkedList[T] {
	return &DoublyLinkedList[T]{}
}

func (l *DoublyLinkedList[T]) Append(value T) {
	node := &LinkedListNode[T]{Value: value}
	if l.Head == nil {
		l.Head = node
		l.Tail = node
		return
	}
	l.Tail.Next = node
	node.Prev = l.Tail
	l.Tail = node
}

func (l *DoublyLinkedList[T]) Prepend(value T) {
	node := &LinkedListNode[T]{Value: value}
	if l.Head == nil {
		l.Head = node
		l.Tail = node
		return
	}
	l.Head.Prev = node
	node.Next = l.Head
	l.Head = node
}

func (l *DoublyLinkedList[T]) Delete(value T) {
	if l.Head == nil {
		return
	}
	if l.Head.Value == value {
		l.Head = l.Head.Next
		if l.Head != nil {
			l.Head.Prev = nil
		}
		return
	}
	node := l.Head
	for node.Next != nil {
		if node.Next.Value == value {
			node.Next = node.Next.Next
			if node.Next != nil {
				node.Next.Prev = node
			} else {
				l.Tail = node
			}
			return
		}
		node = node.Next
	}
}

func (l *DoublyLinkedList[T]) Search(value T) bool {
	node := l.Head
	for node != nil {
		if node.Value == value {
			return true
		}
		node = node.Next
	}
	return false
}

func (l *DoublyLinkedList[T]) ToSlice() []T {
	var result []T
	node := l.Head
	for node != nil {
		result = append(result, node.Value)
		node = node.Next
	}
	return result
}

func (l *DoublyLinkedList[T]) Reverse() {
	node := l.Head
	var prev *LinkedListNode[T]
	for node != nil {
		next := node.Next
		node.Next = prev
		node.Prev = next
		prev = node
		node = next
	}
	l.Head, l.Tail = l.Tail, l.Head
}

func (l *DoublyLinkedList[T]) IsEmpty() bool {
	return l.Head == nil
}

func (l *DoublyLinkedList[T]) Size() int {
	count := 0
	node := l.Head
	for node != nil {
		count++
		node = node.Next
	}
	return count
}

func (l *DoublyLinkedList[T]) Clear() {
	l.Head = nil
	l.Tail = nil
}

func (l *DoublyLinkedList[T]) GetHead() *LinkedListNode[T] {
	return l.Head
}

func (l *DoublyLinkedList[T]) GetTail() *LinkedListNode[T] {
	return l.Tail
}

func (l *DoublyLinkedList[T]) GetNodeAt(index int) *LinkedListNode[T] {
	if index < 0 {
		return nil
	}
	node := l.Head
	for i := 0; i < index; i++ {
		if node == nil {
			return nil
		}
		node = node.Next
	}
	return node
}

func (l *DoublyLinkedList[T]) InsertAt(index int, value T) {
	if index < 0 {
		return
	}
	node := l.Head
	for i := 0; i < index; i++ {
		if node == nil {
			return
		}
		node = node.Next
	}
	if node == nil {
		return
	}
	newNode := &LinkedListNode[T]{Value: value, Next: node, Prev: node.Prev}
	if node.Prev != nil {
		node.Prev.Next = newNode
	}
	node.Prev = newNode
	if l.Head == node {
		l.Head = newNode
	}
}

func (l *DoublyLinkedList[T]) RemoveAt(index int) {
	if index < 0 {
		return
	}
	node := l.Head
	for i := 0; i < index; i++ {
		if node == nil {
			return
		}
		node = node.Next
	}
	if node == nil {
		return
	}
	if node.Prev != nil {
		node.Prev.Next = node.Next
	}
	if node.Next != nil {
		node.Next.Prev = node.Prev
	}
	if l.Head == node {
		l.Head = node.Next
	}
	if l.Tail == node {
		l.Tail = node.Prev
	}
}

func (l *DoublyLinkedList[T]) InsertAfter(node *LinkedListNode[T], value T) {
	if node == nil {
		return
	}
	newNode := &LinkedListNode[T]{Value: value, Next: node.Next, Prev: node}
	if node.Next != nil {
		node.Next.Prev = newNode
	}
	node.Next = newNode
	if l.Tail == node {
		l.Tail = newNode
	}
}

func (l *DoublyLinkedList[T]) InsertBefore(node *LinkedListNode[T], value T) {
	if node == nil {
		return
	}
	newNode := &LinkedListNode[T]{Value: value, Next: node, Prev: node.Prev}
	if node.Prev != nil {
		node.Prev.Next = newNode
	}
	node.Prev = newNode
	if l.Head == node {
		l.Head = newNode
	}
}

func (l *DoublyLinkedList[T]) RemoveNode(node *LinkedListNode[T]) {
	if node == nil {
		return
	}
	if node.Prev != nil {
		node.Prev.Next = node.Next
	}
	if node.Next != nil {
		node.Next.Prev = node.Prev
	}
	if l.Head == node {
		l.Head = node.Next
	}
	if l.Tail == node {
		l.Tail = node.Prev
	}
}

func (l *DoublyLinkedList[T]) FindNode(value T) *LinkedListNode[T] {
	node := l.Head
	for node != nil {
		if node.Value == value {
			return node
		}
		node = node.Next
	}
	return nil
}

type DoublyLinkedListIterator[T comparable] struct {
	current *LinkedListNode[T]
}

func (l *DoublyLinkedList[T]) Iterator() *DoublyLinkedListIterator[T] {
	return &DoublyLinkedListIterator[T]{current: l.Head}
}

func (it *DoublyLinkedListIterator[T]) Next() *LinkedListNode[T] {
	if it.current == nil {
		return nil
	}
	node := it.current
	it.current = it.current.Next
	return node
}

func (it *DoublyLinkedListIterator[T]) HasNext() bool {
	return it.current != nil
}
