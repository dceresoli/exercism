package bob // package name must match the package name in bob_test.go

import (
	"strings"
	"unicode"
)

const testVersion = 2 // same as targetTestVersion

func isQuestion(s string) bool {
	return len(s) > 0 && s[len(s)-1] == '?'
}

func hasLetter(s string) bool {
	for _, c := range s {
		if unicode.IsLetter(c) {
			return true
		}
	}
	return false
}

func isYelling(s string) bool {
	return hasLetter(s) && s == strings.ToUpper(s)
}

func Hey(t string) string {
	t = strings.TrimSpace(t)

        if len(t) == 0 {
                return "Fine. Be that way!"
	}
	if isYelling(t) {
		return "Whoa, chill out!"
	}
	if isQuestion(t) {
		return "Sure."
	}
	return "Whatever."
}
