package lsproduct

import (
	"errors"
	"unicode"
)

const testVersion = 4

func LargestSeriesProduct(digits string, span int) (int, error) {
	if span < 0 {
		return 0, errors.New("negative span")
	}

	// convert to bytes
	bytes := make([]byte, 0)
	for _, c := range digits {
		if unicode.IsDigit(c) {
			bytes = append(bytes, byte(c)-'0')
		} else {
			return 0, errors.New("non digit in string")
		}
	}

	// sanity checks
	if len(bytes) == 0 && span == 0 {
		return 1, nil
	}
	if len(bytes) == 0 || len(bytes) < span {
		return 0, errors.New("not enough digits")
	}

	// compute max
	max := 0
	for i := 0; i <= len(bytes)-span; i++ {
		prod := 1
		for j := 0; j < span; j++ {
			prod *= int(bytes[j+i])
		}
		if prod > max {
			max = prod
		}
	}
	return max, nil
}
