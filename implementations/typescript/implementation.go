package typescript

import (
	"fmt"
	"github.com/JosephNaberhaus/agnostic"
	"github.com/JosephNaberhaus/agnostic/implementations/code"
	"path/filepath"
	"strconv"
	"strings"
)

type modelWriter struct {
	referencedModels []agnostic.Model
	model            agnostic.Model
}

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

func operatorString(operator agnostic.Operator) string {
	switch operator {
	case agnostic.Add:
		return "+"
	case agnostic.Subtract:
		return "-"
	case agnostic.Multiply:
		return "*"
	case agnostic.Divide:
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
		return string(v)
	case agnostic.ArrayElementValue:
		return m.valueString(v.Array) + "[" + m.valueString(v.Index) + "]"
	case agnostic.MapElementValue:
		return m.valueString(v.Map) + "[" + m.valueString(v.Key) + "]"
	case agnostic.FieldValue:
		return m.valueString(v.Model) + "." + v.FieldName
	case agnostic.VariableValue:
		return string(v)
	case agnostic.ThisValue:
		return "this"
	case agnostic.ComputedValue:
		return m.valueString(v.Left) + " " + operatorString(v.Operator) + " " + m.valueString(v.Right)
	default:
		panic(fmt.Errorf("unkown value: \"%v\"", value))
	}
}

func (m *modelWriter) statementCode(statement agnostic.Statement) writer.Code {
	switch s := statement.(type) {
	case agnostic.Declare:
		return writer.Line("let " + s.Name + " = " + m.valueString(s.Value) + ";")
	case agnostic.AssignVar:
		return writer.Line(m.valueString(s.Var) + " = " + m.valueString(s.Value) + ";")
	case agnostic.AssignField:
		return writer.Line(m.valueString(s.Model) + "." + s.FieldName + " = " + m.valueString(s.Value))
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

func (m *modelWriter) fieldCode(field agnostic.Field) writer.Code {
	getterAccessModifier := "public"
	if field.Access == agnostic.Private {
		getterAccessModifier = "private"
	}

	setterAccessModifier := "private"
	if field.Access == agnostic.Public {
		setterAccessModifier = "public"
	}

	return writer.Group{
		writer.Line("private _" + field.Name + ": " + m.typeString(field.Type) + " = " + m.typeZeroValue(field.Type) + ";"),
		writer.Line(""),
		writer.Line(getterAccessModifier + " get " + field.Name + "() {"),
		writer.Block{
			writer.Line("return this._" + field.Name + ";"),
		},
		writer.Line("}"),
		writer.Line(""),
		writer.Line(setterAccessModifier + " set " + field.Name + "(newValue: " + m.typeString(field.Type) + ") {"),
		writer.Block{
			writer.Line("this._" + field.Name + " = newValue;"),
		},
		writer.Line("}"),
		writer.Line(""),
	}
}

func ModelCode(model agnostic.Model) (writer.Code, error) {
	modelWriter := modelWriter{
		model: model,
	}

	fieldsCode := make([]writer.Code, 0, len(model.Fields))
	for _, field := range model.Fields {
		fieldsCode = append(fieldsCode, modelWriter.fieldCode(field))
	}

	methodsCode := make([]writer.Code, 0, len(model.Methods))
	for _, method := range model.Methods {
		methodsCode = append(methodsCode, modelWriter.methodCode(method))
	}

	importsCode := make([]writer.Code, 0, len(modelWriter.referencedModels))
	for _, referencedModel := range modelWriter.referencedModels {
		code, err := modelImport(model, referencedModel)
		if err != nil {
			return nil, fmt.Errorf("error finding import: %w", err)
		}

		importsCode = append(importsCode, code)
	}

	return writer.Group{
		writer.Group(importsCode),
		writer.Line(""),
		writer.Line("export class " + model.Name + " {"),
		writer.Block{
			writer.Group(fieldsCode),
			writer.Group(methodsCode),
		},
		writer.Line("}"),
		writer.Line(""),
	}, nil
}

func modelImport(cur, target agnostic.Model) (writer.Code, error) {
	relativePath, err := filepath.Rel(cur.Path, target.Path)
	if err != nil {
		return nil, fmt.Errorf("error finding a path from \"%s\" to \"%s\": %w", cur.Path, target.Path, err)
	}

	return writer.Line("import {" + target.Name + "} from \"" + filepath.Join(relativePath, target.Name) + "\";"), nil
}
