package implementations

import (
	"fmt"
	"github.com/JosephNaberhaus/agnostic"
	"github.com/JosephNaberhaus/agnostic/implementations/code"
	"github.com/JosephNaberhaus/agnostic/implementations/golang"
	"github.com/JosephNaberhaus/agnostic/implementations/typescript"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Language int

const (
	Golang Language = iota
	Typescript
)

func WritePackage(p *agnostic.Package, language Language) error {
	err := p.Clean()
	if err != nil {
		return err
	}

	switch language {
	case Golang:
		return writeModels(p, "go", golang.ModelCode)
	case Typescript:
		return writeModels(p, "ts", typescript.ModelCode)
	default:
		return fmt.Errorf("unknown langauge \"%v\"", language)
	}
}

func writeModels(p *agnostic.Package, extension string, modelCodeGenerator func(model agnostic.Model) (writer.Code, error)) error {
	err := os.MkdirAll(p.Path, os.ModePerm)
	if err != nil {
		return fmt.Errorf("error creating package directory: %w", err)
	}

	for _, model := range p.Models {
		code, err := modelCodeGenerator(model)
		if err != nil {
			return fmt.Errorf("error generating code for \"%s\": %w", model.Name, err)
		}

		outputPath := filepath.Join(p.Path, model.Name+"."+extension)
		err = ioutil.WriteFile(outputPath, []byte(writer.CodeString(code, 0)), os.ModePerm)
		if err != nil {
			return fmt.Errorf("error writing to the file \"%s\": %w", outputPath, err)
		}
	}

	return nil
}
