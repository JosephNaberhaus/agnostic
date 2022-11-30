package mappers

import (
	"errors"
	"fmt"
	"github.com/JosephNaberhaus/agnostic/code"
	"regexp"
	"strings"
)

var nonAlphaMatcher = regexp.MustCompile("[^a-zA-Z]")

func validateModuleName(name string) error {
	if nonAlphaMatcher.MatchString(name) {
		return fmt.Errorf("module names should contain only alphabetic characters: '%s'", name)
	}

	if strings.ToUpper(name)[0] != name[0] {
		return fmt.Errorf("module names should start with a capital letter: '%s'", name)
	}

	return nil
}

func validateModelName(name string) error {
	if nonAlphaMatcher.MatchString(name) {
		return fmt.Errorf("model names should contain only alphabetic characters: '%s'", name)
	}

	if strings.ToUpper(name)[0] != name[0] {
		return fmt.Errorf("model names should start with a capital letter: '%s'", name)
	}

	return nil
}

func validateMethodName(name string) error {
	if nonAlphaMatcher.MatchString(name) {
		return fmt.Errorf("method names should contain only alphabetic characters: '%s'", name)
	}

	if strings.ToLower(name)[0] != name[0] {
		return fmt.Errorf("method names should start with a lowercase letter: '%s'", name)
	}

	return nil
}

func validateVariableName(name string) error {
	if nonAlphaMatcher.MatchString(name) {
		return fmt.Errorf("variable names should contain only alphabetic characters: '%s'", name)
	}

	if strings.ToLower(name)[0] != name[0] {
		return fmt.Errorf("variable names should start with a lowercase letter: '%s'", name)
	}

	return nil
}

func validateCondition(value code.Value) error {
	valueType, err := mapValueToType(value)
	if err != nil {
		return err
	}

	if valueType != code.Boolean {
		return fmt.Errorf("conditional must be boolean type")
	}

	return nil
}

func validateReturn(ret *code.Return) error {
	valueType, err := mapValueToType(ret.Value)
	if err != nil {
		return err
	}

	if valueType != ret.StatementMetadata.Parent.ReturnType {
		return errors.New("mismatched return types")
	}

	return nil
}

func validateFunction(function *code.FunctionDef) error {
	if function.ReturnType != code.Void {
		if len(function.Statements) == 0 {
			return errors.New("non-void function doesn't end with a return")
		}

		_, lastStatementIsReturn := function.Statements[len(function.Statements)-1].(*code.Return)
		if !lastStatementIsReturn {
			return errors.New("non-void function doesn't end with a return")
		}
	}

	return nil
}
