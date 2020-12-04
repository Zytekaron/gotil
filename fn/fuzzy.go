package fn

// Test if a given haystack contains a given needle.
//
// If a the needle can be found one character at a time (sequentially)
// in the haystack, even if they haystack has extra characters
// omitted in the needle, this function returns true.
//
// fuzzy("NASA", "National Aeronautics and Space Administration")
// // -> true (N, A, S, and A are found sequentially in the string)
// fuzzy("135", "12345") // true (2 and 4 are omitted, this is ok)
// fuzzy("531", "12345") // false (wrong order)
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
