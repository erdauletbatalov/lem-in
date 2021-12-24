package funcs

import (
	"fmt"
	"strconv"
	"strings"
)

func (g *Graph) findStartIndex(example []string) (int, error) {
	for i, val := range example {
		if val == "##start" && isThereVertexInNextLine(example, "##start", i) {
			g.Start = example[i+1]
			return i, nil
		}
	}
	return 0, fmt.Errorf("ERROR: invalid data format")
}

func (g *Graph) findEndIndex(example []string) (int, error) {
	for i, val := range example {
		if val == "##end" && isThereVertexInNextLine(example, "##end", i) {
			g.End = example[i+1]
			return i, nil
		}
	}
	return 0, fmt.Errorf("ERROR: invalid data format")
}

func isThereVertexInNextLine(example []string, str string, strIndex int) bool {
	if len(example) <= strIndex+1 {
		return false
	}
	vertexLines := strings.Split(example[strIndex+1], " ")
	if len(vertexLines) != 3 {
		return false
	}
	for index := 1; index < 3; index++ {
		_, errLine := strconv.Atoi(vertexLines[index])
		if errLine != nil {
			return false
		}
	}
	return true
}
