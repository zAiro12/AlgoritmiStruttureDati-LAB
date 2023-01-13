package main

import (
	"fmt"
	"strconv"
)

const DimensioneGriglia = 30

func main() {
	scaleSerpenti := make(map[int]int)

	var in1, in2 string
	var num1, num2 int
	for {
		_, err := fmt.Scan(&in1)

		if err != nil {
			break
		}
		fmt.Scan(&in2)
		num1, _ = strconv.Atoi(in1)
		num2, _ = strconv.Atoi(in2)

		scaleSerpenti[num1] = num2
	}

	griglia := creaGriglia(scaleSerpenti)

	minLanci, lanci := minLanci(griglia, 3)
	fmt.Println(minLanci, lanci)

	minLanciSenzaSerpenti, lancisenza := minLanciSenzaSerpenti(scaleSerpenti, griglia, 1)
	fmt.Println(minLanciSenzaSerpenti, lancisenza)

}

func creaGriglia(scaleSerpenti map[int]int) map[int][]int {
	griglia := make(map[int][]int)

	for i := 1; i <= DimensioneGriglia; i++ {
		for j := 1; j <= 6; j++ {

			if scaleSerpenti[i+j] == 0 {
				if i+j <= 30 {
					griglia[i] = append(griglia[i], i+j)
				} else {
					griglia[i] = append(griglia[i], -1)
				}

			} else {
				griglia[i] = append(griglia[i], scaleSerpenti[i+j])
			}
		}
	}

	return griglia
}

func minLanci(griglia map[int][]int, posizioneCorrente int) (int, []int) {
	nLanci := 0
	var lanci []int
	var pos, temp, maxNum int

	if posizioneCorrente == DimensioneGriglia {
		return 0, lanci
	}

	for {
		for i := 0; i < len(griglia[posizioneCorrente]); i++ {
			temp, pos = max(griglia[posizioneCorrente])
			if temp > maxNum {
				maxNum = temp
			}
		}

		lanci = append(lanci, pos+1)
		nLanci++
		posizioneCorrente = maxNum

		if temp == 30 {
			break
		} else {
			temp = 0
			maxNum = 0
		}
	}

	return nLanci, lanci
}

func max(arr []int) (int, int) {
	var max, maxi int
	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
			maxi = i
		}
	}
	return max, maxi
}

func minLanciSenzaSerpenti(scaleSerpenti map[int]int, griglia map[int][]int, posCorrente int) (int, []int) {

	var nLanci, pos int
	var lanci []int

	if posCorrente == DimensioneGriglia {
		return 0, lanci
	}

	for {
		fmt.Println(posCorrente, scaleSerpenti)

		posCorrente, pos = maxSenzaScaleSerpenti(griglia[posCorrente], scaleSerpenti)

		nLanci++
		lanci = append(lanci, pos)

		if posCorrente == DimensioneGriglia {
			break
		}
	}

	return nLanci, lanci
}

func maxSenzaScaleSerpenti(arr []int, scaleSerpenti map[int]int) (int, int) {
	var max, maxi int
	var flag bool
	for i := 0; i < len(arr); i++ {
		if arr[i] > max && arr[i] != -1 {
			for _, v := range scaleSerpenti {
				if v == arr[i] {
					flag = true
				}
			}
			if !flag {
				max = arr[i]
				maxi = i + 1
			}
		}
		flag = false
	}

	return max, maxi
}
