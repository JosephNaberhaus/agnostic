package lexer

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateError(t *testing.T) {
	const expectedPos = 42
	const expectedMessage = "test message"

	r := runes{numConsumed: expectedPos}
	err := createError(r, expectedMessage)

	assert.Equal(t, expectedPos, err.pos)
	assert.Equal(t, expectedMessage, err.message)
}

func TestTakeFurthest(t *testing.T) {
	first := Error{
		pos:     4,
		message: "first",
	}
	second := Error{
		pos:     4000,
		message: "second",
	}

	assert.Equal(t, second, takeFurthest(first, second))
	assert.Equal(t, second, takeFurthest(second, first))
}

func TestTakeFurthest_HandlesUnexpectedErrorTypes(t *testing.T) {
	err := Error{
		pos:     4,
		message: "first",
	}
	unexpected := errors.New("unexpected error type")

	assert.Equal(t, err, takeFurthest(unexpected, err))
	assert.Equal(t, err, takeFurthest(err, unexpected))
	assert.Equal(t, unexpected, takeFurthest(unexpected, unexpected))
}
