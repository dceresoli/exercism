package perfect

import (
	"errors"
)

const testVersion = 1

type Classification int

const (
	ClassificationDeficient Classification = iota
	ClassificationAbundant
	ClassificationPerfect
)

var ErrOnlyPositive = errors.New("Only positive")

func Factorize(n uint64) []int {
	factors := make([]int, 1)
	for i := 1; uint64(i) < n; i++ {
		if n%uint64(i) == 0 {
			factors = append(factors, i)
		}
	}
	return factors
}

func Classify(n uint64) (Classification, error) {
	if n <= 0 {
		return -1, ErrOnlyPositive
	}
	factors := Factorize(n)

	sum := 0
	for _, i := range factors {
		sum += i
	}
	if uint64(sum) == n {
		return ClassificationPerfect, nil
	}
	if uint64(sum) > n {
		return ClassificationAbundant, nil
	}
	return ClassificationDeficient, nil
}
