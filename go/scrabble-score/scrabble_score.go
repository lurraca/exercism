// Package scrabble implements the function to get the score of a given word.
package scrabble

import "strings"

// Score returns the scrabble result of a word
func Score(s string) int {
	score := 0

	splitString := sanitizedInput(s)

	letterValues := mapOfLetterValues()

	for _, letter := range splitString {
		score += letterValues[letter]
	}

	return score
}

func sanitizedInput(s string) []string {
	return strings.Split(strings.ToUpper(s), "")
}

func mapOfLetterValues() map[string]int {
	letterValues := map[string]int{
		"A": 1,
		"E": 1,
		"I": 1,
		"O": 1,
		"U": 1,
		"L": 1,
		"N": 1,
		"R": 1,
		"S": 1,
		"T": 1,
		"D": 2,
		"G": 2,
		"B": 3,
		"C": 3,
		"M": 3,
		"P": 3,
		"F": 4,
		"H": 4,
		"V": 4,
		"W": 4,
		"Y": 4,
		"K": 5,
		"J": 8,
		"X": 8,
		"Q": 10,
		"Z": 10,
	}
	return letterValues
}
