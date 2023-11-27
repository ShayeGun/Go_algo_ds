package dynamic

import (
	"fmt"
	"math"
)

type Edge struct {
	src    int
	dest   int
	weight float64
}

type Graph struct {
	vertices int
	edges    []Edge
}

func InitGraph(vertices int) *Graph {
	return &Graph{
		vertices: vertices,
		edges:    make([]Edge, 0),
	}
}

func (g *Graph) AddEdge(src, dest int, weight float64) {
	g.edges = append(g.edges, Edge{src, dest, weight})
}

func BellmanFord(g *Graph, source int) ([]float64, []int) {
	// each array's index represents a vertex in graph
	dist := make([]float64, g.vertices)
	prev := make([]int, g.vertices)

	// relaxing all vertexes
	for i := 0; i < g.vertices; i++ {
		dist[i] = math.Inf(1)
		prev[i] = -1
	}
	dist[source] = 0

	// finding shortest path
	for i := 0; i < g.vertices-1; i++ {
		for _, edge := range g.edges {
			u := edge.src
			v := edge.dest
			w := edge.weight

			if dist[u]+w < dist[v] {
				dist[v] = dist[u] + w
				prev[v] = u
			}
		}
	}

	// if graph has negative cycle
	for _, edge := range g.edges {
		u := edge.src
		v := edge.dest
		w := edge.weight
		if dist[u]+w < dist[v] {
			fmt.Println("Graph contains a negative weight cycle")
			return nil, nil
		}
	}

	return dist, prev
}

func PrintShortestPaths(dist []float64, prev []int, source int) {
	fmt.Println("Shortest Paths from vertex", source)
	for i := 0; i < len(dist); i++ {
		if dist[i] == math.Inf(1) {
			fmt.Printf("Vertex %d is not reachable\n", i)
		} else {
			path := []int{}
			j := i
			for j != -1 {
				// append the current value of j to the beginning of the path slice
				path = append([]int{j}, path...) // < path... > ==> break down slice to its elements
				j = prev[j]
			}
			fmt.Printf("Vertex %d: Distance=%f, Path=%v\n", i, dist[i], path)
		}
	}
}

func Printy() {
	g := InitGraph(3)

	g.AddEdge(0, 1, 6)
	g.AddEdge(0, 2, 7)
	g.AddEdge(1, 2, 8)
	// g.AddEdge(1, 3, -4)
	// g.AddEdge(1, 4, 5)
	// g.AddEdge(2, 3, 9)
	// g.AddEdge(2, 4, -3)
	// g.AddEdge(3, 1, 7)
	// g.AddEdge(4, 0, 2)
	// g.AddEdge(4, 3, 7)

	source := 0
	dist, prev := BellmanFord(g, source)
	fmt.Println(dist, prev)
	if dist != nil && prev != nil {
		PrintShortestPaths(dist, prev, source)
	}
}
