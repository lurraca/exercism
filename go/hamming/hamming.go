// Package hamming implements the function for determine the distance of 2 dna strands
package hamming

import (
	"errors"
	"strings"
)

// Distance returns the distance between 2 dna strands
func Distance(a, b string) (int, error) {
	var distance int

	if isDnaStrandUnequal(a, b) {
		return -1, errors.New("Lenght of the DNA strands must be equal.")
	}

	splitA := strings.Split(a, "")
	splitB := strings.Split(b, "")

	if a == b {
		distance = 0
	} else {
		for index, nucleotides := range splitA {
			if nucleotides != splitB[index] {
				distance++
			}
		}
	}
	return distance, nil
}

func isDnaStrandUnequal(strandA string, strandB string) bool {
	return len(strandA) != len(strandB)
}
