package romannumerals

import (
	"errors"
)

var romans = []struct {
	roman string
	value int
}{
	{"I", 1},
	{"IV", 4},
	{"V", 5},
	{"IX", 9},
	{"X", 10},
	{"XL", 40},
	{"L", 50},
	{"XC", 90},
	{"C", 100},
	{"CD", 400},
	{"D", 500},
	{"CM", 900},
	{"M", 1000},
}

// ToRomanNumeral converts an arabic number to its roman representation.
func ToRomanNumeral(input int) (res string, err error) {
	if input <= 0 || input > 3000 {
		return "", errors.New("invalid input number")
	}

	for romanIndex, arabic := len(romans)-1, input; romanIndex >= 0; romanIndex-- {
		r := romans[romanIndex]
		for arabic >= r.value {
			arabic -= r.value
			res += r.roman
		}
	}

	return res, nil
}
