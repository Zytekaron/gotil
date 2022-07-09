package rl

import (
	"os"
	"path"
	"reflect"
	"testing"
	"time"
)

var filePath = path.Join(os.TempDir(), "test_bucketmanager.dat")

func TestBucketManager(t *testing.T) {
	bm := New(5, 10*time.Minute)

	bucket := bm.Get("123")
	if bucket == nil {
		t.Error("bucket manager should create buckets in all methods")
	}

	bucket = NewBucket(100, 10*time.Minute)
	bm.Add("456", bucket)
	bucket = bm.Get("456")
	if bucket == nil || bucket.Limit != 100 {
		t.Error("bucket manager should allow manually added buckets")
	}

	WriteFile(t, bm)
	ReadFile(t, bm)
}

func WriteFile(t *testing.T, bm *BucketManager) {
	err := bm.SaveFile(filePath)
	if err != nil {
		t.Error(err)
	}
}

func ReadFile(t *testing.T, bm *BucketManager) {
	loaded, err := LoadFile(filePath)
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(bm, loaded) {
		t.Error("bucket managers are not the same")
	}
}
