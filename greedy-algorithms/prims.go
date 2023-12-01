package greedy

import (
	"math"
)

func prims(graph [][]int, start int) (int, []int) {
	// make a copy of graph
	g := make([][]int, len(graph), len(graph))

	for i := range g {
		g[i] = make([]int, len(graph), len(graph))
	}

	for n := range graph {
		for m := range graph[n] {
			g[m][n] = graph[m][n]
		}
	}

	// remove edge to the same vertex (remove single vertex loop)
	for n := range graph {
		for m := range graph[n] {
			if m == n {
				g[m][n] = 0
			}
		}
	}

	visited := make([]bool, len(g), len(g))
	parent := make([]int, len(g), len(g))
	minCost := 0

	for i := range g {
		visited[i] = false
		parent[i] = -1
	}

	visited[start] = true

	// iterate for vertices number
	for i := 0; i != len(g)-1; i++ {
		min := math.MaxInt32
		var minVertexColumn int
		var minVertexRow int

		// only check already visited vertices
		for j := range visited {
			if visited[j] {
				for k := range g[j] {
					// only check edges of unvisited vertices
					if !visited[k] && g[j][k] < min && g[j][k] != 0 {
						min = g[j][k]
						minVertexColumn = k
						minVertexRow = j
					}
				}
			}
		}

		visited[minVertexColumn] = true
		minCost += min
		parent[minVertexColumn] = minVertexRow
	}

	return minCost, parent
}
