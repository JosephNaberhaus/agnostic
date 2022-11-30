package lexer

import (
	"github.com/JosephNaberhaus/agnostic/ast"
)

func modelDefConsumer() consumer[ast.ModelDef] {
	var result ast.ModelDef
	return attempt(
		&result,
		inOrder(
			optional(handleNoError(
				commentConsumer(),
				func(comment string) {
					// TODO: use this
				},
			)),
			skip(stringConsumer("model")),
			allWhitespaceConsumer(),
			handleNoError(
				alphaConsumer(),
				func(modelName string) {
					result.Name = modelName
				},
			),
			anyWhitespaceConsumer(),
			skip(stringConsumer("{")),
			emptyLineConsumer(),
			repeat(
				first(
					emptyLineConsumer(),
					handleNoError(
						fieldDefConsumer(),
						func(field ast.FieldDef) {
							result.Fields = append(result.Fields, field)
						},
					),
					handleNoError(
						methodDefConsumer(),
						func(method ast.MethodDef) {
							result.Methods = append(result.Methods, method)
						},
					),
				),
			),
			anyWhitespaceConsumer(),
			skip(stringConsumer("}")),
			optional(emptyLineConsumer()),
		),
	)
}

func fieldDefConsumer() consumer[ast.FieldDef] {
	var result ast.FieldDef
	return attempt(
		&result,
		inOrder(
			anyWhitespaceConsumer(),
			handleNoError(
				alphaConsumer(),
				func(fieldName string) {
					result.Name = fieldName
				},
			),
			anyWhitespaceConsumer(),
			skip(stringConsumer(":")),
			anyWhitespaceConsumer(),
			handleNoError(
				typeConsumer(),
				func(fieldType ast.Type) {
					result.Type = fieldType
				},
			),
			anyWhitespaceConsumer(),
			skip(stringConsumer(";")),
			emptyLineConsumer(),
		),
	)
}

func methodDefConsumer() consumer[ast.MethodDef] {
	return mapResult(
		functionDefConsumer(),
		func(functionDef ast.FunctionDef) (ast.MethodDef, error) {
			return ast.MethodDef{
				Function: functionDef,
			}, nil
		},
	)
}

func functionDefConsumer() consumer[ast.FunctionDef] {
	var result ast.FunctionDef
	return attempt(
		&result,
		inOrder(
			anyWhitespaceConsumer(),
			handleNoError(
				typeConsumer(),
				func(returnType ast.Type) {
					result.ReturnType = returnType
				},
			),
			allWhitespaceConsumer(),
			handleNoError(
				alphaConsumer(),
				func(name string) {
					result.Name = name
				},
			),
			anyWhitespaceConsumer(),
			skip(stringConsumer("(")),
			// TODO arguments
			skip(stringConsumer(")")),
			anyWhitespaceConsumer(),
			skip(stringConsumer("{")),
			emptyLineConsumer(),
			handleNoError(
				statementsConsumer(),
				func(statements []ast.Statement) {
					result.Statements = statements
				},
			),
			anyWhitespaceConsumer(),
			skip(stringConsumer("}")),
			emptyLineConsumer(),
		),
	)
}
