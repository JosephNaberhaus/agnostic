package find_function_definition

import (
	"errors"
	"github.com/JosephNaberhaus/agnostic/code"
)

// TODO: this isn't a mapper. Find somewhere else for it.
func InStack(targetName string, stack code.Stack) (*code.FunctionDef, error) {
	for stack.IsNotEmpty() {
		switch node := stack.Pop().(type) {
		case *code.Module:
			if function, ok := node.FunctionMap[targetName]; ok {
				return function, nil
			}
		case *code.ModelDef:
			if method, ok := node.MethodMap[targetName]; ok {
				return method, nil
			}
		}
	}

	return nil, errors.New("no function definition found for " + targetName)
}
