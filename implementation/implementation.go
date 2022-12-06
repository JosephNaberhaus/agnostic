package implementation

import (
	"github.com/JosephNaberhaus/agnostic/code"
	"github.com/JosephNaberhaus/agnostic/implementation/text"
)

// TODO comment
type Config struct {
	// The string to indent text one level.
	Indent string
}

type Mapper interface {
	Config() Config
	code.NodeMapperNoError[text.Node]
}
