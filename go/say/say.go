package say

import (
	"strings"
)

func Say(n uint64) string {
	if n == 0 {
		return "zero"
	}
	res := make([]string, 0)

	var powers = []uint64{1e18, 1e15, 1e12, 1e9, 1e6, 1e3}
	var suffix = []string{"quintillion", "quadrillion", "trillion", "billion", "million", "thousand"}

	for i, pow := range powers {
		if n >= pow {
			res = append(res, say999(int(n/pow)))
			res = append(res, suffix[i])
			n %= pow
		}
	}
	if n != 0 {
		res = append(res, say999(int(n)))
	}

	return strings.Join(res, " ")
}

// this handle the case 1-999
func say999(n int) string {
	var units = []string{
		"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten",
		"eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen",
	}
	var tens = []string{
		"zero", "ten", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety",
	}

	res := make([]string, 0)

	h := n / 100
	t := (n / 10) % 10
	u := n % 10

	if h > 0 {
		res = append(res, units[h]+" hundred")
	}
	if t >= 2 {
		if u != 0 {
			res = append(res, tens[t]+"-"+units[u])
		} else {
			res = append(res, tens[t])
		}
	} else {
		m := t*10 + u
		if m != 0 {
			res = append(res, units[m])
		}
	}
	return strings.Join(res, " ")
}
