package main

import (
	"fmt"
	"math/rand"
	"time"
)

type bitreeNode struct {
	left  *bitreeNode
	right *bitreeNode
	val   int
}

type bitree struct {
	root *bitreeNode
}

func main() {
	rand.Seed(time.Now().UnixNano())
	//var albero bitree
	in := gen(rand.Intn(100))
	fmt.Println(in, len(in))

	/* t := &bitree{nil}
	t.root = &bitreeNode{nil, nil, 78}
	t.root.left = newNode(54)
	t.root.right = newNode(21)
	t.root.left.right = newNode(90)
	t.root.left.right.left = newNode(19)
	t.root.left.right.right = newNode(95)
	t.root.right.left = newNode(16)
	t.root.right.left.left = newNode(5)
	t.root.right.right = newNode(19)
	t.root.right.right.left = newNode(56)
	t.root.right.right.right = newNode(43)

	stampaAlberoASommario(t.root, 0)
	fmt.Println()*/
	stampaAlberoASommario(arr2tree(in, 0), 0)
}

func gen(x int) []int {
	var out []int
	for i := 0; i < x; i++ {
		out = append(out, rand.Intn(100))
	}
	return out
}

func newNode(val int) *bitreeNode {
	return &bitreeNode{nil, nil, val}
}

func preorder(node *bitreeNode) {
	//deepSearch
	if node == nil {
		return
	}
	fmt.Println(node.val)
	preorder(node.left)
	preorder(node.right)

}
func inorder(node *bitreeNode) {
	//figlio-padre-figlio
	if node == nil {
		return
	}
	inorder(node.left)
	fmt.Println(node.val, " ")
	inorder(node.right)
}
func postorder(node *bitreeNode) {
	//figli, radice
	if node == nil {
		return
	}
	inorder(node.left)
	inorder(node.right)
	fmt.Println(node.val, " ")
}

func stampaAlberoASommario(node *bitreeNode, spaces int) {
	for i := 0; i < spaces; i++ {
		fmt.Print(" ")
	}
	fmt.Print("*")
	fmt.Println(node.val)

	if node.left != nil {
		stampaAlberoASommario(node.left, spaces+1)
	} else if node.right != nil {
		for i := 0; i < spaces+1; i++ {
			fmt.Print(" ")
		}
		fmt.Println("*")
	}
	if node.right != nil {
		stampaAlberoASommario(node.right, spaces+1)
	} else if node.left != nil {
		for i := 0; i < spaces+1; i++ {
			fmt.Print(" ")
		}
		fmt.Println("*")
	}
}

func stampaAlbero(node *bitreeNode) {
	if node == nil {
		return
	}
	fmt.Print(node.val, "")
	if node.left == nil && node.right == nil {
		return
	}
	fmt.Print(" [")
	if node.left != nil {
		stampaAlbero(node.left)
	} else {
		fmt.Print("-")
	}
	fmt.Print(", ")
	if node.right != nil {
		stampaAlbero(node.right)
	} else {
		fmt.Print("-")
	}
	fmt.Print("]")
}
func arr2tree(a []int, i int) (root *bitreeNode) {
	if i >= len(a) {
		return nil
	}
	root = newNode(a[i])
	root.left = arr2tree(a, 2*i+1)
	root.right = arr2tree(a, 2*i+2)
	return root
}
