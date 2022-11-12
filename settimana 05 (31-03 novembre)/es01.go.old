package main

import (
	"fmt"
	"strconv"
	"strings"
)

type node struct {
	val  int
	next *node
}
type testa = *node

func main() {
	fmt.Println(valuta("5 3 - 2 *"))
	fmt.Println(valuta("2 5 3 - *"))
	fmt.Println(converti("( 5 - 3 ) * 2"))
	fmt.Println(converti("2 * ( 5 - 3 )"))
}

func valuta(espressione string) int {
	x := strings.Split(espressione, " ")

	var pila testa
	for _, k := range x {
		v, err := strconv.Atoi(k)
		if err == nil {
			push(&pila, v)
		} else {
			o2 := pop(&pila)
			o1 := pop(&pila)

			var ris int
			switch k {
			case "+":
				ris = o1 + o2
			case "-":
				ris = o1 - o2
			case "*":
				ris = o1 * o2
			case "/":
				ris = o1 / o2
			}
			push(&pila, ris)
		}
	}
	return pop(&pila)
}

func push(p *testa, x int) {
	var n node
	n.val = x
	n.next = *p
	*p = &n
}
func pop(p *testa) int {
	var x int
	appoggio := *p
	if *p != nil {
		x = appoggio.val
		*p = appoggio.next
	}
	return x
}

func converti(espressione string) string {
	var s string
	x := strings.Split(espressione, " ")
	var pila testa
	for _, k := range x {
		_, err := strconv.Atoi(k)
		if err == nil {
			s += k + " "
		} else if k == ")" {
			s += string(rune(pop(&pila))) + " "
		} else if k != "(" {
			push(&pila, int(k[0]))
		}
	}
	if !vuota(&pila) {
		s += string(rune(pop(&pila))) + " "
	}
	return s
}

func vuota(p *testa) bool {
	return p == nil
}
