package rl

import (
	"github.com/zytekaron/gotil/v2/fn"
	"sync"
	"time"
)

type SyncBucketManager struct {
	Buckets    map[string]*SyncBucket
	Limit      int
	ResetAfter time.Duration
	mutex      sync.RWMutex
}

// NewSync creates a new SyncBucketManager.
//
// SyncBucket instances are automatically inserted as needed.
func NewSync(limit int, resetAfter time.Duration) *SyncBucketManager {
	return &SyncBucketManager{
		Buckets:    make(map[string]*SyncBucket),
		Limit:      limit,
		ResetAfter: resetAfter,
	}
}

// Get gets a SyncBucket from this SyncBucketManager by its ID.
//
// Creates a new SyncBucket if one does not exist.
func (bm *SyncBucketManager) Get(id string) *SyncBucket {
	return bm.getOrCreate(id)
}

// CanDraw checks if a certain number of tokens can be drawn from a SyncBucket.
func (bm *SyncBucketManager) CanDraw(id string, amount int) bool {
	return bm.getOrCreate(id).CanDraw(amount)
}

// Draw draws tokens from a SyncBucket, returning false and doing nothing if there are not enough.
func (bm *SyncBucketManager) Draw(id string, amount int) bool {
	return bm.getOrCreate(id).Draw(amount)
}

// DrawMax draws as many tokens from a SyncBucket as possible up to `amount`, returning the number of drawn tokens.
func (bm *SyncBucketManager) DrawMax(id string, amount int) int {
	return bm.getOrCreate(id).DrawMax(amount)
}

// ForceDraw forcefully draw a certain number of tokens from a SyncBucket
// and returns the number of remaining uses, which may be negative.
//
// The number of uses be reset to the limit at the next reset,
// even if this returns a negative number due to excess drawing.
func (bm *SyncBucketManager) ForceDraw(id string, amount int) int {
	return bm.getOrCreate(id).ForceDraw(amount)
}

// RemainingUses returns the remaining uses until the SyncBucket is depleted, which may be negative.
func (bm *SyncBucketManager) RemainingUses(id string) int {
	return bm.getOrCreate(id).RemainingUses()
}

// RemainingTime returns the remaining time until the SyncBucket resets in milliseconds.
func (bm *SyncBucketManager) RemainingTime(id string) int64 {
	return bm.getOrCreate(id).RemainingTime()
}

// Reset resets a SyncBucket's uses.
func (bm *SyncBucketManager) Reset(id string) {
	bm.getOrCreate(id).Reset()
}

// Purge purges buckets that are no longer needed, as they have reset.
func (bm *SyncBucketManager) Purge() {
	bm.mutex.RLock()
	for id, bucket := range bm.Buckets {
		if bucket.RemainingTime() == 0 {
			bm.mutex.RUnlock()
			bm.mutex.Lock()
			delete(bm.Buckets, id)
			bm.mutex.Unlock()
			bm.mutex.RLock()
		}
	}
	bm.mutex.RUnlock()
}

// Add manually adds a SyncBucket to this SyncBucketManager.
func (bm *SyncBucketManager) Add(id string, bucket *SyncBucket) {
	bm.mutex.Lock()
	bm.Buckets[id] = bucket
	bm.mutex.Unlock()
}

// SaveFile saves this SyncBucketManager to a file as JSON.
func (bm *SyncBucketManager) SaveFile(path string) error {
	bm.mutex.RLock()
	defer bm.mutex.RUnlock()
	return fn.WriteJsonFile(path, bm)
}

// LoadSyncFile loads a SyncBucketManager from a JSON file.
func LoadSyncFile(path string) (*SyncBucketManager, error) {
	var bm *SyncBucketManager
	return bm, fn.ReadJsonFile(path, &bm)
}

func (bm *SyncBucketManager) getOrCreate(id string) *SyncBucket {
	bm.mutex.RLock()
	bucket, ok := bm.Buckets[id]
	bm.mutex.RUnlock()
	if !ok {
		bucket = NewSyncBucket(bm.Limit, bm.ResetAfter)
		bm.Add(id, bucket)
	}
	return bucket
}
