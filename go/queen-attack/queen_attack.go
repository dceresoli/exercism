package queenattack

import "errors"

const testVersion = 1

func ChessToXY(w string) (int, int, bool) {
	if len(w) != 2 {
		return 0, 0, false
	}
	col := int(w[0] - 'a')
	row := int(w[1] - '1')
	if row < 0 || row > 7 || col < 0 || col > 7 {
		return row, col, false
	}
	return row, col, true
}

func Abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func CanQueenAttack(w, b string) (bool, error) {
	var rw, cw, rb, cb int
	var ok bool

	if rw, cw, ok = ChessToXY(w); !ok {
		return false, errors.New("White queen invalid position")
	}
	if rb, cb, ok = ChessToXY(b); !ok {
		return false, errors.New("Black queen invalid position")
	}

	if rw == rb && cw == cb {
		return false, errors.New("Queens in same position")
	}

	// horizontal or vertical attack
	if rw == rb || cw == cb {
		return true, nil
	}

	// diagonal attack
	if Abs(rw-rb) == Abs(cw-cb) {
		return true, nil
	}

	// no attack
	return false, nil
}
