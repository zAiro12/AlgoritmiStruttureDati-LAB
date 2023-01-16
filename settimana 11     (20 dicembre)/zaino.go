package main

import (
	"fmt"

	"golang.org/x/text/encoding/japanese"
)

type item struct {
	peso, valore int
}

func main() {
	var maxPesi int

	fmt.Print("inserire il peso massimo: ")
	fmt.Scan(&maxPesi)

	pesi := make([]int, maxPesi+1)
	oggetti := []item{{5, 2}, {3, 3}, {7, 8}}

	oggettiGenerali := make(map[item]int)
	oggettiGenerali[item{5,2}] = 3
	oggettiGenerali[item{3,3}] = 3
	oggettiGenerali[item{7,8}] = 3

	matricePesi := make([][]int, 9)

	for i := 0; i <= maxPesi; i++ {
		aggiornaPesi(pesi, oggetti, i)
		aggiornaMatrice(matricePesi, oggetti, i)
	}

	fmt.Println("risultato:", pesi[maxPesi])
}

func aggiornaPesi(pesi []int, oggetti []item, pesoCorrente int){
	var tempVal int
	
	for i := 0; i < len(oggetti); i++ {
		
		tempVal = pesi[pesoCorrente] + oggetti[i].valore
		
		if pesoCorrente+oggetti[i].peso < len(pesi) && pesi[pesoCorrente+oggetti[i].peso] < tempVal{
			pesi[pesoCorrente+oggetti[i].peso] = tempVal
		}
	}
}

func aggiornaMatrice(matrice [][]int, oggetti []item, pesoCorrente int){

	for i := 0; i < len(oggetti); i++ {
		
	}
}