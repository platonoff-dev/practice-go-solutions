package chess

import (
	"errors"
)

func intAbs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func CanKnightAttack(white, black string) (bool, error) {
	// Validate input
	if len(white) != 2 || len(black) != 2 {
		return false, errors.New("position should be 2 symbols length")
	}

	if white == black {
		return false, errors.New("two horses on the same cell")
	}

	xWhite := white[0]
	yWhite := white[1]
	xBlack := black[0]
	yBlack := black[1]

	if !('a' <= xWhite && xWhite <= 'h') || !('a' <= xBlack && xBlack <= 'h') {
		return false, errors.New("first char should be between a and h")
	}

	if !('1' <= yWhite && yWhite <= '8') || !('1' <= yBlack && yBlack <= '8') {
		return false, errors.New("second char should be between 1 and 8")
	}

	// Calculate distance between horses
	xDiff := intAbs(int(xBlack) - int(xWhite))
	yDiff := intAbs(int(yBlack) - int(yWhite))
	return xDiff == 2 && yDiff == 1 || xDiff == 1 && yDiff == 2, nil
}
