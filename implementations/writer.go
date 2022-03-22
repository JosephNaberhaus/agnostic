package implementations

import (
	"fmt"
	"github.com/JosephNaberhaus/agnostic"
	"github.com/JosephNaberhaus/agnostic/implementations/golang"
	"github.com/JosephNaberhaus/agnostic/implementations/typescript"
	"github.com/JosephNaberhaus/agnostic/implementations/writer"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type Language string

const (
	Golang     Language = "golang"
	Typescript          = "typescript"
)

var Languages = []Language{
	Golang,
	Typescript,
}

func WriteModel(p agnostic.Model, language Language) error {
	switch language {
	case Golang:
		return writeModel(p, "go", golang.ModelCode)
	case Typescript:
		return writeModel(p, "ts", typescript.ModelCode)
	default:
		return fmt.Errorf("unknown langauge \"%v\"", language)
	}
}

func writeModel(model agnostic.Model, extension string, modelCodeGenerator func(model agnostic.Model) (writer.Code, error)) error {
	err := os.MkdirAll(model.Path, os.ModePerm)
	if err != nil {
		return fmt.Errorf("error creating package directory: %w", err)
	}

	code, err := modelCodeGenerator(model)
	if err != nil {
		return fmt.Errorf("error generating code for \"%s\": %w", model.Name, err)
	}

	outputPath := filepath.Join(model.Path, strings.ToLower(model.Name)+"."+extension)
	err = ioutil.WriteFile(outputPath, []byte(writer.CodeString(code, 0)), os.ModePerm)
	if err != nil {
		return fmt.Errorf("error writing to the file \"%s\": %w", outputPath, err)
	}

	return nil
}
