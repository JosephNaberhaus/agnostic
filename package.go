package agnostic

import (
	"fmt"
	"path/filepath"
)

type Package struct {
	Path   string
	Models []Model
}

func (p *Package) Clean() error {
	// Add a link back to this package to each model
	for _, model := range p.Models {
		model.Package = p
	}

	// Make the package path an absolute path
	absPath, err := filepath.Abs(p.Path)
	if err != nil {
		fmt.Errorf("error translating \"%s\" to an absolute path", p.Path)
	}
	p.Path = absPath

	return nil
}
