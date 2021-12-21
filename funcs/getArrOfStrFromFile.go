package funcs

import (
	"bufio"
	"fmt"
	"os"
)

func GetArrOfStrFromFile() ([]string, error) {
	err := fmt.Errorf("ERROR: invalid data format")
	args := os.Args[1:]
	if len(args) != 1 {
		return []string{}, err
	}
	file, ferr := os.Open(args[0])
	if ferr != nil {
		return []string{}, err
	}

	scanner := bufio.NewScanner(file)
	var example []string

	for scanner.Scan() {
		line := scanner.Text()
		example = append(example, line)
	}
	return example, nil
}
