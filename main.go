package main

import (
	"fmt"
)

// Graph structure
type Graph struct {
	vertices []*Vertex
}

// Vertex structure
type Vertex struct {
	key      int
	adjacent []*Vertex
	prev     *Vertex
}

// AddVertex adds a Vertex to the Graph// Vertex structure
func (g *Graph) AddVertex(k int) {
	if contains(g.vertices, k) {
		err := fmt.Errorf("Vertex %v not added because it is an existing key", k)
		fmt.Println(err.Error())
	} else {
		g.vertices = append(g.vertices, &Vertex{key: k})
	}
}

// getVertex returns a pointer to the Vertex with a key integer
func (g *Graph) getVertex(k int) *Vertex {
	for _, v := range g.vertices {
		if v.key == k {
			return v
		}
	}
	return nil
}

// Add Edge
func (g *Graph) AddEdge(from, to int) {
	//get vertex
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)
	//check error
	if fromVertex == nil || toVertex == nil {
		err := fmt.Errorf("Invalid edge (%v -> %v)", from, to)
		fmt.Println(err.Error())
	} else if contains(fromVertex.adjacent, to) {
		err := fmt.Errorf("Existing edge (%v -> %v)", from, to)
		fmt.Println(err.Error())
	} else {
		//add edge
		fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
	}
}

// contains
func contains(s []*Vertex, k int) bool {
	for _, v := range s {
		if k == v.key {
			return true
		}
	}
	return false
}

func (g *Graph) Print() {
	for _, val := range g.vertices {
		fmt.Printf("\nVertex %v : ", val.key)
		for _, v := range val.adjacent {
			fmt.Printf("%v ", v.key)
		}
	}
	fmt.Println()
}

func (g *Graph) BreadthFirstSearch(from, to int) {
	//get Vertex
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)
	//check error
	if fromVertex == nil || toVertex == nil {
		err := fmt.Errorf("Invalid Vertexes (%v or %v)", from, to)
		fmt.Println(err.Error())
	} else {
		var visitedVertex map[*Vertex]bool
		visitedVertex = make(map[*Vertex]bool)
		current := fromVertex
		queue := current.adjacent
		for {
			if current == toVertex {
				fmt.Println("I've found the path!")
				end := current
				path := []*Vertex{}
				for end != nil {
					path = append(path, end)
					end = end.prev
				}
				path = reverseArray(path)
				for _, v := range path {
					if v == path[len(path)-1] {
						fmt.Printf("%v", v.key)
						fmt.Println()
					} else {
						fmt.Printf("%v -> ", v.key)
					}
				}
				return
			}
			for _, v := range current.adjacent {
				if !visitedVertex[v] {
					queue = append(queue, v)
					v.prev = current
					visitedVertex[v] = true
				}
			}
			if current != queue[0] {
				current = queue[0]
			} else {
				fmt.Println("There is no such path =(")
				return
			}
			if len(queue) != 1 {
				queue = queue[1:]
			}
		}
	}
}

func reverseArray(arr []*Vertex) []*Vertex {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

func main() {
	test := &Graph{}

	for i := 1; i < 9; i++ {
		test.AddVertex(i)
	}

	test.AddEdge(1, 2)
	test.AddEdge(2, 3)
	test.AddEdge(1, 4)
	test.AddEdge(4, 3)
	test.AddEdge(4, 5)
	test.AddEdge(1, 6)
	test.AddEdge(6, 7)
	test.AddEdge(6, 8)
	test.Print()
	test.BreadthFirstSearch(1, 8)
}
