package lexer

import "github.com/JosephNaberhaus/agnostic/language/token"

func isWhitespace(r rune) bool {
	switch r {
	case ' ', '\t':
		return true
	}

	return false
}

func isNewline(r rune) bool {
	return r == '\n'
}

// allWhitespaceConsumer creates a consumer that consumes one or more whitespace character that are next.
func allWhitespaceConsumer() consumer {
	return func(r runes) (runes, []token.Token, error) {
		if !isWhitespace(r.consumeOne()) {
			return runes{}, nil, createError(r, "expected whitespace")
		}

		r.consumeWhile(isWhitespace)

		return r, nil, nil
	}
}

// anyWhitespaceConsumer creates a consumer that consumes any whitespace characters that are next.
func anyWhitespaceConsumer() consumer {
	return func(r runes) (runes, []token.Token, error) {
		r.consumeWhile(isWhitespace)
		return r, nil, nil
	}
}

// newlineConsumer creates a consumer that consumes a whitespace character.
func newlineConsumer() consumer {
	return func(r runes) (runes, []token.Token, error) {
		if !isNewline(r.consumeOne()) {
			return runes{}, nil, createError(r, "expected newline")
		}

		return r, nil, nil
	}
}

// emptyLineConsumer creates a consumer that consumes an empty line.
func emptyLineConsumer() consumer {
	return inOrder(
		anyWhitespaceConsumer(),
		newlineConsumer(),
	)
}
