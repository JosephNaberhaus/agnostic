package node_types

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

// FindAll parses the node types file and returns the node types found in it.
func FindAll(nodeTypesPath string) ([]NodeType, error) {
	file, err := parser.ParseFile(token.NewFileSet(), nodeTypesPath, nil, parser.SkipObjectResolution)
	if err != nil {
		return nil, fmt.Errorf("getCodeBlocks failed to parse file: %w", err)
	}

	var codeBlocks []NodeType
	for _, decl := range file.Decls {
		if genDecl, ok := decl.(*ast.GenDecl); ok {
			for _, spec := range genDecl.Specs {
				if typeSpec, ok := spec.(*ast.TypeSpec); ok {
					if _, ok := typeSpec.Type.(*ast.InterfaceType); ok {
						codeBlockName := typeSpec.Name.String()
						codeBlocks = append(codeBlocks, NodeType{
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
