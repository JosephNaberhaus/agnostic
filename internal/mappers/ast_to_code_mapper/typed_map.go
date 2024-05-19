package ast_to_code_mapper

import (
	"fmt"

	"github.com/JosephNaberhaus/agnostic/ast"
	"github.com/JosephNaberhaus/agnostic/code"
)

func mapAstNodeTo[T code.Node](original ast.Node, mapper *Mapper) (T, error) {
	result, err := ast.MapNode(original, mapper)
	if err != nil {
		var zero T
		return zero, err
	}

	typedResult, ok := result.(T)
	if !ok {
		var zero T
		return zero, fmt.Errorf("expected %T but got %T", zero, result)
	}

	return typedResult, nil
}

func mapAstNodesTo[T code.Node, V ast.Node](originalNodes []V, mapper *Mapper) ([]T, error) {
	results := make([]T, 0, len(originalNodes))
	for _, original := range originalNodes {
		result, err := mapAstNodeTo[T](original, mapper)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}

	return results, nil
}
