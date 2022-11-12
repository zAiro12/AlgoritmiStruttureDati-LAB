package main

import (
	"fmt"
	"strings"
)

type node struct {
	val  string
	next *node
}
type lista struct {
	testa *node
}

func main() {
	fmt.Println(check("<a> <b> </b> <c> <d> </d> </c> </a>"))
	fmt.Println(check("<a> <b> </a> </c>"))
	fmt.Println(check("<a> <b> </a> </b>"))
}

func check(input string) bool {
	var pila lista
	in := strings.Split(input, " ")

	for i := 0; i < len(in); i++ {
		caso := in[i]
		if len(caso) == 3 {
			add(&pila, caso)
			//fmt.Println("i:", i) 						//DEBUG
			//stampa(pila)         						//DEBUG
		} else {
			caso = caso[:1] + caso[2:]
			controllo := pop(&pila)
			//fmt.Println("AAAA", in, controllo) 		//DEBUG
			//stampa(pila)                      		//DEBUG
			if controlla(caso, controllo) == false {
				//fmt.Println("error pos", i) 			//DEBUG
				return false
			}
		}
	}
	return true
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
	if pila.testa.next != nil {
		*pila.testa = *pila.testa.next
	} else {
		*pila.testa = node{"", nil}
	}
	return x
}

func controlla(a, b string) bool {
	return a == b
}

func stampa(pila lista) {
	appoggio := pila.testa

	for appoggio != nil {
		fmt.Print(appoggio.val, " ")
		appoggio = appoggio.next
	}
	fmt.Println()
}
