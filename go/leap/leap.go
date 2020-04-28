/*
Package leap checks if a certain year is a leap year or not.
Occurs:
	on every year that is evenly divisible by 4
		except every year that is evenly divisible by 100
			unless the year is also evenly divisible by 400
*/
package leap

func isDivisibleBy(dividend int, divisor int) bool {
	return (dividend % divisor) == 0
}

// IsLeapYear checks if a certain year is a leap year or not.
func IsLeapYear(year int) bool {
	var byFour = isDivisibleBy(year, 4)
	var byOneHundred = isDivisibleBy(year, 100)
	var byFourHundred = isDivisibleBy(year, 400)

	if byFour && !byOneHundred ||
		byFourHundred {
		return true
	}

	return false
}

/*
Thought it would be fun to implement a high order function in Go,
so here is another option.
*/
func isDivisible(divisor int) func(int) bool {
	return func(dividend int) bool {
		return (dividend % divisor) == 0
	}
}

var isDivisibleByFour = isDivisible(4)
var isDivisibleByOneHundred = isDivisible(100)
var isDivisibleByFourHundred = isDivisible(400)

// IsLeapYearHighOrder checks if a certain year is a leap year or not,
// using high order functions.
func IsLeapYearHighOrder(year int) bool {
	if isDivisibleByFour(year) && !isDivisibleByOneHundred(year) ||
		isDivisibleByFourHundred(year) {
		return true
	}

	return false
}
