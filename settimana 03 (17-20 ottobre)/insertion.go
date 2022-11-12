package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	rnd := rand.Intn(100)
	flag:= false
	var arr []int
	arr = append(arr, rnd)
	for i := 0; i < 99; i++ {
		rnd = rand.Intn(100)
		if rnd > arr[len(arr)-1] {
			arr = append(arr, rnd)
		}else{
			for j := len(arr)-2; j >= 0; j-- {
				if rnd > arr[j]{
					arr = mettiMezzo(arr, rnd, j+1)
					flag = true
					break
				}
			}
			if flag{
				flag = false
			}else{
				arr = mettiMezzo(arr, rnd, 0)
			}
		}
	}
	fmt.Println(len(arr), arr)
}

func mettiMezzo(a []int, x int, pos int) []int {
	
	var b []int
	b = append(b, a[:pos]...)
	b = append(b, x)
	b = append(b, a[pos:]...)
	return b
}
