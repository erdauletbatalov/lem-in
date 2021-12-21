package funcs

import "fmt"

func (g *Graph) BreadthFirstSearch(from, to string) {
	//get Vertex
	fromVertex := g.GetVertex(from)
	toVertex := g.GetVertex(to)
	//check error
	if fromVertex == nil || toVertex == nil {
		err := fmt.Errorf("Invalid Vertexes (%v or %v)", from, to)
		fmt.Println(err.Error())
	} else {
		var visitedVertex map[*Vertex]bool
		visitedVertex = make(map[*Vertex]bool)
		current := fromVertex
		queue := current.Adjacent
		for {
			if current == toVertex {
				fmt.Println("I've found the path!")
				end := current
				path := []*Vertex{}
				for end != nil {
					path = append(path, end)
					end = end.Prev
				}
				path = reverseArray(path)
				for _, v := range path {
					if v == path[len(path)-1] {
						fmt.Printf("%v", v.Key)
						fmt.Println()
					} else {
						fmt.Printf("%v -> ", v.Key)
					}
				}
				return
			}
			for _, v := range current.Adjacent {
				if !visitedVertex[v] {
					queue = append(queue, v)
					v.Prev = current
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
