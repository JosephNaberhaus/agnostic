package golang

import (
	"fmt"
	"github.com/JosephNaberhaus/agnostic"
	"github.com/JosephNaberhaus/agnostic/implementations/code"
	"github.com/JosephNaberhaus/agnostic/utils"
	"io/ioutil"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

type modelWriter struct {
	model            agnostic.Model
	referencedModels []agnostic.Model
}

func (m modelWriter) receiverName() string {
	return strings.ToLower(m.model.Name[:1])
}

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
		return m.valueString(v.Model) + "." + v.FieldName + "()"
	case agnostic.VariableValue:
		return string(v)
	case agnostic.ThisValue:
		return m.receiverName()
	case agnostic.ComputedValue:
		return m.valueString(v.Left) + " " + operatorString(v.Operator) + " " + m.valueString(v.Right)
	default:
		panic(fmt.Errorf("unkown value: \"%v\"", value))
	}
}

func (m *modelWriter) statementCode(statement agnostic.Statement) writer.Code {
	switch s := statement.(type) {
	case agnostic.Declare:
		return writer.Line(s.Name + " := " + m.valueString(s.Value))
	case agnostic.AssignVar:
		return writer.Line(m.valueString(s.Var) + " = " + m.valueString(s.Value))
	case agnostic.AssignField:
		return writer.Line(m.valueString(s.Model) + ".Set" + s.FieldName + "(" + m.valueString(s.Value) + ")")
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
			writer.Block(m.statementsCode(s.Statements)),
			writer.Line("}"),
		}
	case agnostic.If:
		return writer.Group{
			writer.Line("if " + m.valueString(s.Condition) + " {"),
			writer.Block(m.statementsCode(s.Statements)),
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
	for i, parameter := range method.Parameters {
		parameters.WriteString(parameter.Name)
		parameters.WriteString(" ")
		parameters.WriteString(m.typeString(parameter.Type))
		if i+1 != len(method.Parameters) {
			parameters.WriteString(", ")
		}
	}

	return writer.Group{
		writer.Line("func (" + m.receiverName() + " *" + m.model.Name + ") " + method.Name + "(" + parameters.String() + ")" + returnTypeSuffix + " {"),
		writer.Block(m.statementsCode(method.Statements)),
		writer.Line("}"),
		writer.Line(""),
	}
}

func fieldName(field agnostic.Field) string {
	return strings.ToLower(field.Name[:1]) + field.Name[1:]
}

func (m *modelWriter) fieldCode(field agnostic.Field) writer.Code {
	return writer.Line(fieldName(field) + " " + m.typeString(field.Type))
}

func (m *modelWriter) fieldGetterSetterCode(field agnostic.Field) writer.Code {
	methodName := strings.ToUpper(field.Name[:1]) + field.Name[1:]

	return writer.Group{
		writer.Line("func (" + m.receiverName() + " *" + m.model.Name + ") " + methodName + "() " + m.typeString(field.Type) + " {"),
		writer.Block{
			writer.Line("return " + m.receiverName() + "." + fieldName(field)),
		},
		writer.Line("}"),
		writer.Line(""),
		writer.Line("func (" + m.receiverName() + " *" + m.model.Name + ") Set" + methodName + "(newValue " + m.typeString(field.Type) + ") " + " {"),
		writer.Block{
			writer.Line(m.receiverName() + "." + fieldName(field) + " = newValue"),
		},
		writer.Line("}"),
	}
}

func ModelCode(model agnostic.Model) (writer.Code, error) {
	modelWriter := modelWriter{
		model: model,
	}

	fieldsCode := make([]writer.Code, 0, len(model.Fields))
	fieldGettersCode := make([]writer.Code, 0)
	for _, field := range model.Fields {
		fieldsCode = append(fieldsCode, modelWriter.fieldCode(field))
		fieldGettersCode = append(fieldGettersCode, modelWriter.fieldGetterSetterCode(field))
	}

	methodsCode := make([]writer.Code, 0, len(model.Methods))
	for _, method := range model.Methods {
		methodsCode = append(methodsCode, modelWriter.methodCode(method))
	}

	importsCode := make([]writer.Code, 0, len(modelWriter.referencedModels))
	for _, referencedModel := range modelWriter.referencedModels {
		path, err := modelImportPath(referencedModel)
		if err != nil {
			return nil, err
		}

		importsCode = append(importsCode, writer.Line("\""+path+"\""))
	}

	return writer.Group{
		writer.Line("package " + filepath.Base(model.Path)),
		writer.Line(""),
		writer.Line("import ("),
		writer.Block(importsCode),
		writer.Line(")"),
		writer.Line(""),
		writer.Line("type " + model.Name + " struct {"),
		writer.Block(fieldsCode),
		writer.Line("}"),
		writer.Line(""),
		writer.Group(fieldGettersCode),
		writer.Line(""),
		writer.Group(methodsCode),
	}, nil
}

func modelImportPath(m agnostic.Model) (string, error) {
	root, err := goModRoot(m.Path)
	if err != nil {
		return "", fmt.Errorf("no go mod root found for model at %s: %w", m.Path, err)
	}

	goModFilePath := filepath.Join(root, "go.mod")
	goModeFileContent, err := ioutil.ReadFile(goModFilePath)
	if err != nil {
		return "", err
	}

	moduleRegex := regexp.MustCompile(`^module (.*)$`)
	match := moduleRegex.Find(goModeFileContent)
	if len(match) != 2 {
		return "", fmt.Errorf("no match found when looking for the module path in \"%s\"", goModFilePath)
	}

	modulePath := string(match[1])
	moduleRelativePath, err := filepath.Rel(root, m.Path)
	if err != nil {
		return "", fmt.Errorf("couldn't find a relative path from \"%s\" to \"%s\"", root, m.Path)
	}

	return modulePath + "/" + moduleRelativePath, nil
}

func goModRoot(path string) (string, error) {
	for !utils.FileExists(filepath.Join(path, "go.mod")) {
		path = filepath.Dir(path)
		if path == "." {
			break
		}
	}

	if !utils.FileExists(filepath.Join(path, "go.mod")) {
		return "", fmt.Errorf("couldn't find a go.mod file")
	}

	return filepath.Join(path), nil
}
