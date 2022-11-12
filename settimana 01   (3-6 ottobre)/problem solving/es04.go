package main

import (
	"fmt"
)

// MAPPA DI BOOLE
func main() {
	permutazione := []int{4, 5, 1, 3, 6, 2}
	c := 0
	visto := make(map[int]bool)
	for _, i := range permutazione {
		visto[i] = true
		if visto[i+1] {
			c++
		}
	}
	fmt.Println(c)
}
