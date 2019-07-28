//Package hamming calculates the hamming distance between two DNA strands
package hamming

import (
	"errors"
)

// Distance calculates the distance (int) between DNA strands (a, b)
func Distance(a, b string) (int, error) {
	// Check both strands are equal length
	if len(a) != len(b) {
		return 0, errors.New("not possible to calculate hamming distance between strands of different lengths")
	}

	// Loop through the strings and compare each char
	count := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			count++
		}
	}

	return count, nil
}
