package parser

import "github.com/JosephNaberhaus/agnostic/language/token"

type tokenQueue struct {
	remaining []token.Token
}

func (t *tokenQueue) isEmpty() bool {
	return len(t.remaining) == 0
}

func (t *tokenQueue) isNotEmpty() bool {
	return !t.isEmpty()
}

func (t *tokenQueue) peekOne() token.Token {
	if t.isEmpty() {
		return token.Token{}
	}

	return t.remaining[0]
}

func (t *tokenQueue) consumeOne() token.Token {
	if t.isEmpty() {
		return token.Token{}
	}

	consumed := t.remaining[0]
	t.remaining = t.remaining[1:]
	return consumed
}

func (t *tokenQueue) consumeUntilTypeIs(tokenType token.Type, handler func(token.Token)) {
	for t.isNotEmpty() {
		if t.peekOne().Type == tokenType {
			return
		}

		handler(t.consumeOne())
	}
}
