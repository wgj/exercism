// Package diffsquares ...
package diffsquares

import "math"

// SquareOfSums ...
func SquareOfSums(n int) (sq int) {
	var sum int
	for i := 1; i <= n; i++ {
		sum += i
	}
	sq = int(math.Pow(float64(sum), 2.0))
	return
}

// SumOfSquares ...
func SumOfSquares(n int) (sum int) {
	for i := 1; i <= n; i++ {
		sum += int(math.Pow(float64(i), 2.0))
	}
	return

}

// Difference ...
func Difference(n int) (d int) {
	return SquareOfSums(n) - SumOfSquares(n)
}
