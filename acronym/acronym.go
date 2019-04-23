package acronym

import (
	"regexp"
	"strings"
)

var rExp = regexp.MustCompile(`\b\w`)

// Abbreviate takes a phrase and converts it to its acronym.
func Abbreviate(s string) string {
	s = strings.Replace(s, "'", "", -1)
	letters := rExp.FindAllString(s, -1)
	return strings.ToUpper(strings.Join(letters, ""))
}
