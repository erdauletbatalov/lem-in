package funcs

// GetVertex returns a pointer to the Vertex with a Key integer
func (g *Graph) GetVertex(k string) *Vertex {
	for _, v := range g.Vertices {
		if v.Key == k {
			return v
		}
	}
	return nil
}
