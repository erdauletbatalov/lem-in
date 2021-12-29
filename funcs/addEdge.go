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
		fmt.Println("from or to nil")
		return err
	} else if contains(fromVertex.Adjacents, to) {
		fmt.Printf("%v - %v edge Already Exists\n", from, to)
		return err
	} else {
		//add edge
		fmt.Printf("%v - %v edge added\n", from, to)
		fromVertex.Adjacents = append(fromVertex.Adjacents, toVertex)
		return nil
	}
}
