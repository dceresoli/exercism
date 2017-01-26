package triangle

import "math"

const testVersion = 3

type Kind int

const (
	NaT Kind = iota // not a triangle
	Equ             // equilateral
	Iso             // isosceles
	Sca             // scalene
)

func isPositive(a float64) bool {
	if math.IsInf(a, 0) || math.IsNaN(a) {
		return false
	}
	return a > 0.0
}

func KindFromSides(a, b, c float64) Kind {
	// sort in ascending order
	if a > b {
		a, b = b, a
	}
	if a > c {
		a, c = c, a
	}
	if b > c {
		b, c = c, b
	}

	// impossible triangle
	if !isPositive(a) || !isPositive(b) || !isPositive(c) {
		return NaT
	}

	// impossible triangle
	if a+b < c {
		return NaT
	}

	// possible triangle
	if a == b && b == c {
		return Equ
	}

	if a == b || b == c {
		return Iso
	}

	return Sca
}
