package rl

import (
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

// Creates a new Bucket
//
// Automatically called by BucketManager on all methods
func NewBucket(limit int, resetAfter time.Duration) *Bucket {
	return &Bucket{0, limit, resetAfter, unixEpoch}
}

// Checks if a certain number of tokens can
// be drawn from this Bucket. Returns a bool
func (b *Bucket) CanDraw(amount int) bool {
	b.ensureReset()
	return b.RemainingUses() >= amount
}

// Draw a specific number of tokens from this bucket
//
// Returns false and does nothing if there
// are not enough tokens in the bucket
func (b *Bucket) Draw(amount int) bool {
	b.ensureReset()
	if b.CanDraw(amount) {
		b.Uses += amount
		return true
	}
	return false
}

// Draw as many tokens from this Bucket as possible
//
// Returns the number of tokens drawn
func (b *Bucket) DrawMax(amount int) int {
	b.ensureReset()
	max := min(float64(amount), float64(b.RemainingUses()))
	b.Uses += max
	return max
}

// Forcefully draw a certain number of tokens from this Bucket
//
// Returns the number of remaining uses (may be negative)
//
// This will be reset to 0 at the next reset, even if it is negative
func (b *Bucket) ForceDraw(amount int) int {
	b.ensureReset()
	b.Uses += amount
	return b.RemainingUses()
}

// Check the remaining uses until this Bucket is depleted
func (b *Bucket) RemainingUses() int {
	b.ensureReset()
	return b.Limit - b.Uses
}

// Check the remaining time until this Bucket resets
func (b *Bucket) RemainingTime() int64 {
	remainingTime := b.NextReset.Unix() - time.Now().Unix()
	maxTime := max(0, float64(remainingTime))
	return int64(maxTime)
}

// Reset this Bucket's uses
func (b *Bucket) Reset() {
	b.Uses = 0
	b.NextReset = time.Now().Add(b.ResetAfter)
}

func (b *Bucket) ensureReset() {
	if b.RemainingTime() == 0 {
		b.Reset()
	}
}

func max(a, b float64) int {
	return int(math.Max(a, b))
}

func min(a, b float64) int {
	return int(math.Min(a, b))
}
