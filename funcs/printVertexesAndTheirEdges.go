package funcs

import "fmt"

func (g *Graph) PrintVertexesAndTheirEdges() {
	for _, val := range g.Vertices {
		fmt.Printf("\nVertex %v : ", val.Key)
		for _, v := range val.Adjacents {
			fmt.Printf("%v ", v.Key)
		}
	}
	fmt.Println()
}
