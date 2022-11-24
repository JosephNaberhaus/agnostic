package parser

import (
	"github.com/JosephNaberhaus/agnostic/ast"
	"github.com/JosephNaberhaus/agnostic/language/lexer"
)

func ToAst(rawText string) (ast.Module, error) {
	tokens, err := lexer.Test(rawText)
	if err != nil {
		return ast.Module{}, err
	}

	parser := &parser{remainingTokens: tokens}
	return parser.parseModule(), nil
}

func modelParser() parser[ast.Model] {
	return func(tokens tokenQueue) (tokenQueue, ast.Model, error) {
		var model ast.Model

		return tokens, model, nil
	}
}

func typeNameToPrimitive(name string) (ast.Primitive, bool) {
	switch name {
	case "bool":
		return ast.Boolean, true
	case "int32":
		return ast.Int32, true
	case "string":
		return ast.String, true
	case "void":
		return ast.Void, true
	}

	return 0, false
}
