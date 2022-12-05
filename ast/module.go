package ast

type Module struct {
	Name      string
	Models    []ModelDef
	Functions []FunctionDef
	Constants []ConstantDef
}

type ConstantDef struct {
	Name  string
	Value ConstantValue
}

func (c ConstantDef) isDefinition() {}

type FunctionDef struct {
	Name       string
	Arguments  []ArgumentDef
	Block      Block
	ReturnType Type
}
