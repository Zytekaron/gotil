package rl

import (
	"fmt"
	"github.com/zytekaron/gotil/v2/rl"
	"testing"
	"time"
)

func TestBucket(t *testing.T) {
	bucket := rl.NewBucket(5, 10*time.Minute)
	fmt.Println(bucket)

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
}
