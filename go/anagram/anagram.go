package anagram

import (
	"reflect"
	"strings"
)

func countLetters(s string) map[rune]int {
	res := make(map[rune]int)
	for _, c := range s {
		res[c] += 1
	}
	return res
}

func isAnagram(a, b string) bool {
	if len(a) != len(b) || a == b {
		return false
	}
	return reflect.DeepEqual(countLetters(a), countLetters(b))
}

func Detect(subject string, candidates []string) []string {
	res := make([]string, 0)
	subject = strings.ToLower(subject)
	for _, candidate := range candidates {
		candidate = strings.ToLower(candidate)
		if isAnagram(subject, candidate) {
			res = append(res, candidate)
		}
	}
	return res
}
