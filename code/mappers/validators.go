package mappers

import (
	"errors"
	"fmt"
	"github.com/JosephNaberhaus/agnostic/code"
	"github.com/JosephNaberhaus/agnostic/code/mappers/callable_def_to_return_type"
	"github.com/JosephNaberhaus/agnostic/code/mappers/value_to_type"
	"reflect"
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
	valueType, err := code.MapValue[code.Type](value, value_to_type.Mapper{})
	if err != nil {
		return err
	}

	if valueType != code.Boolean {
		return fmt.Errorf("conditional must be boolean type")
	}

	return nil
}

func validateReturn(ret *code.Return) error {
	valueType, err := code.MapValue[code.Type](ret.Value, value_to_type.Mapper{})
	if err != nil {
		return err
	}

	if !reflect.DeepEqual(valueType, code.MapCallableDefNoError[code.Type](ret.CallableDef, callable_def_to_return_type.Mapper{})) {
		return errors.New("mismatched return types")
	}

	return nil
}

func validateFunction(function *code.FunctionDef) error {
	if function.ReturnType != code.Void {
		if len(function.Block.Statements) == 0 {
			return errors.New("non-void function doesn't end with a return")
		}

		_, lastStatementIsReturn := function.Block.Statements[len(function.Block.Statements)-1].(*code.Return)
		if !lastStatementIsReturn {
			return errors.New("non-void function doesn't end with a return")
		}
	}

	return nil
}
