package value_to_type

import (
	"errors"
	"github.com/JosephNaberhaus/agnostic/code"
	"github.com/JosephNaberhaus/agnostic/code/mappers/definition_to_type"
	"runtime/debug"
)

// mapValueToType finds the base code.Type of a code.Value.
func mapValueToType(value code.Value) (code.Type, error) {
	return code.MapValue[code.Type](value, &Mapper{})
}

// TODO remove the need for this
func MapValueToPrimitiveType(value code.Value) (code.Primitive, error) {
	valueType, err := mapValueToType(value)
	if err != nil {
		return 0, err
	}

	primitiveType, ok := valueType.(code.Primitive)
	if !ok {
		// TODO improve
		return 0, errors.New("is not primitive\n" + string(debug.Stack()))
	}

	return primitiveType, nil
}

// TODO remove the need for this
func MapValueModelType(value code.Value) (*code.Model, error) {
	valueType, err := mapValueToType(value)
	if err != nil {
		return nil, err
	}

	modelType, ok := valueType.(*code.Model)
	if !ok {
		// TODO improve
		return nil, errors.New("value is not a model")
	}

	return modelType, nil
}

type Mapper struct{}

type binaryOperatorPair struct {
	left, right code.Primitive
}

func (m Mapper) MapLiteralInt(original *code.LiteralInt) (code.Type, error) {
	return code.Int, nil
}

func (m Mapper) MapLiteralRune(original *code.LiteralRune) (code.Type, error) {
	return code.Rune, nil
}

func (m Mapper) MapLiteralString(original *code.LiteralString) (code.Type, error) {
	return code.String, nil
}

// unaryOperatorMappings maps from a primitive type into another map containing the unary operations that can be done
// and the output primitive that they produce.
var unaryOperatorMappings = map[code.Primitive]map[code.UnaryOperator]code.Primitive{
	code.Boolean: {
		code.Not: code.Boolean,
	},
	code.Int: {
		code.Negate: code.Int,
	},
	code.Rune: {
		code.CastToInt: code.Int,
	},
}

// binaryOperatorMappings maps from a primitive type into another map containing the binary operations that can be done
// and the output primitive that they produce.
var binaryOperatorMappings = map[binaryOperatorPair]map[code.BinaryOperator]code.Primitive{
	{left: code.Int, right: code.Int}: {
		code.Equals:      code.Boolean,
		code.LessThan:    code.Boolean,
		code.GreaterThan: code.Boolean,
		code.Multiply:    code.Int,
		code.Divide:      code.Int,
		code.Add:         code.Int,
		code.Subtract:    code.Int,
	},
	{left: code.Rune, right: code.Rune}: {
		code.Equals:      code.Boolean,
		code.LessThan:    code.Boolean,
		code.GreaterThan: code.Boolean,
		code.Multiply:    code.Rune,
		code.Divide:      code.Rune,
		code.Add:         code.Rune,
		code.Subtract:    code.Rune,
	},
	{left: code.String, right: code.String}: {
		code.Equals: code.Boolean,
		code.Add:    code.String,
	},
	{left: code.Boolean, right: code.Boolean}: {
		code.And: code.Boolean,
		code.Or:  code.Boolean,
	},
}

func (m Mapper) MapUnaryOperation(original *code.UnaryOperation) (code.Type, error) {
	inputPrimitive, err := MapValueToPrimitiveType(original.Value)
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

func (m Mapper) MapBinaryOperation(original *code.BinaryOperation) (code.Type, error) {
	leftPrimitive, err := MapValueToPrimitiveType(original.Left)
	if err != nil {
		return nil, err
	}

	rightPrimitive, err := MapValueToPrimitiveType(original.Right)
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

func (m Mapper) MapProperty(original *code.Property) (code.Type, error) {
	ofModel, err := MapValueModelType(original.Of)
	if err != nil {
		return nil, err
	}

	return code.MapDefinitionNoError[code.Type](ofModel.ModelMetadata.Definition.FieldMap[original.Name], definition_to_type.Mapper{}), nil
}

func (m Mapper) MapVariable(original *code.Variable) (code.Type, error) {
	return code.MapDefinitionNoError[code.Type](original.Definition, definition_to_type.Mapper{}), nil
}

func (m Mapper) MapLookup(original *code.Lookup) (code.Type, error) {
	fromType, err := mapValueToType(original.From)
	if err != nil {
		return nil, err
	}

	switch original.LookupType {
	case code.LookupTypeList:
		return fromType.(*code.List).Base, nil
	case code.LookupTypeMap:
		return fromType.(*code.Map).Value, nil
	case code.LookupTypeString:
		return code.Rune, nil
	}

	panic("unreachable")
}

func (m Mapper) MapCall(original *code.Call) (code.Type, error) {
	return original.Definition.ReturnType, nil
}

func (m Mapper) MapNew(original *code.New) (code.Type, error) {
	return original.Model, nil
}

func (m Mapper) MapLiteralList(original *code.LiteralList) (code.Type, error) {
	if len(original.Items) == 0 {
		return nil, errors.New("literal list must have at least one element")
	}

	firstElementType, err := code.MapValue[code.Type](original.Items[0], m)
	if err != nil {
		return nil, err
	}

	for _, item := range original.Items[1:] {
		elementType, err := code.MapValue[code.Type](item, m)
		if err != nil {
			return nil, err
		}

		if elementType != firstElementType {
			return nil, errors.New("literal list must have elements of the same type")
		}
	}

	return &code.List{Base: firstElementType}, nil
}

func (m Mapper) MapLength(original *code.Length) (code.Type, error) {
	return code.Int, nil
}

func (m Mapper) MapLiteralMap(original *code.LiteralMap) (code.Type, error) {
	if len(original.Entries) == 0 {
		return nil, errors.New("literal map must have at least one entry")
	}

	firstEntryKeyType, err := code.MapValue[code.Type](original.Entries[0].Key, m)
	if err != nil {
		return nil, err
	}

	firstEntryValueType, err := code.MapValue[code.Type](original.Entries[0].Value, m)
	if err != nil {
		return nil, err
	}

	for _, item := range original.Entries[1:] {
		keyType, err := code.MapValue[code.Type](item.Key, m)
		if err != nil {
			return nil, err
		}

		if keyType != firstEntryKeyType {
			return nil, errors.New("literal map must have elements of the same key type")
		}

		valueType, err := code.MapValue[code.Type](item.Value, m)
		if err != nil {
			return nil, err
		}

		if valueType != firstEntryValueType {
			return nil, errors.New("literal map must have elements of the same value type")
		}
	}

	return &code.Map{
		Key:   firstEntryKeyType,
		Value: firstEntryValueType,
	}, nil
}
