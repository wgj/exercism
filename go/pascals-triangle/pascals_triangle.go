package pascal

// Triangle ...
func Triangle(n int) [][]int {
	/*
		[ [1] ]
		[ [1], [1,1] ]
		[ [1], [1,1], [1,2,1] ]
	*/
	s := makeSlice(n)
	// Iterate through outer slice.
	for i := 0; i < n; i++ {
		// Iterate through inner slice.
		for j := 0; j < len(s[i]); j++ {
			// If position is at the beginning of a row, or the end of a row, the value is 1.
			if j == 0 || j+1 == len(s[i]) {
				s[i][j] = 1
				continue
			}
			// Find the 'parents' of s[i][j], and add them together for the current value.
			left := s[i-1][j-1]
			right := s[i-1][j]
			s[i][j] = left + right
		}
	}
	return s
}

func makeSlice(n int) [][]int {
	s := make([][]int, n)
	for i := 0; i < n; i++ {
		s[i] = make([]int, i+1)
	}
	return s
}
