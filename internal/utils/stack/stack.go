package stack

type Stack[T any] []T

// Push adds an element to the top of the stack.
func (s *Stack[T]) Push(value T) {
	*s = append(*s, value)
}

// Pop removes the top element of the stack and returns it.
func (s *Stack[T]) Pop() T {
	value := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return value
}

// Peek returns the top element of the stack without removing.
func (s *Stack[T]) Peek() T {
	return (*s)[len(*s)-1]
}

// Copy creates a shallow copy of the stack.
func (s *Stack[T]) Copy() Stack[T] {
	newStack := make(Stack[T], len(*s))
	copy(newStack, *s)
	return newStack
}
