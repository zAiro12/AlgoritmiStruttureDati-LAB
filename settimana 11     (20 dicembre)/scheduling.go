package main

import (
	"bufio"
	"fmt"
	"os"
)

type intervallo struct {
	inizio, fine, peso int
}

func main() {
	file, _ := os.Open("inputScheduling.txt")
	sc := bufio.NewScanner(file)

	var inizio, fine, val int
	var intervalli []intervallo

	for sc.Scan() {
		fmt.Sscanf(sc.Text(), "%d-%d %d", &inizio, &fine, &val)
		intervalli = append(intervalli, intervallo{inizio: inizio, fine: fine, peso: val})
	}

	sortaIntervalli(intervalli)

	var intervalliFinali []intervallo
	maxval := intervalli[0].fine
	valTot := intervalli[0].peso
	intervalliFinali = append(intervalliFinali, intervalli[0])

	for i := 1; i < len(intervalli); i++ {
		if intervalli[i].inizio > maxval {
			intervalliFinali = append(intervalliFinali, intervalli[i])
			maxval = intervalli[i].fine
			valTot += intervalli[i].peso
		}
	}

	fmt.Println(valTot, intervalliFinali)
}

func sortaIntervalli(arr []intervallo) {
	var min int
	for i := 0; i < len(arr); i++ {
		min = cercaMin(arr[i:])
		min += i
		arr[i], arr[min] = arr[min], arr[i]
	}
}

func cercaMin(arr []intervallo) int {
	min := arr[0].fine
	var pos int
	for i := 1; i < len(arr); i++ {
		if arr[i].fine < min {
			min = arr[i].fine
			pos = i
		}
	}

	return pos
}
