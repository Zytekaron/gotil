package maths

import (
	"math"
)

// SumInt calculates the sum of the elements in a slice
func SumInt(nums []int) int {
	sum := 0
	for _, e := range nums {
		sum += e
	}
	return sum
}

// SumFloat64 calculates the sum of the elements in a slice
func SumFloat64(nums []float64) float64 {
	sum := 0.0
	for _, e := range nums {
		sum += e
	}
	return sum
}

// AverageInt calculates the average value of all the elements in a slice
//
// This uses rolling averages to prevent issues related to integer overflow
func AverageInt(nums []int) float64 {
	x := 0
	y := 0
	length := len(nums)
	for i := 0; i < length; i++ {
		x += nums[i] / length
		b := nums[i] % length
		if y >= length-b {
			x++
			y -= length - b
		} else {
			y += b
		}
	}
	return float64(x) + float64(y)/float64(length)
}

// AverageFloat64 calculates the average value of all the elements in a slice
//
// This uses rolling averages to prevent issues related to number overflow
func AverageFloat64(nums []float64) float64 {
	x := 0
	y := 0.0
	length := float64(len(nums))
	for i := 0; i < len(nums); i++ {
		x += int(nums[i]) / int(length)
		b := math.Mod(nums[i], length)
		if y >= length-b {
			x++
			y -= length - b
		} else {
			y += b
		}
	}
	return float64(x) + y/length
}
