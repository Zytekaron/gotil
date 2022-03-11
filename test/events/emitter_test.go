package events

import (
	"fmt"
	. "github.com/zytekaron/gotil/v2/events"
	"testing"
	"time"
)

func TestGlobalEvent(*testing.T) {
	emitter := New()
	emitter.On("eventName", func(args ...interface{}) {
		fmt.Println(args...)
	})
	emitter.Emit("something", 1, 2, 3)
}

func TestResponses(t *testing.T) {
	emitter := New()

	emitter.On("a", func() string { return "a1" })
	emitter.On("a", func() string { return "a2" })
	emitter.On("b", func() string { return "b" })

	emitter.On("*", func() string { return "catch-all" })

	// cancel func is called
	emitter.On("c", func() string { return "c (cancelled)" })()
	emitter.On(All, func() string { return "catch-all (cancelled)" })()

	emitter.OnConditional(All, func() bool { return true }, func() string { return "catch-all-predicate" })
	emitter.OnConditional(All, func() bool { return false }, func() string { return "catch-all-predicate (disabled)" })

	// calls a (a1), a (a2), and * (catch-all)
	ch := emitter.Dispatch("a")
	values := make([]interface{}, 0)
	for val := range ch {
		values = append(values, val[0])
	}
	if len(values) != 4 {
		t.Error("expected 3 values, got", len(values), values)
	}

	// calls * (catch-all)
	// (c event was cancelled)
	ch = emitter.Dispatch("c")
	values = make([]interface{}, 0)
	for val := range ch {
		values = append(values, val)
	}
	if len(values) != 2 {
		t.Error("expected 1 value, got", len(values), values)
	}

	emitter.Emit("d")

	<-time.After(time.Millisecond)
}
