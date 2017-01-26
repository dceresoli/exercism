package cryptosquare

import (
	"math"
	"unicode"
	//"fmt"
)

const testVersion = 2

// strip spaces, punctuation, and lower
func Normalize(s string) string {
	res := make([]rune, 0)
	for _, c := range s {
		if unicode.IsLetter(c) || unicode.IsDigit(c) {
			res = append(res, unicode.ToLower(c))
		}
	}
	return string(res)
}

// perform square encoding
func Encode(s string) string {
	s = Normalize(s)
	l := len(s)

	//fmt.Printf("Normalized:%s\n", s)

	// number of columns and row
	rows := int(math.Floor(math.Sqrt(float64(l))))
	cols := rows
	if rows*cols < l {
		cols++
	}
	if rows*cols < l {
		rows++
	}
	//fmt.Printf("len=%v cols=%v rows=%v\n", l, cols, rows)

	// transform into array
	runes := []rune(s)
	res := make([]rune, 0)

	// output transposed
	for c := 0; c < cols; c++ {
		for r := 0; r < rows; r++ {
			i := c + r*cols
			if i >= l {
				break
			}
			res = append(res, runes[i])
		}
		if c < cols-1 {
			res = append(res, ' ')
		}
	}

	return string(res)
}
