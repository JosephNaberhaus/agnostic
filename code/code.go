//go:generate go run ../tool/code_generator
//go:generate go run ../tool/mapper_generator -usePointersForStructs -nodeTypesFile=node_types.g.go -exclude=metadata.go,optional.g.go,stack.go

package code

import "github.com/JosephNaberhaus/agnostic/ast"

// Parse converts an ast.Module into a Module.
func Parse(astModule *ast.Module) (*Module, error) {
	return nil, nil
}
