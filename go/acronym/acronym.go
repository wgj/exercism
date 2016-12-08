package acronym

import (
	"strings"
	"unicode"
)

const testVersion = 1

func abbreviate(s string) string {
	if len(s) == 0 {
		return ""
	}

	initials := make([]string, 0)

	// Split string into separate 'words'.
	words := strings.FieldsFunc(s, myFields)

	for _, w := range words {
		// Take first character from each word, add to initials.
		initials = append(initials, strings.ToUpper(w[0:1]))
		// Skip words that are already initialisms.
		if w == strings.ToUpper(w) {
			continue
		}
		// Look at the rest of the word for CamelCase compound words.
		for _, c := range w[1:] {
			if unicode.IsUpper(c) {
				initials = append(initials, string(c))
			}
		}
	}

	return strings.Join(initials, "")
}

func myFields(c rune) bool {
	return unicode.IsSpace(c) || c == '-'
}
