package acronym

import (
	"strings"
	"unicode"
)

const testVersion = 1

func abbreviate(s string) string {
	var acronym string
	s = strings.Trim(s, " \t")
	runes := []rune(s)
	for i, r := range runes {
		if i == 0 {
			acronym += string(r)
			continue
		}
		rp := runes[i-1]
		if (unicode.IsUpper(r) && unicode.IsLower(rp)) ||
			(unicode.IsLetter(r) && (unicode.IsPunct(rp) || unicode.IsSpace(rp))) {
			acronym += string(r)
			continue
		}
	}
	return strings.ToUpper(acronym)
}
