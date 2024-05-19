package ast

type Optional[T any] struct {
	value T
	set   bool
}

func OptionalWithValue[T any](value T) Optional[T] {
	return Optional[T]{
		value: value,
		set:   true,
	}
}

func (o Optional[T]) Value() (T, bool) {
	return o.value, o.set
}

func (o Optional[T]) IsSet() bool {
	return o.set
}
