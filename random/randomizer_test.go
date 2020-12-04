package random

import (
	"testing"
)

func TestNewRandomizer(t *testing.T) {
	rand := NewRandomizer()
	err := rand.Add(10, 0)
	if err != nil {
		t.Error("error occurred whilst adding 1 to randomizer:", err.Error())
	}

	err = rand.AddElement(&RandomizerElement{20, 1})
	if err != nil {
		t.Error("error occurred whilst adding 2 to randomizer:", err.Error())
	}

	err = rand.AddMany([]RandomizerElement{
		{30, 2},
		{40, 3},
	})
	if err != nil {
		t.Error("error occurred whilst adding 3, 4 to randomizer:", err.Error())
	}

	rand.Prepare()

	for i := 0; i < 1000000; i++ {
		_, err := rand.Sample()
		if err != nil {
			t.Error(err)
		}
	}
}

func TestNewSecureRandomizer(t *testing.T) {
	rand := NewSecureRandomizer()
	err := rand.Add(10, 0)
	if err != nil {
		t.Error("error occurred whilst adding 1 to randomizer:", err.Error())
	}

	err = rand.AddElement(&RandomizerElement{20, 1})
	if err != nil {
		t.Error("error occurred whilst adding 2 to randomizer:", err.Error())
	}

	err = rand.AddMany([]RandomizerElement{
		{30, 2},
		{40, 3},
	})
	if err != nil {
		t.Error("error occurred whilst adding 3, 4 to randomizer:", err.Error())
	}

	rand.Prepare()

	for i := 0; i < 1000000; i++ {
		_, err := rand.Sample()
		if err != nil {
			t.Error(err)
		}
	}
}