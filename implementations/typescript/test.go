package typescript

import (
	"fmt"
	writer "github.com/JosephNaberhaus/agnostic/implementations/code"
	"github.com/JosephNaberhaus/agnostic/test"
	"github.com/JosephNaberhaus/agnostic/utils"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

// testModelWriter is a fake model writer that is used to access the code generation tools of modelWriter
var testModelWriter modelWriter

func assertionCode(assertion test.Assertion) writer.Code {
	switch a := assertion.(type) {
	case test.IsEqual:
		return writer.Line("expect(" + testModelWriter.valueString(a.Actual) + ").to.equal(" + testModelWriter.valueString(a.Expected) + ");")
	case test.IsNotEqual:
		return writer.Line("expect(" + testModelWriter.valueString(a.Actual) + ").to.not.equal(" + testModelWriter.valueString(a.NotExpected) + ");")
	case test.IsTrue:
		return writer.Line("expect(" + testModelWriter.valueString(a.Actual) + ").to.be.true;")
	case test.IsFalse:
		return writer.Line("expect(" + testModelWriter.valueString(a.Actual) + ").to.be.false;")
	default:
		panic(fmt.Errorf("unkown assertion type: \"%v\"", assertion))
	}
}

func assertionsCode(assertions []test.Assertion) []writer.Code {
	code := make([]writer.Code, 0, len(assertions))
	for _, assertion := range assertions {
		code = append(code, assertionCode(assertion))
	}

	return code
}

func TestCode(suite test.Suite, t test.Test) writer.Code {
	return writer.Group{
		writer.Line("import { expect } from 'chai';"),
		writer.Line(""),
		writer.Line("describe('TypeScript', () => {"),
		writer.Block{
			writer.Line("describe('" + suite.Name + "', () => {"),
			writer.Block{
				writer.Line("it('" + t.Name + "', () => {"),
				writer.Block{
					writer.Group(testModelWriter.statementsCode(t.Before)),
					writer.Group(assertionsCode(t.Assertions)),
				},
				writer.Line("});"),
			},
			writer.Line("});"),
		},
		writer.Line("});"),
	}
}

func TestFileName(t test.Test) string {
	pascalCaseSeparator := regexp.MustCompile(`[A-Z][^A-Z]*`)
	pascalCaseParts := pascalCaseSeparator.FindAllString(t.Name, -1)

	lowercaseParts := make([]string, 0, len(pascalCaseParts))
	for _, pascalCasePart := range pascalCaseParts {
		lowercaseParts = append(lowercaseParts, strings.ToLower(pascalCasePart))
	}

	return strings.Join(lowercaseParts, "_") + ".spec.ts"
}

func InitTestSuiteDirectory(suite test.Suite) error {
	templateFiles := []string{
		"implementations/typescript/templates/package.json.template",
		"implementations/typescript/templates/tsconfig.json.template",
	}

	data := map[string]string{
		"testName": strings.ToLower(suite.Name),
	}

	for _, templatePath := range templateFiles {
		fullTemplatePath := filepath.Join(utils.GitRootDir(), templatePath)
		err := utils.ExecuteTemplate(fullTemplatePath, data)
		if err != nil {
			return err
		}
	}

	err := exec.Command("npm", "i").Run()
	if err != nil {
		return fmt.Errorf("error installing npm packages: %w", err)
	}

	return nil
}

func RunTests() error {
	testCommand := exec.Command("npm", "run", "test")
	testCommand.Stdout = os.Stdout
	testCommand.Stderr = os.Stderr

	return testCommand.Run()
}
