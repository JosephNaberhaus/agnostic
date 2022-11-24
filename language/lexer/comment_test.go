package lexer

import (
	"github.com/JosephNaberhaus/agnostic/language/token"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestComment(t *testing.T) {
	consumer := commentConsumer(token.ModelComment)

	r, tokens, err := consumer(newRunes("  //  test <test> 42 $!<>\n"))
	assert.Empty(t, r.text)
	assert.Equal(t, []token.Token{{Type: token.ModelComment, Text: "test <test> 42 $!<>"}}, tokens)
	assert.NoError(t, err)
}
