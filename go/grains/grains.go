package grains

import (
	"errors"
	"fmt"
	"log"
	"math"
)

const testVersion = 1

func Square(i int) (uint64, error) {
	// Test for too many squares.
	if i > 64 {
		err := fmt.Sprintf("there are only 64 squares on a chessboard, input was %d", i)
		return 0, errors.New(err)
	}
	// Test for not enough squares.
	if i < 1 {
		err := fmt.Sprintf("not squares occupied, input was %d", i)
		return 0, errors.New(err)
	}
	// return 2**i-1
	f := float64(i)
	u := uint64(math.Pow(2, f-1))
	return u, nil
}

func Total() (grains uint64) {
	for i := 1; i < 65; i++ {
		u, err := Square(i)
		if err != nil {
			log.Println("How did you get here?")
		}
		grains += u
	}
	return grains
}
