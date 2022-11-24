package lexer

import "github.com/JosephNaberhaus/agnostic/language/token"

// TODO rename
func Test(rawText string) ([]token.Token, error) {
	r := newRunes(rawText)

	var tokens []token.Token
	for r.isNotEmpty() {
		newR, newTokens, err := first(
			emptyLineConsumer(),
			modelConsumer(),
		)(r)
		if err != nil {
			return nil, contextualize(err, []rune(rawText))
		}

		r = newR
		tokens = append(tokens, newTokens...)
	}

	return tokens, nil
}
