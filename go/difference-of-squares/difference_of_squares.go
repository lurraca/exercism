// Package diffsquares implements a few methods to determine the difference between Square of Sums and Sum of Squares
package diffsquares

// SquareOfSums return the square of the sums of an input
func SquareOfSums(input int) int {
	output := 0
	for i := 1; i <= input; i++ {
		output = i + output
	}
	return output * output
}

// SumOfSquares return the sum of the squares of an input
func SumOfSquares(input int) int {
	output := 0
	for i := 1; i <= input; i++ {
		output = (i * i) + output
	}
	return output
}

// Difference returns the difference between SquareOfSums and SquareOfSums
func Difference(input int) int {
	return SquareOfSums(input) - SumOfSquares(input)
}
