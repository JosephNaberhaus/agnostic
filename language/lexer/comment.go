package lexer

import "github.com/JosephNaberhaus/agnostic/language/token"

func commentConsumer(tokenType token.Type) consumer {
	return inOrder(
		anyWhitespaceConsumer(),
		stringConsumer("//"),
		anyWhitespaceConsumer(),
		restOfLineTokenConsumer(tokenType),
	)
}
