package optional

import (
	. "github.com/zytekaron/gotil/v2/optional"
	"strconv"
	"testing"
)

func TestCreation(t *testing.T) {
	present := Of(123)
	if !present.IsPresent() {
		t.Error("present optional fails presence check")
	}

	empty := Empty[int]()
	if empty.IsPresent() {
		t.Error("empty optional fails presence check")
	}

	presentPtr := new(int)
	present = OfPointer(presentPtr)
	if !present.IsPresent() {
		t.Error("present pointer optional fails presence check")
	}

	var emptyPtr *int
	present = OfPointer(emptyPtr)
	if empty.IsPresent() {
		t.Error("empty pointer optional fails presence check")
	}
}

func TestStateMethods(t *testing.T) {
	present := Of(123)
	if present.Get() != 123 {
		t.Error("expected 123 but got", present.Get())
	}
	if present.GetOrZero() != 123 {
		t.Error("expected 123 but got", present.GetOrZero())
	}
	if present.OrElse(0) != 123 {
		t.Error("expected 123 but got", present.OrElse(0))
	}

	var val int
	present.IfPresent(func(i int) {
		val = i
	})
	if val != 123 {
		t.Error("expected 123 but got", val)
	}

	empty := Empty[int]()
	if empty.GetOrZero() != 0 {
		t.Error("expected 0 but got", empty.GetOrZero())
	}

	val = 0
	empty.IfPresent(func(i int) {
		val = i
	})
	if val != 0 {
		t.Error("expected 0 but got", val)
	}
}

func TestMap(t *testing.T) {
	opt := Of(123)
	out := Map(opt, strconv.Itoa)
	if !out.IsPresent() {
		t.Error("mapped optional fails presence check")
	}
	if out.Get() != "123" {
		t.Error("mapped optional fails value check")
	}
}

func TestEquals(t *testing.T) {
	a1 := Of(123)
	a2 := Of(123)
	b := Of(456)
	c1 := Empty[int]()
	c2 := Empty[int]()

	if !Equals(a1, a2) {
		t.Error("expected a1 to equal a2")
	}
	if !Equals(c1, c2) {
		t.Error("expected c1 to equal c2")
	}
	if Equals(a1, b) {
		t.Error("expected a1 to not equal b")
	}
	if Equals(a1, c1) {
		t.Error("expected a1 to not equal c1")
	}
	if Equals(b, c1) {
		t.Error("expected b to not equal c1")
	}
}

func TestFunctions(t *testing.T) {
	opt := Of(123)
	out1 := Map(opt, strconv.Itoa)
	if !out1.IsPresent() {
		t.Error("expected out1 to be present")
	}
	if out1.Get() != "123" {
		t.Error("expected out1 to be '123'")
	}

	out2 := FlatMap(opt, func(i int) Optional[string] {
		return Of(strconv.Itoa(i))
	})
	if !out2.IsPresent() {
		t.Error("expected out2 to be present")
	}
	if out2.Get() != "123" {
		t.Error("expected out2 to be '123'")
	}
}
