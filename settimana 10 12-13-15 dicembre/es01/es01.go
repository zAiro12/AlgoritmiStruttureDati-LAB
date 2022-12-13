package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type galleria struct {
	nome       string
	luminosita int
}

func main() {
	file, _ := os.Open(os.Args[1])
	sc := bufio.NewScanner(file)
	defer file.Close()

	gallerie := make(map[string][]galleria)
	visitato := make(map[string]bool)
	var in string
	var split []string

	var H, S string
	var lum int
	for sc.Scan() {
		in = sc.Text()
		split = strings.Split(in, " ")

		if len(split) == 4 {
			H = split[2]
			S = split[3]

		} else {
			lum, _ = strconv.Atoi(split[2])
			gallerie[split[0]] = append(gallerie[split[0]], galleria{split[1], lum})
			gallerie[split[1]] = append(gallerie[split[1]], galleria{split[0], lum})

		}
	}

	var c int

	for {
		if H == "-1" || H == S || visitato[H] {
			if H == "-1" || visitato[H] {
				c = -1
			}
			break
		}
		visitato[H] = true
		H = trovaGalleria(gallerie[H])
		c++
	}

	fmt.Println(c)

}

func trovaGalleria(vicini []galleria) string {
	if len(vicini) == 1 {
		return vicini[0].nome
	}

	min := vicini[0].luminosita
	for i := 0; i < len(vicini); i++ {
		if vicini[i].luminosita < min {
			min = vicini[i].luminosita
			vicini = []galleria{{vicini[i].nome, min}}
			return vicini[0].nome
		}
	}

	return "-1"
}
