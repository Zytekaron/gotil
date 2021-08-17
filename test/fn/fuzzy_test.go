package fn

import (
	. "github.com/zytekaron/gotil/fn"
	"testing"
)

func TestFuzzy(t *testing.T) {
	TestFuzzyString(t)
}

func TestFuzzyString(t *testing.T) {
	var needle, haystack string

	needle = "NASA"
	haystack = "National Aeronautics and Space Administration"
	if r := Fuzzy(needle, haystack); !r {
		t.Error("Invalid result for Test 1 - NASA")
	}

	needle = "Hero Gang"
	haystack = "Hermione Granger"
	if r := Fuzzy(needle, haystack); !r {
		t.Error("Invalid result for Test 2 - Hero Gang")
	}

	needle = "Zyk"
	haystack = "Zytekaron"
	if r := Fuzzy(needle, haystack); !r {
		t.Error("Invalid result for Test 3 - Zyk")
	}

	needle = "Zyke"
	haystack = "Zytekaron"
	if r := Fuzzy(needle, haystack); r {
		t.Error("Invalid result for Test 4 - Zyke")
	}

	needle = "13579"
	haystack = "12345678"
	if r := Fuzzy(needle, haystack); r {
		t.Error("Invalid result for Test 5 - 13579")
	}

	needle = "23"
	haystack = "4321"
	if r := Fuzzy(needle, haystack); r {
		t.Error("Invalid result for Test 6 - 23")
	}
}
