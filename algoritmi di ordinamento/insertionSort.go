package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	array := creaArray()
	fmt.Println("prima:", array)

	var x, j int
	for i := 1; i < len(array); i++ {
		x = array[i]
		j = i - 1

		for j >= 0 && array[j] > x {
			array[j+1] = array[j]
			j--
		}
		array[j+1] = x
	}

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
