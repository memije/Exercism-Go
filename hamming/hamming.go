// Package hamming calculates the Hamming difference between two DNA strands
package hamming

import "errors"

// Distance counts the number of differences between two homologous DNA strands
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("non equal length sequences")
	}

	distance := 0
	runeA, runeB := []rune(a), []rune(b)
	for i := range runeA {
		if runeA[i] != runeB[i] {
			distance++
		}
	}
	return distance, nil
}
