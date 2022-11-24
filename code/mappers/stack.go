package mappers

import "github.com/JosephNaberhaus/agnostic/code"

// codeStack is an implementation of a stack data structure that contains code.Node elements.
//
// The zero value of a codeStack is ready to be used.
type codeStack []code.Node

// push appends a value to the end of the codeStack,
func (c *codeStack) push(value code.Node) {
	*c = append(*c, value)
}

// pop removes the last value from the codeStack and returns it.
func (c *codeStack) pop() code.Node {
	lastValue := c.peek()
	*c = (*c)[:len(*c)-1]
	return lastValue
}

// peek returns the last element of the codeStack without removing it from the codeStack.
func (c *codeStack) peek() code.Node {
	return (*c)[len(*c)-1]
}

// peekParent returns the second to last element of the codeStack without removing it from the codeStack.
func (c *codeStack) peekParent() code.Node {
	return (*c)[len(*c)-2]
}

// isEmpty returns whether the codeStack is empty.
func (c *codeStack) isEmpty() bool {
	return len(*c) == 0
}

// isEmpty returns whether the codeStack is not empty
func (c *codeStack) isNotEmpty() bool {
	return !c.isEmpty()
}

// copy creates a shallow copy of the codeStack.
//
// The copy will use a different backing array, so it can be modified without effecting the original stack.
func (c *codeStack) copy() codeStack {
	var newStack codeStack
	for _, item := range *c {
		newStack.push(item)
	}
	return newStack
}

func firstOfType[T code.Node](stack codeStack) (T, bool) {
	stack = stack.copy()
	for stack.isNotEmpty() {
		converted, ok := stack.pop().(T)
		if ok {
			return converted, true
		}
	}

	var zero T
	return zero, false
}
