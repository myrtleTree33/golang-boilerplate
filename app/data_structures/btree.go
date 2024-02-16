package datastructures

import "golang.org/x/exp/constraints"

type BTree[T constraints.Ordered] struct {
	Root *BTreeNode[T]
}

func NewBTree[T constraints.Ordered]() *BTree[T] {
	return &BTree[T]{}
}

func (b *BTree[T]) Insert(value T) {
	if b.Root == nil {
		b.Root = &BTreeNode[T]{Value: value}
		return
	}
	b.insert(b.Root, value)
}

func (b *BTree[T]) insert(node *BTreeNode[T], value T) {
	if value < node.Value {
		if node.Left == nil {
			node.Left = &BTreeNode[T]{Value: value}
		} else {
			b.insert(node.Left, value)
		}
	} else {
		if node.Right == nil {
			node.Right = &BTreeNode[T]{Value: value}
		} else {
			b.insert(node.Right, value)
		}
	}
}

func (b *BTree[T]) Inorder() *InorderIterator[T] {
	return NewInorderIterator[T](b.Root)
}

func (b *BTree[T]) Search(value T) bool {
	return b.search(b.Root, value)
}

func (b *BTree[T]) search(node *BTreeNode[T], value T) bool {
	if node == nil {
		return false
	}
	if value == node.Value {
		return true
	}
	if value < node.Value {
		return b.search(node.Left, value)
	}
	return b.search(node.Right, value)
}

func (b *BTree[T]) Delete(value T) {
	b.Root = b.delete(b.Root, value)
}

func (b *BTree[T]) delete(node *BTreeNode[T], value T) *BTreeNode[T] {
	if node == nil {
		return nil
	}
	if value < node.Value {
		node.Left = b.delete(node.Left, value)
	} else if value > node.Value {
		node.Right = b.delete(node.Right, value)
	} else {
		if node.Left == nil {
			return node.Right
		} else if node.Right == nil {
			return node.Left
		}
		node.Value = b.minValue(node.Right)
		node.Right = b.delete(node.Right, node.Value)
	}
	return node
}

func (b *BTree[T]) minValue(node *BTreeNode[T]) T {
	minValue := node.Value
	for node.Left != nil {
		minValue = node.Left.Value
		node = node.Left
	}
	return minValue
}

func (b *BTree[T]) MaxDepth() int {
	return b.maxDepth(b.Root)
}

func (b *BTree[T]) maxDepth(node *BTreeNode[T]) int {
	if node == nil {
		return 0
	}
	leftDepth := b.maxDepth(node.Left)
	rightDepth := b.maxDepth(node.Right)
	if leftDepth > rightDepth {
		return leftDepth + 1
	}
	return rightDepth + 1
}

func (b *BTree[T]) MinDepth() int {
	return b.minDepth(b.Root)
}

func (b *BTree[T]) minDepth(node *BTreeNode[T]) int {
	if node == nil {
		return 0
	}
	leftDepth := b.minDepth(node.Left)
	rightDepth := b.minDepth(node.Right)
	if leftDepth == 0 || rightDepth == 0 {
		return leftDepth + rightDepth + 1
	}
	if leftDepth < rightDepth {
		return leftDepth + 1
	}
	return rightDepth + 1
}

func (b *BTree[T]) IsBalanced() bool {
	return b.MaxDepth()-b.MinDepth() <= 1
}

func (b *BTree[T]) IsSymmetric() bool {
	return b.isSymmetric(b.Root, b.Root)
}

func (b *BTree[T]) isSymmetric(left *BTreeNode[T], right *BTreeNode[T]) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	return left.Value == right.Value && b.isSymmetric(left.Left, right.Right) && b.isSymmetric(left.Right, right.Left)
}

func (b *BTree[T]) IsSameTree(other *BTree[T]) bool {
	return b.isSameTree(b.Root, other.Root)
}

func (b *BTree[T]) isSameTree(node *BTreeNode[T], other *BTreeNode[T]) bool {
	if node == nil && other == nil {
		return true
	}
	if node == nil || other == nil {
		return false
	}
	return node.Value == other.Value && b.isSameTree(node.Left, other.Left) && b.isSameTree(node.Right, other.Right)
}

func (b *BTree[T]) IsSubtree(other *BTree[T]) bool {
	return b.isSubtree(b.Root, other.Root)
}

func (b *BTree[T]) isSubtree(node *BTreeNode[T], other *BTreeNode[T]) bool {
	if node == nil {
		return false
	}
	if b.isSameTree(node, other) {
		return true
	}
	return b.isSubtree(node.Left, other) || b.isSubtree(node.Right, other)
}

func (b *BTree[T]) IsCousins(x T, y T) bool {
	return b.isCousins(b.Root, x, y, 0, 0)
}

func (b *BTree[T]) isCousins(node *BTreeNode[T], x T, y T, xDepth int, yDepth int) bool {
	if node == nil {
		return false
	}
	if node.Left != nil && node.Right != nil {
		if (node.Left.Value == x && node.Right.Value == y) || (node.Left.Value == y && node.Right.Value == x) {
			return false
		}
	}
	if node.Value == x {
		xDepth = yDepth + 1
	}
	if node.Value == y {
		yDepth = yDepth + 1
	}
	return (xDepth == yDepth) && (b.isCousins(node.Left, x, y, xDepth, yDepth) || b.isCousins(node.Right, x, y, xDepth, yDepth))
}

func (b *BTree[T]) IsUnivalTree() bool {
	return b.isUnivalTree(b.Root, b.Root.Value)
}

func (b *BTree[T]) isUnivalTree(node *BTreeNode[T], value T) bool {
	if node == nil {
		return true
	}
	if node.Value != value {
		return false
	}
	return b.isUnivalTree(node.Left, value) && b.isUnivalTree(node.Right, value)
}

func (b *BTree[T]) IsComplete() bool {
	return b.isComplete(b.Root, 0, b.CountNodes())
}

func (b *BTree[T]) isComplete(node *BTreeNode[T], index int, count int) bool {
	if node == nil {
		return true
	}
	if index >= count {
		return false
	}
	return b.isComplete(node.Left, 2*index+1, count) && b.isComplete(node.Right, 2*index+2, count)
}

func (b *BTree[T]) CountNodes() int {
	return b.countNodes(b.Root)
}

func (b *BTree[T]) countNodes(node *BTreeNode[T]) int {
	if node == nil {
		return 0
	}
	return 1 + b.countNodes(node.Left) + b.countNodes(node.Right)
}

func (b *BTree[T]) IsSubtreeOfAnotherTree(other *BTree[T]) bool {
	return b.isSubtreeOfAnotherTree(b.Root, other.Root)
}

func (b *BTree[T]) isSubtreeOfAnotherTree(node *BTreeNode[T], other *BTreeNode[T]) bool {
	if node == nil {
		return false
	}
	if b.isSameTree(node, other) {
		return true
	}
	return b.isSubtreeOfAnotherTree(node.Left, other) || b.isSubtreeOfAnotherTree(node.Right, other)
}

func (b *BTree[T]) IsSymmetricTree() bool {
	return b.isSymmetricTree(b.Root, b.Root)
}

func (b *BTree[T]) isSymmetricTree(left *BTreeNode[T], right *BTreeNode[T]) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	return left.Value == right.Value && b.isSymmetricTree(left.Left, right.Right) && b.isSymmetricTree(left.Right, right.Left)
}

func (b *BTree[T]) IsBalancedTree() bool {
	return b.isBalancedTree(b.Root)
}

func (b *BTree[T]) isBalancedTree(node *BTreeNode[T]) bool {
	if node == nil {
		return true
	}
	return b.isBalancedTree(node.Left) && b.isBalancedTree(node.Right) && (b.maxDepth(node.Left)-b.minDepth(node.Right) <= 1)
}

func (b *BTree[T]) IsSameTreeAs(other *BTree[T]) bool {
	return b.isSameTreeAs(b.Root, other.Root)
}

func (b *BTree[T]) isSameTreeAs(node *BTreeNode[T], other *BTreeNode[T]) bool {
	if node == nil && other == nil {
		return true
	}
	if node == nil || other == nil {
		return false
	}
	return node.Value == other.Value && b.isSameTreeAs(node.Left, other.Left) && b.isSameTreeAs(node.Right, other.Right)
}

func (b *BTree[T]) IsSubtreeOf(other *BTree[T]) bool {
	return b.isSubtreeOf(b.Root, other.Root)
}

func (b *BTree[T]) isSubtreeOf(node *BTreeNode[T], other *BTreeNode[T]) bool {
	if node == nil {
		return false
	}
	if b.isSameTreeAs(node, other) {
		return true
	}
	return b.isSubtreeOf(node.Left, other) || b.isSubtreeOf(node.Right, other)
}

func (b *BTree[T]) IsCousin(x T, y T) bool {
	return b.isCousin(b.Root, x, y, 0, 0)
}

func (b *BTree[T]) isCousin(node *BTreeNode[T], x T, y T, xDepth int, yDepth int) bool {
	if node == nil {
		return false
	}
	if node.Left != nil && node.Right != nil {
		if (node.Left.Value == x && node.Right.Value == y) || (node.Left.Value == y && node.Right.Value == x) {
			return false
		}
	}
	if node.Value == x {
		xDepth = yDepth + 1
	}
	if node.Value == y {
		yDepth = yDepth + 1
	}
	return (xDepth == yDepth) && (b.isCousin(node.Left, x, y, xDepth, yDepth) || b.isCousin(node.Right, x, y, xDepth, yDepth))
}

func (b *BTree[T]) IsUnival() bool {
	return b.isUnival(b.Root, b.Root.Value)
}

func (b *BTree[T]) isUnival(node *BTreeNode[T], value T) bool {
	if node == nil {
		return true
	}
	if node.Value != value {
		return false
	}
	return b.isUnival(node.Left, value) && b.isUnival(node.Right, value)
}

func (b *BTree[T]) IsCompleteTree() bool {
	return b.isCompleteTree(b.Root, 0, b.CountNodes())
}

func (b *BTree[T]) isCompleteTree(node *BTreeNode[T], index int, count int) bool {
	if node == nil {
		return true
	}
	if index >= count {
		return false
	}
	return b.isCompleteTree(node.Left, 2*index+1, count) && b.isCompleteTree(node.Right, 2*index+2, count)
}

func (b *BTree[T]) Count() int {
	return b.count(b.Root)
}

func (b *BTree[T]) count(node *BTreeNode[T]) int {
	if node == nil {
		return 0
	}
	return 1 + b.count(node.Left) + b.count(node.Right)
}

func (b *BTree[T]) IsSubtreeOfAnother(other *BTree[T]) bool {
	return b.isSubtreeOfAnother(b.Root, other.Root)
}

func (b *BTree[T]) isSubtreeOfAnother(node *BTreeNode[T], other *BTreeNode[T]) bool {
	if node == nil {
		return false
	}
	if b.isSameTreeAs(node, other) {
		return true
	}
	return b.isSubtreeOfAnother(node.Left, other) || b.isSubtreeOfAnother(node.Right, other)
}
