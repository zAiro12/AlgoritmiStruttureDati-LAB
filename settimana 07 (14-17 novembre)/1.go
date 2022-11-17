package main

import "fmt"

type colleghi struct {
	supervisore string
	sotto       []string
}

func main() {
	azienda := make(map[string]colleghi)

	azienda["radici"] = colleghi{"", []string{"a", "g"}}

	azienda["a"] = colleghi{"", []string{"b", "c", "d"}}
	azienda["b"] = colleghi{"a", []string{"e", "f"}}
	azienda["c"] = colleghi{"a", []string{}}
	azienda["d"] = colleghi{"a", []string{}}
	azienda["e"] = colleghi{"b", []string{}}
	azienda["f"] = colleghi{"b", []string{"i"}}
	azienda["i"] = colleghi{"f", []string{}}

	azienda["g"] = colleghi{"", []string{"h"}}
	azienda["h"] = colleghi{"g", []string{}}

	stampaImpiegatiSopra("i", azienda)
	stampaImpiegatiConSupervisore(azienda)
}
func stampaSubordinati(dipendente string, azienda map[string]colleghi) []string {
	return azienda[dipendente].sotto
}

func quantiSenzaSubordinati(azienza map[string]colleghi) int {
	c := 0
	for _, v := range azienza {
		if len(v.sotto) == 0 {
			c++
		}
	}
	return c
}

func supervisore(dipendente string, azienda map[string]colleghi) string {
	return azienda[dipendente].supervisore
}

func stampaImpiegatiSopra(dipendente string, azienda map[string]colleghi) {
	if azienda[dipendente].supervisore == "" {
		fmt.Println(dipendente)
		return
	}
	sup := azienda[dipendente].supervisore
	for _, v := range azienda[sup].sotto {
		fmt.Print(v, " ")
	}
	stampaImpiegatiSopra(sup, azienda)
}

func stampaImpiegatiConSupervisore(azienda map[string]colleghi) {
	for k, v := range azienda {
		if v.supervisore != "" {
			fmt.Print(k, " > ", v.supervisore, "  ")
		}
	}
	fmt.Println()
}
