package main

import (
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	arr := rand.Perm(10)
}

func mergeSort(arr []int) []int {
	return nil
}
