package ast

// Argument

type ArgumentDef struct {
	Meta
	Name string
	Type Type
}

func (a ArgumentDef) isDefinition() {}

// FunctionDef

type FunctionDef struct {
	Meta
	Name       string
	Arguments  []ArgumentDef
	Block      Block
	ReturnType Type
}

func (f FunctionDef) isCallableDef() {}
