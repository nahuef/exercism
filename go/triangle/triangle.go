// Package triangle has helper functions of the geometrical figure.
package triangle

import "math"

// Kind of triangle.
type Kind int

const (
	// NaT not a triangle
	NaT = iota
	// Equ equilateral
	Equ
	// Iso isosceles
	Iso
	// Sca scalene
	Sca
)

func hasThreeSides(a, b, c float64) bool {
	if a > 0 && !math.IsInf(a, 0) &&
		b > 0 && !math.IsInf(b, 0) &&
		c > 0 && !math.IsInf(c, 0) {
		return true
	}

	return false
}

func hasInequality(a, b, c float64) bool {
	if (a+b) >= c && (a+c) >= b && (b+c) >= a {
		return false
	}

	return true
}

func isTriangle(a, b, c float64) bool {
	if hasThreeSides(a, b, c) && !hasInequality(a, b, c) {
		return true
	}

	return false
}

func isEquilateral(a, b, c float64) bool {
	side := (a + b + c) / 3

	if side == a && side == b && side == c {
		return true
	}

	return false
}

func isIsosceles(a, b, c float64) bool {
	if a == b || a == c || b == c {
		return true
	}

	return false
}

func isScalene(a, b, c float64) bool {
	if !(a == b || a == c || b == c) {
		return true
	}

	return false
}

// KindFromSides returns the type of triangle given its sides lengths.
func KindFromSides(a, b, c float64) Kind {
	if !isTriangle(a, b, c) {
		return NaT
	}

	if isEquilateral(a, b, c) {
		return Equ
	}

	if isIsosceles(a, b, c) {
		return Iso
	}

	if isScalene(a, b, c) {
		return Sca
	}

	return NaT
}
