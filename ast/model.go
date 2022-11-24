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
	Name       string
	Arguments  []ArgumentDef
	Statements []Statement
	ReturnType Type
}

// Model

type ModelDef struct {
	Name    string
	Fields  []FieldDef
	Methods []MethodDef
}

// Property

type Property struct {
	Of   Value
	Name string
}

func (p Property) isValue() {}

func (p Property) isAssignable() {}

// This

type This struct{}

func (t This) isValue() {}
