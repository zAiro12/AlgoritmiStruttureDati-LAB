package main

import "fmt"

type vettore struct {
	n int
	s string
}

func main() {
	var v [9]vettore
	var in vettore
	for i := 0; i < 9; i++ {
		fmt.Scanf("%d %s", &in.n, &in.s)
		v[9-in.n] = in
	}
	stampa(v)
}

func stampa(v [9]vettore) {
	fmt.Print("\n\n")
	for i := 0; i < 9; i++ {
		fmt.Println(v[i].n, v[i].s)
	}
}
