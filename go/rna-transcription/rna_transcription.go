package strand

const testVersion = 3

func ToRNA(dna string) string {
	rna := make([]byte, len(dna))
	for i, c := range []byte(dna) {
		switch c {
		case 'C':
			rna[i] = 'G'
		case 'G':
			rna[i] = 'C'
		case 'T':
			rna[i] = 'A'
		case 'A':
			rna[i] = 'U'
		}
	}
	return string(rna)
}
