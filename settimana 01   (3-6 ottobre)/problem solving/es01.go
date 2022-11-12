package main

import "fmt"

func main() {
	c := 0
	var uno, due, tre int
	fmt.Scan(&uno, &due)
	if due < uno {
		c++
	}
	for i := 2; i < 8; i++ {
		fmt.Scan(&tre)
		if due > uno && due < tre {
			c++
		}
		uno, due = due, tre
	}
	fmt.Println(c)
}
