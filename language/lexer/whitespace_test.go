package lexer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAllWhitespaceConsumer(t *testing.T) {
	consumer := allWhitespaceConsumer()

	r, tokens, err := consumer(newRunes("notwhitespace"))
	assert.Error(t, err)

	r, tokens, err = consumer(newRunes("      notwhitespace"))
	assert.Equal(t, []rune("notwhitespace"), r.text)
	assert.Empty(t, tokens)
	assert.NoError(t, err)
}

func TestAnyWhitespaceConsumer(t *testing.T) {
	consumer := anyWhitespaceConsumer()

	r, tokens, err := consumer(newRunes("notwhitespace"))
	assert.Equal(t, []rune("notwhitespace"), r.text)
	assert.Empty(t, tokens)
	assert.NoError(t, err)

	r, tokens, err = consumer(newRunes("      notwhitespace"))
	assert.Equal(t, []rune("notwhitespace"), r.text)
	assert.Empty(t, tokens)
	assert.NoError(t, err)
}

func TestNewlineConsumer(t *testing.T) {
	consumer := newlineConsumer()

	r, tokens, err := consumer(newRunes("notnewline"))
	assert.Error(t, err)

	r, tokens, err = consumer(newRunes("\n\nnotnewline"))
	assert.Equal(t, []rune("\nnotnewline"), r.text)
	assert.Empty(t, tokens)
	assert.NoError(t, err)
}

func TestEmptyLineConsumer(t *testing.T) {
	consumer := emptyLineConsumer()

	r, tokens, err := consumer(newRunes("notemptyline"))
	assert.Error(t, err)

	r, tokens, err = consumer(newRunes("\n\nnotemptyline"))
	assert.Equal(t, []rune("\nnotemptyline"), r.text)
	assert.Empty(t, tokens)
	assert.NoError(t, err)

	r, tokens, err = consumer(newRunes("    \n\nnotemptyline"))
	assert.Equal(t, []rune("\nnotemptyline"), r.text)
	assert.Empty(t, tokens)
	assert.NoError(t, err)
}
