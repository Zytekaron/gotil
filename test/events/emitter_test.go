package events

import (
	"fmt"
	. "github.com/zytekaron/gotil/events"
	"testing"
	"time"
)

func TestGlobalEvent(t *testing.T) {
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

	var cancel func()
	cancel = emitter.On("c", func() string { return "c (cancelled)" })
	cancel()
	cancel = emitter.On(All, func() string { return "catch-all (cancelled)" })
	cancel()

	// calls a (a1), a (a2), and * (catch-all)
	ch := emitter.Emit("a")
	values := make([]interface{}, 0)
	for val := range ch {
		values = append(values, val[0])
	}
	if len(values) != 3 {
		t.Error("Expected 3 values, got", len(values), values)
	}

	// calls * (catch-all)
	// (c event was cancelled)
	ch = emitter.Emit("c")
	values = make([]interface{}, 0)
	for val := range ch {
		values = append(values, val)
	}
	if len(values) != 1 {
		t.Error("Expected 1 values, got", len(values), values)
	}

	<-time.After(time.Millisecond)
}
