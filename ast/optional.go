package ast

type Optional[T any] struct {
	value T
	isSet bool
}

func (o *Optional[T]) Set(value T) {
	o.isSet = true
	o.value = value
}

func (o *Optional[T]) Unset() {
	o.isSet = false
	var zero T
	o.value = zero
}

func (o *Optional[T]) IsSet() bool {
	return o.isSet
}

func (o *Optional[T]) Value() T {
	if !o.isSet {
		panic("cannot get value of optional when a value is not set")
	}

	return o.value
}
