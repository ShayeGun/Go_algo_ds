package graph

import "fmt"

func (g *GraphList) Dfs(start int) []int {

	startVertex := g.getVertex(start)
	visitedList := make(map[int]bool)
	dfsOrder := []int{}

	return traverseDfs(startVertex, visitedList, dfsOrder)
}

func traverseDfs(startVertex *Vertex, visitedList map[int]bool, dfsOrder []int) []int {

	if !visitedList[startVertex.key] {

		visitedList[startVertex.key] = true
		dfsOrder = append(dfsOrder, startVertex.key)

		for _, v := range startVertex.neighbors {
			dfsOrder = traverseDfs(v, visitedList, dfsOrder)
		}
	}

	return dfsOrder
}

func (g *GraphList) Bfs(start int) []int {

	startVertex := g.getVertex(start)
	visitedList := make(map[int]bool)
	stack := []*Vertex{startVertex}
	bfsOrder := []int{}

	return traverseBfs(stack, visitedList, bfsOrder)
}

func traverseBfs(stack []*Vertex, visitedList map[int]bool, bfsOrder []int) []int {

	if len(stack) == 0 {
		return bfsOrder
	}

	startVertex := stack[0]
	stack = stack[1:]

	if !visitedList[startVertex.key] {

		visitedList[startVertex.key] = true
		bfsOrder = append(bfsOrder, startVertex.key)

		for _, v := range startVertex.neighbors {
			stack = append(stack, v)
		}
	}

	return traverseBfs(stack, visitedList, bfsOrder)
}

func makeVisitedList(g *GraphList) map[int]bool {
	visitedList := make(map[int]bool)

	for _, v := range g.vertices {
		visitedList[v.key] = false
	}

	return visitedList
}

func Printy() {
	g := &GraphList{}
	g.addVertex(1)
	g.addVertex(2)
	g.addVertex(3)
	g.addVertex(4)
	g.addVertex(5)
	g.addVertex(6)
	g.addVertex(7)

	g.addEdge(1, 2)
	g.addEdge(1, 3)
	g.addEdge(2, 3)
	g.addEdge(2, 4)
	g.addEdge(3, 6)
	g.addEdge(3, 5)
	g.addEdge(5, 6)
	g.addEdge(5, 7)

	fmt.Println(g.Bfs(1))

	// g.print()
}
