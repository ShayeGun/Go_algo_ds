package graph

import (
	"fmt"
)

type GraphList struct{
	vertices []*vertex
}

type vertex struct {
	key int
	neighbors []*vertex
}

func (g *GraphList) addVertex (k int){
	g.vertices = append(g.vertices, &vertex{key:k})
}

func (g *GraphList) print (){
	for _,k := range g.vertices{
		fmt.Printf("vertex ==> %v : ",k.key)
		for _,e := range k.neighbors{
		fmt.Printf("%v, ",e.key)
		}
	fmt.Printf("\n")
	}
}

func contains (v []*vertex,k int) bool{
	for _,val := range v {
		if val.key == k {
			return true
		}
	}
	return false
}

// type GraphMatrix struct{
// 	size int
// 	arr [][] int
// }

func Printy(){
	gr := GraphList{}
	gr.addVertex(1)
	gr.addVertex(2)
	gr.addVertex(4)
	
	gr.print()
}