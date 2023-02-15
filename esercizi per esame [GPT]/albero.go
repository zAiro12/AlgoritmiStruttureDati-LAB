package main

import (
	"fmt"
)

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func main() {

	var root *Node
	//root := &Node{Value: 5}

	root = root.Insert(5)
	root.Insert(3)
	root.Insert(7)
	root.Insert(1)
	root.Insert(4)
	root.Insert(6)
	root.Insert(8)

	fmt.Println("Tree created from main:")
	// Implement code to print the tree here
	// Example:
	// fmt.Println(root)
	// fmt.Println(root.Value)
	// fmt.Println(root.Left.Value)
	// fmt.Println(root.Right.Value)
	// ...

	parent := make(map[int]int)
	createParentTable(root, parent)
	root.PrintTree()
	fmt.Println(parent)
}

func (root *Node) Insert(value int) *Node {
	if root == nil {
		return &Node{value, nil, nil}
	} else if value <= root.Value {
		root.Left = root.Left.Insert(value)
	} else {
		root.Right = root.Right.Insert(value)
	}
	return root
}

func createParentTable(root *Node, parent map[int]int) {
	if root == nil {
		return
	}

	if root.Left != nil {
		parent[root.Left.Value] = root.Value
		createParentTable(root.Left, parent)
	}

	if root.Right != nil {
		parent[root.Right.Value] = root.Value
		createParentTable(root.Right, parent)
	}
}

func (root *Node) PrintTree() {
	if root == nil {
		return
	}
	root.Left.PrintTree()
	fmt.Println(root.Value)
	root.Right.PrintTree()
}
