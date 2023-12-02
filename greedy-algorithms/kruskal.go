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
	/*
		negative sign means that vertex is root and the negative number is number of children vertices of that vertex
		for example [-2,0,0,-1] means vertex 0 is parent of 2 vertices (1,2) and vertex 3 is a single vertex (-1)
	*/
	// in this kind of representation we cluster vertices and show it by a representative vertex (root vertex like 0)
	for parent[a] >= 0 {
		a = parent[a]
	}
	return a
}

func union(a, b int, parent []int) bool {
	p1 := find(a, parent)
	p2 := find(b, parent)

	// if parent of two vertices are same which means there is a cycle
	if p1 == p2 {
		return true
	}

	// vertex with more children must be root
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
