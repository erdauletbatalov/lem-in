package funcs

import (
	"fmt"
	"strconv"
	"strings"
)

func (g *Graph) GetDataFromArrOfStr(example []string) error {
	if len(example) < 6 {
		return fmt.Errorf("ERROR: invalid data format")
	}

	ants, err := strconv.Atoi(example[0])
	if err != nil {
		return fmt.Errorf("ERROR: invalid data format")
	}
	if ants == 0 {
		return fmt.Errorf("ERROR: invalid data format")
	}
	g.AddAnts(ants)

	if example[1] != "##start" {
		return fmt.Errorf("ERROR: invalid data format")
	}

	indexStartOfEdges := 0

	for i := 2; i < len(example); i++ {
		if example[i] == "##end" {
			if i+2 >= len(example) {
				return fmt.Errorf("ERROR: invalid data format")
			}
			line := strings.Split(example[i+1], " ")
			if len(line) != 3 {
				return fmt.Errorf("ERROR: invalid data format")
			}
			g.AddVertex(line[0])
			g.End = line[0]
			indexStartOfEdges = i + 2
			for index := 1; index < 3; index++ {
				_, errLine := strconv.Atoi(line[index])
				if errLine != nil {
					return fmt.Errorf("ERROR: invalid data format")
				}
			}
			// add edges
			for j := indexStartOfEdges; j < len(example); j++ {
				edgeLines := strings.Split(example[j], "-")
				if len(edgeLines) != 2 {
					return fmt.Errorf("ERROR: invalid data format")
				}
				errAddEdge := g.AddEdge(edgeLines[0], edgeLines[1])
				if errAddEdge != nil {
					return fmt.Errorf("ERROR: invalid data format")
				}
			}
			break

		}
		line := strings.Split(example[i], " ")
		if len(line) != 3 {
			return fmt.Errorf("ERROR: invalid data format2 %v", i)
		}
		g.AddVertex(line[0])
		g.Start = line[0]
		for index := 1; index < 3; index++ {
			_, errLine := strconv.Atoi(line[index])
			if errLine != nil {
				return fmt.Errorf("ERROR: invalid data format")
			}
		}
	}
	return nil
}
