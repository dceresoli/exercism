package protein

const testVersion = 1

/*
AUG                   | Methionine
UUU, UUC              | Phenylalanine
UUA, UUG              | Leucine
UCU, UCC, UCA, UCG    | Serine
UAU, UAC              | Tyrosine
UGU, UGC              | Cysteine
UGG                   | Tryptophan
UAA, UAG, UGA         | STOP
*/

func FromCodon(input string) string {
	var residue string
	switch input {
	case "AUG":
		residue = "Methionine"

	case "UUU":
		fallthrough
	case "UUC":
		residue = "Phenylalanine"

	case "UUA":
		fallthrough
	case "UUG":
		residue = "Leucine"

	case "UCU":
		fallthrough
	case "UCC":
		fallthrough
	case "UCA":
		fallthrough
	case "UCG":
		residue = "Serine"

	case "UAU":
		fallthrough
	case "UAC":
		residue = "Tyrosine"

	case "UGU":
		fallthrough
	case "UGC":
		residue = "Cysteine"

	case "UGG":
		residue = "Tryptophan"

	case "UAA":
		fallthrough
	case "UAG":
		fallthrough
	case "UGA":
		residue = "STOP"

	}
	return residue
}

func FromRNA(input string) []string {
	protein := make([]string, 0)
	for i := 0; i < len(input); i += 3 {
		residue := FromCodon(input[i : i+3])
		if residue == "STOP" {
			break
		}
		protein = append(protein, residue)
	}
	return protein
}
