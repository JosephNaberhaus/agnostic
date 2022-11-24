package lexer

import (
	"github.com/JosephNaberhaus/agnostic/language/token"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringConsumer(t *testing.T) {
	consumer := stringConsumer("<this>")

	r, tokens, err := consumer(newRunes("not the expected string"))
	assert.Error(t, err)

	r, tokens, err = consumer(newRunes("<this> not the expected string"))
	assert.Equal(t, []rune(" not the expected string"), r.text)
	assert.Empty(t, tokens)
	assert.NoError(t, err)
}

func TestStringTokenConsumer(t *testing.T) {
	consumer := stringTokenConsumer("<this>", token.Add)

	r, tokens, err := consumer(newRunes("not the expected string"))
	assert.Error(t, err)

	r, tokens, err = consumer(newRunes("<this> not the expected string"))
	assert.Equal(t, []rune(" not the expected string"), r.text)
	assert.Equal(t, []token.Token{{Type: token.Add, Text: "<this>"}}, tokens)
	assert.NoError(t, err)
}

func TestAllWhileTokenConsumer(t *testing.T) {
	trueNTimes := func(n int) func(r rune) bool {
		count := 0
		return func(r rune) bool {
			count++
			return count <= n
		}
	}

	consumer := allWhileTokenConsumer(trueNTimes(4), token.Add, "test")
	r, tokens, err := consumer(newRunes(""))
	assert.Error(t, err)

	consumer = allWhileTokenConsumer(trueNTimes(4), token.Add, "test")
	r, tokens, err = consumer(newRunes("fourfour"))
	assert.Equal(t, []rune("four"), r.text)
	assert.Equal(t, []token.Token{{Type: token.Add, Text: "four"}}, tokens)
	assert.NoError(t, err)
}

func TestAnyWhileTokenConsumer(t *testing.T) {
	alwaysFalse := func(r rune) bool {
		return false
	}
	trueNTimes := func(n int) func(r rune) bool {
		count := 0
		return func(r rune) bool {
			count++
			return count <= n
		}
	}

	consumer := allWhileTokenConsumer(alwaysFalse, token.Add, "test")
	r, tokens, err := consumer(newRunes("four"))
	assert.Error(t, err)

	consumer = allWhileTokenConsumer(trueNTimes(4), token.Add, "test")
	r, tokens, err = consumer(newRunes("fourfour"))
	assert.Equal(t, []rune("four"), r.text)
	assert.Equal(t, []token.Token{{Type: token.Add, Text: "four"}}, tokens)
	assert.NoError(t, err)
}

func TestAlphaTokenConsumer(t *testing.T) {
	consumer := alphaTokenConsumer(token.Add)

	_, _, err := consumer(newRunes("42"))
	assert.Error(t, err)

	r, tokens, err := consumer(newRunes("thing42"))
	assert.Equal(t, []rune("42"), r.text)
	assert.Equal(t, []token.Token{{Type: token.Add, Text: "thing"}}, tokens)
	assert.NoError(t, err)
}

func TestNumericTokenConsumer(t *testing.T) {
	consumer := numericTokenConsumer(token.Add)

	_, _, err := consumer(newRunes("thing"))
	assert.Error(t, err)

	r, tokens, err := consumer(newRunes("42thing"))
	assert.Equal(t, []rune("thing"), r.text)
	assert.Equal(t, []token.Token{{Type: token.Add, Text: "42"}}, tokens)
	assert.NoError(t, err)
}

func TestAlphanumericTokenConsumer(t *testing.T) {
	consumer := alphanumericTokenConsumer(token.Add)

	_, _, err := consumer(newRunes("{}"))
	assert.Error(t, err)

	r, tokens, err := consumer(newRunes("thing42{}"))
	assert.Equal(t, []rune("{}"), r.text)
	assert.Equal(t, []token.Token{{Type: token.Add, Text: "thing42"}}, tokens)
	assert.NoError(t, err)
}

func TestRestOfLineConsumer(t *testing.T) {
	consumer := restOfLineTokenConsumer(token.Add)

	r, tokens, err := consumer(newRunes(""))
	assert.Equal(t, []rune(""), r.text)
	assert.Equal(t, []token.Token{{Type: token.Add, Text: ""}}, tokens)
	assert.NoError(t, err)

	r, tokens, err = consumer(newRunes("fourfour\nfour\n"))
	assert.Equal(t, []rune("four\n"), r.text)
	assert.Equal(t, []token.Token{{Type: token.Add, Text: "fourfour"}}, tokens)
	assert.NoError(t, err)

	r, tokens, err = consumer(newRunes("four"))
	assert.Equal(t, []rune(""), r.text)
	assert.Equal(t, []token.Token{{Type: token.Add, Text: "four"}}, tokens)
	assert.NoError(t, err)
}
