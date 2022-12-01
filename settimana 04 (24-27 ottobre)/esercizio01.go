package main

import "fmt"

type listNode struct {
	item int
	next *listNode
}
type linkedList struct {
	head *listNode
}

func main() {
}

func newNode(n int) *listNode {
	return &listNode{n, nil}
}

func addNewNode(l *linkedList, n int) {
	nuovoNodo := newNode(n)
	nuovoNodo.next = l.head
	l.head = nuovoNodo
}

func printList(l linkedList) {
	x := l.head
	for x != nil {
		fmt.Print(x.item, " ")
		x = x.next
	}
	fmt.Println()
}

func searchList(l linkedList, x int) (bool, *listNode) {
	p := l.head
	for p != nil {
		if p.item == x {
			return true, p
		}
		p = p.next
	}
	return false, nil
}

func removeItem(l linkedList, x int) bool {
	var curr, prev *listNode = l.head, nil
	for curr != nil {
		if curr.item == x {
			if prev == nil {
				l.head = l.head.next
			} else {
				prev.next = curr.next
			}
			return true
		}
		prev = curr
		curr = curr.next
	}
	return false
}
