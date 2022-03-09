package maths

import (
	"golang.org/x/exp/constraints"
	"math"
)

// Collisions calculates the number of samples
// in a given set required for there to be a
// certain probability that there will be at
// least one collision. See: Birthday Paradox
func Collisions(size int, probability float64) float64 {
	if probability > 1 {
		probability /= 100
	}
	return math.Sqrt(2.0 * float64(size) * math.Log(1/(1-probability)))
}

// Factorial calculates the factorial of a given number
func Factorial(n int) uint64 {
	total := uint64(n)
	for i := 1; i < n; i++ {
		total *= uint64(i)
	}
	return total
}

// Root calculates the nth root of a number
func Root(num, root float64) float64 {
	return math.Pow(num, 1/root)
}

// Round rounds a number to the nearest n trailing zeroes
func Round[T constraints.Float](num T, n int) T {
	pow := T(math.Pow10(n))
	return T(math.Round(float64(num/pow))) * pow
}

// Floor rounds a number down to n trailing zeroes
func Floor[T constraints.Float](num T, n int) T {
	pow := T(math.Pow10(n))
	return T(math.Floor(float64(num/pow))) * pow
}

// Ceil rounds a number up to n trailing zeroes
func Ceil[T constraints.Float](num T, n int) T {
	pow := T(math.Pow10(n))
	return T(math.Ceil(float64(num/pow))) * pow
}

// RoundDecimal rounds a number to the nearest n decimal places
func RoundDecimal[T constraints.Float](num T, n int) T {
	pow := T(math.Pow10(n))
	return T(math.Round(float64(num*pow))) / pow
}

// FloorDecimal rounds a number down to n decimal places
func FloorDecimal[T constraints.Float](num T, n int) T {
	pow := T(math.Pow10(n))
	return T(math.Floor(float64(num*pow))) / pow
}

// CeilDecimal rounds a number up to n decimal places
func CeilDecimal[T constraints.Float](num T, n int) T {
	pow := T(math.Pow10(n))
	return T(math.Ceil(float64(num*pow))) / pow
}
