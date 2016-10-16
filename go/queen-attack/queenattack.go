package queenattack

import "fmt"

const testVersion = 1

func CanQueenAttack(w, b string) (bool, error) {
	/* Yes, I realize this is upside down.
	 * [a1][a2][a3][a4][a5][a6][a7][a8]
	 * [b1][b2][b3][b4][b5][b6][b7][b8]
	 * [c1][c2][c3][c4][c5][c6][c7][c8]
	 * [d1][d2][d3][d4][d5][d6][d7][d8]
	 * [e1][e2][e3][e4][e5][e6][e7][e8]
	 * [f1][f2][f3][f4][f5][f6][f7][f8]
	 * [g1][g2][g3][g4][g5][g6][g7][g8]
	 * [h1][h2][h3][h4][h5][h6][h7][h8]
	 */

	if w == b {
		return false, fmt.Errorf("white and black occupy the same square, white: %s; black %s", w, b)
	}
	if string(w[0]) > "h" || string(w[1]) > "8" || string(b[0]) > "h" || string(b[1]) > "8" {
		return false, fmt.Errorf("invalid position")
	}

	if sameRank(w, b) {
		return true, nil
	}

	if sameFile(w, b) {
		return true, nil
	}

	if sameDiagnal(w, b) {
		return true, nil
	}

	// Catch all.
	return false, nil
}

func sameRank(w, b string) bool {
	return string(w[0]) == string(b[0])
}

func sameFile(w, b string) bool {
	return string(w[1]) == string(b[1])
}

func sameDiagnal(w, b string) bool {
	var s []string
	s = append(s, "leftdown", "rightup", "leftup", "rightdown")
	for _, dir := range s {
		if diagnal(w, b, dir) {
			return true
		}
	}
	return false
}

func diagnal(w, b, dir string) bool {
	var x, y int
	switch dir {
	case "rightdown":
		x, y = 1, 1
	case "leftdown":
		x, y = 1, -1
	case "rightup":
		x, y = -1, 1
	case "leftup":
		x, y = -1, -1
	}
	for {
		if w == b {
			return true
		}
		// Check adjacent space.
		w = fmt.Sprintf("%s%s", string(int(w[0])+x), string(int(w[1])+y))
		// Check if we are off the board.
		if w[0] < byte(97) || w[0] > byte(104) || w[1] < byte(49) || w[1] > byte(56) {
			break
		}
	}
	return false
}
