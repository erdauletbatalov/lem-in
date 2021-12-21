package funcs

import "fmt"

// Add Edge
func (g *Graph) AddEdge(from, to string) error {
	err := fmt.Errorf("Error")
	//get vertex
	fromVertex := g.GetVertex(from)
	toVertex := g.GetVertex(to)
	//check error
	if fromVertex == nil || toVertex == nil {
		return err
	} else if contains(fromVertex.Adjacent, to) {
		return err
	} else {
		//add edge
		fromVertex.Adjacent = append(fromVertex.Adjacent, toVertex)
		return nil
	}
}
