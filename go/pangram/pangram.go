package pangram

import (
	"unicode"
)

const testVersion = 1

func IsPangram(s string) bool {
	letters := make(map[rune]bool)
	for r := 'a'; r <= 'z'; r++ {
		letters[r] = false
	}

	for _, r := range s {
		if !unicode.IsLetter(r) {
			continue
		}
		r = unicode.ToLower(r)
		letters[r] = true
	}

	res := true
	for _, v := range letters {
		res = res && v
	}
	return res
}
