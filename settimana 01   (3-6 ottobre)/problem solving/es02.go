package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	var a, c int
	for i := 0; i < len(s); i++ {
		if s[i] == 'a' {
			a++
		}
		if s[i] == 'b' {
			c += a
		}
	}
	fmt.Println(c)
}
