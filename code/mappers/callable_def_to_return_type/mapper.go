package callable_def_to_return_type

import "github.com/JosephNaberhaus/agnostic/code"

type Mapper struct{}

func (m Mapper) MapFunctionDef(original *code.FunctionDef) code.Type {
	return original.ReturnType
}

func (m Mapper) MapEqualOverride(_ *code.EqualOverride) code.Type {
	return code.Boolean
}

func (m Mapper) MapHashOverride(_ *code.HashOverride) code.Type {
	return code.Int
}
