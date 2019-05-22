package luhn

import "strings"
import "strconv"
import "unicode"
import "errors"


func Valid(input string) bool {
	sanitizedInput, err := sanitize(input)
	if err != nil {
		return false
	}
	return isValid(checksum(sanitizedInput))
}


func sanitize(input string) ([]int, error) {
	nonWhiteSpaceInput := strings.Split(strings.Join(strings.Fields(input), ""), "")

	if (len(nonWhiteSpaceInput) <= 1){
		return nil, errors.New("Input needs to be at least two chars")
	}

  nonWhiteSpaceInput = reverse(nonWhiteSpaceInput)
	sanitizedSlice := make([]int, len(nonWhiteSpaceInput))

	for i, el := range nonWhiteSpaceInput {
		if (isDigit(el)){
			return nil, errors.New("Element is not a digit")
		}
		sanitizedSlice[i], _ = strconv.Atoi(el)
	}
	return sanitizedSlice, nil
}

func isValid(digits []int) bool {
	sumOfAllDigits := 0
	isValid := false

	for _, digit :=  range digits {
		sumOfAllDigits =+ sumOfAllDigits + digit
	}

	if (sumOfAllDigits % 10 == 0) {
		isValid = true
	}
	return isValid
}

func checksum(digits []int) []int {
	checksum := make([]int, len(digits))
	for i, digit := range digits {
		if ((i + 1) % 2 == 0) {
			if (digit * 2 > 9) {
				checksum[i] = (digit * 2) - 9
			} else {
				checksum[i] = digit * 2
			}
		} else {
			checksum[i] = digit
		}
	}

	return checksum
}

func reverse(input []string) []string {
	s := make([]string, len(input))
	copy(s, input)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
    s[i], s[j] = s[j], s[i]
	}
	return s
}

func isDigit(element string) bool {
	return !unicode.IsDigit([]rune(element)[0])
}
