package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	array := creaArray()
	fmt.Println("prima:", array)

	var min int
	for i := 0; i < len(array); i++ {
		min = cercaIndexMin(array[i:]) + i
		array[i], array[min] = array[min], array[i]
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

func cercaIndexMin(array []int) int {
	var min, indexMin int

	min = array[0]
	for i := 1; i < len(array); i++ {
		if array[i] < min {
			min = array[i]
			indexMin = i
		}
	}

	return indexMin
}
