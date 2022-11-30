package ast

// Function

type Function struct {
	Name string
}

func (f Function) isCallable() {}

// Function Property

type FunctionProperty struct {
	Of   Value
	Name string
}

func (f FunctionProperty) isCallable() {}
