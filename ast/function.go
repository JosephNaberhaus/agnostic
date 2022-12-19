package ast

// Argument

type ArgumentDef struct {
	Name string
	Type Type
}

func (a ArgumentDef) isDefinition() {}

// FunctionDef

type FunctionDef struct {
	Name       string
	Arguments  []ArgumentDef
	Block      Block
	ReturnType Type
}

func (f FunctionDef) isCallableDef() {}
