package code

type OptionalMetadata struct{}

type FunctionMetadata struct {
	Definition   *FunctionDef
	ParentModule *Module
}

type FunctionPropertyMetadata struct{}

type LiteralIntMetadata struct{}

type LiteralRuneMetadata struct{}

type LiteralStringMetadata struct{}

type LiteralBoolMetadata struct{}

type LiteralListMetadata struct{}

type KeyValueMetadata struct{}

type LiteralMapMetadata struct{}

type LiteralSetMetadata struct{}

type EmptyListMetadata struct{}

type EmptySetMetadata struct{}

type ArgumentDefMetadata struct{}

type FieldDefMetadata struct {
	ParentModel *ModelDef
}

type MethodDefMetadata struct {
	ParentModel *ModelDef
}

type EqualOverrideMetadata struct {
	ParentModel *ModelDef
}

type HashOverrideMetadata struct {
	ParentModel *ModelDef
}

type ModelDefMetadata struct {
	FieldMap  map[string]*FieldDef
	MethodMap map[string]*FunctionDef
}

type VariableMetadata struct {
	Definition Definition
	IsConstant bool
	Type       Type
}

type PropertyMetadata struct {
	Type Type
}

type ModuleMetadata struct {
	ModelMap     map[string]*ModelDef
	FunctionMap  map[string]*FunctionDef
	ConstantsMap map[string]*ConstantDef
}

type ConstantDefMetadata struct {
	ParentModule *Module
	Type         Type
}

type FunctionDefMetadata struct {
	IsMethod bool
}

type BinaryOperationMetadata struct {
	OutputType Primitive
}

type UnaryOperationMetadata struct {
	OutputType Primitive
}

type BlockMetadata struct{}

type AssignmentMetadata struct{}

type IfMetadata struct{}

type ElseIfMetadata struct{}

type ElseMetadata struct{}

type ConditionalMetadata struct{}

type ReturnMetadata struct {
	CallableDef CallableDef
}

type DeclareMetadata struct{ Type Type }

type ForMetadata struct{}

type ForInType int

const (
	ForInTypeList ForInType = iota + 1
	ForInTypeSet
)

type ForInMetadata struct {
	ForInType ForInType
	ItemType  Type
}

type CallFunctionMetadata struct{}

type AddToSetMetadata struct{}

type PushMetadata struct{}

type ModelMetadata struct {
	Definition *ModelDef
}

type VoidMetadata struct{}

type ListMetadata struct{}

type MapMetadata struct{}

type SetMetadata struct{}

type CallMetadata struct {
	Definition *FunctionDef
}

type LookupFromType int

const (
	LookupTypeList LookupFromType = iota + 1
	LookupTypeMap
	LookupTypeString
)

type LookupMetadata struct {
	LookupType LookupFromType
	OutputType Type
}

type NewMetadata struct{}

type LengthType int

const (
	LengthTypeString LengthType = iota + 1
	LengthTypeList
	LengthTypeMap
	LengthTypeSet
)

type LengthMetadata struct {
	LengthType LengthType
}

type SetContainsMetadata struct{}

type PopMetadata struct{}

type NullMetadata struct {
	Parent Node
}

type SelfMetadata struct {
	ParentModel *ModelDef
}

type DeclareNullMetadata struct{}

type BreakMetadata struct{}

type ContinueMetadata struct{}

type LiteralStructMetadata struct {
	Definition *ModelDef
}

type LiteralPropertyMetadata struct{}
