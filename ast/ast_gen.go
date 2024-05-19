package ast

type AddToSet struct {
	Set Value

	Value Value
}

func (AddToSet) isNode() {}

// isStatement is just a inteface guard to restrict what can be used as a Statement.
func (AddToSet) isStatement() {}

type ArgumentDef struct {
	Name string

	Type Type
}

func (ArgumentDef) isNode() {}

// isDefinition is just a inteface guard to restrict what can be used as a Definition.
func (ArgumentDef) isDefinition() {}

type Assignment struct {
	From Value

	To Value
}

func (Assignment) isNode() {}

// isStatement is just a inteface guard to restrict what can be used as a Statement.
func (Assignment) isStatement() {}

type Block struct {
	Statements []Statement
}

func (Block) isNode() {}

type Bool struct {
}

func (Bool) isNode() {}

// isType is just a inteface guard to restrict what can be used as a Type.
func (Bool) isType() {}

type Break struct {
}

func (Break) isNode() {}

// isStatement is just a inteface guard to restrict what can be used as a Statement.
func (Break) isStatement() {}

type Call struct {
	Arguments []Value

	Function Callable
}

func (Call) isNode() {}

// isStatement is just a inteface guard to restrict what can be used as a Statement.
func (Call) isStatement() {}

// isValue is just a inteface guard to restrict what can be used as a Value.
func (Call) isValue() {}

type Conditional struct {
	Else Optional[Block]

	Ifs []If
}

func (Conditional) isNode() {}

// isStatement is just a inteface guard to restrict what can be used as a Statement.
func (Conditional) isStatement() {}

type ConstantDef struct {
	Name string

	Value ConstantValue
}

func (ConstantDef) isNode() {}

// isDefinition is just a inteface guard to restrict what can be used as a Definition.
func (ConstantDef) isDefinition() {}

type Continue struct {
}

func (Continue) isNode() {}

// isStatement is just a inteface guard to restrict what can be used as a Statement.
func (Continue) isStatement() {}

type Declare struct {
	Name string

	Value Value
}

func (Declare) isNode() {}

// isStatement is just a inteface guard to restrict what can be used as a Statement.
func (Declare) isStatement() {}

// isDefinition is just a inteface guard to restrict what can be used as a Definition.
func (Declare) isDefinition() {}

type EmptyList struct {
	Type Type
}

func (EmptyList) isNode() {}

// isConstantValue is just a inteface guard to restrict what can be used as a ConstantValue.
func (EmptyList) isConstantValue() {}

// isValue is just a inteface guard to restrict what can be used as a Value.
func (EmptyList) isValue() {}

type EqualOverride struct {
	Block Block

	OtherName string
}

func (EqualOverride) isNode() {}

type FieldDef struct {
	Name string

	Type Type
}

func (FieldDef) isNode() {}

// isDefinition is just a inteface guard to restrict what can be used as a Definition.
func (FieldDef) isDefinition() {}

type For struct {
	AfterEach Optional[Statement]

	Block Block

	Condition Value

	Initialization Optional[Statement]
}

func (For) isNode() {}

// isStatement is just a inteface guard to restrict what can be used as a Statement.
func (For) isStatement() {}

type ForEach struct {
	Block Block

	ItemName string

	Iterable Value
}

func (ForEach) isNode() {}

// isDefinition is just a inteface guard to restrict what can be used as a Definition.
func (ForEach) isDefinition() {}

// isStatement is just a inteface guard to restrict what can be used as a Statement.
func (ForEach) isStatement() {}

type FunctionDef struct {
	Arguments []ArgumentDef

	Block Block

	Name string

	ReturnTYpe Type
}

func (FunctionDef) isNode() {}

// isCallable is just a inteface guard to restrict what can be used as a Callable.
func (FunctionDef) isCallable() {}

type HashOverride struct {
	Block Block
}

func (HashOverride) isNode() {}

type If struct {
	Block Block

	Condition Value
}

func (If) isNode() {}

type Int64 struct {
}

func (Int64) isNode() {}

// isType is just a inteface guard to restrict what can be used as a Type.
func (Int64) isType() {}

type KeyValue struct {
	Key Value

	Value Value
}

func (KeyValue) isNode() {}

type Length struct {
	Of Value
}

func (Length) isNode() {}

// isValue is just a inteface guard to restrict what can be used as a Value.
func (Length) isValue() {}

type List struct {
	Item Type
}

func (List) isNode() {}

// isType is just a inteface guard to restrict what can be used as a Type.
func (List) isType() {}

type LiteralBool struct {
	Value bool
}

func (LiteralBool) isNode() {}

// isConstantValue is just a inteface guard to restrict what can be used as a ConstantValue.
func (LiteralBool) isConstantValue() {}

// isValue is just a inteface guard to restrict what can be used as a Value.
func (LiteralBool) isValue() {}

type LiteralInt struct {
	Value int64
}

func (LiteralInt) isNode() {}

// isConstantValue is just a inteface guard to restrict what can be used as a ConstantValue.
func (LiteralInt) isConstantValue() {}

// isValue is just a inteface guard to restrict what can be used as a Value.
func (LiteralInt) isValue() {}

type LiteralList struct {
	Values []Value
}

func (LiteralList) isNode() {}

// isConstantValue is just a inteface guard to restrict what can be used as a ConstantValue.
func (LiteralList) isConstantValue() {}

// isValue is just a inteface guard to restrict what can be used as a Value.
func (LiteralList) isValue() {}

type LiteralMap struct {
	Keys []KeyValue
}

func (LiteralMap) isNode() {}

// isConstantValue is just a inteface guard to restrict what can be used as a ConstantValue.
func (LiteralMap) isConstantValue() {}

// isValue is just a inteface guard to restrict what can be used as a Value.
func (LiteralMap) isValue() {}

type LiteralRune struct {
	Value rune
}

func (LiteralRune) isNode() {}

// isConstantValue is just a inteface guard to restrict what can be used as a ConstantValue.
func (LiteralRune) isConstantValue() {}

// isValue is just a inteface guard to restrict what can be used as a Value.
func (LiteralRune) isValue() {}

type LiteralSet struct {
	Values []Value
}

func (LiteralSet) isNode() {}

// isConstantValue is just a inteface guard to restrict what can be used as a ConstantValue.
func (LiteralSet) isConstantValue() {}

// isValue is just a inteface guard to restrict what can be used as a Value.
func (LiteralSet) isValue() {}

type LiteralString struct {
	Value string
}

func (LiteralString) isNode() {}

// isConstantValue is just a inteface guard to restrict what can be used as a ConstantValue.
func (LiteralString) isConstantValue() {}

// isValue is just a inteface guard to restrict what can be used as a Value.
func (LiteralString) isValue() {}

type Lookup struct {
	From Value

	Key Value
}

func (Lookup) isNode() {}

// isValue is just a inteface guard to restrict what can be used as a Value.
func (Lookup) isValue() {}

type Map struct {
	Key Type

	Value Type
}

func (Map) isNode() {}

// isType is just a inteface guard to restrict what can be used as a Type.
func (Map) isType() {}

type Model struct {
	Name string
}

func (Model) isNode() {}

// isType is just a inteface guard to restrict what can be used as a Type.
func (Model) isType() {}

type ModelDef struct {
	EqualOverride EqualOverride

	Fields []FieldDef

	HashOverride HashOverride

	Methods []FunctionDef

	Name string
}

func (ModelDef) isNode() {}

type Module struct {
	Constants []ConstantDef

	Functions []FunctionDef

	Models []ModelDef

	Name string
}

func (Module) isNode() {}

type New struct {
	Model Model
}

func (New) isNode() {}

// isValue is just a inteface guard to restrict what can be used as a Value.
func (New) isValue() {}

type Null struct {
	Type Type
}

func (Null) isNode() {}

// isConstantValue is just a inteface guard to restrict what can be used as a ConstantValue.
func (Null) isConstantValue() {}

// isValue is just a inteface guard to restrict what can be used as a Value.
func (Null) isValue() {}

type Pop struct {
	List Value
}

func (Pop) isNode() {}

// isStatement is just a inteface guard to restrict what can be used as a Statement.
func (Pop) isStatement() {}

// isValue is just a inteface guard to restrict what can be used as a Value.
func (Pop) isValue() {}

type Property struct {
	Name string

	Of Value
}

func (Property) isNode() {}

// isAssignable is just a inteface guard to restrict what can be used as a Assignable.
func (Property) isAssignable() {}

// isValue is just a inteface guard to restrict what can be used as a Value.
func (Property) isValue() {}

type Push struct {
	List Value

	Value Value
}

func (Push) isNode() {}

// isStatement is just a inteface guard to restrict what can be used as a Statement.
func (Push) isStatement() {}

type Return struct {
	Value Value
}

func (Return) isNode() {}

// isStatement is just a inteface guard to restrict what can be used as a Statement.
func (Return) isStatement() {}

type Rune struct {
}

func (Rune) isNode() {}

// isType is just a inteface guard to restrict what can be used as a Type.
func (Rune) isType() {}

type Self struct {
}

func (Self) isNode() {}

// isValue is just a inteface guard to restrict what can be used as a Value.
func (Self) isValue() {}

type Set struct {
	Item Type
}

func (Set) isNode() {}

// isType is just a inteface guard to restrict what can be used as a Type.
func (Set) isType() {}

type SetContains struct {
	Set Set

	Value Value
}

func (SetContains) isNode() {}

// isValue is just a inteface guard to restrict what can be used as a Value.
func (SetContains) isValue() {}

type String struct {
}

func (String) isNode() {}

// isType is just a inteface guard to restrict what can be used as a Type.
func (String) isType() {}

type Variable struct {
	Name string
}

func (Variable) isNode() {}

// isAssignable is just a inteface guard to restrict what can be used as a Assignable.
func (Variable) isAssignable() {}

// isValue is just a inteface guard to restrict what can be used as a Value.
func (Variable) isValue() {}

type Void struct {
}

func (Void) isNode() {}

// isType is just a inteface guard to restrict what can be used as a Type.
func (Void) isType() {}
