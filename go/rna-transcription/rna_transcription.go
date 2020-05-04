package strand

var complements = map[rune]string{
	'G': "C",
	'C': "G",
	'T': "A",
	'A': "U",
}

// ToRNA Given a DNA strand, return its RNA complement.
func ToRNA(dna string) (RNA string) {
	for _, r := range dna {
		RNA = RNA + complements[r]
	}
	return
}
