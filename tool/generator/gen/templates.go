package gen

import (
	_ "embed"
	"html/template"
	"os"
	"path/filepath"
	"strings"

	"github.com/JosephNaberhaus/agnostic/tool/generator/find"
	"github.com/JosephNaberhaus/agnostic/tool/generator/model"
)

const (
	astPackage   = "ast"
	astDirectory = "../../ast"

	codePackage   = "code"
	codeDirectory = "../../code"

	astFilename      = "ast_gen.go"
	codeFilename     = "code_gen.go"
	nodeTypeFilename = "node_type_gen.go"
	optionalFilename = "optional_gen.go"
)

//go:embed ast.go.tmpl
var astTemplate string

//go:embed code.go.tmpl
var codeTemplate string

//go:embed node_types.go.tmpl
var nodeTypesTemplate string

//go:embed optional.go.tmpl
var optionalTemplate string

func WriteAST(specs []model.Spec) error {
	astFile := filepath.Join(astDirectory, astFilename)
	err := executeTemplate(astTemplate, astFile, specs)
	if err != nil {
		return err
	}

	err = writeNodeTypes(specs, astPackage, astDirectory)
	if err != nil {
		return err
	}

	err = writeOptional(astPackage, astDirectory)
	if err != nil {
		return err
	}

	return nil
}

func WriteCode(specs []model.Spec) error {
	codeFile := filepath.Join(codeDirectory, codeFilename)
	err := executeTemplate(codeTemplate, codeFile, specs)
	if err != nil {
		return err
	}

	err = writeNodeTypes(specs, codePackage, codeDirectory)
	if err != nil {
		return err
	}

	err = writeOptional(codePackage, codeDirectory)
	if err != nil {
		return err
	}

	return nil
}

func writeNodeTypes(specs []model.Spec, packageName, outputDir string) error {
	data := struct {
		Package   string
		NodeTypes []string
	}{
		Package:   packageName,
		NodeTypes: find.AllNodeTypes(specs),
	}

	nodeTypesFile := filepath.Join(outputDir, nodeTypeFilename)
	return executeTemplate(nodeTypesTemplate, nodeTypesFile, data)
}

func writeOptional(packageName, outputDir string) error {
	data := struct {
		Package string
	}{
		Package: packageName,
	}

	optionalFile := filepath.Join(outputDir, optionalFilename)
	return executeTemplate(optionalTemplate, optionalFile, data)
}

func executeTemplate(templateText, outputFile string, data any) error {
	err := os.MkdirAll(filepath.Dir(outputFile), os.ModePerm)
	if err != nil {
		return err
	}

	tmpl := template.New("template ")
	tmpl.Funcs(template.FuncMap{
		"title": title,
	})

	tmpl, err = tmpl.Parse(templateText)
	if err != nil {
		return err
	}

	file, err := os.OpenFile(outputFile, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}

	err = tmpl.Execute(file, data)
	if err != nil {
		return err
	}

	return nil
}

func title(str string) string {
	return strings.ToUpper(str[:1]) + str[1:]
}
