// Package acronym contains the implemntation of taking a phrase and converting it to an acronym.
package acronym

import "strings"

// Abbreviate takes a phrase and convert it to an acronum
func Abbreviate(phrase string) string {
	acronym := ""

	splitString := strings.FieldsFunc(phrase, split)

	for _, word := range splitString {
		acronym = acronym + strings.ToUpper(word[:1])
	}
	return acronym
}

func split(r rune) bool {
	return r == ' ' || r == '-'
}
