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

// New creates a new BucketManager
//
// Bucket instances are automatically created and inserted
// when they don't already exist, but you can create and
// add your own to specify per-bucket rate limits
func New(limit int, resetAfter time.Duration) *BucketManager {
	return &BucketManager{
		Buckets:    make(map[string]*Bucket),
		Limit:      limit,
		ResetAfter: resetAfter,
	}
}

// Get gets a Bucket from this BucketManager by its ID
//
// Creates a new Bucket if one does not exist
func (bm *BucketManager) Get(id string) *Bucket {
	return bm.getOrCreate(id)
}

// CanDraw checks if a certain number of tokens
// can be drawn from a Bucket. Returns a bool
func (bm *BucketManager) CanDraw(id string, amount int) bool {
	return bm.getOrCreate(id).CanDraw(amount)
}

// Draw draws a specific number of tokens from a bucket
//
// Returns false and does nothing if there
// are not enough tokens in the bucket
func (bm *BucketManager) Draw(id string, amount int) bool {
	return bm.getOrCreate(id).Draw(amount)
}

// DrawMax draws as many tokens from a Bucket as possible
//
// Returns the number of tokens drawn
func (bm *BucketManager) DrawMax(id string, amount int) int {
	return bm.getOrCreate(id).DrawMax(amount)
}

// ForceDraw forcefully draw a certain number of tokens from a Bucket
//
// Returns the number of remaining uses (may be negative)
//
// This will be reset to 0 at the next reset, even if it is negative
func (bm *BucketManager) ForceDraw(id string, amount int) int {
	return bm.getOrCreate(id).ForceDraw(amount)
}

// RemainingUses returns the remaining uses until a Bucket is depleted.
func (bm *BucketManager) RemainingUses(id string) int {
	return bm.getOrCreate(id).RemainingUses()
}

// RemainingTime returns the remaining time until a Bucket resets
func (bm *BucketManager) RemainingTime(id string) int64 {
	return bm.getOrCreate(id).RemainingTime()
}

// Reset resets a Bucket's uses
func (bm *BucketManager) Reset(id string) {
	bm.getOrCreate(id).Reset()
}

// Add manually adds a Bucket to this BucketManager
func (bm *BucketManager) Add(id string, bucket *Bucket) {
	bm.Buckets[id] = bucket
}

// SaveFile saves this BucketManager to a file as JSON
func (bm *BucketManager) SaveFile(path string) error {
	return fn.WriteJsonFile(path, bm)
}

// LoadFile loads a BucketManager from a JSON file
func LoadFile(path string) (*BucketManager, error) {
	var bm BucketManager
	return &bm, fn.ReadJsonFile(path, &bm)
}

func (bm *BucketManager) getOrCreate(id string) *Bucket {
	if _, err := bm.Buckets[id]; !err {
		bm.Buckets[id] = NewBucket(bm.Limit, bm.ResetAfter)
	}
	return bm.Buckets[id]
}
