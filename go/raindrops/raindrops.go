// Package raindrops contains the implementation of converting a number into string depending on the number's factors.
package raindrops

import "strconv"

// Convert a number into a string which depending on the number's factors
func Convert(number int) string {
	output := ""

	if factorOf3(number) {
		output += "Pling"
	}

	if factorOf5(number) {
		output += "Plang"
	}

	if factorOf7(number) {
		output += "Plong"
	}

	if output == "" {
		output = strconv.Itoa(number)

	}

	return output
}

func factorOf3(number int) bool {
	return number%3 == 0
}

func factorOf5(number int) bool {
	return number%5 == 0
}

func factorOf7(number int) bool {
	return number%7 == 0
}
