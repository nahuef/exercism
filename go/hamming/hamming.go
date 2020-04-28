// Package hamming calculates the Hamming Distance between two strings.
package hamming

import "errors"

// Distance calculates the Hamming Distance between two strings.
func Distance(a string, b string) (int, error) {
	var err = validateLenghts(a, b)
	if err != nil {
		return -1, err
	}

	distance := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
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
