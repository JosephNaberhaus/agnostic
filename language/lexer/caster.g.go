package lexer

import "github.com/JosephNaberhaus/agnostic/ast"


func castToDefinition[T ast.Definition](consumer consumer[T]) consumer[ast.Definition] {
	return func(state parserState) (parserState, ast.Definition, error) {
		newState, result, err := consumer(state)
		if err != nil {
			return parserState{}, nil, err
		}

		return newState, result, nil
	}
}

func castToStatement[T ast.Statement](consumer consumer[T]) consumer[ast.Statement] {
	return func(state parserState) (parserState, ast.Statement, error) {
		newState, result, err := consumer(state)
		if err != nil {
			return parserState{}, nil, err
		}

		return newState, result, nil
	}
}

func castToType[T ast.Type](consumer consumer[T]) consumer[ast.Type] {
	return func(state parserState) (parserState, ast.Type, error) {
		newState, result, err := consumer(state)
		if err != nil {
			return parserState{}, nil, err
		}

		return newState, result, nil
	}
}

func castToValue[T ast.Value](consumer consumer[T]) consumer[ast.Value] {
	return func(state parserState) (parserState, ast.Value, error) {
		newState, result, err := consumer(state)
		if err != nil {
			return parserState{}, nil, err
		}

		return newState, result, nil
	}
}
