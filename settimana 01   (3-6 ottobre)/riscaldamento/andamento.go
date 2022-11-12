package main

import "fmt"

func main() {
	var prima, dopo int
	fmt.Scan(&dopo)
	for dopo !=0{
		prima = dopo
		fmt.Scan(&prima)
		if dopo >= prima{
			fmt.Println("+")
		}else{
			fmt.Println("-")
		}
	}
}