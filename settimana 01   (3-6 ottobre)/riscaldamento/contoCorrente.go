package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	conto, _ := strconv.Atoi(os.Args[1])
	var in int
	for conto > 0 {
		fmt.Scan(&in)
		conto -= in
	}
	fmt.Println(conto)
}
