package isogram

import (
	"sort"
	"strings"
)

// IsIsogram determines if a word or phrase is an isogram.
func IsIsogram(input string) (result bool) {
	// Remove non valid characters and transform to uppercase
	replacer := strings.NewReplacer(" ", "", "-", "")
	input = replacer.Replace(input)
	input = strings.ToUpper(input)

	// Sort letters
	letters := strings.Split(input, "")
	sort.Strings(letters)

	// Compare each letter with the following one
	for i := 0; i < len(letters)-1; i++ {
		if letters[i] == letters[i+1] {
			return false
		}
	}

	return true
}
