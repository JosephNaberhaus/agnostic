package ast

// Field

type FieldDef struct {
	Name string
	Type Type
}

func (f FieldDef) isDefinition() {}

// Method

type ArgumentDef struct {
	Name string
	Type Type
}

func (a ArgumentDef) isDefinition() {}

type MethodDef struct {
	Function FunctionDef
}

// Model

type ModelDef struct {
	Name    string
	Fields  []FieldDef
	Methods []MethodDef
}

// Variable

type Variable struct {
	Name string
}

func (v Variable) isValue() {}

func (v Variable) isAssignable() {}

// Property

type Property struct {
	Of   Value
	Name string
}

func (p Property) isValue() {}

func (p Property) isAssignable() {}
