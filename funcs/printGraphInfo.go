package funcs

import "fmt"

func (g *Graph) PrintGraphInfo() {
	fmt.Println()
	fmt.Printf("Ants count = %v\n", g.Ants)
	fmt.Printf("Start key is %v\n", g.Start)
	fmt.Printf("End key is %v\n", g.End)
	fmt.Printf("Vertices count = %v\n", len(g.Vertices))
	fmt.Printf("Edges count = %v\n", g.Edges)
	fmt.Println()
}
