package graph

type GraphMatrix struct {
	matrix [][]int
}

func create(n int) *GraphMatrix {
	var g GraphMatrix
	g.matrix = make([][]int, n)
	for i := 0; i < n; i++ {
		g.matrix[i] = make([]int, n)
	}
	return &g
}

func (g *GraphMatrix) addEdge(v1 int, v2 int) {
	g.matrix[v1][v2] = 1
	g.matrix[v2][v1] = 1
}

func (g *GraphMatrix) removeEdge(v1 int, v2 int) {
	g.matrix[v1][v2] = 0
	g.matrix[v2][v1] = 0
}
