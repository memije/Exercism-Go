// Package accumulate allows to apply a function to each member of a collection
package accumulate

// Accumulate returns a new collection containing the result of applying
// an operation to each element of the input collection
func Accumulate(input []string, converter func(string) string) (result []string) {
	result = input[:0]
	for _, element := range input {
		result = append(result, converter(element))
	}
	return result
}
