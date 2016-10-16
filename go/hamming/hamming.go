package hamming

import "errors"

const testVersion = 4

// Distance returns Hamming Distance of a and b, and an error if len(a) is not equal to len(b).
func Distance(a, b string) (int, error) {
	d := 0
	if len(a) != len(b) {
		return d, errors.New("sequences must be equal length")
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			d++
		}
	}

	return d, nil
}
