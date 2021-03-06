package rl

import (
	"encoding/json"
	"github.com/zytekaron/gotil/fn"
	"time"
)

type BucketManager struct {
	Buckets    map[string]*Bucket
	Limit      int
	ResetAfter time.Duration
}

// Create a new BucketManager
//
// Bucket instances are automatically created and inserted
// when they don't already exist, but you can create and
// add your own to specify per-bucket rate limits
func New(limit int, resetAfter time.Duration) *BucketManager {
	return &BucketManager{make(map[string]*Bucket), limit, resetAfter}
}

// Get a Bucket from this BucketManager by its ID
//
// Creates a new Bucket if one does not exist
func (bm *BucketManager) Get(id string) *Bucket {
	return bm.getOrCreate(id)
}

// Checks if a certain number of tokens can
// be drawn from a Bucket. Returns a bool
func (bm *BucketManager) CanDraw(id string, amount int) bool {
	return bm.getOrCreate(id).CanDraw(amount)
}

// Draw a specific number of tokens from a bucket
//
// Returns false and does nothing if there
// are not enough tokens in the bucket
func (bm *BucketManager) Draw(id string, amount int) bool {
	return bm.getOrCreate(id).Draw(amount)
}

// Draw as many tokens from a Bucket as possible
//
// Returns the number of tokens drawn
func (bm *BucketManager) DrawMax(id string, amount int) int {
	return bm.getOrCreate(id).DrawMax(amount)
}

// Forcefully draw a certain number of tokens from a Bucket
//
// Returns the number of remaining uses (may be negative)
//
// This will be reset to 0 at the next reset, even if it is negative
func (bm *BucketManager) ForceDraw(id string, amount int) int {
	return bm.getOrCreate(id).ForceDraw(amount)
}

// Check the remaining uses until a Bucket is depleted.
func (bm *BucketManager) RemainingUses(id string) int {
	return bm.getOrCreate(id).RemainingUses()
}

// Check the remaining time until a Bucket resets
func (bm *BucketManager) RemainingTime(id string) int64 {
	return bm.getOrCreate(id).RemainingTime()
}

// Reset a Bucket's uses
func (bm *BucketManager) Reset(id string) {
	bm.getOrCreate(id).Reset()
}

// Manually add a Bucket to this BucketManager
func (bm *BucketManager) Add(id string, bucket *Bucket) {
	bm.Buckets[id] = bucket
}

// Save this BucketManager as JSON
func (bm *BucketManager) Save() ([]byte, error) {
	b, err := json.Marshal(bm)
	return b, err
}

// Save this BucketManager to a file as a Go Object
func (bm *BucketManager) SaveFile(file string) error {
	err := fn.WriteGobFile(file, bm)
	return err
}

// Load a BucketManager from a saved JSON String
func Load(bytes []byte) (*BucketManager, error) {
	var bm BucketManager
	err := json.Unmarshal(bytes, &bm)
	return &bm, err
}

// Load a BucketManager from a Go Object file
func LoadFile(file string) (*BucketManager, error) {
	var bm BucketManager
	err := fn.ReadGobFile(file, &bm)
	return &bm, err
}

func (bm *BucketManager) getOrCreate(id string) *Bucket {
	if _, err := bm.Buckets[id]; !err {
		bm.Buckets[id] = NewBucket(bm.Limit, bm.ResetAfter)
	}
	return bm.Buckets[id]
}
