package lexer

import (
	"fmt"
	"math/big"
)

// stringConsumer creates a consumer that consumes only the given string.
func stringConsumer(str string) consumer[string] {
	return func(r parserState) (parserState, string, error) {
		if r.peekStr(len(str)) == str {
			r.consumeStr(len(str))
			return r, str, nil
		}

		return parserState{}, "", createError(r, fmt.Sprintf("expected '%s'", str))
	}
}

func allWhileConsumer(condition func(r rune) bool, noMatchErrorMessage string) consumer[string] {
	return func(r parserState) (parserState, string, error) {
		consumed := r.consumeWhileStr(condition)
		if consumed == "" {
			return parserState{}, "", createError(r, noMatchErrorMessage)
		}

		return r, consumed, nil
	}
}

func anyWhileConsumer(condition func(r rune) bool) consumer[string] {
	return func(r parserState) (parserState, string, error) {
		consumed := r.consumeWhileStr(condition)
		return r, consumed, nil
	}
}

func isAlpha(r rune) bool {
	return ('a' <= r && r <= 'z') || ('A' <= r && r <= 'Z')
}

func alphaConsumer() consumer[string] {
	return allWhileConsumer(isAlpha, "expected one or more alphabetic characters")
}

func isNumeric(r rune) bool {
	return '0' <= r && r <= '9'
}

func intConsumer() consumer[*big.Int] {
	return mapResult(
		reduce(
			func(prev, new string) string {
				return prev + new
			},
			optional(stringConsumer("-")),
			allWhileConsumer(isNumeric, "expected one ore more numeric characters"),
		),
		func(numberStr string) (*big.Int, error) {
			number := new(big.Int)
			number, ok := number.SetString(numberStr, 10)
			if !ok {
				panic("unreachable")
			}

			return number, nil
		},
	)
}

func isAlphanumeric(r rune) bool {
	return isAlpha(r) || isNumeric(r)
}

func alphanumericTokenConsumer() consumer[string] {
	return allWhileConsumer(isAlphanumeric, "expected one or more alphanumeric characters")
}

func restOfLineConsumer() consumer[string] {
	var result string
	return attempt(
		&result,
		inOrder(
			handleNoError(
				anyWhileConsumer(func(r rune) bool {
					return r != '\n'
				}),
				func(value string) {
					result = value
				},
			),
			// Optional because their might not be a newline at the end of a file.
			optional(newlineConsumer()),
		),
	)
}
