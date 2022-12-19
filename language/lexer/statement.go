package lexer

import (
	"errors"
	"github.com/JosephNaberhaus/agnostic/ast"
)

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
		castToStatement(incrementConsumer()),
		castToStatement(decrementConsumer()),
		castToStatement(assignmentConsumer()),
		castToStatement(returnConsumer()),
		castToStatement(declareConsumer()),
		castToStatement(declareNullConsumer()),
		castToStatement(callConsumer()),
		castToStatement(breakConsumer()),
		castToStatement(continueConsumer()),
	)
}

func incrementConsumer() consumer[ast.Assignment] {
	var result ast.Assignment
	return attempt(
		&result,
		inOrder(
			anyWhitespaceConsumer(),
			handleNoError(
				valueConsumer(),
				func(value ast.Value) {
					result = ast.Assignment{
						To: value,
						From: ast.BinaryOperation{
							Left:     value,
							Operator: ast.Add,
							Right:    ast.LiteralInt{Value: 1},
						},
					}
				},
			),
			anyWhitespaceConsumer(),
			skip(stringConsumer("++")),
		),
	)
}

func decrementConsumer() consumer[ast.Assignment] {
	var result ast.Assignment
	return attempt(
		&result,
		inOrder(
			anyWhitespaceConsumer(),
			handleNoError(
				valueConsumer(),
				func(value ast.Value) {
					result = ast.Assignment{
						To: value,
						From: ast.BinaryOperation{
							Left:     value,
							Operator: ast.Subtract,
							Right:    ast.LiteralInt{Value: 1},
						},
					}
				},
			),
			anyWhitespaceConsumer(),
			skip(stringConsumer("--")),
		),
	)
}

func assignmentConsumer() consumer[ast.Assignment] {
	type resultType struct {
		to       ast.Value
		from     ast.Value
		operator ast.BinaryOperator
	}
	var result resultType
	return mapResult(
		attempt(
			&result,
			inOrder(
				handleNoError(
					valueConsumer(),
					func(to ast.Value) {
						result.to = to
					},
				),
				anyWhitespaceConsumer(),
				optional(handleNoError(
					first(
						mapResultToConstant(stringConsumer("+"), ast.Add),
						mapResultToConstant(stringConsumer("-"), ast.Subtract),
						mapResultToConstant(stringConsumer("*"), ast.Multiply),
						mapResultToConstant(stringConsumer("/"), ast.Divide),
						mapResultToConstant(stringConsumer("%"), ast.Modulo),
					),
					func(operator ast.BinaryOperator) {
						result.operator = operator
					},
				)),
				skip(stringConsumer("=")),
				anyWhitespaceConsumer(),
				handleNoError(
					valueConsumer(),
					func(from ast.Value) {
						result.from = from
					},
				),
			),
		),
		func(result resultType) (ast.Assignment, error) {
			from := result.from
			if result.operator != 0 {
				from = ast.BinaryOperation{
					Left:     result.to,
					Operator: result.operator,
					Right:    result.from,
				}
			}

			return ast.Assignment{
				To:   result.to,
				From: from,
			}, nil
		},
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

func declareNullConsumer() consumer[ast.DeclareNull] {
	var result ast.DeclareNull
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
			handleNoError(
				typeConsumer(),
				func(declareType ast.Type) {
					result.Type = declareType
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
			anyWhitespaceConsumer(),
			optional(inOrder(
				handleNoError(
					inlineStatementConsumer(),
					func(initialization ast.Statement) {
						result.Initialization.Set(initialization)
					},
				),
				anyWhitespaceConsumer(),
				skip(stringConsumer(";")),
				anyWhitespaceConsumer(),
			)),
			optional(handleNoError(
				valueConsumer(),
				func(condition ast.Value) {
					result.Condition = condition
				},
			)),
			anyWhitespaceConsumer(),
			optional(inOrder(
				skip(stringConsumer(";")),
				anyWhitespaceConsumer(),
				handleNoError(
					inlineStatementConsumer(),
					func(afterEach ast.Statement) {
						result.AfterEach.Set(afterEach)
					},
				),
				anyWhitespaceConsumer(),
			)),
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

func callConsumer() consumer[ast.Call] {
	return mapResult(
		valueConsumer(),
		func(value ast.Value) (ast.Call, error) {
			if call, ok := value.(ast.Call); ok {
				return call, nil
			}

			return ast.Call{}, errors.New("expected function call")
		},
	)
}

func breakConsumer() consumer[ast.Break] {
	return mapResultToConstant(stringConsumer("break"), ast.Break{})
}

func continueConsumer() consumer[ast.Continue] {
	return mapResultToConstant(stringConsumer("continue"), ast.Continue{})
}
