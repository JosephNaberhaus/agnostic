package lexer

func commentConsumer() consumer[string] {
	var result string
	return meta(attempt(
		&result,
		inOrder(
			anyWhitespaceConsumer(),
			skip(stringConsumer("//")),
			anyWhitespaceConsumer(),
			handleNoError(
				restOfLineConsumer(),
				func(value string) {
					result = value
				},
			),
		),
	), TokenKind_comment)
}
