package ast

// Field

type FieldDef struct {
	Name string
	Type Type
}

func (f FieldDef) isDefinition() {}

// EqualOverride

type EqualOverride struct {
	OtherName string
	Block     Block
}

func (e EqualOverride) isCallableDef() {}

func (e EqualOverride) isDefinition() {}

// HashOverride

type HashOverride struct {
	Block Block
}

func (h HashOverride) isCallableDef() {}

// Model

type ModelDef struct {
	Name          string
	Fields        []FieldDef
	Methods       []FunctionDef
	EqualOverride Optional[EqualOverride]
	HashOverride  Optional[HashOverride]
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
