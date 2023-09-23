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
	// studentGrades := map[string]int{
	// 	"John":  85,
	// 	"Alice": 92,
	// 	"Bob":   78,
	// }
	// arr := []int{
	// 	1,3,4,
	// }

	// name := "Ali"

	// if grade, ok := studentGrades[name]; ok {
	// 	fmt.Println("Grade for", name, grade,ok)
	// } else {
	// 	fmt.Println("No Grade for", name, grade,ok)
	// }
}