package golang

import (
	"fmt"
	"github.com/JosephNaberhaus/agnostic"
	"path/filepath"
)

func primitiveTypeString(primitiveType agnostic.PrimitiveType) string {
	switch primitiveType {
	case agnostic.BooleanType:
		return "bool"
	case agnostic.IntType:
		return "int"
	case agnostic.FloatType:
		return "float"
	case agnostic.StringType:
		return "string"
	default:
		panic(fmt.Errorf("unkown primitive type: \"%v\"", primitiveType))
	}
}

func (m *modelWriter) typeString(_type agnostic.Type) string {
	switch t := _type.(type) {
	case agnostic.PrimitiveType:
		return primitiveTypeString(t)
	case agnostic.ArrayType:
		return "[]" + m.typeString(t.ElementType)
	case agnostic.MapType:
		return "map[" + m.typeString(t.KeyType) + "]" + m.typeString(t.ValueType)
	case agnostic.ModelType:
		m.referencedModels = append(m.referencedModels, t.Model)
		return filepath.Base(t.Model.Path) + "." + t.Model.Name
	default:
		panic(fmt.Errorf("unkown type: \"%v\"", _type))
	}
}
