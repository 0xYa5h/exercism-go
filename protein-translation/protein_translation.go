package protein

import (
	"errors"
)

var (
	ErrStop        = errors.New("terminating codon")
	ErrInvalidBase = errors.New("invalid base")
)

func FromRNA(rna string) ([]string, error) {
	var polypeptide []string

	for i := 0; i < len(rna); i += 3 {
		codon := rna[i:i+3]

		protein, err := FromCodon(codon)
		if err != nil {
			if err == ErrStop {
				return polypeptide, nil
			}

			return polypeptide, err
		}

		polypeptide = append(polypeptide, protein)
	}

	return polypeptide, nil
}

func FromCodon(codon string) (string, error) {
	switch codon {
	case "AUG":
		return "Methionine", nil
	case "UUU", "UUC":
		return "Phenylalanine", nil
	case "UUA", "UUG":
		return "Leucine", nil
	case "UCU", "UCC", "UCA", "UCG":
		return "Serine", nil
	case "UAU", "UAC":
		return "Tyrosine", nil
	case "UGU", "UGC":
		return "Cysteine", nil
	case "UGG":
		return "Tryptophan", nil
	case "UAA", "UAG", "UGA":
		return "", ErrStop
	}

	return "", ErrInvalidBase
}
