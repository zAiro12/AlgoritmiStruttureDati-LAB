package main

import (
	"fmt"
	"strings"
)

func main() {
	var in string
	mappa := map[byte]int{}
	for {
		_, err := fmt.Scan(&in)
		if err != nil {
			var lettera string
			fmt.Scan(&lettera)
			lettera = strings.ToUpper(lettera)
			fmt.Println(mappa[lettera[0]])
		}
		in = strings.ToUpper(in)
		mappa[in[0]]++
	}
}
