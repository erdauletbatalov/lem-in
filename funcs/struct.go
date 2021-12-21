package funcs

// Graph structure

type Graph struct {
	Vertices []*Vertex
	Ants     int
	Start    string
	End      string
}

// Vertex structure
type Vertex struct {
	Key      string
	Adjacent []*Vertex
	Prev     *Vertex
}
