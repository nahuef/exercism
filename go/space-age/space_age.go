// Package space contains an age calculator based on planet.
package space

// Planet is planet name string.
type Planet string

const secondsPerEarthYear float64 = 31557600

var orbitalPeriodsInEarthYears = map[Planet]float64{
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Earth":   1.0,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

func calculateSecondsPerYear(planetName Planet) float64 {
	return (orbitalPeriodsInEarthYears[planetName] * secondsPerEarthYear)
}

func calculateAge(secondsOld float64, secondsPerYear float64) float64 {
	return (secondsOld / secondsPerYear)
}

// Age calculates the age of a person on a certain planet, given an amount of seconds.
func Age(secondsOld float64, planet Planet) float64 {
	var secondsPerYear = calculateSecondsPerYear(planet)
	var age = calculateAge(secondsOld, secondsPerYear)

	return age
}
