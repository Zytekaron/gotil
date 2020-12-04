package maths

import (
	. "github.com/zytekaron/gotil/maths"
	"math"
	"math/rand"
	"testing"
)

func TestAverageInt(t *testing.T) {
	arr := make([]int, 32)
	for i := range arr {
		arr[i] = rand.Intn(65536)
	}
	avg1 := AverageInt(arr)
	avg2 := float64(SumInt(arr)) / float64(len(arr))
	if avg1 != avg2 {
		t.Error("average values for []int are not equal")
	}
}

func TestAverageFloat(t *testing.T) {
	arr := make([]float64, 32)
	for i := range arr {
		arr[i] = rand.Float64() * 65536
	}
	avg1 := AverageFloat64(arr)
	avg2 := SumFloat64(arr) / float64(len(arr))
	if math.Abs(avg1-avg2) >= .1 {
		t.Error("average values for []float64 are not similar", avg1, avg2)
	}
}
