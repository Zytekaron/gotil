package fn

import (
	. "github.com/zytekaron/gotil/v2/fn"
	"testing"
)

func TestFuzzy(t *testing.T) {
	TestFuzzyString(t)
}

func TestFuzzyString(t *testing.T) {
	var needle, haystack string

	needle = "NASA"
	haystack = "National Aeronautics and Space Administration"
	if !Fuzzy(needle, haystack) {
		t.Error("invalid result for:", needle, haystack)
	}

	needle = "Hero Gang"
	haystack = "Hermione Granger"
	if !Fuzzy(needle, haystack) {
		t.Error("invalid result for:", needle, haystack)
	}

	needle = "Zyk"
	haystack = "Zytekaron"
	if !Fuzzy(needle, haystack) {
		t.Error("invalid result for:", needle, haystack)
	}

	needle = "Zyke"
	haystack = "Zytekaron"
	if Fuzzy(needle, haystack) {
		t.Error("invalid result for:", needle, haystack)
	}

	needle = "13579"
	haystack = "12345678"
	if Fuzzy(needle, haystack) {
		t.Error("invalid result for:", needle, haystack)
	}

	needle = "23"
	haystack = "4321"
	if Fuzzy(needle, haystack) {
		t.Error("invalid result for:", needle, haystack)
	}
}

func TestFuzzySlice(t *testing.T) {
	haystack := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	needle := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	if !FuzzySlice(needle, haystack) {
		t.Error("invalid result for:", needle, haystack)
	}

	needle = []int{1, 2, 4, 8}
	if !FuzzySlice(needle, haystack) {
		t.Error("invalid result for:", needle, haystack)
	}

	needle = []int{0, 5, 9}
	if !FuzzySlice(needle, haystack) {
		t.Error("invalid result for:", needle, haystack)
	}

	needle = []int{1, 4, 2, 8}
	if FuzzySlice(needle, haystack) {
		t.Error("invalid result for:", needle, haystack)
	}

	needle = []int{15, 5, 8, 15, 20}
	if FuzzySlice(needle, haystack) {
		t.Error("invalid result for:", needle, haystack)
	}

	needle = []int{1, 2, 4, 14}
	if FuzzySlice(needle, haystack) {
		t.Error("invalid result for:", needle, haystack)
	}
}
