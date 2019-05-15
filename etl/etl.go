package etl

import "strings"

// Transform transforms scrabble scores from a legacy system.
func Transform(legacy map[int][]string) (transformed map[string]int) {
	transformed = make(map[string]int)
	for score, letters := range legacy {
		for _, letter := range letters {
			transformed[strings.ToLower(letter)] = score
		}
	}
	return transformed
}
