package maths

import (
	"golang.org/x/exp/constraints"
)

// Sum calculates the sum of the elements in a slice.
func Sum[T constraints.Integer | constraints.Float](nums []T) T {
	var sum T
	for _, e := range nums {
		sum += e
	}
	return sum
}

// Average calculates the average of the elements in a slice.
//
// This function uses rolling average calculation to sidestep
// issues related to overflow with large numbers in big slices.
func Average[T constraints.Float](input []T) T {
	var result T
	for i := 0; i < len(input); i++ {
		result = averageRoller(result, T(i), input[i])
	}
	return result
}

func averageRoller[T constraints.Float](current, n, value T) T {
	return (value + n*current) / (n + 1)
}
