package optional

// Map maps an optional from one type to another using a transform function.
//
// If the optional is not present, no operation occurs on the internal value.
func Map[T, U any](t Optional[T], transform func(T) U) Optional[U] {
	if t.present {
		val := transform(*t.value)
		return Of[U](val)
	}
	return Empty[U]()
}

// FlatMap maps an optional from one type to another using a transform
// function which returns another optional instead of the output type.
// This allows for more complex operations regarding whether the result is present.
//
// If the optional is not present, no operation occurs on the internal value.
func FlatMap[T, U any](t Optional[T], transform func(T) Optional[U]) Optional[U] {
	if t.present {
		return transform(*t.value)
	}
	return Empty[U]()
}

// Equals returns whether two optionals are both empty, or both present with the same value.
func Equals[T comparable](a, b Optional[T]) bool {
	if a.present && b.present {
		return *a.value == *b.value
	}
	return !a.present && !b.present
}
