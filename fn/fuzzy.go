package fn

// Fuzzy tests if a given haystack contains a given needle.
//
// If a the needle can be found one character at a time (sequentially)
// in the haystack, even if they haystack has extra characters
// omitted in the needle, this function returns true.
func Fuzzy(needle, haystack string) bool {
	nl := len(needle)
	hl := len(haystack)

	if nl > hl {
		return false
	}

	if nl == hl {
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
