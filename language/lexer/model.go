package lexer

import "github.com/JosephNaberhaus/agnostic/language/token"

func modelConsumer() consumer {
	return inOrder(
		optional(commentConsumer(token.ModelComment)),
		stringTokenConsumer("model", token.Model),
		allWhitespaceConsumer(),
		alphaTokenConsumer(token.ModelName),
		anyWhitespaceConsumer(),
		stringTokenConsumer("{", token.StartBlock),
		emptyLineConsumer(),
		repeatAndThen(
			first(
				emptyLineConsumer(),
				fieldConsumer(),
				funcConsumer(),
			),
			inOrder(
				anyWhitespaceConsumer(),
				stringTokenConsumer("}", token.EndBlock),
			),
		),
		optional(emptyLineConsumer()),
	)
}

func fieldConsumer() consumer {
	return inOrder(
		optional(commentConsumer(token.FieldComment)),
		anyWhitespaceConsumer(),
		alphaTokenConsumer(token.FieldName),
		anyWhitespaceConsumer(),
		stringConsumer(":"),
		anyWhitespaceConsumer(),
		alphanumericTokenConsumer(token.TypeName),
		anyWhitespaceConsumer(),
		stringTokenConsumer(";", token.Semicolon),
		emptyLineConsumer(),
	)
}

func funcConsumer() consumer {
	return inOrder(
		optional(commentConsumer(token.FuncComment)),
		anyWhitespaceConsumer(),
		stringTokenConsumer("func", token.Func),
		allWhitespaceConsumer(),
		alphaTokenConsumer(token.FuncName),
		anyWhitespaceConsumer(),
		stringTokenConsumer("(", token.StartArguments),
		repeatAndThen(
			inOrder(
				anyWhitespaceConsumer(),
				argConsumer(),
				anyWhitespaceConsumer(),
				stringConsumer(","),
			),
			inOrder(
				anyWhitespaceConsumer(),
				optional(argConsumer()),
				anyWhitespaceConsumer(),
				stringTokenConsumer(")", token.EndArguments),
			),
		),
		anyWhitespaceConsumer(),
		stringTokenConsumer("{", token.StartBlock),
		emptyLineConsumer(),
		repeatAndThen(
			inOrder(
				statementConsumer(),
			),
			inOrder(
				anyWhitespaceConsumer(),
				stringTokenConsumer("}", token.EndBlock),
			),
		),
		optional(emptyLineConsumer()),
	)
}

func argConsumer() consumer {
	return inOrder(
		alphaTokenConsumer(token.ArgName),
		anyWhitespaceConsumer(),
		stringConsumer(":"),
		anyWhitespaceConsumer(),
		alphanumericTokenConsumer(token.TypeName),
	)
}

func statementConsumer() consumer {
	return inOrder(
		anyWhitespaceConsumer(),
		assignmentConsumer(),
		anyWhitespaceConsumer(),
		stringTokenConsumer(";", token.Semicolon),
		emptyLineConsumer(),
	)
}

func assignmentConsumer() consumer {
	return inOrder(
		alphaTokenConsumer(token.VariableName),
		anyWhitespaceConsumer(),
		stringTokenConsumer("=", token.Assign),
		anyWhitespaceConsumer(),
		valueConsumer(),
	)
}

func valueConsumer() consumer {
	return first(
		alphaTokenConsumer(token.VariableName),
		boolLiteralConsumer(),
	)
}

func boolLiteralConsumer() consumer {
	return first(
		stringTokenConsumer("false", token.BooleanLit),
		stringTokenConsumer("true", token.BooleanLit),
	)
}
