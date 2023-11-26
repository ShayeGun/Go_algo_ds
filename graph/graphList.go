package graph

import (
	"fmt"
)

type GraphList struct {
	vertices []*Vertex
}

type Vertex struct {
	key       int
	neighbors []*Vertex
}

func (g *GraphList) addVertex(k int) {
	if contains(g.vertices, k) {
		fmt.Println("duplicate Vertex !")
		return
	}
	g.vertices = append(g.vertices, &Vertex{key: k})
}

func (g *GraphList) getVertex(k int) *Vertex {
	for _, val := range g.vertices {
		if val.key == k {
			return val
		}
	}
	return nil
}

func (g *GraphList) addEdge(from, to int) {
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)

	if fromVertex == nil || toVertex == nil {
		fmt.Println("at least 1 of the vertexes is invalid")
		return
	}

	if contains(fromVertex.neighbors, to) {
		fmt.Println("edge already exists")
		return
	}

	fromVertex.neighbors = append(fromVertex.neighbors, toVertex)
}

func (g *GraphList) print() {
	for _, k := range g.vertices {
		fmt.Printf("Vertex ==> %v : ", k.key)
		for _, e := range k.neighbors {
			fmt.Printf("%v, ", e.key)
		}
		fmt.Printf("\n")
	}
}

func contains(v []*Vertex, k int) bool {
	for _, val := range v {
		if val.key == k {
			return true
		}
	}
	return false
}
