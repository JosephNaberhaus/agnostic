package ast

type Module struct {
	Name      string
	Models    []ModelDef
	Functions []FunctionDef
}

type FunctionDef struct {
	Name       string
	Arguments  []ArgumentDef
	Statements []Statement
	ReturnType Type
}
