package code

// TODO move this somewhere else entirely

// Stack is an implementation of a stack data structure that contains code.Node elements.
//
// The zero value of a Stack is ready to be used.
type Stack []Node

// Push appends a value to the end of the Stack,
func (s *Stack) Push(value Node) {
	*s = append(*s, value)
}

// Pop removes the last value from the Stack and returns it.
func (s *Stack) Pop() Node {
	lastValue := s.Peek()
	*s = (*s)[:len(*s)-1]
	return lastValue
}

// Peek returns the last element of the Stack without removing it from the Stack.
func (s *Stack) Peek() Node {
	return (*s)[len(*s)-1]
}

// PeekParent returns the second to last element of the Stack without removing it from the Stack.
func (s *Stack) PeekParent() Node {
	return (*s)[len(*s)-2]
}

// IsEmpty returns whether the Stack is empty.
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// IsNotEmpty returns whether the Stack is not empty
func (s *Stack) IsNotEmpty() bool {
	return !s.IsEmpty()
}

// Copy creates a shallow Copy of the Stack.
//
// The Copy will use a different backing array, so it can be modified without effecting the original stack.
func (s *Stack) Copy() Stack {
	var newStack Stack
	for _, item := range *s {
		newStack.Push(item)
	}
	return newStack
}

func FirstOfType[T Node](stack Stack) (T, bool) {
	stack = stack.Copy()
	for stack.IsNotEmpty() {
		converted, ok := stack.Pop().(T)
		if ok {
			return converted, true
		}
	}

	var zero T
	return zero, false
}
