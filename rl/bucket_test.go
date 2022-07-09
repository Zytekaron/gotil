package rl

import (
	"testing"
	"time"
)

func TestBucket(t *testing.T) {
	bucket := NewBucket(5, 10*time.Minute)

	ok := bucket.Draw(3)
	if !ok {
		t.Error("failed to draw 3 of 5")
	}

	n := bucket.DrawMax(5)
	if n != 2 {
		t.Error("failed to draw remaining 2 of 5")
	}

	ok = bucket.CanDraw(1)
	if ok {
		t.Error("should not be able to draw 1 from empty bucket")
	}

	bucket.Reset()
	ok = bucket.CanDraw(5)
	if !ok {
		t.Error("should be able to draw 5 from full bucket")
	}

	n = bucket.ForceDraw(10)
	if n != -5 {
		t.Error("force draw should return remaining uses, even if negative")
	}

	n = bucket.DrawMax(5)
	if n != 0 {
		t.Error("bucket should return 0 drawn tokens due to a previous overdraw")
	}
	n = bucket.RemainingUses()
	if n != -5 {
		t.Error("expected remaining uses to be -5 due to previous overdraw")
	}
}
