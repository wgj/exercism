// Package bob boilerplate
package bob

import "strings"

const testVersion = 2 // same as targetTestVersion

// Hey is an angsty chatbot.
func Hey(s string) string {

	// Bob answers 'Sure.' if you ask him a question.
	if isQuestion(s) && !isYelling(s) {
		return "Sure."
	}
	// Bob answers 'Whoa, chill out!' if you yell at him.
	if isYelling(s) {
		return "Whoa, chill out!"
	}

	// Bob says 'Fine. Be that way!' if you address him without actually saying anything.
	if isSilent(s) {
		return "Fine. Be that way!"
	}

	// Bob answers 'Whatever.' to anything else.
	return "Whatever."
}

func isYelling(s string) bool {
	// If string is the same lower'd and upper'd, it doesn't have alpha characters.
	if s == strings.ToUpper(s) && !(strings.ToUpper(s) == strings.ToLower(s)) {
		return true
	}
	return false
}

func isSilent(s string) bool {
	s = strings.TrimSpace(s)
	return s == ""
}

func isQuestion(s string) bool {
	s = strings.TrimSpace(s)
	return strings.HasSuffix(s, "?")
}
