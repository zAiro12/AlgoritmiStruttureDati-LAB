package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	array := creaArray()
	fmt.Println("prima:", array)

	array = mergeSort(array)

	fmt.Println("dopo:", array)
}

func creaArray() []int {
	var array []int

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < rand.Intn(100); i++ {
		array = append(array, rand.Intn(100))
	}

	return array
}

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	m := len(arr) / 2
	sx := mergeSort(arr[:m])
	dx := mergeSort(arr[m:])

	return merge(sx, dx)
}

func merge(sx, dx []int) []int {
	x := make([]int, 0, len(sx)+len(dx))

	for len(sx) > 0 || len(dx) > 0 {
		if len(sx) == 0 {
			return append(x, dx...)
		}
		if len(dx) == 0 {
			return append(x, sx...)
		}
		if sx[0] <= dx[0] {
			x = append(x, sx[0])
			sx = sx[1:]
		} else {
			x = append(x, dx[0])
			dx = dx[1:]
		}
	}

	return x
}
