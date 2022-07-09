package rl

import (
	"sync"
	"time"
)

type SyncBucket struct {
	Uses       int
	Limit      int
	ResetAfter time.Duration
	NextReset  time.Time
	mutex      sync.RWMutex
}

// NewSyncBucket creates a new SyncBucket.
func NewSyncBucket(limit int, resetAfter time.Duration) *SyncBucket {
	return &SyncBucket{
		Uses:       0,
		Limit:      limit,
		ResetAfter: resetAfter,
		NextReset:  unixEpoch,
	}
}

// CanDraw checks if a certain number of tokens can be drawn.
func (b *SyncBucket) CanDraw(amount int) bool {
	b.ensureReset()

	return b.remainingUsesInternal() >= amount
}

// Draw draws tokens from a SyncBucket, returning false and doing nothing if there are not enough.
func (b *SyncBucket) Draw(amount int) bool {
	b.ensureReset()

	if !b.CanDraw(amount) {
		return false
	}

	b.mutex.Lock()
	b.Uses += amount
	b.mutex.Unlock()
	return true
}

// DrawMax draws as many tokens from as possible up to `amount`, returning the number of drawn tokens.
func (b *SyncBucket) DrawMax(amount int) int {
	b.ensureReset()

	available := max(0, b.remainingUsesInternal())
	count := min(amount, available)
	b.mutex.Lock()
	b.Uses += count
	b.mutex.Unlock()
	return count
}

// ForceDraw forcefully draw a certain number of tokens and
// returns the number of remaining uses, which may be negative.
//
// The number of uses be reset to the limit at the next reset,
// even if this returns a negative number due to excess drawing.
func (b *SyncBucket) ForceDraw(amount int) int {
	b.ensureReset()

	b.mutex.Lock()
	b.Uses += amount
	b.mutex.Unlock()
	return b.remainingUsesInternal()
}

// RemainingUses returns the remaining uses until the SyncBucket is depleted, which may be negative.
func (b *SyncBucket) RemainingUses() int {
	b.ensureReset()

	b.mutex.RLock()
	defer b.mutex.RUnlock()
	return b.Limit - b.Uses
}

// RemainingTime returns the remaining time until the next reset, in milliseconds.
func (b *SyncBucket) RemainingTime() int64 {
	b.mutex.RLock()
	remainingTime := b.NextReset.Unix() - time.Now().Unix()
	b.mutex.RUnlock()
	return max(0, remainingTime)
}

// Reset resets this SyncBucket's uses.
func (b *SyncBucket) Reset() {
	b.mutex.Lock()
	b.Uses = 0
	b.NextReset = time.Now().Add(b.ResetAfter)
	b.mutex.Unlock()
}

// used to prevent a duplicate ensureReset call when used internally
func (b *SyncBucket) remainingUsesInternal() int {
	b.mutex.RLock()
	defer b.mutex.RUnlock()
	return b.Limit - b.Uses
}

// reset the bucket whenever the next reset time has passed
func (b *SyncBucket) ensureReset() {
	if b.RemainingTime() == 0 {
		b.Reset()
	}
}
