package wordcount

import (
	"regexp"
	"strings"
)

const testVersion = 2

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	freq := make(Frequency)
	re := regexp.MustCompile(`[a-zA-Z0-9_]+`)
	words := re.FindAllString(phrase, -1)
	for _, word := range words {
		word = strings.ToLower(word)
		freq[word] += 1
	}
	return freq
}
