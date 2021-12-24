package main

import (
	"fmt"
	. "lem-in/funcs"
)

func main() {
	test := &Graph{}

	example, err := GetArrOfStrFromFile()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = test.GetDataFromArrOfStr(example)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	PrintExample(example)

	test.PrintVertexesAndTheirEdges()

	test.PrintGraphInfo()

	test.BreadthFirstSearch(test.Start, test.End)
}
