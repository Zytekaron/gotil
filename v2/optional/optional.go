package optional

type Optional[T any] struct {
	value   *T
	present bool
}

func Empty[T any]() Optional[T] {
	return Optional[T]{
		value:   nil,
		present: false,
	}
}

func Of[T any](t T) Optional[T] {
	return Optional[T]{
		value:   &t,
		present: true,
	}
}

func OfPointer[T any](t *T) Optional[T] {
	return Optional[T]{
		value:   t,
		present: t != nil,
	}
}

func (o Optional[T]) IsPresent() bool {
	return o.present
}

func (o Optional[T]) IfPresent(callback func(T)) {
	if o.present {
		callback(*o.value)
	}
}

func (o Optional[T]) Get() T {
	return *o.value
}

func (o Optional[T]) GetOrZero() T {
	if o.present {
		return *o.value
	}
	var zero T
	return zero
}

func (o Optional[T]) OrElseZero() T {
	if o.present {
		return *o.value
	}
	var zero T
	return zero
}

func (o Optional[T]) OrElse(other T) T {
	if o.present {
		return *o.value
	}
	return other
}

func (o Optional[T]) OrElseGet(other func() T) T {
	if o.present {
		return *o.value
	}
	return other()
}
