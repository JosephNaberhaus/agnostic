package golang

import (
	"fmt"
	writer "github.com/JosephNaberhaus/agnostic/implementations/code"
	"github.com/JosephNaberhaus/agnostic/test"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

// testModelWriter is a fake model writer that is used to access the code generation tools of modelWriter
var testModelWriter modelWriter

func assertionCode(assertion test.Assertion) writer.Code {
	switch a := assertion.(type) {
	case test.IsEqual:
		return writer.Line("assert.Equal(t, " + testModelWriter.valueString(a.Expected) + ", " + testModelWriter.valueString(a.Actual) + ")")
	case test.IsNotEqual:
		return writer.Line("assert.NotEqual(t, " + testModelWriter.valueString(a.NotExpected) + ", " + testModelWriter.valueString(a.Actual) + ")")
	case test.IsTrue:
		return writer.Line("assert.True(t, " + testModelWriter.valueString(a.Actual) + ")")
	case test.IsFalse:
		return writer.Line("assert.False(t, " + testModelWriter.valueString(a.Actual) + ")")
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
		writer.Line("package " + strings.ToLower(suite.Name)),
		writer.Line(""),
		writer.Line("import ("),
		writer.Block{
			writer.Line("\"github.com/stretchr/testify/assert\""),
			writer.Line("\"testing\""),
		},
		writer.Line(")"),
		writer.Line(""),
		writer.Line("func TestGolang_" + t.Name + "(t *testing.T) {"),
		writer.Block{
			writer.Group(testModelWriter.statementsCode(t.Before)),
			writer.Group(assertionsCode(t.Assertions)),
		},
		writer.Line("}"),
	}
}

func TestFileName(t test.Test) string {
	pascalCaseSeparator := regexp.MustCompile(`[A-Z][^A-Z]*`)
	pascalCaseParts := pascalCaseSeparator.FindAllString(t.Name, -1)

	lowercaseParts := make([]string, 0, len(pascalCaseParts))
	for _, pascalCasePart := range pascalCaseParts {
		lowercaseParts = append(lowercaseParts, strings.ToLower(pascalCasePart))
	}

	return strings.Join(lowercaseParts, "_") + "_test.go"
}

func InitTestSuiteDirectory(suite test.Suite) error {
	const modulePrefix = "github.com/JosephNaberhaus/agnostic/implementations/golang/tests/"
	err := exec.Command("go", "mod", "init", modulePrefix+strings.ToLower(suite.Name)).Run()
	if err != nil {
		return fmt.Errorf("error initializing go mod: %w", err)
	}

	err = exec.Command("go", "get", "github.com/stretchr/testify@v1.7.1").Run()
	if err != nil {
		return fmt.Errorf("error getting testify: %w", err)
	}

	return nil
}

func RunTests() error {
	testCommand := exec.Command("go", "test", "./...")
	testCommand.Stdout = os.Stdout
	testCommand.Stderr = os.Stderr

	return testCommand.Run()
}
