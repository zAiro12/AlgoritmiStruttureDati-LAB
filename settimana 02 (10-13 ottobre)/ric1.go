package main

import "fmt"

func main() {
	a := 0
	b := 7

	fmt.Println(multiply(a, b))
}

func multiply(x, y int) int {
	if y == 0 {
		return y
	} else {
		return x + multiply(x, y-1)
	}
}
