package greedy

import (
	"fmt"
	"math"
)

func dijkstra(graph [][]int, start int) ([]int, []int) {
	visited := make([]bool, len(graph), len(graph))
	parent := make([]int, len(graph), len(graph))
	weight := make([]int, len(graph), len(graph))

	for v := range graph {
		visited[v] = false
		parent[v] = -1
		weight[v] = math.MaxInt32
	}

	weight[start] = 0

	// for simplicity we repeat this till vertices count
	for i := 0; i < len(graph); i++ {
		u := pickLowestWeight(weight, visited)

		if u != -1 {
			visited[u] = true
			for i, v := range graph[u] {
				if v+weight[u] < weight[i] && v > 0 {
					weight[i] = v + weight[u]
					parent[i] = u
				}
			}
		}
	}

	return parent, weight
}

func pickLowestWeight(weightArr []int, visited []bool) int {
	// if all graph's vertices weight is infinite function returns -1 --> which means there is no way
	index := -1
	max := math.MaxInt32

	for i, v := range weightArr {
		if v <= max && !visited[i] {
			max = v
			index = i
		}
	}

	return index
}

func PrintShortestPath(start int, finish int, parent []int, weight []int) {
	var paths []int

	for v := finish; v != start; v = parent[v] {
		u := parent[v]

		if weight[u] == math.MaxInt32 {
			fmt.Printf("Vertex %d is not reachable\n", finish)
			return
		}

		paths = append(paths, u)
	}

	paths = append([]int{finish}, paths...)

	fmt.Println(paths)
}
