package strand

var transcriptions = map[rune]rune{
	'G': 'C',
	'C': 'G',
	'T': 'A',
	'A': 'U',
}

// ToRNA transforms a DNA strand to its RNA complement.
func ToRNA(dna string) (rna string) {
	for _, nucleotide := range []rune(dna) {
		if value, ok := transcriptions[nucleotide]; ok {
			rna += string(value)
		}
	}
	return rna
}
