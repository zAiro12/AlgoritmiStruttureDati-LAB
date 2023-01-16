package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("inputJetPoints.txt")
	sc := bufio.NewScanner(file)

	var split []string
	var numSwindles, numPoint int
	swindles := make(map[int]int)

	for sc.Scan() {
		split = strings.Split(sc.Text(), " ")
		numSwindles, _ = strconv.Atoi(split[0])
		numPoint, _ = strconv.Atoi(split[1])

		swindles[numPoint] = numSwindles
	}

	var numJetPoints int
	fmt.Print("inserire numero di jet-points: ")
	fmt.Scan(&numJetPoints)

	
}
