package main

import (
	"embed"
	"flag"
	"fmt"
	"github.com/JosephNaberhaus/agnostic/tool/node_types"
	"log"
	"os"
	"text/template"
)

var (
	nodeTypesPath = flag.String("nodeTypesFile", "./node_types.go", "the node types file")
)

//go:embed caster.go.tmpl
var embeddedFiles embed.FS

const casterTemplateFilename = "caster.go.tmpl"

func main() {
	flag.Parse()

	nodeTypes, err := node_types.FindAll(*nodeTypesPath)
	if err != nil {
		log.Fatalf("failed to get node types: %s", err.Error())
	}

	err = writeCasterFile(nodeTypes)
	if err != nil {
		log.Fatalf("failed to write caster file: %s", err.Error())
	}
}

func writeCasterFile(nodeTypes []node_types.NodeType) error {
	casterTemplateText, err := embeddedFiles.ReadFile(casterTemplateFilename)
	if err != nil {
		return fmt.Errorf("writeCasterFile failed to read caster template file: %w", err)
	}

	casterTemplate, err := template.New(casterTemplateFilename).Parse(string(casterTemplateText))
	if err != nil {
		return fmt.Errorf("writeCasterFile failed to parse caster template: %w", err)
	}

	outputFile, err := os.Create("caster.g.go")
	if err != nil {
		return fmt.Errorf("writeCasterFile failed to create output file: %w", err)
	}
	defer outputFile.Close()

	err = casterTemplate.Execute(outputFile, nodeTypes)
	if err != nil {
		return fmt.Errorf("writeCasterFile failed to execute template: %w", err)
	}

	return nil
}
