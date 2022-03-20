package typescript

import (
	"fmt"
	"github.com/JosephNaberhaus/agnostic"
	"github.com/JosephNaberhaus/agnostic/implementations/code"
	"path/filepath"
	"strings"
)

type modelWriter struct {
	referencedModels []agnostic.Model
	model            agnostic.Model
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
		writer.Line("private _" + fieldName(field.Name) + ": " + m.typeString(field.Type) + " = " + m.typeZeroValue(field.Type) + ";"),
		writer.Line(""),
		writer.Line(getterAccessModifier + " get " + fieldName(field.Name) + "() {"),
		writer.Block{
			writer.Line("return this._" + fieldName(field.Name) + ";"),
		},
		writer.Line("}"),
		writer.Line(""),
		writer.Line(setterAccessModifier + " set " + fieldName(field.Name) + "(newValue: " + m.typeString(field.Type) + ") {"),
		writer.Block{
			writer.Line("this._" + fieldName(field.Name) + " = newValue;"),
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

func fieldName(name string) string {
	return strings.ToLower(name[:1]) + name[1:]
}

func modelImport(cur, target agnostic.Model) (writer.Code, error) {
	relativePath, err := filepath.Rel(cur.Path, target.Path)
	if err != nil {
		return nil, fmt.Errorf("error finding a path from \"%s\" to \"%s\": %w", cur.Path, target.Path, err)
	}

	return writer.Line("import {" + target.Name + "} from \"" + filepath.Join(relativePath, target.Name) + "\";"), nil
}
