package main

import (
	datastructures "app/data_structures"
	"fmt"
)

func main() {

	// // Create a new stack
	// a := datastructures.NewStack[int]()
	// a.Push(1)
	// a.Push(2)
	// a.Push(3).Push(4)

	// for !a.IsEmpty() {
	// 	fmt.Println(*a.Pop())
	// }

	binaryTree := datastructures.NewBTree[int]()
	binaryTree.Insert(2)

	iterator := binaryTree.Inorder()
	for iterator.HasNext() {
		fmt.Print(iterator.Next(), " ")
	}

}
