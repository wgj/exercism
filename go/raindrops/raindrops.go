// Package raindrops ...
package raindrops

const testVersion = 2

/*
If the number contains 3 as a factor, output 'Pling'.
If the number contains 5 as a factor, output 'Plang'.
If the number contains 7 as a factor, output 'Plong'.
If the number does not contain 3, 5, or 7 as a factor, just pass the number's digits straight through.
*/

func Convert(i int) string {
	var s string

	switch {
	case i%3 == 0:
		s += "Pling"
	case i%5 == 0:
		s += "Plang"
	case i%7 == 0:
		s += "Plong"
	default:
		s = string(i)
	}
	return s
}
