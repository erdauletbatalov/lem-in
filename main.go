package main

import (
	"fmt"
	"lem-in/funcs"
)

func main() {
	test := &funcs.Graph{}

	example, err := funcs.GetArrOfStrFromFile()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = test.GetDataFromArrOfStr(example)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	funcs.PrintExample(example)

	// test := &funcs.Graph{}

	// for i := 1; i < 9; i++ {
	// 	test.AddVertex(i)
	// }

	// test.AddEdge(1, 2)
	// test.AddEdge(2, 3)
	// test.AddEdge(1, 4)
	// test.AddEdge(4, 3)
	// test.AddEdge(4, 5)
	// test.AddEdge(1, 6)
	// test.AddEdge(6, 7)
	// test.AddEdge(6, 8)
	test.Print()
	// test.BreadthFirstSearch(1, 8)
}
