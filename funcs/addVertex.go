package funcs

import "fmt"

// AddVertex adds a Vertex to the Graph// Vertex structure
func (g *Graph) AddVertex(k string) {
	if contains(g.Vertices, k) {
		err := fmt.Errorf("Vertex %v not added because it is an existing Key", k)
		fmt.Println(err.Error())
	} else {
		g.Vertices = append(g.Vertices, &Vertex{Key: k})
	}
}

func (g *Graph) Print() {
	for _, val := range g.Vertices {
		fmt.Printf("\nVertex %v : ", val.Key)
		for _, v := range val.Adjacent {
			fmt.Printf("%v ", v.Key)
		}
	}
	fmt.Println()
}
