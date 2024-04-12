package ast

type Module struct {
	Meta
	Name      string
	Models    []ModelDef
	Functions []FunctionDef
	Constants []ConstantDef
}

type ConstantDef struct {
	Meta
	Name  string
	Value ConstantValue
}

func (c ConstantDef) isDefinition() {}
