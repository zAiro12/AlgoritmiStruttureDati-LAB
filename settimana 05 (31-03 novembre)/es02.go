package main

import "fmt"

type node struct {
	val  string
	next *node
}
type lista struct {
	testa *node
}

func main() {
	var pila lista

	
}

func newNode(val string) *node {
	return &node{val, nil}
}

func add(pila *lista, val string) {
	node := newNode(val)
	node.next = pila.testa
	pila.testa = node
}

func pop(pila *lista) string {
	var x string
	x = pila.testa.val
	*pila.testa = *pila.testa.next
	return x
}

func stampa(pila lista) {
	appoggio := pila.testa

	for appoggio != nil {
		fmt.Print(appoggio.val, " ")
		appoggio = appoggio.next
	}
	fmt.Println()
}
