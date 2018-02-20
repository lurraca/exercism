// Package isogram implements a function that returns a boolean whether a word is an isogram or not
package isogram

import (
	"strings"
)

// IsIsogram returns true if the input word is an Isogram
func IsIsogram(input string) bool {
	isIsogram := true

	sanitizedInput := strings.Split(strings.ToUpper(removeSpecialCharacters(input)), "")

	for index, letter := range sanitizedInput {
		if contains(letter, sanitizedInput[index+1:]) {
			isIsogram = false
			break
		}
	}
	return isIsogram
}

func removeSpecialCharacters(input string) string {
	return strings.Replace(removeWhiteSpaces(input), "-", "", -1)
}
func removeWhiteSpaces(input string) string {
	return strings.Replace(input, " ", "", -1)
}
func contains(letter string, slice []string) bool {
	for _, v := range slice {
		if v == letter {
			return true
		}
	}
	return false
}
