package callable_to_function_definition

import (
	"errors"
	"github.com/JosephNaberhaus/agnostic/code"
	"github.com/JosephNaberhaus/agnostic/code/mappers/value_to_type"
)

type Mapper struct{}

func (m Mapper) MapFunction(original *code.Function) (*code.FunctionDef, error) {
	return original.Definition, nil
}

func (m Mapper) MapFunctionProperty(original *code.FunctionProperty) (*code.FunctionDef, error) {
	ofType, err := code.MapValue[code.Type](original.Of, value_to_type.Mapper{})
	if err != nil {
		return nil, err
	}

	ofTypeAsModel, ok := ofType.(*code.Model)
	if !ok {
		return nil, errors.New("function property must be of a model type")
	}

	method, ok := ofTypeAsModel.Definition.MethodMap[original.Name]
	if !ok {
		return nil, errors.New("no method found")
	}

	return method.Function, nil
}
