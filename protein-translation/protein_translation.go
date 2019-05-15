package protein

import "errors"

// ErrStop is raised when the protein is stop
// ErrInvalidBase represents an error when the base is invalid
var (
	ErrStop        = errors.New("stop")
	ErrInvalidBase = errors.New("invalid base")
)

var proteins = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine",
	"UUC": "Phenylalanine",
	"UUA": "Leucine",
	"UUG": "Leucine",
	"UCU": "Serine",
	"UCC": "Serine",
	"UCA": "Serine",
	"UCG": "Serine",
	"UAU": "Tyrosine",
	"UAC": "Tyrosine",
	"UGU": "Cysteine",
	"UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "STOP",
	"UAG": "STOP",
	"UGA": "STOP",
}

// FromCodon translates a codon to a protein.
func FromCodon(codon string) (string, error) {
	protein, ok := proteins[codon]
	if !ok {
		return "", ErrInvalidBase
	}
	if protein == "STOP" {
		return "", ErrStop
	}
	return protein, nil
}

// FromRNA translates a sequence of codons to a sequence proteins.
func FromRNA(rna string) ([]string, error) {
	result := []string{}
	for i := 0; i <= len(rna)-3; i += 3 {
		protein, err := FromCodon(rna[i : i+3])
		if err == ErrStop {
			break
		}
		if err != nil {
			return result, err
		}
		result = append(result, protein)
	}
	return result, nil
}
