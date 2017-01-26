package etl

//type given map[int][]string
//type expectation map[string]int
import "strings"

func Transform(g given) (e expectation) {
	e = make(expectation)
	for score, clist := range g {
		for _, c := range clist {
			e[strings.ToLower(c)] = score
		}
	}
	return
}
