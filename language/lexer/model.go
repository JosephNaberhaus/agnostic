package lexer

import (
	"errors"
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
			meta(skip(stringConsumer("model")), TokenKind_keyword),
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
						functionDefConsumer(),
						func(method ast.FunctionDef) {
							result.Methods = append(result.Methods, method)
						},
					),
					handle(
						equalOverrideConsumer(),
						func(equalOverride ast.EqualOverride) error {
							if result.EqualOverride.IsSet() {
								return errors.New("only one definition of @== is allowed")
							}

							result.EqualOverride.Set(equalOverride)

							return nil
						},
					),
					handle(
						hashOverrideConsumer(),
						func(hashOverride ast.HashOverride) error {
							if result.HashOverride.IsSet() {
								return errors.New("only one definition of @hash is allowed")
							}

							result.HashOverride.Set(hashOverride)

							return nil
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

func argumentsConsumer() consumer[[]ast.ArgumentDef] {
	var result []ast.ArgumentDef
	return attempt(
		&result,
		inOrder(
			repeat(inOrder(
				anyWhitespaceConsumer(),
				handleNoError(
					argumentConsumer(),
					func(argument ast.ArgumentDef) {
						result = append(result, argument)
					},
				),
				anyWhitespaceConsumer(),
				skip(stringConsumer(",")),
			)),
			anyWhitespaceConsumer(),
			optional(
				handleNoError(
					argumentConsumer(),
					func(argument ast.ArgumentDef) {
						result = append(result, argument)
					},
				),
			),
		),
	)
}

func argumentConsumer() consumer[ast.ArgumentDef] {
	var result ast.ArgumentDef
	return attempt(
		&result,
		inOrder(
			handleNoError(
				alphaConsumer(),
				func(name string) {
					result.Name = name
				},
			),
			anyWhitespaceConsumer(),
			skip(stringConsumer(":")),
			anyWhitespaceConsumer(),
			handleNoError(
				typeConsumer(),
				func(argumentType ast.Type) {
					result.Type = argumentType
				},
			),
		),
	)
}

func functionDefConsumer() consumer[ast.FunctionDef] {
	var result ast.FunctionDef
	return attempt(
		&result,
		inOrder(
			anyWhitespaceConsumer(),
			handleNoError(
				meta(alphaConsumer(), TokenKind_function),
				func(name string) {
					result.Name = name
				},
			),
			anyWhitespaceConsumer(),
			skip(stringConsumer("(")),
			handleNoError(
				argumentsConsumer(),
				func(arguments []ast.ArgumentDef) {
					result.Arguments = arguments
				},
			),
			skip(stringConsumer(")")),
			anyWhitespaceConsumer(),
			skip(stringConsumer(":")),
			anyWhitespaceConsumer(),
			handleNoError(
				typeConsumer(),
				func(returnType ast.Type) {
					result.ReturnType = returnType
				},
			),
			anyWhitespaceConsumer(),
			skip(stringConsumer("{")),
			emptyLineConsumer(),
			handleNoError(
				blockConsumer(),
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

func equalOverrideConsumer() consumer[ast.EqualOverride] {
	var result ast.EqualOverride
	return attempt(
		&result,
		inOrder(
			anyWhitespaceConsumer(),
			skip(stringConsumer("@==")),
			anyWhitespaceConsumer(),
			skip(stringConsumer("(")),
			anyWhitespaceConsumer(),
			handleNoError(
				alphaConsumer(),
				func(otherName string) {
					result.OtherName = otherName
				},
			),
			anyWhitespaceConsumer(),
			skip(stringConsumer(")")),
			anyWhitespaceConsumer(),
			skip(stringConsumer("{")),
			emptyLineConsumer(),
			handleNoError(
				blockConsumer(),
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

func hashOverrideConsumer() consumer[ast.HashOverride] {
	var result ast.HashOverride
	return attempt(
		&result,
		inOrder(
			anyWhitespaceConsumer(),
			skip(stringConsumer("@hash")),
			anyWhitespaceConsumer(),
			skip(stringConsumer("(")),
			anyWhitespaceConsumer(),
			skip(stringConsumer(")")),
			anyWhitespaceConsumer(),
			skip(stringConsumer("{")),
			emptyLineConsumer(),
			handleNoError(
				blockConsumer(),
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
