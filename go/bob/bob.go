// Package bob implements the Hey function
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"strings"
)

// Hey takes in a remark and answers like a teenager.
func Hey(remark string) string {
	trimmedRemark := strings.TrimSpace(remark)

	answer := "Whatever."

	if isShouted(trimmedRemark) {
		if isQuestion(trimmedRemark) {
			answer = "Calm down, I know what I'm doing!"
		} else {
			answer = "Whoa, chill out!"
		}
	} else if isQuestion(trimmedRemark) {
		answer = "Sure."
	} else if isSilence(trimmedRemark) {
		answer = "Fine. Be that way!"
	}
	return answer
}

func isSilence(remark string) bool {
	return remark == ""
}
func isShouted(remark string) bool {
	return strings.ToUpper(remark) == remark && strings.ToLower(remark) != remark
}
func isQuestion(remark string) bool {
	return strings.HasSuffix(remark, "?")
}
