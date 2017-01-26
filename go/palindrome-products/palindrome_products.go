package palindrome

import (
	"errors"
	"math"
	"strconv"
)

const testVersion = 1

type Product struct {
	Product        int      // palindromic, of course
	Factorizations [][2]int // list of all possible two-factor factorizations of Product, within given limits, in order
}

func NewProduct(prod int) Product {
	return Product{prod, make([][2]int, 0)}
}

func (p *Product) Clear() {
	p.Factorizations = make([][2]int, 0)
}

func (p *Product) AddPair(i, j int) {
	p.Product = i * j
	p.Factorizations = append(p.Factorizations, [2]int{i, j})
}

func (p *Product) Len() int {
	return len(p.Factorizations)
}

// transform to string and check if palindrome
func isPalindrome(n int) bool {
	s := strconv.Itoa(n)
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

func Products(fmin, fmax int) (pmin, pmax Product, err error) {
	pmin = NewProduct(math.MaxInt64)
	pmax = NewProduct(math.MinInt64)

	if fmin > fmax {
		return pmin, pmax, errors.New("fmin > fmax")
	}

	for i := fmin; i <= fmax; i++ {
		for j := i; j <= fmax; j++ {
			prod := i * j
			if isPalindrome(prod) {
				if prod <= pmin.Product {
					if prod < pmin.Product {
						pmin.Clear()
					}
					pmin.AddPair(i, j)
				}
				if prod >= pmax.Product {
					if prod > pmax.Product {
						pmax.Clear()
					}
					pmax.AddPair(i, j)
				}
			}
		}
	}
	if pmin.Len() == 0 && pmax.Len() == 0 {
		return pmin, pmax, errors.New("No palindromes...")
	}
	return pmin, pmax, nil
}
