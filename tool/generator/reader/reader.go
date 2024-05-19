package reader

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/JosephNaberhaus/agnostic/tool/generator/model"
	"gopkg.in/yaml.v3"
)

func FindAllSpecs() ([]model.Spec, error) {
	var specs []model.Spec
	err := filepath.Walk("./spec", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if filepath.Ext(path) == ".yaml" {
			spec, err := loadSpec(path)
			if err != nil {
				return err
			}

			specs = append(specs, spec)
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return specs, nil
}

func loadSpec(path string) (model.Spec, error) {
	file, err := os.Open(path)
	if err != nil {
		return model.Spec{}, err
	}

	data, err := io.ReadAll(file)
	if err != nil {
		return model.Spec{}, err
	}

	var spec model.Spec
	err = yaml.Unmarshal(data, &spec)
	if err != nil {
		return model.Spec{}, fmt.Errorf("error parsing %s: %w", path, err)
	}

	return spec, nil
}
