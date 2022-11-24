package lexer

import (
	"github.com/JosephNaberhaus/agnostic/language/token"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func failsAtPosConsumer(pos int) consumer {
	return func(_ runes) (runes, []token.Token, error) {
		return runes{}, nil, Error{pos: pos}
	}
}

func returnsTokensConsumer(tokens []token.Token) consumer {
	return func(r runes) (runes, []token.Token, error) {
		return r, tokens, nil
	}
}

func TestFirst(t *testing.T) {
	expected := []token.Token{
		{Type: token.Add, Text: "+"},
		{Type: token.Subtract, Text: "-"},
	}

	consumer := first(
		failsAtPosConsumer(100),
		returnsTokensConsumer(expected),
		failsAtPosConsumer(100),
	)

	r, tokens, err := consumer(newRunes("testing"))
	assert.Equal(t, []rune("testing"), r.text)
	assert.Equal(t, expected, tokens)
	assert.NoError(t, err)
}

func TestFirst_Error_ReturnsTheFurthestError(t *testing.T) {
	consumer := first(
		failsAtPosConsumer(200),
		failsAtPosConsumer(4000),
		failsAtPosConsumer(100),
	)

	_, _, err := consumer(newRunes("testing"))
	var lexerError Error
	require.ErrorAs(t, err, &lexerError)
	assert.Equal(t, 4000, lexerError.pos)
}

func TestInOrder(t *testing.T) {
	firstExpected := []token.Token{
		{Type: token.Add, Text: "+"},
		{Type: token.Subtract, Text: "-"},
	}
	secondExpected := []token.Token{
		{Type: token.Multiply, Text: "*"},
	}
	expected := append(firstExpected, secondExpected...)

	consumer := inOrder(
		returnsTokensConsumer(firstExpected),
		returnsTokensConsumer(secondExpected),
	)
	r, tokens, err := consumer(newRunes("testing"))
	assert.Equal(t, []rune("testing"), r.text)
	assert.Equal(t, expected, tokens)
	assert.NoError(t, err)
}

func TestInOrder_Error_ReturnsFirstError(t *testing.T) {
	consumer := inOrder(
		returnsTokensConsumer(nil),
		failsAtPosConsumer(100),
		failsAtPosConsumer(104),
	)

	_, _, err := consumer(newRunes("testing"))
	var lexerError Error
	require.ErrorAs(t, err, &lexerError)
	assert.Equal(t, 100, lexerError.pos)
}

func TestOptional(t *testing.T) {
	expected := []token.Token{
		{Type: token.Add, Text: "+"},
		{Type: token.Subtract, Text: "-"},
	}

	consumer := optional(returnsTokensConsumer(expected))

	r, tokens, err := consumer(newRunes("testing"))
	assert.Equal(t, []rune("testing"), r.text)
	assert.Equal(t, expected, tokens)
	assert.NoError(t, err)
}

func TestOptional_Error(t *testing.T) {
	consumer := optional(failsAtPosConsumer(100))

	r, tokens, err := consumer(newRunes("testing"))
	assert.Equal(t, []rune("testing"), r.text)
	assert.Empty(t, tokens)
	assert.NoError(t, err)
}
