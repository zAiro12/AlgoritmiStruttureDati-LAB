package main

import (
	"fmt"
)

func main() {
	const N = 10
	parole := make([]string, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&parole[i])
	}
	fmt.Println(piuCorta(parole))
}

func quanteConA(parole []string) int {
	c := 0
	for i := 0; i < len(parole); i++ {
		for _, k := range parole[i] {
			if k == 'a' {
				c++
				break
			}
		}
	}
	return c
}

func primaConA(parole []string) string {
	for i := 0; i < len(parole); i++ {
		if parole[i][0] == 'a' {
			return parole[i]
		}
	}
	var s string
	return s
}

func piuCorta(parole []string) int {
	min := len(parole[0])
	for i := 0; i < len(parole); i++ {
		if len(parole[i]) < min {
			min = len(parole[i])
			fmt.Println(len(parole[i]))
		}
	}
	return min
}
