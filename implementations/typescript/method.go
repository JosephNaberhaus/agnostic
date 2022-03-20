package typescript

import (
	"fmt"
	"github.com/JosephNaberhaus/agnostic"
	writer "github.com/JosephNaberhaus/agnostic/implementations/code"
	"strings"
)

func (m *modelWriter) statementCode(statement agnostic.Statement) writer.Code {
	switch s := statement.(type) {
	case agnostic.Declare:
		return writer.Line("let " + s.Name + " = " + m.valueString(s.Value) + ";")
	case agnostic.AssignVar:
		return writer.Line(m.valueString(s.Var) + " = " + m.valueString(s.Value) + ";")
	case agnostic.AssignField:
		return writer.Line(m.valueString(s.Model) + "." + fieldName(s.FieldName) + " = " + m.valueString(s.Value))
	case agnostic.AssignOwnField:
		return writer.Line("this." + fieldName(s.FieldName) + " = " + m.valueString(s.Value))
	case agnostic.AppendValue:
		return writer.Line(m.valueString(s.Array) + ".push(" + m.valueString(s.ToAppend) + ");")
	case agnostic.AppendArray:
		return writer.Line(m.valueString(s.Array) + ".push(..." + m.valueString(s.ToAppend) + ");")
	case agnostic.RemoveValue:
		return writer.Line(m.valueString(s.Array) + ".slice(" + m.valueString(s.Index) + ", 1);")
	case agnostic.MapPut:
		return writer.Line(m.valueString(s.Map)+".set("+m.valueString(s.Key)+", "+m.valueString(s.Value)) + ");"
	case agnostic.MapDelete:
		return writer.Line(m.valueString(s.Map) + ".delete(" + m.valueString(s.Key) + ");")
	case agnostic.ForEach:
		return writer.Group{
			writer.Line(m.valueString(s.Array) + ".foreach((" + s.ValueVariableName + ", " + s.IndexVariableName + ") => {"),
			m.statementsCode(s.Statements),
			writer.Line("});"),
		}
	case agnostic.If:
		return writer.Group{
			writer.Line("if (" + m.valueString(s.Condition) + ") {"),
			m.statementsCode(s.Statements),
			writer.Line("}"),
		}
	case agnostic.IfElse:
		return writer.Group{
			writer.Line("if (" + m.valueString(s.Condition) + ") {"),
			m.statementsCode(s.TrueStatements),
			writer.Line("} else {"),
			m.statementsCode(s.FalseStatements),
			writer.Line("}"),
		}
	case agnostic.Return:
		return writer.Line("return " + m.valueString(s.ToReturn))
	default:
		panic(fmt.Errorf("unkown statement \"%v\"", statement))
	}
}

func (m *modelWriter) statementsCode(statements []agnostic.Statement) writer.Code {
	code := make([]writer.Code, 0, len(statements))
	for _, statement := range statements {
		code = append(code, m.statementCode(statement))
	}

	return writer.Group(code)
}

func (m *modelWriter) methodCode(method agnostic.Method) writer.Group {
	returnType := "void"
	if method.Returns != nil {
		returnType = m.typeString(method.Returns)
	}

	parameters := strings.Builder{}
	for i, parameter := range method.Parameters {
		parameters.WriteString(parameter.Name)
		parameters.WriteString(": ")
		parameters.WriteString(m.typeString(parameter.Type))
		if i+1 != len(method.Parameters) {
			parameters.WriteString(", ")
		}
	}

	return writer.Group{
		writer.Line(method.Name + "(" + parameters.String() + "): " + returnType + " {"),
		writer.Block{
			m.statementsCode(method.Statements),
		},
		writer.Line("}"),
		writer.Line(""),
	}
}
