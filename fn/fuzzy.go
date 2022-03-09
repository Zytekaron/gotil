package fn

import "golang.org/x/exp/slices"

// Fuzzy tests if a given haystack contains a given needle.
//
// If the needle can be found one character at a time in
// the haystack, even if the haystack has extra characters
// omitted in the needle, this function returns true.
func Fuzzy(needle, haystack string) bool {
	nl := len(needle)
	hl := len(haystack)

	if nl >= hl {
		return needle == haystack
	}

	needleRunes := []rune(needle)
	haystackRunes := []rune(haystack)
outer:
	for i, j := 0, 0; i < nl; i++ {
		nch := needleRunes[i]
		for j < hl {
			if haystackRunes[j] == nch {
				j++
				continue outer
			}
			j++
		}
		return false
	}
	return true
}

// FuzzySlice tests if a given haystack contains a given needle.
//
// If the needle can be found one byte at a time in
// the haystack, even if the haystack has extra bytes
// omitted in the needle, this function returns true.
func FuzzySlice[T comparable](needle, haystack []T) bool {
	nl := len(needle)
	hl := len(haystack)

	if nl >= hl {
		return slices.Equal(needle, haystack)
	}

outer:
	for i, j := 0, 0; i < nl; i++ {
		nch := needle[i]
		for j < hl {
			if haystack[j] == nch {
				j++
				continue outer
			}
			j++
		}
		return false
	}
	return true
}
