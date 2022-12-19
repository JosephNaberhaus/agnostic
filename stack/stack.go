package stack

type Stack[T any] []T

func (s *Stack[T]) Push(value T) {
	*s = append(*s, value)
}

func (s *Stack[T]) Pop() T {
	value := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return value
}

func (s *Stack[T]) Peek() T {
	return (*s)[len(*s)-1]
}

func (s *Stack[T]) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack[T]) IsNotEmpty() bool {
	return len(*s) != 0
}

func (s *Stack[T]) Copy() Stack[T] {
	newStack := make(Stack[T], len(*s))
	copy(newStack, *s)
	return newStack
}
