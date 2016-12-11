package pythagorean

import "math"

// The three elements of each returned triplet must be in order,
// t[0] <= t[1] <= t[2], and the list of triplets must be in lexicographic
// order
type Triplet [3]int

// Range returns a list of all Pythagorean triplets with sides in the
// range min to max inclusive.
func Range(min, max int) []Triplet {
	var aSquare, aSquarePlusBSquare float64
	var triplets []Triplet
	// A Pythagorean Triple is:
	// a**2 + b**2 = c**2
	// For every a, while c is less than max
	for i := float64(min); i <= float64(max); i++ {
		// For every b, where b starts with a, while c is less than max
		aSquare = math.Pow(i, 2)
		for j := i; j <= float64(max); j++ {
			// Clear aSquarePlusBSquare
			// Add a**2 + b**2 to aSquarePlusBSquare
			aSquarePlusBSquare = aSquare + math.Pow(j, 2)
			// For every c, where c starts with b, while c is less than max
			for k := j; k <= float64(max); k++ {
				//Test if aSquarePlusBSquare == c**2 is true,
				if aSquarePlusBSquare == math.Pow(k, 2) {
					//convert triplet to int
					//create a Triplet
					// Add to triplets
					triplets = append(triplets, Triplet{int(i), int(j), int(k)})
				}
			}
		}
	}
	return triplets
}

// Sum returns a list of all Pythagorean triplets where the sum a+b+c
// (the perimeter) is equal to p.
func Sum(p int) []Triplet {
	var s []Triplet
	tripletsToSum := Range(1, p/2)
	for _, t := range tripletsToSum {
		if t[0]+t[1]+t[2] == p {
			s = append(s, t)
		}
	}
	return s
}
