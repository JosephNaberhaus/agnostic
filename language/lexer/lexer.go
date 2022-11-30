//go:generate go run ../../tool/consumer_caster_generator/main.go -nodeTypesFile ../../ast/node_types.go

package lexer

import (
	"github.com/JosephNaberhaus/agnostic/ast"
)

// TODO rename
func Test(rawText string) (ast.Module, error) {
	r := newRunes(rawText)

	var module ast.Module

	newR, _, err := inOrder(
		anyWhitespaceConsumer(),
		skip(stringConsumer("module")),
		allWhitespaceConsumer(),
		handleNoError(
			alphaConsumer(),
			func(name string) {
				module.Name = name
			},
		),
		emptyLineConsumer(),
		repeat(first(
			emptyLineConsumer(),
			handleNoError(
				modelDefConsumer(),
				func(model ast.ModelDef) {
					module.Models = append(module.Models, model)
				},
			),
			handleNoError(
				functionDefConsumer(),
				func(function ast.FunctionDef) {
					module.Functions = append(module.Functions, function)
				},
			),
		)),
	)(r)
	if err != nil {
		return ast.Module{}, contextualize(err, []rune(rawText))
	}

	if newR.isNotEmpty() {
		return ast.Module{}, contextualize(createError(newR, "expected end of module"), []rune(rawText))
	}

	return module, nil
}
