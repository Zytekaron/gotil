package rl

import (
	"github.com/zytekaron/gotil/v2/fn"
	"time"
)

type BucketManager struct {
	Buckets    map[string]*Bucket
	Limit      int
	ResetAfter time.Duration
}

// New creates a new BucketManager.
//
// Bucket instances are automatically inserted as needed.
func New(limit int, resetAfter time.Duration) *BucketManager {
	return &BucketManager{
		Buckets:    make(map[string]*Bucket),
		Limit:      limit,
		ResetAfter: resetAfter,
	}
}

// Get gets a Bucket from this BucketManager by its ID.
//
// Creates a new Bucket if one does not exist.
func (bm *BucketManager) Get(id string) *Bucket {
	return bm.getOrCreate(id)
}

// CanDraw checks if a certain number of tokens can be drawn from a Bucket.
func (bm *BucketManager) CanDraw(id string, amount int) bool {
	return bm.getOrCreate(id).CanDraw(amount)
}

// Draw draws tokens from a bucket, returning false and doing nothing if there are not enough.
func (bm *BucketManager) Draw(id string, amount int) bool {
	return bm.getOrCreate(id).Draw(amount)
}

// DrawMax draws as many tokens from a Bucket as possible up to `amount`, returning the number of drawn tokens.
func (bm *BucketManager) DrawMax(id string, amount int) int {
	return bm.getOrCreate(id).DrawMax(amount)
}

// ForceDraw forcefully draw a certain number of tokens from a Bucket
// and returns the number of remaining uses, which may be negative.
//
// The number of uses be reset to the limit at the next reset,
// even if this returns a negative number due to excess drawing.
func (bm *BucketManager) ForceDraw(id string, amount int) int {
	return bm.getOrCreate(id).ForceDraw(amount)
}

// RemainingUses returns the remaining uses until the Bucket is depleted, which may be negative.
func (bm *BucketManager) RemainingUses(id string) int {
	return bm.getOrCreate(id).RemainingUses()
}

// RemainingTime returns the remaining time until the Bucket resets in milliseconds.
func (bm *BucketManager) RemainingTime(id string) int64 {
	return bm.getOrCreate(id).RemainingTime()
}

// Reset resets a Bucket's uses.
func (bm *BucketManager) Reset(id string) {
	bm.getOrCreate(id).Reset()
}

// Add manually adds a Bucket to this BucketManager.
func (bm *BucketManager) Add(id string, bucket *Bucket) {
	bm.Buckets[id] = bucket
}

// SaveFile saves this BucketManager to a file as JSON.
func (bm *BucketManager) SaveFile(path string) error {
	return fn.WriteJsonFile(path, bm)
}

// LoadFile loads a BucketManager from a JSON file.
func LoadFile(path string) (*BucketManager, error) {
	var bm *BucketManager
	return bm, fn.ReadJsonFile(path, &bm)
}

func (bm *BucketManager) getOrCreate(id string) *Bucket {
	if _, err := bm.Buckets[id]; !err {
		bm.Buckets[id] = NewBucket(bm.Limit, bm.ResetAfter)
	}
	return bm.Buckets[id]
}
