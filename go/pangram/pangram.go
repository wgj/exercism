package pangram

import (
	"strings"
)

const testVersion = 1

func IsPangram(s string) bool {
	// If empty string, return true.
	if s == "" {
		return false
	}
	s = strings.ToLower(s)
	// For every letter in the alphabet
	for i := 'a'; i < 'z'+1; i++ {
		// Check if s contains the letter, if false, return false.
		if !strings.Contains(s, string(i)) {
			return false
		}
	}

	return true
}
