package main

import "fmt"

func main() {
	// Crea un grafo non orientato
	graph := make(map[int][]int)

	// Aggiungi nodi al grafo
	var n int
	fmt.Println("Quanti numeri vuoi inserire?")
	fmt.Scanln(&n)
	fmt.Println("Inserisci i numeri:")
	for i := 0; i < n; i++ {
		var node int
		fmt.Scanln(&node)
		graph[node] = []int{}
	}

	// Aggiungi archi tra i nodi se la differenza tra i loro valori è <= 2
	for node1 := range graph {
		for node2 := range graph {
			if node1 != node2 && abs(node1-node2) <= 2 {
				addEdge(graph, node1, node2)
			}
		}
	}

	// Stampa il grafo
	fmt.Println("Grafo:")
	for node, neighbors := range graph {
		fmt.Printf("%d -> ", node)
		for _, neighbor := range neighbors {
			fmt.Printf("%d ", neighbor)
		}
		fmt.Println()
	}
}

// Aggiungi un arco tra due nodi in un grafo non orientato
func addEdge(graph map[int][]int, node1 int, node2 int) {
	if _, ok := graph[node1]; !ok {
		graph[node1] = []int{}
	}
	if _, ok := graph[node2]; !ok {
		graph[node2] = []int{}
	}
	if !contains(graph[node1], node2) {
		graph[node1] = append(graph[node1], node2)
	}
	if !contains(graph[node2], node1) {
		graph[node2] = append(graph[node2], node1)
	}
}

// Verifica se un elemento è presente in un array
func contains(array []int, element int) bool {
	for _, v := range array {
		if v == element {
			return true
		}
	}
	return false
}

// Calcola il valore assoluto di un intero
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
