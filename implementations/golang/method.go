package golang

import (
	"fmt"
	"github.com/JosephNaberhaus/agnostic"
	writer "github.com/JosephNaberhaus/agnostic/implementations/writer"
	"strings"
)

func (m *modelWriter) statementCode(statement agnostic.Statement) writer.Code {
	switch s := statement.(type) {
	case agnostic.Declare:
		return writer.Group{
			writer.Line(s.Name + " := " + m.valueString(s.Value)),
			writer.Line("_ = " + s.Name),
		}
	case agnostic.AssignVar:
		return writer.Line(m.valueString(s.Var) + " = " + m.valueString(s.Value))
	case agnostic.AssignField:
		return writer.Line(m.valueString(s.Model) + ".Set" + fieldMethodName(s.FieldName) + "(" + m.valueString(s.Value) + ")")
	case agnostic.AssignOwnField:
		return writer.Line(m.receiverName() + "." + fieldName(s.FieldName) + " = " + m.valueString(s.Value))
	case agnostic.AppendValue:
		arrayValueString := m.valueString(s.Array)
		return writer.Line(arrayValueString + " = append(" + arrayValueString + ", " + m.valueString(s.ToAppend) + ")")
	case agnostic.AppendArray:
		arrayValueString := m.valueString(s.Array)
		return writer.Line(arrayValueString + " = append(" + arrayValueString + ", " + m.valueString(s.ToAppend) + "...)")
	case agnostic.RemoveValue:
		arrayValueString := m.valueString(s.Array)
		indexValueString := m.valueString(s.Index)
		left := arrayValueString + "[:" + indexValueString + "]"
		right := arrayValueString + "[" + indexValueString + "+1:]"
		return writer.Line(arrayValueString + " = append(" + left + ", " + right + "...)")
	case agnostic.MapPut:
		return writer.Line(m.valueString(s.Map) + "[" + m.valueString(s.Key) + "] = " + m.valueString(s.Value))
	case agnostic.MapDelete:
		return writer.Line("delete(" + m.valueString(s.Map) + ", " + m.valueString(s.Key) + ")")
	case agnostic.ForEach:
		return writer.Group{
			writer.Line("for "+s.IndexVariableName+", "+s.ValueVariableName+" := range "+m.valueString(s.Array)) + " {",
			writer.Block{
				writer.Line("_, _ = " + s.IndexVariableName + ", " + s.ValueVariableName),
				writer.Group(m.statementsCode(s.Statements)),
			},
			writer.Line("}"),
		}
	case agnostic.If:
		return writer.Group{
			writer.Line("if " + m.valueString(s.Condition) + " {"),
			writer.Block(m.statementsCode(s.Statements)),
			writer.Line("}"),
		}
	case agnostic.IfElse:
		return writer.Group{
			writer.Line("if " + m.valueString(s.Condition) + " {"),
			writer.Block(m.statementsCode(s.TrueStatements)),
			writer.Line("} else {"),
			writer.Block(m.statementsCode(s.FalseStatements)),
			writer.Line("}"),
		}
	case agnostic.Return:
		return writer.Line("return " + m.valueString(s.ToReturn))
	default:
		panic(fmt.Errorf("unkown statement \"%v\"", statement))
	}
}

func (m *modelWriter) statementsCode(statements []agnostic.Statement) []writer.Code {
	code := make([]writer.Code, 0, len(statements))
	for _, statement := range statements {
		code = append(code, m.statementCode(statement))
	}

	return code
}

func (m *modelWriter) methodCode(method agnostic.Method) writer.Group {
	returnTypeSuffix := ""
	if method.Returns != nil {
		returnTypeSuffix = " " + m.typeString(method.Returns)
	}

	parameters := strings.Builder{}
	useParametersCode := make([]writer.Code, 0, len(method.Parameters))
	for i, parameter := range method.Parameters {
		parameters.WriteString(parameter.Name)
		parameters.WriteString(" ")
		parameters.WriteString(m.typeString(parameter.Type))
		if i+1 != len(method.Parameters) {
			parameters.WriteString(", ")
		}

		useParametersCode = append(useParametersCode, writer.Line("_ = "+parameter.Name))
	}

	return writer.Group{
		writer.Line("func (" + m.receiverName() + " *" + m.model.Name + ") " + method.Name + "(" + parameters.String() + ")" + returnTypeSuffix + " {"),
		writer.Block{
			writer.Group(useParametersCode),
			writer.Group(m.statementsCode(method.Statements)),
		},
		writer.Line("}"),
		writer.Line(""),
	}
}
