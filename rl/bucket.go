package rl

import (
	"golang.org/x/exp/constraints"
	"math"
	"time"
)

type Bucket struct {
	Uses       int
	Limit      int
	ResetAfter time.Duration
	NextReset  time.Time
}

var unixEpoch = time.Unix(0, 0)

// NewBucket creates a new Bucket.
func NewBucket(limit int, resetAfter time.Duration) *Bucket {
	return &Bucket{0, limit, resetAfter, unixEpoch}
}

// CanDraw checks if a certain number of tokens can be drawn.
func (b *Bucket) CanDraw(amount int) bool {
	b.ensureReset()
	return b.RemainingUses() >= amount
}

// Draw draws tokens from a bucket, returning false and doing nothing if there are not enough.
func (b *Bucket) Draw(amount int) bool {
	b.ensureReset()
	if b.CanDraw(amount) {
		b.Uses += amount
		return true
	}
	return false
}

// DrawMax draws as many tokens from as possible up to `amount`, returning the number of drawn tokens.
func (b *Bucket) DrawMax(amount int) int {
	b.ensureReset()
	count := min(amount, b.RemainingUses())
	b.Uses += count
	return count
}

// ForceDraw forcefully draw a certain number of tokens and
// returns the number of remaining uses, which may be negative.
//
// The number of uses be reset to the limit at the next reset,
// even if this returns a negative number due to excess drawing.
func (b *Bucket) ForceDraw(amount int) int {
	b.ensureReset()
	b.Uses += amount
	return b.RemainingUses()
}

// RemainingUses returns the remaining uses until the bucket is depleted, which may be negative.
func (b *Bucket) RemainingUses() int {
	b.ensureReset()
	return b.Limit - b.Uses
}

// RemainingTime returns the remaining time until the next reset, in milliseconds.
func (b *Bucket) RemainingTime() int64 {
	remainingTime := b.NextReset.Unix() - time.Now().Unix()
	return max(0, remainingTime)
}

// Reset resets this Bucket's uses.
func (b *Bucket) Reset() {
	b.Uses = 0
	b.NextReset = time.Now().Add(b.ResetAfter)
}

func (b *Bucket) ensureReset() {
	if b.RemainingTime() == 0 {
		b.Reset()
	}
}

func max[T constraints.Integer | constraints.Float](a T, b T) T {
	return T(math.Max(float64(a), float64(b)))
}

func min[T constraints.Integer | constraints.Float](a T, b T) T {
	return T(math.Min(float64(a), float64(b)))
}
