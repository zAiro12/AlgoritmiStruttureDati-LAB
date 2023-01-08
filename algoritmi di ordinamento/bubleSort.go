package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	array := creaArray()
	fmt.Println("prima:", array)

	var scambiato bool
	for{
		scambiato = false
		for i := 0; i < len(array)-1; i++ {
			if array[i] > array[i+1]{
				scambiato = true
				array[i], array[i+1] = array[i+1], array[i]
			}
		}
		if !scambiato{
			break
		}
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
