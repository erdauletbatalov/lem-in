package funcs

import "fmt"

func findTriggerIndex(str string, example []string) (int, error) {
	for i, val := range example {
		if val == str {
			return i, nil
		}
	}
	return 0, fmt.Errorf("ERROR: invalid data format")
}
