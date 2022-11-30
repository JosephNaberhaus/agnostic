package lexer

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
func allWhitespaceConsumer() consumer[void] {
	return func(r parserState) (parserState, void, error) {
		if !isWhitespace(r.consumeOne()) {
			return parserState{}, nil, createError(r, "expected whitespace")
		}

		r.consumeWhile(isWhitespace)

		return r, nil, nil
	}
}

// anyWhitespaceConsumer creates a consumer that consumes any whitespace characters that are next.
func anyWhitespaceConsumer() consumer[void] {
	return func(r parserState) (parserState, void, error) {
		r.consumeWhile(isWhitespace)
		return r, nil, nil
	}
}

// newlineConsumer creates a consumer that consumes a whitespace character.
func newlineConsumer() consumer[void] {
	return func(r parserState) (parserState, void, error) {
		if !isNewline(r.peekOne()) {
			return parserState{}, nil, createError(r, "expected newline")
		}

		r.consumeOne()
		return r, nil, nil
	}
}

// emptyLineConsumer creates a consumer that consumes an empty line.
func emptyLineConsumer() consumer[void] {
	return inOrder(
		anyWhitespaceConsumer(),
		newlineConsumer(),
	)
}
