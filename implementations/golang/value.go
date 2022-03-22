package golang

import (
	"fmt"
	"github.com/JosephNaberhaus/agnostic"
	"strconv"
)

func operatorString(operator agnostic.Operator) string {
	switch operator {
	case agnostic.Add:
		return "+"
	case agnostic.Subtract:
		return "-"
	case agnostic.Multiply:
		return "*"
	case agnostic.IntegerDivision:
		return "/"
	case agnostic.Modulo:
		return "%"
	case agnostic.Equal:
		return "=="
	case agnostic.NotEqual:
		return "!="
	case agnostic.GreaterThan:
		return ">"
	case agnostic.GreaterThanOrEqualTo:
		return ">="
	case agnostic.LessThan:
		return "<"
	case agnostic.LessThanOrEqualTo:
		return "<="
	default:
		panic(fmt.Errorf("unkown operator: \"%v\"", operator))
	}
}

func (m *modelWriter) valueString(value agnostic.Value) string {
	switch v := value.(type) {
	case agnostic.BooleanLiteralValue:
		return strconv.FormatBool(bool(v))
	case agnostic.IntLiteralValue:
		return strconv.FormatInt(int64(v), 10)
	case agnostic.FloatLiteralValue:
		return strconv.FormatFloat(float64(v), 'f', -1, 64)
	case agnostic.StringLiteralValue:
		return "\"" + string(v) + "\""
	case agnostic.ArrayElementValue:
		return m.valueString(v.Array) + "[" + m.valueString(v.Index) + "]"
	case agnostic.MapElementValue:
		return m.valueString(v.Map) + "[" + m.valueString(v.Key) + "]"
	case agnostic.FieldValue:
		return m.valueString(v.Model) + "." + fieldMethodName(v.FieldName) + "()"
	case agnostic.OwnField:
		return m.receiverName() + "." + fieldName(string(v))
	case agnostic.VariableValue:
		return string(v)
	case agnostic.ComputedValue:
		return m.valueString(v.Left) + " " + operatorString(v.Operator) + " " + m.valueString(v.Right)
	default:
		panic(fmt.Errorf("unkown value: \"%v\"", value))
	}
}
