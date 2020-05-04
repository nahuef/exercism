// Package raindrops converts numbers to string acording to factor numbers.
package raindrops

import "strconv"

// Convert converts a number into a string using corresponding raidrop sounds.
func Convert(number int) (result string) {
	if number%3 == 0 {
		result += "Pling"
	}

	if number%5 == 0 {
		result += "Plang"
	}

	if number%7 == 0 {
		result += "Plong"
	}

	if len(result) == 0 {
		result = strconv.Itoa(number)
	}

	return
}
