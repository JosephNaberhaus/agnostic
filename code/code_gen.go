package code

type AddToSet struct {
	Set Value

	Value Value

	AddToSetMetadata
}

func (AddToSet) isNode() {}

func (AddToSet) isStatement() {}

type ArgumentDef struct {
	Name string

	Type Type

	ArgumentDefMetadata
}

func (ArgumentDef) isNode() {}

func (ArgumentDef) isDefinition() {}

type Assignment struct {
	From Value

	To Value

	AssignmentMetadata
}

func (Assignment) isNode() {}

func (Assignment) isStatement() {}

type Block struct {
	Statements []Statement

	BlockMetadata
}

func (Block) isNode() {}

type Bool struct {
	BoolMetadata
}

func (Bool) isNode() {}

func (Bool) isType() {}

type Break struct {
	BreakMetadata
}

func (Break) isNode() {}

func (Break) isStatement() {}

type Call struct {
	Arguments []Value

	Function Callable

	CallMetadata
}

func (Call) isNode() {}

func (Call) isStatement() {}

func (Call) isValue() {}

type Conditional struct {
	Else Optional[Block]

	Ifs []If

	ConditionalMetadata
}

func (Conditional) isNode() {}

func (Conditional) isStatement() {}

type ConstantDef struct {
	Name string

	Value ConstantValue

	ConstantDefMetadata
}

func (ConstantDef) isNode() {}

func (ConstantDef) isDefinition() {}

type Continue struct {
	ContinueMetadata
}

func (Continue) isNode() {}

func (Continue) isStatement() {}

type Declare struct {
	Name string

	Value Value

	DeclareMetadata
}

func (Declare) isNode() {}

func (Declare) isStatement() {}

func (Declare) isDefinition() {}

type EmptyList struct {
	Type Type

	EmptyListMetadata
}

func (EmptyList) isNode() {}

func (EmptyList) isConstantValue() {}

func (EmptyList) isValue() {}

type EqualOverride struct {
	Block Block

	OtherName string

	EqualOverrideMetadata
}

func (EqualOverride) isNode() {}

type FieldDef struct {
	Name string

	Type Type

	FieldDefMetadata
}

func (FieldDef) isNode() {}

func (FieldDef) isDefinition() {}

type For struct {
	AfterEach Optional[Statement]

	Block Block

	Condition Value

	Initialization Optional[Statement]

	ForMetadata
}

func (For) isNode() {}

func (For) isStatement() {}

type ForEach struct {
	Block Block

	ItemName string

	Iterable Value

	ForEachMetadata
}

func (ForEach) isNode() {}

func (ForEach) isDefinition() {}

func (ForEach) isStatement() {}

type FunctionDef struct {
	Arguments []ArgumentDef

	Block Block

	Name string

	ReturnTYpe Type

	FunctionDefMetadata
}

func (FunctionDef) isNode() {}

func (FunctionDef) isCallable() {}

type HashOverride struct {
	Block Block

	HashOverrideMetadata
}

func (HashOverride) isNode() {}

type If struct {
	Block Block

	Condition Value

	IfMetadata
}

func (If) isNode() {}

type Int64 struct {
	Int64Metadata
}

func (Int64) isNode() {}

func (Int64) isType() {}

type KeyValue struct {
	Key Value

	Value Value

	KeyValueMetadata
}

func (KeyValue) isNode() {}

type Length struct {
	Of Value

	LengthMetadata
}

func (Length) isNode() {}

func (Length) isValue() {}

type List struct {
	Item Type

	ListMetadata
}

func (List) isNode() {}

func (List) isType() {}

type LiteralBool struct {
	Value bool

	LiteralBoolMetadata
}

func (LiteralBool) isNode() {}

func (LiteralBool) isConstantValue() {}

func (LiteralBool) isValue() {}

type LiteralInt struct {
	Value int64

	LiteralIntMetadata
}

func (LiteralInt) isNode() {}

func (LiteralInt) isConstantValue() {}

func (LiteralInt) isValue() {}

type LiteralList struct {
	Values []Value

	LiteralListMetadata
}

func (LiteralList) isNode() {}

func (LiteralList) isConstantValue() {}

func (LiteralList) isValue() {}

type LiteralMap struct {
	Keys []KeyValue

	LiteralMapMetadata
}

func (LiteralMap) isNode() {}

func (LiteralMap) isConstantValue() {}

func (LiteralMap) isValue() {}

type LiteralRune struct {
	Value rune

	LiteralRuneMetadata
}

func (LiteralRune) isNode() {}

func (LiteralRune) isConstantValue() {}

func (LiteralRune) isValue() {}

type LiteralSet struct {
	Values []Value

	LiteralSetMetadata
}

func (LiteralSet) isNode() {}

func (LiteralSet) isConstantValue() {}

func (LiteralSet) isValue() {}

type LiteralString struct {
	Value string

	LiteralStringMetadata
}

func (LiteralString) isNode() {}

func (LiteralString) isConstantValue() {}

func (LiteralString) isValue() {}

type Lookup struct {
	From Value

	Key Value

	LookupMetadata
}

func (Lookup) isNode() {}

func (Lookup) isValue() {}

type Map struct {
	Key Type

	Value Type

	MapMetadata
}

func (Map) isNode() {}

func (Map) isType() {}

type Model struct {
	Name string

	ModelMetadata
}

func (Model) isNode() {}

func (Model) isType() {}

type ModelDef struct {
	EqualOverride EqualOverride

	Fields []FieldDef

	HashOverride HashOverride

	Methods []FunctionDef

	Name string

	ModelDefMetadata
}

func (ModelDef) isNode() {}

type Module struct {
	Constants []ConstantDef

	Functions []FunctionDef

	Models []ModelDef

	Name string

	ModuleMetadata
}

func (Module) isNode() {}

type New struct {
	Model Model

	NewMetadata
}

func (New) isNode() {}

func (New) isValue() {}

type Null struct {
	Type Type

	NullMetadata
}

func (Null) isNode() {}

func (Null) isConstantValue() {}

func (Null) isValue() {}

type Pop struct {
	List Value

	PopMetadata
}

func (Pop) isNode() {}

func (Pop) isStatement() {}

func (Pop) isValue() {}

type Property struct {
	Name string

	Of Value

	PropertyMetadata
}

func (Property) isNode() {}

func (Property) isAssignable() {}

func (Property) isValue() {}

type Push struct {
	List Value

	Value Value

	PushMetadata
}

func (Push) isNode() {}

func (Push) isStatement() {}

type Return struct {
	Value Value

	ReturnMetadata
}

func (Return) isNode() {}

func (Return) isStatement() {}

type Rune struct {
	RuneMetadata
}

func (Rune) isNode() {}

func (Rune) isType() {}

type Self struct {
	SelfMetadata
}

func (Self) isNode() {}

func (Self) isValue() {}

type Set struct {
	Item Type

	SetMetadata
}

func (Set) isNode() {}

func (Set) isType() {}

type SetContains struct {
	Set Set

	Value Value

	SetContainsMetadata
}

func (SetContains) isNode() {}

func (SetContains) isValue() {}

type String struct {
	StringMetadata
}

func (String) isNode() {}

func (String) isType() {}

type Variable struct {
	Name string

	VariableMetadata
}

func (Variable) isNode() {}

func (Variable) isAssignable() {}

func (Variable) isValue() {}

type Void struct {
	VoidMetadata
}

func (Void) isNode() {}

func (Void) isType() {}
