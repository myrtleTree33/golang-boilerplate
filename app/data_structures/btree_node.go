package datastructures

import "golang.org/x/exp/constraints"

type BTreeNode[T constraints.Ordered] struct {
	Value T
	Left  *BTreeNode[T]
	Right *BTreeNode[T]
}

type InorderIterator[T constraints.Ordered] struct {
	stack Stack[*BTreeNode[T]]
	node  *BTreeNode[T]
}

func NewInorderIterator[T constraints.Ordered](node *BTreeNode[T]) *InorderIterator[T] {
	return &InorderIterator[T]{node: node}
}

func (i *InorderIterator[T]) HasNext() bool {
	return i.node != nil || !i.stack.IsEmpty()
}

func (i *InorderIterator[T]) Next() T {
	// Push all left nodes to the stack
	for i.node != nil {
		i.stack.Push(i.node)
		i.node = i.node.Left
	}

	// Pop the top node from the stack
	node := *i.stack.Pop()

	// Set the next node to the right node of the popped node
	i.node = node.Right
	return node.Value
}
