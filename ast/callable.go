package ast

// Function

type Function struct {
	Meta
	Name string
}

func (f Function) isCallable() {}

// Function Property

type FunctionProperty struct {
	Meta
	Of   Value
	Name string
}

func (f FunctionProperty) isCallable() {}
