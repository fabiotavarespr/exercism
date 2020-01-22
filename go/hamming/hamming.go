package hamming

import "errors"

// Distance - Calculate the Hamming Distance between two DNA strands.
func Distance(a, b string) (int, error) {

	// Check if a and b have the same length
	if len(a) != len(b) {
		return 0, errors.New("DNA strands must be the same lengths")
	}

	// Iterate between the two DNA strands and calculate hamming distance
	hammingDistance := 0
	for i := range a {
		if a[i] != b[i] {
			hammingDistance++
		}
	}
	return hammingDistance, nil

}
