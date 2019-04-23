package leap

// IsLeapYear calculates if a given year is leap or not
func IsLeapYear(year int) bool {
	return year%4 == 0 && year%100 != 0 || year%400 == 0
}
