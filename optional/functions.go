package optional

func Map[T, U any](t Optional[T], transform func(T) U) Optional[U] {
	if t.present {
		val := transform(*t.value)
		return Of[U](val)
	}
	return Empty[U]()
}

func FlatMap[T, U any](t Optional[T], transform func(T) Optional[U]) Optional[U] {
	if t.present {
		return transform(*t.value)
	}
	return Empty[U]()
}

func Equals[T comparable](a, b Optional[T]) bool {
	if a.present && b.present {
		return *a.value == *b.value
	}
	return !a.present && !b.present
}
