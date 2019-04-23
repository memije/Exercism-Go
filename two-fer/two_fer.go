package twofer

import "fmt"

// ShareWith returns a string containing "One for <name>, one for me."
// or "One for you, one for me." when name is empty
func ShareWith(name string) string {
	if len(name) == 0 {
		name = "you"
	}
	return fmt.Sprintf("One for %v, one for me.", name)
}
