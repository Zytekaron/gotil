package maths

import (
	"testing"
)

func TestCollisions(t *testing.T) {
	bday := Collisions(365, 0.5)
	if int(bday) < 22 || int(bday) > 23 {
		t.Error("expected roughly 23, got", bday)
	}

	min := Collisions(100, 0)
	if min != 0 {
		t.Error("expected 0, got", min)
	}
}

func TestFactorial(t *testing.T) {
	five := Factorial(5)
	if five != 120 {
		t.Error("expected 120, got", five)
	}

	fifteen := Factorial(15)
	if fifteen != 1307674368000 {
		t.Error("expected 1307674368000, got", fifteen)
	}
}

func TestRoot(t *testing.T) {
	five := Root(625, 4)
	if five != 5 {
		t.Error("expected 5, got", five)
	}

	twelve := Root(248832, 5)
	if twelve != 12 {
		t.Error("expected 12, got", twelve)
	}
}

func TestRound(t *testing.T) {
	const num = float64(123_456_789)

	if Round(num, 6) != 123_000_000 {
		t.Error("expected 123,000,000, got", Round(num, 6))
	}
	if Round(num, 3) != 123_457_000 {
		t.Error("expected 123,457,000, got", Round(num, 3))
	}
}

func TestFloor(t *testing.T) {
	const num = float64(123_456_789)

	if Floor(num, 6) != 123_000_000 {
		t.Error("expected 123,000,000, got", Floor(num, 6))
	}
	if Floor(num, 3) != 123_456_000 {
		t.Error("expected 123,456,000, got", Floor(num, 3))
	}
}

func TestCeil(t *testing.T) {
	const num = float64(123_456_789)

	if Ceil(num, 6) != 124_000_000 {
		t.Error("expected 124,000,000, got", Ceil(num, 6))
	}
	if Ceil(num, 3) != 123_457_000 {
		t.Error("expected 123,457,000, got", Ceil(num, 3))
	}
}

func TestRoundDecimal(t *testing.T) {
	const num = 1.123456789

	if RoundDecimal(num, 3) != 1.123 {
		t.Error("expected 1.123, got", RoundDecimal(num, 3))
	}
	if RoundDecimal(num, 5) != 1.12346 {
		t.Error("expected 1.12346, got", RoundDecimal(num, 5))
	}
}

func TestFloorDecimal(t *testing.T) {
	const num = 1.123456789

	if FloorDecimal(num, 3) != 1.123 {
		t.Error("expected 1.123, got", FloorDecimal(num, 3))
	}
	if FloorDecimal(num, 5) != 1.12345 {
		t.Error("expected 1.12345, got", FloorDecimal(num, 5))
	}
}

func TestCeilDecimal(t *testing.T) {
	const num = 1.123456789

	if CeilDecimal(num, 3) != 1.124 {
		t.Error("expected 1.124, got", CeilDecimal(num, 3))
	}
	if CeilDecimal(num, 5) != 1.12346 {
		t.Error("expected 1.12346, got", CeilDecimal(num, 5))
	}
}
