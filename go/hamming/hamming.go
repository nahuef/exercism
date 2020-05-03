// Package hamming calculates the Hamming Distance between two strings.
package hamming

import "errors"

// Distance calculates the Hamming Distance between two strings.
func Distance(a string, b string) (int, error) {
	var err = validateLenghts(a, b)
	if err != nil {
		return 0, err
	}

	distance := 0
	bRunes := []rune(b)

	for i, v := range a {
		if v != bRunes[i] {
			distance++
		}
	}

	return distance, nil
}

func validateLenghts(a string, b string) error {
	if len(a) != len(b) {
		return errors.New("Cannot compare different lengths")
	}

	return nil
}
