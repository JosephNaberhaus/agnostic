package interpreter

import (
	"github.com/JosephNaberhaus/agnostic/code"
)

type valueMapper struct {
	runtime runtime
}

func (v *valueMapper) MapLiteralInt(original *code.LiteralInt) Value {
	return Value{Value: original.Value}
}

func (v *valueMapper) MapLiteralRune(original *code.LiteralRune) Value {
	return Value{Value: original.Value}
}

func (v *valueMapper) MapLiteralString(original *code.LiteralString) Value {
	return Value{Value: original.Value}
}

func (v *valueMapper) MapLiteralBool(original *code.LiteralBool) Value {
	return Value{Value: original.Value}
}

func (v *valueMapper) MapLiteralList(original *code.LiteralList) Value {
	list := List(code.MapValuesNoError[Value](original.Items, v))
	return Value{Value: &list}
}

func (v *valueMapper) MapLiteralMap(original *code.LiteralMap) Value {
	mapValue := map[Value]Value{}
	for _, entry := range original.Entries {
		key := code.MapValueNoError[Value](entry.Key, v)
		value := code.MapValueNoError[Value](entry.Value, v)
		mapValue[key] = value
	}

	return Value{Value: mapValue}
}

func (v *valueMapper) MapLiteralSet(original *code.LiteralSet) Value {
	setValue := map[Value]struct{}{}
	for _, item := range original.Items {
		setValue[code.MapValueNoError[Value](item, v)] = struct{}{}
	}

	return Value{Value: setValue}
}

func (v *valueMapper) MapEmptyList(_ *code.EmptyList) Value {
	return Value{Value: &List{}}
}

func (v *valueMapper) MapEmptySet(_ *code.EmptySet) Value {
	return Value{Value: map[Value]struct{}{}}
}

func (v *valueMapper) MapVariable(original *code.Variable) Value {
	return v.runtime.variableByName(original.Name)
}

func (v *valueMapper) MapProperty(original *code.Property) Value {
	return code.MapValueNoError[Value](original.Of, v).Value.(Model).Properties[original.Name]
}

func (v *valueMapper) MapUnaryOperation(original *code.UnaryOperation) Value {
	value := code.MapValueNoError[Value](original.Value, v)

	switch original.Operator {
	case code.Not:
		return Value{Value: !value.Value.(bool)}
	case code.CastToInt:
		switch value := value.Value.(type) {
		case rune:
			return Value{Value: int64(value)}
		}
	case code.CastToString:
		switch value := value.Value.(type) {
		case rune:
			return Value{Value: string(value)}
		}
	}

	panic("unreachable")
}

func (v *valueMapper) MapBinaryOperation(original *code.BinaryOperation) Value {
	left := code.MapValueNoError[Value](original.Left, v)
	right := code.MapValueNoError[Value](original.Right, v)

	switch original.Operator {
	case code.Add:
		switch left := left.Value.(type) {
		case rune:
			switch right := right.Value.(type) {
			case rune:
				return Value{Value: left + right}
			}
		case int64:
			switch right := right.Value.(type) {
			case int64:
				return Value{Value: left + right}
			}
		case string:
			switch right := right.Value.(type) {
			case string:
				return Value{Value: left + right}
			}
		}
	case code.Subtract:
		switch left := left.Value.(type) {
		case rune:
			switch right := right.Value.(type) {
			case rune:
				return Value{Value: left - right}
			}
		case int64:
			switch right := right.Value.(type) {
			case int64:
				return Value{Value: left - right}
			}
		}
	case code.Multiply:
		switch left := left.Value.(type) {
		case rune:
			switch right := right.Value.(type) {
			case rune:
				return Value{Value: left * right}
			}
		case int64:
			switch right := right.Value.(type) {
			case int64:
				return Value{Value: left * right}
			}
		}
	case code.Divide:
		switch left := left.Value.(type) {
		case rune:
			switch right := right.Value.(type) {
			case rune:
				return Value{Value: left / right}
			}
		case int64:
			switch right := right.Value.(type) {
			case int64:
				return Value{Value: left / right}
			}
		}
	case code.Equal:
		switch left := left.Value.(type) {
		case rune:
			switch right := right.Value.(type) {
			case rune:
				return Value{Value: left == right}
			}
		case int64:
			switch right := right.Value.(type) {
			case int64:
				return Value{Value: left == right}
			}
		case string:
			switch right := right.Value.(type) {
			case string:
				return Value{Value: left == right}
			}
		case bool:
			switch right := right.Value.(type) {
			case bool:
				return Value{Value: left == right}
			}
		}
	case code.NotEqual:
		switch left := left.Value.(type) {
		case rune:
			switch right := right.Value.(type) {
			case rune:
				return Value{Value: left != right}
			}
		case int64:
			switch right := right.Value.(type) {
			case int64:
				return Value{Value: left != right}
			}
		case string:
			switch right := right.Value.(type) {
			case string:
				return Value{Value: left != right}
			}
		case bool:
			switch right := right.Value.(type) {
			case bool:
				return Value{Value: left != right}
			}
		}
	case code.LessThan:
		switch left := left.Value.(type) {
		case rune:
			switch right := right.Value.(type) {
			case rune:
				return Value{Value: left < right}
			}
		case int64:
			switch right := right.Value.(type) {
			case int64:
				return Value{Value: left < right}
			}
		}
	case code.LessThanOrEqualTo:
		switch left := left.Value.(type) {
		case rune:
			switch right := right.Value.(type) {
			case rune:
				return Value{Value: left <= right}
			}
		case int64:
			switch right := right.Value.(type) {
			case int64:
				return Value{Value: left <= right}
			}
		}
	case code.GreaterThan:
		switch left := left.Value.(type) {
		case rune:
			switch right := right.Value.(type) {
			case rune:
				return Value{Value: left > right}
			}
		case int64:
			switch right := right.Value.(type) {
			case int64:
				return Value{Value: left > right}
			}
		}
	case code.GreaterThanOrEqualTo:
		switch left := left.Value.(type) {
		case rune:
			switch right := right.Value.(type) {
			case rune:
				return Value{Value: left >= right}
			}
		case int64:
			switch right := right.Value.(type) {
			case int64:
				return Value{Value: left >= right}
			}
		}
	case code.Or:
		switch left := left.Value.(type) {
		case bool:
			switch right := right.Value.(type) {
			case bool:
				return Value{Value: left || right}
			}
		}
	case code.And:
		switch left := left.Value.(type) {
		case bool:
			switch right := right.Value.(type) {
			case bool:
				return Value{Value: left && right}
			}
		}
	}

	panic("unreachable")
}

func (v *valueMapper) MapCall(original *code.Call) Value {
	callable := code.MapCallableNoError[callable](original.Function, &callableMapper{runtime: v.runtime})
	arguments := code.MapValuesNoError[Value](original.Arguments, v)

	return v.runtime.execute(
		callable.function,
		newCallScope(callable.model, callable.function, arguments),
	)
}

func (v *valueMapper) MapLookup(original *code.Lookup) Value {
	fromValue := code.MapValueNoError[Value](original.From, v)
	keyValue := code.MapValueNoError[Value](original.Key, v)

	switch original.LookupType {
	case code.LookupTypeList:
		return fromValue.Value.(*List).Get(keyValue.Value.(int64))
	case code.LookupTypeMap:
		return fromValue.Value.(*Map).Get(keyValue)
	case code.LookupTypeString:
		return Value{Value: []rune(fromValue.Value.(string))[keyValue.Value.(int64)]}
	}

	panic("unreachable")
}

func (v *valueMapper) MapNew(original *code.New) Value {
	return NewModel(original.Model.Definition)
}

func (v *valueMapper) MapLength(original *code.Length) Value {
	valueValue := code.MapValueNoError[Value](original.Value, v)

	switch original.LengthType {
	case code.LengthTypeString:
		return Value{Value: int64(len([]rune(valueValue.Value.(string))))}
	case code.LengthTypeList:
		return valueValue.Value.(*List).Length()
	}

	panic("unreachable")
}

func (v *valueMapper) MapSetContains(original *code.SetContains) Value {
	setValue := code.MapValueNoError[Value](original.Set, v)
	valueValue := code.MapValueNoError[Value](original.Value, v)

	return Value{Value: setValue.Value.(*Set).Contains(valueValue)}
}

func (v *valueMapper) MapPop(original *code.Pop) Value {
	valueValue := code.MapValueNoError[Value](original.Value, v)
	return valueValue.Value.(*List).Pop()
}

func (v *valueMapper) MapNull(_ *code.Null) Value {
	return Value{}
}

func (v *valueMapper) MapSelf(_ *code.Self) Value {
	return Value{Value: v.runtime.self()}
}

func (v *valueMapper) MapLiteralStruct(original *code.LiteralStruct) Value {
	modelValue := NewModel(original.Definition)

	for _, property := range original.Properties {
		modelValue.Value.(*Model).Properties[property.Name] = code.MapValueNoError[Value](property.Value, v)
	}

	return modelValue
}
