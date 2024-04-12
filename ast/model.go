package ast

// Field

type FieldDef struct {
	Meta
	Name string
	Type Type
}

func (f FieldDef) isDefinition() {}

// EqualOverride

type EqualOverride struct {
	Meta
	OtherName string
	Block     Block
}

func (e EqualOverride) isCallableDef() {}

func (e EqualOverride) isDefinition() {}

// HashOverride

type HashOverride struct {
	Meta
	Block Block
}

func (h HashOverride) isCallableDef() {}

// Model

type ModelDef struct {
	Meta
	Name          string
	Fields        []FieldDef
	Methods       []FunctionDef
	EqualOverride Optional[EqualOverride]
	HashOverride  Optional[HashOverride]
}

// Variable

type Variable struct {
	Meta
	Name string
}

func (v Variable) isValue() {}

func (v Variable) isAssignable() {}

// Property

type Property struct {
	Meta
	Of   Value
	Name string
}

func (p Property) isValue() {}

func (p Property) isAssignable() {}
