package rl

import (
	"os"
	"path"
	"reflect"
	"testing"
	"time"
)

var syncFilePath = path.Join(os.TempDir(), "test_bucketmanager.dat")

func TestSyncBucketManager(t *testing.T) {
	bm := NewSync(5, 10*time.Minute)

	bucket := bm.Get("123")
	if bucket == nil {
		t.Error("bucket manager should create buckets in all methods")
	}

	bucket = NewSyncBucket(100, 10*time.Minute)
	bm.Add("456", bucket)
	bucket = bm.Get("456")
	if bucket == nil || bucket.Limit != 100 {
		t.Error("bucket manager should allow manually added buckets")
	}

	WriteSyncFile(t, bm)
	ReadSyncFile(t, bm)

	bm.Purge() // should not block
}

func WriteSyncFile(t *testing.T, bm *SyncBucketManager) {
	err := bm.SaveFile(syncFilePath)
	if err != nil {
		t.Error(err)
	}
}

func ReadSyncFile(t *testing.T, bm *SyncBucketManager) {
	loaded, err := LoadSyncFile(syncFilePath)
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(bm, loaded) {
		t.Error("bucket managers are not the same")
	}
}
