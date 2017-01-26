package pythagorean

type Triplet [3]int

func Range(min, max int) []Triplet {
	res := make([]Triplet, 0)
	for a := min; a <= max; a++ {
		for b := a; b <= max; b++ {
			for c := b; c <= max; c++ {
				if a*a+b*b == c*c {
					res = append(res, Triplet{a, b, c})
				}
			}
		}
	}
	return res
}

func Sum(p int) []Triplet {
	res := make([]Triplet, 0)
	triplets := Range(1, p)
	for _, t := range triplets {
		if t[0]+t[1]+t[2] == p {
			res = append(res, t)
		}
	}
	return res
}
