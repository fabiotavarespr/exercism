/*
Package hamming implements

Calculate the Hamming Distance between two DNA strands.

The Hamming distance is only defined for sequences of equal length, so
an attempt to calculate it between sequences of different lengths should
not work. The general handling of this situation (e.g., raising an
exception vs returning a special value) may differ between languages.
*/
package hamming

import "errors"

// Distance - Calculate the Hamming Distance between two DNA strands.
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("Strings lengths don't match")
	}

	aChars := []rune(a)
	bChars := []rune(b)
	dist := 0

	for i, v := range aChars {
		if v != bChars[i] {
			dist++
		}
	}

	return dist, nil
}
