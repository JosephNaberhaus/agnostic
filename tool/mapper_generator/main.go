package main

import (
	"embed"
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"
)

var (
	sourcesDirPath        = flag.String("source", ".", "the source directory")
	nodeTypesPath         = flag.String("nodeTypesFile", "./node_types.go", "the node types file")
	outputFilePath        = flag.String("output", "node_mapper.g.go", "the output file")
	usePointersForStructs = flag.Bool("usePointersForStructs", false, "whether to use pointers for struct nodes")
	fileFilter            = flag.String("filterOut", "$^", "regex used to filter out files")
)

//go:embed mapper.go.tmpl
var embeddedFiles embed.FS

const mapperTemplateFilename = "mapper.go.tmpl"

func main() {
	flag.Parse()

	codeBlocks, err := findNodeTypes()
	if err != nil {
		log.Fatalf("failed to get node types: %s", err.Error())
	}

	nodes, err := findNodes()
	if err != nil {
		log.Fatalf("failed to find nodes: %s", err.Error())
	}

	implementations, err := findNodeTypeImplementations(codeBlocks)
	if err != nil {
		log.Fatalf("failed to find implementations: %s", err.Error())
	}

	err = writeMapperFile(nodes, implementations)
	if err != nil {
		log.Fatalf("failed to write mapper file: %s", err.Error())
	}

	err = formatMapperFile()
	if err != nil {
		log.Fatalf("failed to format mapper file: %s", err.Error())
	}
}

// isImplementationFile returns whether the file can contain implementations of code block interfaces.
func isImplementationFile(entry os.DirEntry) bool {
	if entry.IsDir() {
		return false
	}

	switch entry.Name() {
	case filepath.Base(*nodeTypesPath), filepath.Base(*outputFilePath):
		return false
	default:
		return true
	}
}

func findImplementationFilePaths(dirPath string) ([]string, error) {
	allFiles, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, fmt.Errorf("findImplementationFilePaths failed to read directory \"%\": %w", dirPath, err)
	}

	filter, err := regexp.Compile(*fileFilter)
	if err != nil {
		return nil, fmt.Errorf("findImplementationFilePaths failed to compiled file filter out regex: %w", err)
	}

	var implementationFilePaths []string
	for _, file := range allFiles {
		if filter.MatchString(file.Name()) {
			continue
		}

		if isImplementationFile(file) {
			implementationFilePaths = append(implementationFilePaths, filepath.Join(dirPath, file.Name()))
		}
	}

	return implementationFilePaths, nil
}

type node struct {
	Name     string
	IsStruct bool
}

// findNodes return all nodes found in the implementation files.
func findNodes() ([]node, error) {
	implementationFilePaths, err := findImplementationFilePaths(*sourcesDirPath)
	if err != nil {
		return nil, fmt.Errorf("findNodeTypeImplementations failed to find implementation files")
	}

	var nodes []node
	for _, implementationFilePath := range implementationFilePaths {
		file, err := parser.ParseFile(token.NewFileSet(), implementationFilePath, nil, parser.SkipObjectResolution)
		if err != nil {
			return nil, fmt.Errorf("findNodeTypeImplementations failed to parse file \"%s\": %w", implementationFilePath, err)
		}

		for _, nodeInFile := range findNodesInFile(file) {
			nodes = append(nodes, nodeInFile)
		}
	}

	return nodes, err
}

func findNodesInFile(file *ast.File) []node {
	var nodes []node
	for _, decl := range file.Decls {
		if genDecl, ok := decl.(*ast.GenDecl); ok {
			for _, spec := range genDecl.Specs {
				if typeSpec, ok := spec.(*ast.TypeSpec); ok {
					_, isStruct := typeSpec.Type.(*ast.StructType)

					nodes = append(nodes, node{
						Name:     typeSpec.Name.String(),
						IsStruct: isStruct,
					})
				}
			}
		}
	}

	return nodes
}

type nodeType struct {
	Name           string
	ExpectedMethod string
}

// findNodeTypes parses the node types file and returns the node types found in it.
func findNodeTypes() ([]nodeType, error) {
	file, err := parser.ParseFile(token.NewFileSet(), *nodeTypesPath, nil, parser.SkipObjectResolution)
	if err != nil {
		return nil, fmt.Errorf("getCodeBlocks failed to parse file: %w", err)
	}

	var codeBlocks []nodeType
	for _, decl := range file.Decls {
		if genDecl, ok := decl.(*ast.GenDecl); ok {
			for _, spec := range genDecl.Specs {
				if typeSpec, ok := spec.(*ast.TypeSpec); ok {
					if _, ok := typeSpec.Type.(*ast.InterfaceType); ok {
						codeBlockName := typeSpec.Name.String()
						codeBlocks = append(codeBlocks, nodeType{
							Name:           codeBlockName,
							ExpectedMethod: createCodeBlockMethodName(codeBlockName),
						})
					}
				}
			}
		}
	}

	return codeBlocks, nil
}

// createCodeBlockMethodName creates the name of the method that must be implemented to satisfy the interface of the
// code block.
//
// For simplicity, this function assumes that the code block was defined according to the standard:
//
// type <CodeBlockName> interface {
//     is<CodeBlockName>()
// }
func createCodeBlockMethodName(codeBlockName string) string {
	return "is" + codeBlockName
}

// findNodeTypeImplementations finds the implementations for the given node types.
func findNodeTypeImplementations(nodeTypes []nodeType) (map[nodeType][]string, error) {
	implementationFilePaths, err := findImplementationFilePaths(*sourcesDirPath)
	if err != nil {
		return nil, fmt.Errorf("findNodeTypeImplementations failed to find implementation files")
	}

	implementations := map[nodeType][]string{}
	for _, implementationFilePath := range implementationFilePaths {
		file, err := parser.ParseFile(token.NewFileSet(), implementationFilePath, nil, parser.SkipObjectResolution)
		if err != nil {
			return nil, fmt.Errorf("findNodeTypeImplementations failed to parse file \"%s\": %w", implementationFilePath, err)
		}

		fileImplementations, err := findNodeTypeImplementationsInFile(file, nodeTypes)
		if err != nil {
			return nil, fmt.Errorf("findNodeTypeImplementations failed to find implementations for file \"%s\": %w", implementationFilePath, err)
		}

		for codeBlock, names := range fileImplementations {
			implementations[codeBlock] = append(implementations[codeBlock], names...)
		}
	}

	return implementations, nil
}

func findNodeTypeImplementationsInFile(file *ast.File, nodeTypes []nodeType) (map[nodeType][]string, error) {
	implementations := map[nodeType][]string{}
	for _, decl := range file.Decls {
		if funcDecl, ok := decl.(*ast.FuncDecl); ok {
			// We are only interested in methods
			if funcDecl.Recv == nil {
				continue
			}

			// Node type interface methods don't take any parameters
			if funcDecl.Type.Params.List != nil {
				continue
			}

			// Node ty[e interface methods don't return anything
			if funcDecl.Type.Results != nil {
				continue
			}

			for _, nodeType := range nodeTypes {
				if nodeType.ExpectedMethod == funcDecl.Name.String() {
					structName := types.ExprString(funcDecl.Recv.List[0].Type)

					implementations[nodeType] = append(implementations[nodeType], structName)
				}
			}
		}
	}

	return implementations, nil
}

func getReceiver(name string) string {
	return strings.ToLower(string([]rune(name)[0]))
}

func removePointer(name string) string {
	return strings.TrimPrefix(name, "*")
}

func getPackageName() (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("getPackageName failed to get the working directory: %w", err)
	}

	return filepath.Base(filepath.Dir(filepath.Join(wd, *outputFilePath))), nil
}

func writeMapperFile(nodes []node, nodeTypeImplementations map[nodeType][]string) error {
	mapperTemplateText, err := embeddedFiles.ReadFile(mapperTemplateFilename)
	if err != nil {
		return fmt.Errorf("writeMapperFile failed to read mapper template file: %w", err)
	}

	mapperTemplate, err := template.New(mapperTemplateFilename).
		Funcs(template.FuncMap{
			"GetReceiver":   getReceiver,
			"RemovePointer": removePointer,
		}).
		Parse(string(mapperTemplateText))
	if err != nil {
		return fmt.Errorf("writeMapperFile failed to parse mapper template: %w", err)
	}

	outputFile, err := os.Create(*outputFilePath)
	if err != nil {
		return fmt.Errorf("writeMapperFile failed to create output file: %w", err)
	}
	defer outputFile.Close()

	packageName, err := getPackageName()
	if err != nil {
		return fmt.Errorf("writeMapperFile failed to get the package name: %w", err)
	}

	templateData := struct {
		PackageName             string
		Nodes                   []node
		NodeTypeImplementations map[nodeType][]string
		UsePointersForStructs   bool
	}{
		PackageName:             packageName,
		Nodes:                   nodes,
		NodeTypeImplementations: nodeTypeImplementations,
		UsePointersForStructs:   *usePointersForStructs,
	}

	err = mapperTemplate.Execute(outputFile, templateData)
	if err != nil {
		return fmt.Errorf("writeMapperFile failed to execute tempalte: %w", err)
	}

	return nil
}

// formatMapperFile runs the go formatter against the generated file.
func formatMapperFile() error {
	err := exec.Command("go", "fmt", *outputFilePath).Run()
	if err != nil {
		return fmt.Errorf("formatMapperFile failed to run format command: %w", err)
	}

	return err
}
