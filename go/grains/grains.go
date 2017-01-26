package grains

import "errors"

const testVersion = 1

func Square(n int) (uint64, error) {
	if n < 1 || n > 64 {
		return 0, errors.New("n must be between 1 and 64")
	}
	return 1 << uint(n-1), nil
}

func Total() (tot uint64) {
	tot = 0
	for n := 1; n <= 64; n++ {
		grains, _ := Square(n)
		tot += grains
	}
	return tot
}
