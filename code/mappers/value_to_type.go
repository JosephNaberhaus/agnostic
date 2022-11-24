package mappers

import (
	"errors"
	"github.com/JosephNaberhaus/agnostic/code"
)

// mapValueToType finds the base code.Type of a code.Value.
//
// It will return an error if the value's type isn't a primitive, or if there is an error in the code tree.
func mapValueToType(value code.Value) (code.Type, error) {
	return code.MapValue[code.Type](value, &valueToTypeMapper{})
}

func mapValueToPrimitiveType(value code.Value) (code.Primitive, error) {
	valueType, err := mapValueToType(value)
	if err != nil {
		return 0, err
	}

	primitiveType, ok := valueType.(code.Primitive)
	if !ok {
		// TODO improve
		return 0, errors.New("is not primitive")
	}

	return primitiveType, nil
}

func mapValueModelType(value code.Value) (*code.Model, error) {
	valueType, err := mapValueToType(value)
	if err != nil {
		return nil, err
	}

	modelType, ok := valueType.(*code.Model)
	if !ok {
		// TODO improve
		return nil, errors.New("is not a model")
	}

	return modelType, nil
}

type valueToTypeMapper struct{}

var _ code.ValueMapper[code.Type] = (*valueToTypeMapper)(nil)

type binaryOperatorPair struct {
	left, right code.Primitive
}

func (v *valueToTypeMapper) MapLiteralInt32(original *code.LiteralInt32) (code.Type, error) {
	return code.Int32, nil
}

func (v *valueToTypeMapper) MapLiteralString(original *code.LiteralString) (code.Type, error) {
	return code.String, nil
}

// unaryOperatorMappings maps from a primitive type into another map containing the unary operations that can be done
// and the output primitive that they produce.
var unaryOperatorMappings = map[code.Primitive]map[code.UnaryOperator]code.Primitive{
	code.Boolean: {
		code.Not: code.Boolean,
	},
	code.Int32: {
		code.Negate: code.Int32,
	},
}

// binaryOperatorMappings maps from a primitive type into another map containing the binary operations that can be done
// and the output primitive that they produce.
var binaryOperatorMappings = map[binaryOperatorPair]map[code.BinaryOperator]code.Primitive{
	{left: code.Int32, right: code.Int32}: {
		code.Add: code.Int32,
	},
	{left: code.String, right: code.String}: {
		code.Add: code.String,
	},
}

func (v *valueToTypeMapper) MapUnaryOperation(original *code.UnaryOperation) (code.Type, error) {
	inputPrimitive, err := mapValueToPrimitiveType(original.Value)
	if err != nil {
		return nil, err
	}

	outputMappings, ok := unaryOperatorMappings[inputPrimitive]
	if !ok {
		// TODO improve
		return nil, errors.New("can't use this unary operator on the provided primitive")
	}

	outputPrimitive, ok := outputMappings[original.Operator]
	if !ok {
		// TODO improve
		return nil, errors.New("invalid unary operator")
	}

	return outputPrimitive, nil
}

func (v *valueToTypeMapper) MapBinaryOperation(original *code.BinaryOperation) (code.Type, error) {
	leftPrimitive, err := mapValueToPrimitiveType(original.Left)
	if err != nil {
		return nil, err
	}

	rightPrimitive, err := mapValueToPrimitiveType(original.Right)
	if err != nil {
		return nil, err
	}

	outputMappings, ok := binaryOperatorMappings[binaryOperatorPair{left: leftPrimitive, right: rightPrimitive}]
	if !ok {
		// TODO improve
		return nil, errors.New("can't use this binary operator on the provided primitive")
	}

	outputPrimitive, ok := outputMappings[original.Operator]
	if !ok {
		// TODO improve
		return nil, errors.New("invalid binary operator")
	}

	return outputPrimitive, nil
}

func (v *valueToTypeMapper) MapProperty(original *code.Property) (code.Type, error) {
	ofModel, err := mapValueModelType(original.Of)
	if err != nil {
		return nil, err
	}

	field, ok := ofModel.ModelMetadata.Definition.FieldMap[original.Name]
	if !ok {
		// TODO improve
		return nil, errors.New("unkown property")
	}

	return field.Type, nil
}

func (v *valueToTypeMapper) MapThis(original *code.This) (code.Type, error) {
	return &code.Model{
		Name: original.This.Name,
		ModelMetadata: code.ModelMetadata{
			Definition: original.This,
		},
	}, nil
}
