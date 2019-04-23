package space

// Planet is a user-defined type representing the planet name
type Planet string

// planetMap is a map that holds the ratio of orbital period
// for all the 8 planets
var planetMap = map[Planet]float64{
	"Earth":   1,
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

// Age calculates how old someone would be on <planet>
// given an age in seconds
func Age(seconds float64, planet Planet) float64 {
	orbitalPeriod := planetMap[planet] * 31557600

	if orbitalPeriod == 0 {
		return 0
	}

	return seconds / orbitalPeriod
}
