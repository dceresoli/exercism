package stringset

import (
	"strconv"
	"strings"
)

const testVersion = 4

type Set map[string]bool

func New() Set {
	return make(map[string]bool)
}

func NewFromSlice(slice []string) Set {
	set := New()
	for _, s := range slice {
		set[s] = true
	}
	return set
}

func (s Set) String() string {
	str := "{"
	var list []string
	for key := range s {
		list = append(list, strconv.Quote(key))
	}
	str += strings.Join(list, ", ")
	str += "}"
	return str
}

func (s Set) IsEmpty() bool {
	return len(s) == 0
}

func (s Set) Has(str string) bool {
	_, ok := s[str]
	return ok
}

func Subset(s1, s2 Set) bool {
	for key := range s1 {
		_, ok := s2[key]
		if !ok {
			return false
		}
	}
	return true
}

func Disjoint(s1, s2 Set) bool {
	for key := range s1 {
		_, ok := s2[key]
		if ok {
			return false
		}
	}
	return true
}

func Equal(s1, s2 Set) bool {
	return Subset(s1, s2) && Subset(s2, s1)
}

func (s Set) Add(str string) {
	s[str] = true
}

func Intersection(s1, s2 Set) Set {
	s := New()
	for key := range s1 {
		_, ok := s2[key]
		if ok {
			s.Add(key)
		}
	}
	return s
}

func Difference(s1, s2 Set) Set {
	s := New()
	for key := range s1 {
		_, ok := s2[key]
		if !ok {
			s.Add(key)
		}
	}
	return s
}

func Union(s1, s2 Set) Set {
	s := New()
	for key := range s1 {
		s.Add(key)
	}
	for key := range s2 {
		s.Add(key)
	}
	return s
}
