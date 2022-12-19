package value_to_type

import (
	"errors"
	"github.com/JosephNaberhaus/agnostic/code"
	"github.com/JosephNaberhaus/agnostic/code/mappers/callable_def_to_return_type"
	"github.com/JosephNaberhaus/agnostic/code/mappers/definition_to_type"
	"reflect"
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
		code.Not:  code.Boolean,
		code.Hash: code.Int,
	},
	code.Int: {
		code.Negate: code.Int,
		code.Hash:   code.Int,
	},
	code.Rune: {
		code.CastToInt:    code.Int,
		code.CastToString: code.String,
		code.Hash:         code.Int,
	},
	code.String: {
		code.Hash: code.Int,
	},
}

// binaryOperatorMappings maps from a primitive type into another map containing the binary operations that can be done
// and the output primitive that they produce.
var binaryOperatorMappings = map[binaryOperatorPair]map[code.BinaryOperator]code.Primitive{
	{left: code.Int, right: code.Int}: {
		code.Equal:                code.Boolean,
		code.NotEqual:             code.Boolean,
		code.LessThan:             code.Boolean,
		code.LessThanOrEqualTo:    code.Boolean,
		code.GreaterThan:          code.Boolean,
		code.GreaterThanOrEqualTo: code.Boolean,
		code.Multiply:             code.Int,
		code.Divide:               code.Int,
		code.Modulo:               code.Int,
		code.Add:                  code.Int,
		code.Subtract:             code.Int,
	},
	{left: code.Rune, right: code.Rune}: {
		code.Equal:                code.Boolean,
		code.NotEqual:             code.Boolean,
		code.LessThan:             code.Boolean,
		code.LessThanOrEqualTo:    code.Boolean,
		code.GreaterThan:          code.Boolean,
		code.GreaterThanOrEqualTo: code.Boolean,
		code.Multiply:             code.Rune,
		code.Divide:               code.Rune,
		code.Modulo:               code.Int,
		code.Add:                  code.Rune,
		code.Subtract:             code.Rune,
	},
	{left: code.String, right: code.String}: {
		code.Equal:    code.Boolean,
		code.NotEqual: code.Boolean,
		code.Add:      code.String,
	},
	{left: code.Boolean, right: code.Boolean}: {
		code.Equal:    code.Boolean,
		code.NotEqual: code.Boolean,
		code.And:      code.Boolean,
		code.Or:       code.Boolean,
	},
}

func (m Mapper) MapUnaryOperation(original *code.UnaryOperation) (code.Type, error) {
	valueType, err := code.MapValue[code.Type](original.Value, m)
	if err != nil {
		return nil, err
	}

	switch valueType := valueType.(type) {
	case *code.Model:
		switch original.Operator {
		case code.Hash:
			return code.Int, nil
		}
	case *code.List:
		switch original.Operator {
		case code.Hash:
			return code.Int, nil
		}
	case *code.Set:
		switch original.Operator {
		case code.Hash:
			return code.Int, nil
		}
	case code.Primitive:
		outputType, ok := unaryOperatorMappings[valueType][original.Operator]
		if !ok {
			// TODO improve
			return nil, errors.New("cannot use the operator with the value")
		}

		return outputType, nil
	}

	return nil, errors.New("invalid type")
}

func (m Mapper) MapBinaryOperation(original *code.BinaryOperation) (code.Type, error) {
	leftType, err := mapValueToType(original.Left)
	if err != nil {
		return nil, err
	}

	rightType, err := mapValueToType(original.Right)
	if err != nil {
		return nil, err
	}

	// Allow any type to be equality compared to itself. The validator will further restrict this later on.
	isEquality := original.Operator == code.Equal || original.Operator == code.NotEqual
	if isEquality && reflect.DeepEqual(leftType, rightType) {
		return code.Boolean, nil
	}

	leftPrimitive, leftIsPrimitive := leftType.(code.Primitive)
	rightPrimitive, rightIsPrimitive := leftType.(code.Primitive)

	if leftIsPrimitive && rightIsPrimitive {
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

	return nil, errors.New("unsupported types for binary operation")
}

func (m Mapper) MapProperty(original *code.Property) (code.Type, error) {
	ofModel, err := MapValueModelType(original.Of)
	if err != nil {
		return nil, err
	}

	field, ok := ofModel.ModelMetadata.Definition.FieldMap[original.Name]
	if !ok {
		return nil, errors.New("unknown field " + original.Name)
	}

	return code.MapDefinitionNoError[code.Type](field, definition_to_type.Mapper{}), nil
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

		if !reflect.DeepEqual(elementType, firstElementType) {
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

func (m Mapper) MapLiteralSet(original *code.LiteralSet) (code.Type, error) {
	if len(original.Items) == 0 {
		return nil, errors.New("literal set must have at least one element")
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
			return nil, errors.New("literal set must have elements of the same type")
		}
	}

	return &code.Set{Base: firstElementType}, nil
}

func (m Mapper) MapEmptyList(original *code.EmptyList) (code.Type, error) {
	return &code.List{Base: original.Type}, nil
}

func (m Mapper) MapEmptySet(original *code.EmptySet) (code.Type, error) {
	return &code.Set{Base: original.Type}, nil
}

func (m Mapper) MapSetContains(original *code.SetContains) (code.Type, error) {
	return code.Boolean, nil
}

func (m Mapper) MapPop(original *code.Pop) (code.Type, error) {
	valueType, err := code.MapValue[code.Type](original.Value, m)
	if err != nil {
		return nil, err
	}

	switch valueType := valueType.(type) {
	case *code.List:
		return valueType.Base, nil
	default:
		return nil, errors.New("pop can only be used on a list")
	}
}

func (m Mapper) MapLiteralBool(original *code.LiteralBool) (code.Type, error) {
	return code.Boolean, nil
}

func (m Mapper) MapNull(original *code.Null) (code.Type, error) {
	switch parent := original.Parent.(type) {
	case *code.Return:
		return code.MapCallableDefNoError[code.Type](parent.CallableDef, callable_def_to_return_type.Mapper{}), nil
	case *code.BinaryOperation:
		if parent.Operator != code.Equal && parent.Operator != code.NotEqual {
			return nil, errors.New("the null literal can only be used for equality comparisons")
		}

		_, leftIsNull := parent.Left.(*code.Null)
		_, rightIsNull := parent.Right.(*code.Null)
		if leftIsNull && rightIsNull {
			return nil, errors.New("the null literal can not be compared to a null literal")
		}

		if leftIsNull {
			return code.MapValue[code.Type](parent.Right, m)
		} else {
			return code.MapValue[code.Type](parent.Left, m)
		}
	default:
		return nil, errors.New("unsupported usage of null")
	}
}

func (m Mapper) MapSelf(original *code.Self) (code.Type, error) {
	return &code.Model{
		Name: original.ParentModel.Name,
		ModelMetadata: code.ModelMetadata{
			Definition: original.ParentModel,
		},
	}, nil
}

func (m Mapper) MapLiteralStruct(original *code.LiteralStruct) (code.Type, error) {
	return &code.Model{
		Name: original.Definition.Name,
		ModelMetadata: code.ModelMetadata{
			Definition: original.Definition,
		},
	}, nil
}
