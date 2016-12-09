package etl

import (
	"strings"
)

func Transform(m map[int][]string) map[string]int {
	results := make(map[string]int)
	// for each key in m
	for k, v := range m {
		// for each value in key
		for _, letter := range v {
			// insert results[value]key
			results[strings.ToLower(letter)] = k
		}
	}
	return results
}
