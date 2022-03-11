package maths

import (
	. "github.com/zytekaron/gotil/v2/maths"
	"math/rand"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func TestSum(t *testing.T) {
	var arr []int
	if Sum(arr) != 0 {
		t.Error("expected 0 but got", Sum(arr))
	}

	arr = []int{1, 4, 5, 5, 6, 6}
	if Sum(arr) != 27 {
		t.Error("expected 27 but got", Sum(arr))
	}
}

func TestAverage(t *testing.T) {
	var arr1 []float64
	if Average(arr1) != 0 {
		t.Error("expected 0 but got", Average(arr1))
	}

	arr2 := []float64{1, 4, 5, 5, 6, 6}
	if Average(arr2) != 4.5 {
		t.Error("expected 4.5 but got", Average(arr2))
	}

	arr3 := make([]float64, 1024)
	for i := range arr3 {
		arr3[i] = rand.Float64()
	}
	expect := Sum(arr3) / float64(len(arr3))
	roundedResult := Round(Average(arr3), 9)
	roundedExpect := Round(expect, 9)
	if roundedResult != roundedExpect {
		t.Error("expected", roundedExpect, "but got", roundedResult)
	}
}
