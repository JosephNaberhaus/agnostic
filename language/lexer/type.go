package lexer

import "github.com/JosephNaberhaus/agnostic/ast"

func typeConsumer() consumer[ast.Type] {
	return first(
		castToType(primitiveConsumer()),
		castToType(listConsumer()),
		castToType(mapConsumer()),
		castToType(modelConsumer()),
	)
}

func primitiveConsumer() consumer[ast.Primitive] {
	return first(
		mapResultToConstant(stringConsumer("bool"), ast.Boolean),
		mapResultToConstant(stringConsumer("int"), ast.Int),
		mapResultToConstant(stringConsumer("string"), ast.String),
		mapResultToConstant(stringConsumer("void"), ast.Void),
	)
}

func modelConsumer() consumer[ast.Model] {
	return mapResult(
		alphaConsumer(),
		func(name string) (ast.Model, error) {
			return ast.Model{
				Name: name,
			}, nil
		},
	)
}

func listConsumer() consumer[ast.List] {
	var result ast.List
	return attempt(
		&result,
		inOrder(
			skip(stringConsumer("list<")),
			anyWhitespaceConsumer(),
			handleNoError(
				deferred(typeConsumer),
				func(base ast.Type) {
					result.Base = base
				},
			),
			anyWhitespaceConsumer(),
			skip(stringConsumer(">")),
		),
	)
}

func mapConsumer() consumer[ast.Map] {
	var result ast.Map
	return attempt(
		&result,
		inOrder(
			skip(stringConsumer("map<")),
			handleNoError(
				deferred(typeConsumer),
				func(key ast.Type) {
					result.Key = key
				},
			),
			anyWhitespaceConsumer(),
			skip(stringConsumer(",")),
			anyWhitespaceConsumer(),
			handleNoError(
				deferred(typeConsumer),
				func(value ast.Type) {
					result.Value = value
				},
			),
			anyWhitespaceConsumer(),
			skip(stringConsumer(">")),
		),
	)
}
