package rl

import (
	. "github.com/zytekaron/gotil/rl"
	"reflect"
	"testing"
	"time"
)

var path = "C:\\Users\\Zytekaron\\AppData\\Local\\Temp\\rl.dat"

// todo test Load Save

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
	err := bm.SaveFile(path)
	if err != nil {
		t.Error(err)
	}
}

func ReadFile(t *testing.T, bm *BucketManager) {
	loaded, err := LoadFile(path)
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(bm, loaded) {
		t.Error("bucket managers are not the same")
	}
}
