package funcs

// Graph structure

type Graph struct {
	Vertices []*Vertex
	Ants     int
	Start    string
	End      string
	Edges    int
}

// Vertex structure
type Vertex struct {
	Key       string
	Adjacents []*Vertex
	Prev      *Vertex
}
