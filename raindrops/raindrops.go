// Package raindrops converts a number to a string
package raindrops

import (
	"strconv"
	"strings"
)

// Convert converts a number to a string. The content depends on the
// numbers factors.
//
// If the number has 3 as factor, output 'Pling'
// If the number has 5 as factor, output 'Plang'
// If the number has 7 as factor, output 'Plong'
//
// If the number does not have 3, 5, or 7 as a factor, just pass the number's
// digits straight through
func Convert(input int) (result string) {
	result = ""
	for _, factor := range getFactors(input) {
		switch factor {
		case 3:
			addRaindrop(&result, "Pling")
		case 5:
			addRaindrop(&result, "Plang")
		case 7:
			addRaindrop(&result, "Plong")
		}
	}

	if result == "" {
		result += strconv.Itoa(input)
	}

	return result
}

// getFactors returns an array with the factors of a number
func getFactors(input int) (factors []int) {
	factor := 2
	factors = []int{}

	for input >= factor*factor {
		if input%factor == 0 {
			factors = append(factors, factor)
			input /= factor
			continue
		}
		factor++
	}

	factors = append(factors, input)
	return factors
}

// addRaindrop adds a raindrop to the rain string if it does not exist already
func addRaindrop(rain *string, raindrop string) {
	if !strings.Contains(*rain, raindrop) {
		*rain += raindrop
	}
}
