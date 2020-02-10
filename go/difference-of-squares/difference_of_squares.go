package diffsquares

// SquareOfSum - Return the square of the sum of the first N number.
func SquareOfSum(n int) int {
	sum := (n * (n + 1)) / 2
	return sum * sum
}

// SumOfSquares - Return the sum of the squares of the first N number.
func SumOfSquares(n int) int {
	sum := (n * (n + 1) * (2*n + 1)) / 6
	return sum
}

// Difference between the square of the sum of the first N natural numbers and the sum of the squares of the first n numbers.
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
