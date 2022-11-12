package main

import "fmt"

func main() {
	fmt.Println(sassi(3))
}

func sassi(height int) int {
	if height == 0 {
		return height
	}
	return height*height + sassi(height-1)
}
