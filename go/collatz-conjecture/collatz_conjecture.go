package collatzconjecture

import "errors"

// CollatzConjecture Given a number n, return the number of steps required to reach 1.
// Following the CollatzConjecture rules.
func CollatzConjecture(number int) (steps int, err error) {
	if number <= 0 {
		return 0, errors.New("Error")
	}

	for number != 1 {
		steps++

		if number%2 == 0 {
			number = number / 2
		} else {
			number = number*3 + 1
		}
	}

	return
}
