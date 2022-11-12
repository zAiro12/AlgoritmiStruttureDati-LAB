package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	arr := rand.Perm(10)
	fmt.Println(iterativo(arr))
	fmt.Println()
	fmt.Println(ricorsivo(arr))
}

func iterativo(a []int) []int {
	for i := 0; i < len(a)-1; i++ {
		for j := i + 1; j < len(a); j++ {
			if a[j] < a[i] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
	return a
}

func ricorsivo(a []int) []int {
	n := len(a) - 1
	if n == 0 {
		return a
	}
	for i := 0; i < len(a); i++ {
		if a[i] > a[n] {
			a[i], a[n] = a[n], a[i]
		}
	}
	ricorsivo(a[:n])
	return a
}
