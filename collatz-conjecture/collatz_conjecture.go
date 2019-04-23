package collatzconjecture

import "errors"

// CollatzConjecture returns the number of steps required to reach 1,
// given a number n.
func CollatzConjecture(input int) (int, error) {
	if input <= 0 {
		return -1, errors.New("invalid input")
	}
	res := 0
	for input != 1 {
		if input%2 == 0 {
			input /= 2
		} else {
			input *= 3
			input++
		}
		res++
	}
	return res, nil
}
