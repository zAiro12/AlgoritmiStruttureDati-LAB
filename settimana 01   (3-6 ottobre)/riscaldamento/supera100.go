package main

import (
	"fmt"
)

func main() {
	var in int
	for{
		fmt.Scan(&in)
		if in == -1 {
			break
		}
		if in > 100 {
			fmt.Println(in)
			return
		}
	}
	fmt.Println("nessun numero maggiore di 100")
}
