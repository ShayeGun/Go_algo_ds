package greedy

import (
	"fmt"
	"sort"
)

type Graph struct {
	vertices int
	edges    []Edge
}

type Edge struct {
	weight int
	src    int
	dest   int
}

func (g *Graph) addEdge(e Edge) *Graph {
	g.edges = append(g.edges, e)
	return g
}

func find(a int, parent []int) int {
	for parent[a] >= 0 {
		a = parent[a]
	}
	return a
}

func union(a, b int, parent []int) bool {
	p1 := find(a, parent)
	p2 := find(b, parent)

	if p1 == p2 {
		return true
	}

	if parent[p1] <= parent[p2] {
		parent[p1] += parent[p2]
		parent[p2] = p1
	} else {
		parent[p2] += parent[p1]
		parent[p1] = p2
	}

	return false
}

func kruskal(graph *Graph) ([]Edge, int) {

	var result []Edge
	cost := 0
	parent := make([]int, graph.vertices)

	for i := range parent {
		parent[i] = -1
	}

	// Sort edges in increasing order
	sort.Slice(graph.edges, func(i, j int) bool {
		return graph.edges[i].weight < graph.edges[j].weight
	})

	for _, edge := range graph.edges {
		if !union(edge.src, edge.dest, parent) {
			result = append(result, edge)
			cost += edge.weight
		}
	}

	return result, cost

}

func Printy() {
	g := &Graph{}

	g.vertices = 5
	g.addEdge(Edge{2, 0, 2})
	g.addEdge(Edge{6, 0, 3})
	g.addEdge(Edge{5, 1, 2})
	g.addEdge(Edge{1, 1, 4})
	g.addEdge(Edge{2, 2, 3})
	g.addEdge(Edge{3, 3, 4})

	fmt.Println(kruskal(g))

}
