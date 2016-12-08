package pangram

import (
	"strings"
)

const testVersion = 1

func IsPangram(s string) bool {
	// Create an array of int, size 26
	var alphabet [26]int
	s = strings.ToLower(s)
	// For every char in s
	for _, v := range s[:] {
		if v >= 'a' && v <= 'z' {
			// 'a' starts at 97 in ASCII
			alphabet[v-97]++
		}
	}

	for _, v := range alphabet {
		if v == 0 {
			// If alphabet[i] is 0, then it wasn't in s.
			return false
		}
	}
	return true
}
