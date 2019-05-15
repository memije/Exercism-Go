package bob

import (
	"strings"
)

func quesiton(s string) bool {
	return s[len(s)-1:] == "?"
}

func shouting(s string) bool {
	return s == strings.ToUpper(s) && s != strings.ToLower(s)
}

func silence(s string) bool {
	return len(s) == 0
}

// Hey gets an answer from Bob.
func Hey(remark string) (answer string) {
	remark = strings.TrimSpace(remark)

	if silence(remark) {
		return "Fine. Be that way!"
	} else if shouting(remark) && quesiton(remark) {
		return "Calm down, I know what I'm doing!"
	} else if shouting(remark) {
		return "Whoa, chill out!"
	} else if quesiton(remark) {
		return "Sure."
	}
	return "Whatever."
}
