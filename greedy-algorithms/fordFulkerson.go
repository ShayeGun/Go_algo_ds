package greedy

import (
	"fmt"
	"math"
)

/*
Returns true if there is a path from source 's' to
sink 't' in residual graph. Also fills parent[] to
store the path
*/
func bfs(g [][]int, s, t int, parent []int) bool {
	visited := make([]bool, len(g), len(g))
	var queue []int

	for i := range g {
		visited[i] = false
		parent[i] = -1
	}
	visited[s] = true
	queue = append(queue, s)

	for len(queue) != 0 {
		u := queue[0]
		queue = queue[1:]

		for i, w := range g[u] {
			if visited[i] == false && w > 0 {

				/*
					If we find a connection to the sink
					node, then there is no point in BFS
					anymore We just have to set its parent
					and can return true
				*/
				if i == t {
					parent[i] = u
					return true
				}

				visited[i] = true
				queue = append(queue, i)
				parent[i] = u
			}
		}
	}

	// Augment the flow while there is path from source to sink
	return false
}

func fordFulkerson(g [][]int, source, sink int) int {

	/*
		Create a residual graph and fill the residual
		graph with given capacities in the original graph
		as residual capacities in residual graph
	*/
	rGraph := make([][]int, len(g), len(g))

	for i := range rGraph {
		rGraph[i] = make([]int, len(g), len(g))
	}

	for i := range g {
		for j := range g[i] {
			rGraph[i][j] = g[i][j]
		}
	}

	// This array is filled by BFS and to store path
	parent := make([]int, len(g), len(g))
	// There is no flow initially
	maxFlow := 0

	// Augment the flow while there is path from source
	// to sink
	for bfs(rGraph, source, sink, parent) {
		pathFlow := math.MaxInt32

		for v := sink; v != source; v = parent[v] {
			u := parent[v]
			pathFlow = min(pathFlow, rGraph[u][v])
		}

		for v := sink; v != source; v = parent[v] {
			u := parent[v]
			// from source flow only can go out and flow only can enter sink not get out of it
			if u == source || v == sink {
				rGraph[u][v] -= pathFlow
			} else {
				rGraph[u][v] -= pathFlow
				rGraph[v][u] += pathFlow
			}
		}

		maxFlow += pathFlow
	}

	return maxFlow
}

func Printy() {
	g := [][]int{
		{0, 16, 13, 0, 0, 0},
		{0, 0, 10, 12, 0, 0},
		{0, 4, 0, 0, 14, 0},
		{0, 0, 9, 0, 0, 20},
		{0, 0, 0, 7, 0, 4},
		{0, 0, 0, 0, 0, 0},
	}

	fmt.Println(fordFulkerson(g, 0, 5))

}
