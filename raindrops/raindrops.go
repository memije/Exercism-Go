// Package raindrops converts a number to a string
package raindrops

import "strconv"

var raindrops = map[int]string{
	3: "Pling",
	5: "Plang",
	7: "Plong",
}

// Convert converts a number to a string. The content depends on the
// numbers factors.
//
// If the number has 3 as factor, output 'Pling'
// If the number has 5 as factor, output 'Plang'
// If the number has 7 as factor, output 'Plong'
//
// If the number does not have 3, 5, or 7 as a factor, just pass the number's
// digits straight through
func Convert(input int) (rain string) {
	rain = ""

	for i, sound := range raindrops {
		if input%i == 0 && i > 0 {
			rain += sound
		}
	}

	if rain == "" {
		rain = strconv.Itoa(input)
	}

	return rain
}
