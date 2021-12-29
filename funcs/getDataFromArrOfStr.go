package funcs

import (
	"fmt"
	"strconv"
	"strings"
)

func (g *Graph) GetDataFromArrOfStr(example []string) error {
	// Check for min arg count
	if len(example) < 6 {
		return fmt.Errorf("ERROR: invalid data format1")
	}
	// Check if first arg is num and > 0
	ants, err := strconv.Atoi(example[0])
	if err != nil {
		return fmt.Errorf("ERROR: invalid data format2")
	}
	if ants <= 0 {
		return fmt.Errorf("ERROR: invalid data format3")
	}
	g.AddAnts(ants)

	startIndex, err := g.findStartIndex(example)
	if err != nil {
		return err
	}
	endIndex, err := g.findEndIndex(example)
	if err != nil {
		return err
	}

	var from int
	if startIndex > endIndex {
		from = endIndex
	} else {
		from = startIndex
	}
	err = g.fromStartOrEnd(example, from)
	if err != nil {
		return err
	}

	return nil
}

func (g *Graph) fromStartOrEnd(example []string, from int) error {
	for i := from + 2; i < len(example); i++ {
		// If comment then skip
		if example[i] == "##end" || example[i] == "##start" {
			var err error
			err = g.fromStartOrEnd(example, i)
			if err != nil {
				fmt.Println(1)
				return fmt.Errorf("ERROR: invalid data formatFromSOE %v", i)
			}
			break
		}
		if example[i][0] == '#' {
			continue
		}
		edgeLines := strings.Split(example[i], "-")
		if len(edgeLines) == 2 || !strings.Contains(example[i], " ") {
			for j := i; j < len(example); j++ {
				edgeLines := strings.Split(example[j], "-")
				if len(edgeLines) == 2 || !strings.Contains(example[i], " ") {
					errAddEdge := g.AddEdge(edgeLines[0], edgeLines[1])
					if errAddEdge != nil {
						fmt.Println(j)
						return fmt.Errorf("ERROR: invalid data format9")
					}
				} else {
					return fmt.Errorf("ERROR: invalid data formatedgesError %v", j)
				}
			}
			break
		}

		vertexLines := strings.Split(example[i], " ")
		if len(vertexLines) != 3 {
			fmt.Println(3)
			return fmt.Errorf("ERROR: invalid data format11 %v", i)
		}
		g.AddVertex(vertexLines[0])
		for index := 1; index < 3; index++ {
			_, errLine := strconv.Atoi(vertexLines[index])
			if errLine != nil {
				fmt.Println(4)
				return fmt.Errorf("ERROR: invalid data format12")
			}
		}

	}
	return nil
}
