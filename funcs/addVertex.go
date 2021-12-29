package funcs

import "fmt"

// AddVertex adds a Vertex to the Graph// Vertex structure
func (g *Graph) AddVertex(k string) {
	if contains(g.Vertices, k) {
		err := fmt.Errorf("Vertex %v not added because it is an existing Key", k)
		fmt.Println(err.Error())
	} else {
		fmt.Printf("%v vertex aded\n", k)
		g.Vertices = append(g.Vertices, &Vertex{Key: k})
	}
}
