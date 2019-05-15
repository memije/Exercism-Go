package proverb

import "fmt"

// Proverb generates the relevant proverb given a list of inputs.
func Proverb(rhyme []string) (proverb []string) {
	if len(rhyme) == 0 {
		return proverb
	}

	// Initialize proverb
	proverb = []string{}

	// Proverb building
	if len(rhyme) > 1 {
		for i := 1; i < len(rhyme); i++ {
			proverb = append(
				proverb,
				fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i-1], rhyme[i]),
			)
		}
	}

	// Proverb ending
	proverb = append(
		proverb,
		fmt.Sprintf("And all for the want of a %s.", rhyme[0]),
	)

	return proverb
}
