package main

import (
	"math"
)

const boardSize = 30

// Node represents a node in the board
type Node struct {
	id        int
	neighbors map[*Node]int
}

// ShortestPath return the shortest path to win the game
func ShortestPath(start *Node, end *Node) int {
	dist := make([]int, boardSize)
	previous := make([]*Node, boardSize)
	visited := make([]bool, boardSize)

	// Initialize the distance to all nodes to be infinite
	// and the start node to have distance 0
	for i := 0; i < boardSize; i++ {
		dist[i] = math.MaxInt32
	}
	dist[start.id] = 0

	for len(visited) != boardSize {
		// find the unvisited node with the smallest distance
		minDist := math.MaxInt32
		var nextNode *Node
		for i, node := range previous {
			if !visited[i] && dist[i] <= minDist {
				nextNode = node
				minDist = dist[i]
			}
		}

		// visit the next node
		visited[nextNode.id] = true
		for neighbor, weight := range nextNode.neighbors {
			newDist := dist[nextNode.id] + weight
			if newDist < dist[neighbor.id] {
				dist[neighbor.id] = newDist
				previous[neighbor.id] = nextNode
			}
		}
	}

	return dist[end.id]
}
