package utils

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"
)

func ExecuteTemplate(templatePath string, data interface{}) error {
	templ, err := template.ParseFiles(templatePath)
	if err != nil {
		return fmt.Errorf("error parsing template: %w", err)
	}

	outputFileName := strings.TrimSuffix(filepath.Base(templatePath), ".template")

	file, err := os.Create(outputFileName)
	if err != nil {
		return fmt.Errorf("error creating output file \"%s\" for template output: %w", outputFileName, err)
	}

	err = templ.Execute(file, data)
	if err != nil {
		return fmt.Errorf("error executing template; %w", err)
	}

	return nil
}
