// Package triangle determines the kind of a triangle based on its sides
package triangle

import (
	"math"
)

// Kind represents the kind of a triangle
type Kind string

const (
	// NaT - Not a triangle
	NaT Kind = "NaT"
	// Equ - Equilateral
	Equ Kind = "Equ"
	// Iso - Isosceles
	Iso Kind = "Iso"
	// Sca - Scalene
	Sca Kind = "Sca"
	// Deg - Degenerate
	Deg Kind = "Deg"
)

func isValidSide(side float64) bool {
	return side > 0 && !math.IsNaN(side) && !math.IsInf(side, 0)
}

func isValidTriangle(a, b, c float64) bool {
	// Triangle is invalid if any of its sides is invalid
	if !isValidSide(a) ||
		!isValidSide(b) ||
		!isValidSide(c) {
		return false
	}

	// The sum of the lengths of any two sides must be greater than or equal
	// to the length of the third side
	if a <= b+c && b <= a+c && c <= a+b {
		return true
	}
	return false
}

// An equilateral triangle has all three sides the same length
func isEquilateral(a, b, c float64) bool {
	if a == b && b == c {
		return true
	}
	return false
}

// A scalene triangle has all sides of different lengths
func isScalene(a, b, c float64) bool {
	if a != b && b != c && a != c {
		return true
	}
	return false
}

// An isoceles triangle has at least to sides of the same length
func isIsoceles(a, b, c float64) bool {
	return (a == b && b != c) ||
		(b == c && c != a) ||
		(a == c && c != b)
}

// A degenerate triangle has the sum of the lengths of two sides equal to
// the third one
func isDegenerate(a, b, c float64) bool {
	return a+b == c || b+c == a || a+c == b
}

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	if !isValidTriangle(a, b, c) {
		return NaT
	}

	if isEquilateral(a, b, c) {
		return Equ
	}
	if isScalene(a, b, c) {
		return Sca
	}
	if isIsoceles(a, b, c) {
		return Iso
	}
	if isDegenerate(a, b, c) {
		return Deg
	}
	return NaT
}
