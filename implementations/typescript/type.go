package typescript

import (
	"fmt"
	"github.com/JosephNaberhaus/agnostic"
	"path/filepath"
)

func primitiveTypeString(primitiveType agnostic.PrimitiveType) string {
	switch primitiveType {
	case agnostic.BooleanType:
		return "boolean"
	case agnostic.IntType:
		fallthrough
	case agnostic.FloatType:
		return "number"
	case agnostic.StringType:
		return "String"
	default:
		panic(fmt.Errorf("unkown primitive type: \"%v\"", primitiveType))
	}
}

func (m *modelWriter) typeString(_type agnostic.Type) string {
	switch t := _type.(type) {
	case agnostic.PrimitiveType:
		return primitiveTypeString(t)
	case agnostic.ArrayType:
		return m.typeString(t.ElementType) + "[]"
	case agnostic.MapType:
		return "Map<" + m.typeString(t.KeyType) + ", " + m.typeString(t.ValueType) + ">"
	case agnostic.ModelType:
		m.referencedModels = append(m.referencedModels, t.Model)
		return filepath.Base(m.model.Path) + "." + t.Model.Name
	default:
		panic(fmt.Errorf("unkown type: \"%v\"", _type))
	}
}

func primitiveTypeZeroValue(primitiveType agnostic.PrimitiveType) string {
	switch primitiveType {
	case agnostic.BooleanType:
		return "false"
	case agnostic.IntType:
		fallthrough
	case agnostic.FloatType:
		return "0"
	case agnostic.StringType:
		return `""`
	default:
		panic(fmt.Errorf("unkown primitive type: \"%v\"", primitiveType))
	}
}

func (m *modelWriter) typeZeroValue(_type agnostic.Type) string {
	switch t := _type.(type) {
	case agnostic.PrimitiveType:
		return primitiveTypeZeroValue(t)
	case agnostic.ArrayType:
		return "[]"
	case agnostic.MapType:
		return "new Map<>()"
	case agnostic.ModelType:
		m.referencedModels = append(m.referencedModels, t.Model)
		return "new " + t.Model.Name + "()"
	default:
		panic(fmt.Errorf("unkown type: \"%v\"", _type))
	}
}
