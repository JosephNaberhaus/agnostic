package lexer

import "github.com/JosephNaberhaus/agnostic/ast"

func typeConsumer() consumer[ast.Type] {
	return first(
		castToType(primitiveConsumer()),
		castToType(listConsumer()),
		castToType(mapConsumer()),
		castToType(setConsumer()),
		meta(castToType(modelConsumer()), TokenKind_type),
	)
}

func primitiveConsumer() consumer[ast.Primitive] {
	return meta(first(
		mapResultToConstant(stringConsumer("bool"), ast.Boolean),
		mapResultToConstant(stringConsumer("int"), ast.Int),
		mapResultToConstant(stringConsumer("rune"), ast.Rune),
		mapResultToConstant(stringConsumer("string"), ast.String),
		mapResultToConstant(stringConsumer("void"), ast.Void),
	), TokenKind_keyword)
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
			handleNoError(
				deferred(typeConsumer),
				func(base ast.Type) {
					result.Base = base
				},
			),
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
			anyWhitespaceConsumer(),
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

func setConsumer() consumer[ast.Set] {
	var result ast.Set
	return attempt(
		&result,
		inOrder(
			skip(stringConsumer("set<")),
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
