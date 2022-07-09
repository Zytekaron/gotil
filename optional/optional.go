package optional

// Optional is a type which can be used as a safer alternative to pointers and
// automatic dereferencing. Optionals can represent a present or absent value,
// like pointers, but may also contain a pointer value, allowing for multiple
// valid present states: present but nil, and present with a value. This can be
// useful in serialization. Optional is directly serializable to and from JSON.
type Optional[T any] struct {
	value   *T
	present bool
}

// Empty creates an optional which is not present and has no value.
func Empty[T any]() Optional[T] {
	return Optional[T]{
		value:   nil,
		present: false,
	}
}

// Of creates a present optional with the provided value.
func Of[T any](t T) Optional[T] {
	return Optional[T]{
		value:   &t,
		present: true,
	}
}

// OfPointer creates an optional based on the passed pointer, returning an
// empty optional if the pointer is nil, and a present optional otherwise.
func OfPointer[T any](t *T) Optional[T] {
	return Optional[T]{
		value:   t,
		present: t != nil,
	}
}

// IsPresent returns whether the optional value is present.
func (o Optional[T]) IsPresent() bool {
	return o.present
}

// IfPresent calls a callback with the value only if it is present.
func (o Optional[T]) IfPresent(callback func(T)) {
	if o.present {
		callback(*o.value)
	}
}

// Get gets the optional value. If the optional value is not present,
// this will panic due to a nil pointer dereference. Check IsPresent
// before calling Get if the presence of the value is not guaranteed.
func (o Optional[T]) Get() T {
	return *o.value
}

// GetOrZero returns the optional value if it is present, otherwise
// returning the zero value for the optional type.
//
// This method will be removed in the future in favor of OrElseZero.
// OrElseZero is strictly equivalent to this method.
func (o Optional[T]) GetOrZero() T {
	if o.present {
		return *o.value
	}
	var null T
	return null
}

// OrElseZero returns the optional value if it is present,
// otherwise returning the zero value for the optional type.
func (o Optional[T]) OrElseZero() T {
	if o.present {
		return *o.value
	}
	var null T
	return null
}

// OrElse returns the optional value if it is present,
// otherwise returning the provided default value.
func (o Optional[T]) OrElse(other T) T {
	if o.present {
		return *o.value
	}
	return other
}

// OrElseGet returns the optional value if it is present,
// otherwise calling the default getter value and returning that.
func (o Optional[T]) OrElseGet(other func() T) T {
	if o.present {
		return *o.value
	}
	return other()
}
