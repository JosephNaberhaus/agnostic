package lexer

import "github.com/JosephNaberhaus/agnostic/ast"

func statementsConsumer() consumer[[]ast.Statement] {
	var result []ast.Statement
	return attempt(
		&result,
		repeat(
			first(
				emptyLineConsumer(),
				handleNoError(
					statementConsumer(),
					func(statement ast.Statement) {
						result = append(result, statement)
					},
				),
			),
		),
	)
}

func statementConsumer() consumer[ast.Statement] {
	return first(
		castToStatement(assignmentConsumer()),
		castToStatement(conditionalConsumer()),
		castToStatement(returnConsumer()),
		castToStatement(declareConsumer()),
	)
}

func assignmentConsumer() consumer[ast.Assignment] {
	var result ast.Assignment
	return attempt(
		&result,
		inOrder(
			anyWhitespaceConsumer(),
			handleNoError(
				valueConsumer(),
				func(to ast.Value) {
					result.To = to
				},
			),
			anyWhitespaceConsumer(),
			skip(stringConsumer("=")),
			anyWhitespaceConsumer(),
			handleNoError(
				valueConsumer(),
				func(from ast.Value) {
					result.From = from
				},
			),
			anyWhitespaceConsumer(),
			skip(stringConsumer(";")),
			emptyLineConsumer(),
		),
	)
}

func conditionalConsumer() consumer[ast.Conditional] {
	var result ast.Conditional
	return attempt(
		&result,
		inOrder(
			anyWhitespaceConsumer(),
			handleNoError(
				ifConsumer(),
				func(ifValue ast.If) {
					result.If = ifValue
				},
			),
			repeat(first(
				handleNoError(
					elseIfConsumer(),
					func(elseIfValue ast.ElseIf) {
						result.ElseIfs = append(result.ElseIfs, elseIfValue)
					},
				),
			)),
			optional(handleNoError(
				elseConsumer(),
				func(elseValue ast.Else) {
					result.Else = elseValue
				},
			)),
			emptyLineConsumer(),
		),
	)
}

func ifConsumer() consumer[ast.If] {
	var result ast.If
	return attempt(
		&result,
		inOrder(
			anyWhitespaceConsumer(),
			skip(stringConsumer("if")),
			anyWhitespaceConsumer(),
			skip(stringConsumer("(")),
			anyWhitespaceConsumer(),
			handleNoError(
				valueConsumer(),
				func(condition ast.Value) {
					result.Condition = condition
				},
			),
			anyWhitespaceConsumer(),
			skip(stringConsumer(")")),
			anyWhitespaceConsumer(),
			skip(stringConsumer("{")),
			emptyLineConsumer(),
			handleNoError(
				deferred(statementsConsumer),
				func(statements []ast.Statement) {
					result.Statements = statements
				},
			),
			anyWhitespaceConsumer(),
			skip(stringConsumer("}")),
		),
	)
}

func elseIfConsumer() consumer[ast.ElseIf] {
	var result ast.ElseIf
	return attempt(
		&result,
		inOrder(
			anyWhitespaceConsumer(),
			skip(stringConsumer("else if")),
			anyWhitespaceConsumer(),
			skip(stringConsumer("(")),
			anyWhitespaceConsumer(),
			skip(stringConsumer("(")),
			handleNoError(
				valueConsumer(),
				func(condition ast.Value) {
					result.Condition = condition
				},
			),
			anyWhitespaceConsumer(),
			skip(stringConsumer(")")),
			anyWhitespaceConsumer(),
			skip(stringConsumer("{")),
			emptyLineConsumer(),
			handleNoError(
				deferred(statementsConsumer),
				func(statements []ast.Statement) {
					result.Statements = statements
				},
			),
			anyWhitespaceConsumer(),
			skip(stringConsumer("}")),
		),
	)
}

func elseConsumer() consumer[ast.Else] {
	var result ast.Else
	return attempt(
		&result,
		inOrder(
			anyWhitespaceConsumer(),
			skip(stringConsumer("else")),
			anyWhitespaceConsumer(),
			skip(stringConsumer("{")),
			emptyLineConsumer(),
			handleNoError(
				deferred(statementsConsumer),
				func(statements []ast.Statement) {
					result.Statements = statements
				},
			),
			anyWhitespaceConsumer(),
			skip(stringConsumer("}")),
		),
	)
}

func returnConsumer() consumer[ast.Return] {
	var result ast.Return
	return attempt(
		&result,
		inOrder(
			anyWhitespaceConsumer(),
			skip(stringConsumer("return")),
			allWhitespaceConsumer(),
			handleNoError(
				valueConsumer(),
				func(value ast.Value) {
					result.Value = value
				},
			),
			anyWhitespaceConsumer(),
			skip(stringConsumer(";")),
			emptyLineConsumer(),
		),
	)
}

func declareConsumer() consumer[ast.Declare] {
	var result ast.Declare
	return attempt(
		&result,
		inOrder(
			anyWhitespaceConsumer(),
			skip(stringConsumer("var")),
			allWhitespaceConsumer(),
			handleNoError(
				alphaConsumer(),
				func(name string) {
					result.Name = name
				},
			),
			anyWhitespaceConsumer(),
			skip(stringConsumer("=")),
			anyWhitespaceConsumer(),
			handleNoError(
				valueConsumer(),
				func(value ast.Value) {
					result.Value = value
				},
			),
			anyWhitespaceConsumer(),
			skip(stringConsumer(";")),
			emptyLineConsumer(),
		),
	)
}
