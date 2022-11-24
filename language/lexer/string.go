package lexer

import (
	"fmt"
	"github.com/JosephNaberhaus/agnostic/language/token"
)

// stringConsumer creates a consumer that consumes only the given string.
func stringConsumer(str string) consumer {
	return func(r runes) (runes, []token.Token, error) {
		if r.consumeStr(len(str)) == str {
			return r, nil, nil
		}

		return runes{}, nil, createError(r, fmt.Sprintf("expected '%s'", str))
	}
}

// stringTokenConsumer creates a consumer that consumes only the given string and returns it as a token.
func stringTokenConsumer(str string, tokenType token.Type) consumer {
	return func(text runes) (runes, []token.Token, error) {
		text, _, err := stringConsumer(str)(text)
		if err != nil {
			return runes{}, nil, err
		}

		return text, []token.Token{{Type: tokenType, Text: str}}, nil
	}
}

func allWhileTokenConsumer(condition func(r rune) bool, tokenType token.Type, noMatchErrorMessage string) consumer {
	return func(r runes) (runes, []token.Token, error) {
		consumed := r.consumeWhileStr(condition)
		if consumed == "" {
			return runes{}, nil, createError(r, noMatchErrorMessage)
		}

		return r, []token.Token{{Type: tokenType, Text: consumed}}, nil
	}
}

func anyWhileTokenConsumer(condition func(r rune) bool, tokenType token.Type) consumer {
	return func(r runes) (runes, []token.Token, error) {
		consumed := r.consumeWhileStr(condition)
		return r, []token.Token{{Type: tokenType, Text: consumed}}, nil
	}
}

func isAlpha(r rune) bool {
	return ('a' <= r && r <= 'z') || ('A' <= r && r <= 'Z')
}

func alphaTokenConsumer(tokenType token.Type) consumer {
	return allWhileTokenConsumer(isAlpha, tokenType, "expected one or more alphabetic characters")
}

func isNumeric(r rune) bool {
	return '0' <= r && r < '9'
}

func numericTokenConsumer(tokenType token.Type) consumer {
	return allWhileTokenConsumer(isNumeric, tokenType, "expected one ore more numeric characters")
}

func isAlphanumeric(r rune) bool {
	return isAlpha(r) || isNumeric(r)
}

func alphanumericTokenConsumer(tokenType token.Type) consumer {
	return allWhileTokenConsumer(isAlphanumeric, tokenType, "expected one or more alphanumeric characters")
}

func restOfLineTokenConsumer(tokenType token.Type) consumer {
	return inOrder(
		anyWhileTokenConsumer(func(r rune) bool {
			return r != '\n'
		}, tokenType),
		// Optional because their might not be a newline at the end of a file.
		optional(newlineConsumer()),
	)
}
