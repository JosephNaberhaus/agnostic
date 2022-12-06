package agnostic

import (
	"github.com/JosephNaberhaus/agnostic/ast"
	"github.com/JosephNaberhaus/agnostic/code"
	"github.com/JosephNaberhaus/agnostic/code/mappers"
	"github.com/JosephNaberhaus/agnostic/implementation"
	"github.com/JosephNaberhaus/agnostic/implementation/text"
)

// Map parses the given AST module and writes it using the given implementation.
func Map(module ast.Module, mapper implementation.Mapper) (string, error) {
	codeModule, err := mappers.AstToCode(module)
	if err != nil {
		return "", err
	}

	return mapCode(codeModule, mapper)
}

// mapCode maps a code module into an implementations code using the given mapper.
func mapCode(codeModule *code.Module, mapper implementation.Mapper) (string, error) {
	config := mapper.Config()

	output := mapper.MapModule(codeModule)

	return output.String(text.Config{Indent: config.Indent}), nil
}
