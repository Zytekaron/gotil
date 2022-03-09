package random

import (
	. "github.com/zytekaron/gotil/random"
	"testing"
)

func TestNewRandomizer(t *testing.T) {
	rand := NewRandomizer[int]()
	err := rand.Add(10, 0)
	if err != nil {
		t.Error("error adding 1 to randomizer:", err.Error())
	}

	err = rand.AddElement(&RandomizerElement[int]{Weight: 20, Result: 1})
	if err != nil {
		t.Error("error adding 2 to randomizer:", err.Error())
	}

	err = rand.AddMany([]RandomizerElement[int]{
		{30, 2},
		{40, 3},
	})
	if err != nil {
		t.Error("error adding 3, 4 to randomizer:", err.Error())
	}

	rand.Prepare()

	for i := 0; i < 1024; i++ {
		_, err = rand.Sample()
		if err != nil {
			t.Error(err)
		}
	}
}

func TestNewSecureRandomizer(t *testing.T) {
	rand := NewSecureRandomizer[int]()
	err := rand.Add(10, 0)
	if err != nil {
		t.Error("error adding 1 to randomizer:", err.Error())
	}

	err = rand.AddElement(&RandomizerElement[int]{Weight: 20, Result: 1})
	if err != nil {
		t.Error("error adding 2 to randomizer:", err.Error())
	}

	err = rand.AddMany([]RandomizerElement[int]{
		{30, 2},
		{40, 3},
	})
	if err != nil {
		t.Error("error adding 3, 4 to randomizer:", err.Error())
	}

	rand.Prepare()

	for i := 0; i < 1024; i++ {
		_, err = rand.Sample()
		if err != nil {
			t.Error(err)
		}
	}
}
