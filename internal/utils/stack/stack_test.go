package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	stack := Stack[int]{}

	stack.Push(42)
	stackCopy := stack.Copy()
	stack.Push(84)

	assert.Equal(t, 84, stack.Peek())
	assert.Equal(t, 84, stack.Pop())
	assert.Equal(t, 42, stack.Pop())

	// Trying to Pop on an empty stack is not supported.
	assert.Panics(t, func() {
		stack.Pop()
	})

	// The copy should have been unaffected by the operations on the original.
	assert.Equal(t, 42, stackCopy.Pop())
}
