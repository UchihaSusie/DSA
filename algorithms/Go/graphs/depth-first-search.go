// author: @UchihaSusie	
// algorithms/Go/graphs/depth_first_search.go
package main

import "fmt"

// Depth First Search Algorithm
func dfs(edges [][]int, vis []int, node int) {
	if vis[node] != 1 {
		vis[node] = 1
		fmt.Print(node, " ")
	}

	for _, i := range edges[node] {
		if vis[i] != 1 {
			dfs(edges, vis, i)
		}
	}
}

func main() {
	// First Example
	graph1 := [][]int{
		{1, 2, 3},
		{0, 2},
		{0, 1, 4},
		{0},
		{2},
	}

	vis := make([]int, 10)

	fmt.Println("Graph 1 的 DFS 结果是:")
	dfs(graph1, vis, 0)

	// Reset visited list elements to 0 for the second graph
	for i := range vis {
		vis[i] = 0
	}

	graph2 := [][]int{
		{1, 2, 3},
		{3},
		{4},
		{5, 6},
		{5, 7},
		{2},
		{},
		{},
	}

	fmt.Println("\n\nGraph 2 的 DFS 结果是:")
	dfs(graph2, vis, 0)
}
