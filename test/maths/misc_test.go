package maths

import (
	. "github.com/zytekaron/gotil/maths"
	"testing"
)

func TestCollisions(t *testing.T) {
	bday := Collisions(365, 0.5)
	if int(bday) < 22 || int(bday) > 23 {
		t.Error("birthday paradox result is not roughly 23; instead,", bday, "=", bday)
	}

	min := Collisions(100, 0)
	if min != 0 {
		t.Error("Collisions(100, 0) should be 0;", min)
	}
}

func TestFactorial(t *testing.T) {
	five := Factorial(5)
	if five != 120 {
		t.Error("Factorial(5) should be 120;", five)
	}

	fifteen := Factorial(15)
	if fifteen != 1307674368000 {
		t.Error("Factorial(15) should be 1307674368000;", fifteen)
	}
}

func TestRoot(t *testing.T) {
	five := Root(625, 4)
	if five != 5 {
		t.Error("Root(625, 4) should be 5;", five)
	}

	twelve := Root(248832, 5)
	if twelve != 12 {
		t.Error("Root(248832, 5) should be 12;", twelve)
	}
}
