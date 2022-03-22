package golang

import (
	"fmt"
	"github.com/JosephNaberhaus/agnostic"
	"github.com/JosephNaberhaus/agnostic/implementations/code"
	"github.com/JosephNaberhaus/agnostic/utils"
	"io/ioutil"
	"path/filepath"
	"regexp"
	"strings"
)

type modelWriter struct {
	model            agnostic.Model
	referencedModels []agnostic.Model
}

func (m modelWriter) receiverName() string {
	return strings.ToLower(m.model.Name[:1])
}

func fieldMethodName(name string) string {
	return strings.ToUpper(name[:1]) + name[1:]
}

func fieldName(name string) string {
	return strings.ToLower(name[:1]) + name[1:]
}

func (m *modelWriter) fieldCode(field agnostic.Field) writer.Code {
	return writer.Line(fieldName(field.Name) + " " + m.typeString(field.Type))
}

func (m *modelWriter) fieldGetterSetterCode(field agnostic.Field) writer.Code {
	getterSetterCode := make([]writer.Code, 0)

	if field.Access != agnostic.Private {
		getterSetterCode = append(getterSetterCode, writer.Group{
			writer.Line("func (" + m.receiverName() + " *" + m.model.Name + ") " + fieldMethodName(field.Name) + "() " + m.typeString(field.Type) + " {"),
			writer.Block{
				writer.Line("return " + m.receiverName() + "." + fieldName(field.Name)),
			},
			writer.Line("}"),
			writer.Line(""),
		})
	}

	if field.Access == agnostic.Public {
		getterSetterCode = append(getterSetterCode, writer.Group{
			writer.Line("func (" + m.receiverName() + " *" + m.model.Name + ") Set" + fieldMethodName(field.Name) + "(newValue " + m.typeString(field.Type) + ") " + " {"),
			writer.Block{
				writer.Line(m.receiverName() + "." + fieldName(field.Name) + " = newValue"),
			},
			writer.Line("}"),
		})
	}

	return writer.Group(getterSetterCode)
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
