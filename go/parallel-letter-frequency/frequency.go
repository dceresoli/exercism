package letter

import "sync"

type FreqMap map[rune]int

func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func ConcurrentFrequency(s []string) FreqMap {
	l := len(s)
	res := make([]FreqMap, l)
	var wg sync.WaitGroup

	for i, str := range s {
		wg.Add(1)
		go func(str string, i int) {
			defer wg.Done()
			res[i] = Frequency(str)
		}(str, i)
	}
	wg.Wait()

	m := make(FreqMap)
	for i := 0; i < l; i++ {
		for key, val := range res[i] {
			m[key] += val
		}
	}
	return m
}
