package slice

import "errors"

// All returns a list of all substrings of s with length n.
func All(n int, s string) []string {
	// iterate over string. increment by n.
	r := make([]string, 0)
	for i := 0; i <= len(s)-n; i++ {
		//		fmt.Println(s[i : i+n])
		r = append(r, s[i:i+n])
	}
	return r
}

// UnsafeFirst returns the first substring of s with length n.
func UnsafeFirst(n int, s string) string {
	if n <= len(s) {
		return s[0:n]
	}
	return ""
}

// First returns the first substring, and error if it's not possible to return the first substring.
func First(n int, s string) (string, error) {
	if n > len(s) {
		return "", errors.New("It's not possible to return the first substring.")
	}
	return s[0:n], nil
}
