// Package reverse implements the String function to reverse a string
package reverse

import "strings"

// String return a reversed string of the input
func String(input string) string {
	output := ""

	for _, v := range strings.Split(input, "") {
		output = v + output
	}

	return output
}
