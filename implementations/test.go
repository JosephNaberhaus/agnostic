package implementations

import (
	"fmt"
	writer "github.com/JosephNaberhaus/agnostic/implementations/code"
	"github.com/JosephNaberhaus/agnostic/implementations/golang"
	"github.com/JosephNaberhaus/agnostic/implementations/typescript"
	"github.com/JosephNaberhaus/agnostic/test"
	"github.com/JosephNaberhaus/agnostic/test/tests"
	"github.com/JosephNaberhaus/agnostic/utils"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func writeSuite(suite test.Suite, target Language) error {
	suiteDir := suiteDirPath(target, suite)
	err := os.MkdirAll(suiteDir, os.ModePerm)
	if err != nil {
		return err
	}

	// Move to testDir so that the test/model code is generated relative to it
	err = os.Chdir(suiteDir)
	if err != nil {
		return err
	}

	for _, model := range suite.Models {
		err = WriteModel(model, target)
		if err != nil {
			return err
		}
	}

	err = writeTestFile(suite, target)
	if err != nil {
		return err
	}

	return nil
}

func writeTestFile(suite test.Suite, language Language) error {
	var testCodeGenerator func(test.Suite, test.Test) writer.Code
	var initTestDir func(test.Suite) error
	var testFileNameGenerator func(test.Test) string

	switch language {
	case Golang:
		testCodeGenerator = golang.TestCode
		initTestDir = golang.InitTestSuiteDirectory
		testFileNameGenerator = golang.TestFileName
	case Typescript:
		testCodeGenerator = typescript.TestCode
		initTestDir = typescript.InitTestSuiteDirectory
		testFileNameGenerator = typescript.TestFileName
	default:
		panic(fmt.Errorf("unkown value: \"%v\"", language))
	}

	err := initTestDir(suite)
	if err != nil {
		return err
	}

	for _, t := range suite.Tests {
		testFileName := testFileNameGenerator(t)
		testFileCode := testCodeGenerator(suite, t)

		err = ioutil.WriteFile(testFileName, []byte(writer.CodeString(testFileCode, 0)), os.ModePerm)
		if err != nil {
			return err
		}
	}

	return nil
}

func testsDirPath(language Language) string {
	return filepath.Join(
		utils.GitRootDir(),
		"implementations",
		string(language),
		"tests",
	)
}

func suiteDirPath(language Language, suite test.Suite) string {
	return filepath.Join(
		testsDirPath(language),
		strings.ToLower(suite.Name),
	)
}

func WriteAllTests(language Language) error {
	err := os.RemoveAll(testsDirPath(language))
	if err != nil {
		return err
	}

	for _, suite := range tests.AllSuites {
		err := writeSuite(suite, language)
		if err != nil {
			return err
		}
	}

	return nil
}

func runSuite(suite test.Suite, language Language) error {
	err := os.Chdir(suiteDirPath(language, suite))
	if err != nil {
		return err
	}

	switch language {
	case Golang:
		return golang.RunTests()
	case Typescript:
		return typescript.RunTests()
	default:
		panic(fmt.Errorf("unkown value: \"%v\"", language))
	}
}

func RunAllTests(language Language) error {
	for _, suite := range tests.AllSuites {
		err := runSuite(suite, language)
		if err != nil {
			return err
		}
	}

	return nil
}
