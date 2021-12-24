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

	startIndex, err := findTriggerIndex("##start", example)
	endIndex, err := findTriggerIndex("##end", example)
	if err != nil {
		return err
	}
	var from int
	if startIndex > endIndex {
		from = endIndex
	} else {
		from = startIndex
	}
	countOfVertexes, err := g.fromStartOrEnd(example, from, 0)
	if err != nil {
		return err
	}
	if countOfVertexes <= 1

	return nil
}

func (g *Graph) fromStartOrEnd(example []string, start int, countOfVertexes int) (int, error) {
	count := countOfVertexes
	for i := start + 2; i < len(example); i++ {
		// If comment then skip
		if example[i] == "##end" || example[i] == "##start" {
			var err error
			count, err = g.fromStartOrEnd(example, i, count)
			if err != nil {
				return 0, fmt.Errorf("ERROR: invalid data forma11 %v", i)
			}
			break
		}
		if example[i][0] == '#' {
			continue
		}
		edgeLines := strings.Split(example[i], "-")
		if len(edgeLines) == 2 || !strings.Contains(example[i], " ") {
			for j := i; j < len(example); j++ {
				errAddEdge := g.AddEdge(edgeLines[0], edgeLines[1])
				if errAddEdge != nil {
					return 0, fmt.Errorf("ERROR: invalid data format9")
				}
			}
			break
		}

		vertexLines := strings.Split(example[i], " ")
		if len(vertexLines) != 3 {
			return 0, fmt.Errorf("ERROR: invalid data forma11 %v", i)
		}
		g.AddVertex(vertexLines[0])
		count++
		for index := 1; index < 3; index++ {
			_, errLine := strconv.Atoi(vertexLines[index])
			if errLine != nil {
				return 0, fmt.Errorf("ERROR: invalid data format12")
			}
		}

	}
	return count, nil
}
