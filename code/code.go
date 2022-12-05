//go:generate go run ../tool/code_generator/main.go
//go:generate go run ../tool/mapper_generator/main.go -usePointersForStructs -nodeTypesFile=node_types.g.go -filterOut=(metadata)|(stack)

package code

import "github.com/JosephNaberhaus/agnostic/ast"

// Parse converts an ast.Module into a Module.
func Parse(astModule *ast.Module) (*Module, error) {
	return nil, nil
}
