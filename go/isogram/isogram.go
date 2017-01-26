package isogram

import "unicode"

const testVersion = 1

func IsIsogram(s string) bool {
	chars := []rune(s)
	for i, c := range chars {
		if !unicode.IsLetter(c) {
			continue
		}
                c = unicode.ToLower(c)
		for j := i + 1; j < len(chars); j++ {
			if unicode.ToLower(chars[j]) == c {
				return false
			}
		}
	}
	return true
}
