package scrabble

import (
	"strings"
)

const testVersion = 4

// 'a' is an int32.
var scorem map[int32]int

func init() {
	scorem = map[int32]int{
		'a': 1,
		'b': 3,
		'c': 3,
		'd': 2,
		'e': 1,
		'f': 4,
		'g': 2,
		'h': 4,
		'i': 1,
		'j': 8,
		'k': 5,
		'm': 3,
		'n': 1,
		'o': 1,
		'p': 3,
		'l': 1,
		'q': 10,
		'r': 1,
		's': 1,
		't': 1,
		'u': 1,
		'v': 4,
		'w': 4,
		'x': 8,
		'y': 4,
		'z': 10,
	}
}

func Score(s string) int {
	var score int
	s = strings.ToLower(s)
	for _, v := range s {
		score += scorem[v]
	}
	return score
}
