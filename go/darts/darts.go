package darts

import "math"

const (
	outer  = 10
	middle = 5
	inner  = 1
)

func outside(x, y float64) bool {
	if x > outer || y > outer {
		return true
	}

	return false
}

// Score calculates if a dart throw earns points and how many.
func Score(x, y float64) int {
	x = math.Abs(x)
	y = math.Abs(y)

	return 1
}
