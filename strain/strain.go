package strain

// Ints is a int slice ([]int)
type Ints []int

// Lists is a bidimentional int slice ([][]int)
type Lists [][]int

// Strings is a string slice ([]string)
type Strings []string

// Keep returns the elements that match certain criteria.
func (i Ints) Keep(fn func(int) bool) (result Ints) {
	for _, val := range i {
		if fn(val) {
			result = append(result, val)
		}
	}
	return result
}

// Discard returns the elements that don't match certain criteria.
func (i Ints) Discard(fn func(int) bool) (result Ints) {
	for _, val := range i {
		if !fn(val) {
			result = append(result, val)
		}
	}
	return result
}

// Keep returns the elements that match certain criteria.
func (l Lists) Keep(fn func([]int) bool) (result Lists) {
	for _, list := range l {
		if fn(list) {
			result = append(result, list)
		}
	}
	return result
}

// Keep returns the elements that match certain criteria.
func (s Strings) Keep(fn func(string) bool) (result Strings) {
	for _, str := range s {
		if fn(str) {
			result = append(result, str)
		}
	}
	return result
}
