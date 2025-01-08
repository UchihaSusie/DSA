// author: @UchihaSusie
// algorithms/Go/graphs/bfs_sequence.go
package main

import (
	"fmt"
)

type Graph struct {
	AdjDict map[int][]int
}

func (g *Graph) ShowGraph() {
	for i, neighbors := range g.AdjDict {
		fmt.Printf("%d -> %v\n", i, neighbors)
	}
}

func (g *Graph) DisplayBFS(start int) {
	visited := make(map[int]bool)
	queue := []int{start}
	visited[start] = true

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:] // dequeue
		fmt.Print(curr, " ")

		for _, neighbor := range g.AdjDict[curr] {
			if !visited[neighbor] {
				queue = append(queue, neighbor)
				visited[neighbor] = true
			}
		}
	}
}

func main() {
	g := Graph{
		AdjDict: map[int][]int{
			1: {2, 4},
			2: {4, 5},
			4: {7, 5},
			5: {1, 3, 6},
			6: {3, 8},
			8: {7},
		},
	}

	fmt.Println("Display Graph")
	g.ShowGraph()

	fmt.Println("BFS Sequence")
	g.DisplayBFS(1) // passing start node
}