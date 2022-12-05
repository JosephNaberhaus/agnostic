package mappers

import (
	"github.com/JosephNaberhaus/agnostic/code"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCodeStack_Push(t *testing.T) {
	expectedValues := []code.Node{
		&code.LiteralInt32{Value: 42},
		&code.LiteralString{Value: "hello"},
		&code.LiteralString{Value: "world"},
	}

	var stack code.Stack
	for _, expectedValue := range expectedValues {
		stack.push(expectedValue)
	}

	assert.Equal(t, expectedValues, []code.Node(stack))
}

func TestCodeStack_Pop(t *testing.T) {
	expectedValues := []code.Node{
		&code.LiteralInt32{Value: 42},
		&code.LiteralString{Value: "hello"},
		&code.LiteralString{Value: "world"},
	}

	stack := code.Stack(expectedValues)
	for i := len(expectedValues) - 1; i >= 0; i-- {
		assert.Equal(t, expectedValues[i], stack.pop())
	}

	assert.Empty(t, stack)
}

func TestCodeStack_Peek(t *testing.T) {
	expectedValue := &code.LiteralInt32{Value: 42}
	stack := code.Stack{expectedValue}

	assert.Equal(t, stack.peek(), expectedValue)
	assert.Len(t, stack, 1)
}

func TestCodeStack_Peek_Empty(t *testing.T) {
	var stack code.Stack

	assert.Panics(t, func() {
		stack.peek()
	})
}

func TestCodeStack_PeekParent(t *testing.T) {
	expectedValue := &code.LiteralInt32{Value: 42}
	stack := code.Stack{
		expectedValue,
		&code.LiteralString{Value: "hello"},
	}

	assert.Equal(t, stack.peekParent(), expectedValue)
	assert.Len(t, stack, 2)
}

func TestCodeStack_PeekParent_OneElement(t *testing.T) {
	stack := code.Stack{&code.LiteralString{Value: "hello"}}

	assert.Panics(t, func() {
		stack.peekParent()
	})
}

func TestCodeStack_IsEmpty(t *testing.T) {
	var stack code.Stack
	assert.True(t, stack.isEmpty())

	stack = code.Stack{&code.LiteralInt32{Value: 42}}
	assert.False(t, stack.isEmpty())
}

func TestCodeStack_IsNotEmpty(t *testing.T) {
	var stack code.Stack
	assert.False(t, stack.isNotEmpty())

	stack = code.Stack{&code.LiteralInt32{Value: 42}}
	assert.True(t, stack.isNotEmpty())
}

func TestCodeStack_Copy(t *testing.T) {
	originalStack := code.Stack{
		&code.LiteralInt32{Value: 42},
		&code.LiteralString{Value: "hello"},
		&code.LiteralString{Value: "world"},
	}

	copiedStack := originalStack.copy()
	assert.Equal(t, originalStack, copiedStack)

	// Modify the copiedStack and verify that the originalStack is unchanged
	copiedStack.pop()
	copiedStack.push(&code.LiteralString{Value: "new world!"})
	assert.NotEqual(t, copiedStack, originalStack)
}

func TestCodeStack_FirstOfType(t *testing.T) {
	var stack code.Stack

	_, ok := code.firstOfType[*code.LiteralInt32](stack)
	require.False(t, ok)

	stack.push(&code.LiteralInt32{Value: 42})
	stack.push(&code.LiteralString{Value: "hello"})
	stack.push(&code.LiteralInt32{Value: 84})

	result, ok := code.firstOfType[*code.LiteralInt32](stack)
	require.True(t, ok)
	assert.Equal(t, result.Value, int32(84))
}
