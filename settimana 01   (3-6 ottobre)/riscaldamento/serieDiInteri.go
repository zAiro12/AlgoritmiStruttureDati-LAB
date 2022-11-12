package main

import "fmt"

func main() {
	const N = 10
	numeri := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&numeri[i])
	}
	fmt.Println(minDispari(numeri))
}

func stranoProdotto(numeri []int) int {
	prodotto := 1
	for i := 0; i < len(numeri); i++ {
		if numeri[i] > 7 && numeri[i]%4 == 0 {
			prodotto *= numeri[i]
		}
	}
	return prodotto
}

func pariDispari(numeri []int) {
	for i := 0; i < len(numeri); i++ {
		if numeri[i]%2 == 0 {
			fmt.Println("pari")
		} else {
			fmt.Println("dispari")
		}
	}
}

func minDispari(numeri []int) int {
	min := numeri[0]
	for i := 0; i < len(numeri); i++ {
		if numeri[i]%2 != 0 && numeri[i] < min {
			min = numeri[i]
		}
	}
	if min%2 == 0 {
		return 0
	} else {
		return min
	}
}
