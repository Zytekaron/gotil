package maths

import "math"

// Calculate the number of samples in a given set
// required for there to be a certain probability
// that there will be at least one collision.
// See: Birthday Paradox
func Collisions(size int, probability float64) float64 {
	if probability > 1 {
		probability /= 100
	}
	return math.Sqrt(2.0 * float64(size) * math.Log(1/(1-probability)))
}

// Calculate the factorial of a given number
func Factorial(n uint64) uint64 {
	total := n
	for i := n - 1; i > 1; i-- {
		total *= i
	}
	return total
}

// Calculate the nth root of a number
func Root(num, root float64) float64 {
	return math.Pow(num, 1/root)
}
