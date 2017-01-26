package atbash

import (
	"strings"
	"unicode"
)

func Atbash(s string) string {
	res := make([]rune, 0, len(s))
	counter := 0
	for _, c := range s {
		if !unicode.IsLetter(c) && !unicode.IsNumber(c) {
			continue
		}
		c = unicode.ToLower(c)
		if unicode.IsLetter(c) {
			c = rune('z' - (int(c) - int('a')))
		}
		res = append(res, c)
		counter++
		if counter == 5 {
			counter = 0
			res = append(res, ' ')
		}
	}
	return strings.TrimSpace(string(res))
}
