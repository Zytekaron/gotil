package fn

// Test if a given haystack contains a given needle.
//
// If a the needle can be found one character at a time (sequentially)
// in the haystack, even if they haystack has extra characters
// omitted in the needle, this function returns true.
func Fuzzy(needle, haystack []rune) bool {
	nl := len(needle)
	hl := len(haystack)

	if nl > hl {
		return false
	}

	if nl == hl {
		return eq(needle, haystack)
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

func FuzzyString(needle, haystack string) bool {
	return Fuzzy([]rune(needle), []rune(haystack))
}

func eq(a, b []rune) bool {
	if len(a) != len(b) {
		return false
	}
	for i, e := range a {
		if e != b[i] {
			return false
		}
	}
	return true
}
