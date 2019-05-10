// Package listops provides handy lists operations
package listops

// IntList represents a list of integers
type IntList []int

type binFunc func(x, y int) int
type predFunc func(n int) bool
type unaryFunc func(x int) int

// Foldr performs a right folding of the list
func (i IntList) Foldr(fn binFunc, n int) int {
	if i.Length() == 0 {
		return n
	}
	return i[0:i.Length()-1].Foldr(fn, fn(i[i.Length()-1], n))
}

// Foldl performs a left folding of the list
func (i IntList) Foldl(fn binFunc, n int) int {
	if i.Length() == 0 {
		return n
	}
	return i[1:i.Length()].Foldl(fn, fn(n, i[0]))
}

// Filter returns the elements that follow a certain criteria
func (i IntList) Filter(fn predFunc) (result IntList) {
	result = IntList{}
	for _, element := range i {
		if fn(element) {
			result = append(result, element)
		}
	}
	return result
}

// Map appliest a function to each member of the list
func (i IntList) Map(fn unaryFunc) IntList {
	for n := 0; n < i.Length(); n++ {
		i[n] = fn(i[n])
	}
	return i
}

// Reverse reverses the list
func (i IntList) Reverse() IntList {
	for n, m := 0, i.Length()-1; n < m; n, m = n+1, m-1 {
		i[n], i[m] = i[m], i[n]
	}
	return i
}

// Append combines two lists together
func (i IntList) Append(list IntList) (result IntList) {
	inputLength := i.Length()
	listLength := list.Length()
	result = make(IntList, inputLength+listLength)
	copy(result, i)
	for n := 0; n < listLength; n++ {
		result[n+inputLength] = list[n]
	}
	return result
}

// Concat combines an array of lists together
func (i IntList) Concat(lists []IntList) (result IntList) {
	result = i
	for _, list := range lists {
		result = result.Append(list)
	}
	return result
}

// Length gets the length of the list
func (i IntList) Length() (length int) {
	for range i {
		length++
	}
	return length
}
