package rl

import (
	"golang.org/x/exp/constraints"
	"math"
	"time"
)

var unixEpoch = time.Unix(0, 0)

func max[T constraints.Integer | constraints.Float](a T, b T) T {
	return T(math.Max(float64(a), float64(b)))
}

func min[T constraints.Integer | constraints.Float](a T, b T) T {
	return T(math.Min(float64(a), float64(b)))
}
