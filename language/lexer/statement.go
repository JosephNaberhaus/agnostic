package lexer

import "github.com/JosephNaberhaus/agnostic/ast"

func blockConsumer() consumer[ast.Block] {
	var result ast.Block
	return attempt(
		&result,
		repeat(
			first(
				emptyLineConsumer(),
				handleNoError(
					statementConsumer(),
					func(statement ast.Statement) {
						result.Statements = append(result.Statements, statement)
					},
				),
			),
		),
	)
}

func statementConsumer() consumer[ast.Statement] {
	return first(
		castToStatement(conditionalConsumer()),
		castToStatement(forConsumer()),
		castToStatement(forInConsumer()),
		singleLineStatementConsumer(),
	)
}

func singleLineStatementConsumer() consumer[ast.Statement] {
	var result ast.Statement
	return attempt(
		&result,
		inOrder(
			anyWhitespaceConsumer(),
			handleNoError(
				inlineStatementConsumer(),
				func(statement ast.Statement) {
					result = statement
				},
			),
			anyWhitespaceConsumer(),
			skip(stringConsumer(";")),
			emptyLineConsumer(),
		),
	)
}

func inlineStatementConsumer() consumer[ast.Statement] {
	return first(
		castToStatement(assignmentConsumer()),
		castToStatement(returnConsumer()),
		castToStatement(declareConsumer()),
	)
}

func assignmentConsumer() consumer[ast.Assignment] {
	var result ast.Assignment
	return attempt(
		&result,
		inOrder(
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
				deferred(blockConsumer),
				func(block ast.Block) {
					result.Block = block
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
				deferred(blockConsumer),
				func(block ast.Block) {
					result.Block = block
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
				deferred(blockConsumer),
				func(block ast.Block) {
					result.Block = block
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
			skip(stringConsumer("return")),
			allWhitespaceConsumer(),
			handleNoError(
				valueConsumer(),
				func(value ast.Value) {
					result.Value = value
				},
			),
		),
	)
}

func declareConsumer() consumer[ast.Declare] {
	var result ast.Declare
	return attempt(
		&result,
		inOrder(
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
		),
	)
}

func forConsumer() consumer[ast.For] {
	var result ast.For
	return attempt(
		&result,
		inOrder(
			anyWhitespaceConsumer(),
			skip(stringConsumer("for")),
			anyWhitespaceConsumer(),
			skip(stringConsumer("(")),
			optional(handleNoError(
				inlineStatementConsumer(),
				func(initialization ast.Statement) {
					result.Initialization = initialization
				},
			)),
			anyWhitespaceConsumer(),
			skip(stringConsumer(";")),
			anyWhitespaceConsumer(),
			optional(handleNoError(
				valueConsumer(),
				func(condition ast.Value) {
					result.Condition = condition
				},
			)),
			anyWhitespaceConsumer(),
			skip(stringConsumer(";")),
			anyWhitespaceConsumer(),
			optional(handleNoError(
				inlineStatementConsumer(),
				func(afterEach ast.Statement) {
					result.AfterEach = afterEach
				},
			)),
			anyWhitespaceConsumer(),
			skip(stringConsumer(")")),
			anyWhitespaceConsumer(),
			skip(stringConsumer("{")),
			anyWhitespaceConsumer(),
			emptyLineConsumer(),
			handleNoError(
				deferred(blockConsumer),
				func(block ast.Block) {
					result.Block = block
				},
			),
			anyWhitespaceConsumer(),
			skip(stringConsumer("}")),
			emptyLineConsumer(),
		),
	)
}

func forInConsumer() consumer[ast.ForIn] {
	var result ast.ForIn
	return attempt(
		&result,
		inOrder(
			anyWhitespaceConsumer(),
			skip(stringConsumer("for")),
			anyWhitespaceConsumer(),
			skip(stringConsumer("(")),
			anyWhitespaceConsumer(),
			skip(stringConsumer("var")),
			allWhitespaceConsumer(),
			handleNoError(
				alphaConsumer(),
				func(itemName string) {
					result.ItemName = itemName
				},
			),
			allWhitespaceConsumer(),
			skip(stringConsumer("in")),
			allWhitespaceConsumer(),
			handleNoError(
				valueConsumer(),
				func(iterable ast.Value) {
					result.Iterable = iterable
				},
			),
			anyWhitespaceConsumer(),
			skip(stringConsumer(")")),
			anyWhitespaceConsumer(),
			skip(stringConsumer("{")),
			anyWhitespaceConsumer(),
			emptyLineConsumer(),
			handleNoError(
				deferred(blockConsumer),
				func(block ast.Block) {
					result.Block = block
				},
			),
			anyWhitespaceConsumer(),
			skip(stringConsumer("}")),
			emptyLineConsumer(),
		),
	)
}
