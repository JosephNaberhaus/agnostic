package ast


type AddToSet struct {

	Set Value

	Value Value

}


func (AddToSet) isStatement() {}



type ArgumentDef struct {

	Name string

	Type Type

}


func (ArgumentDef) isDefinition() {}



type Assignment struct {

	From Value

	To Value

}


func (Assignment) isStatement() {}



type Block struct {

	Statements []Statement

}




type Bool struct {

}


func (Bool) isType() {}



type Break struct {

}


func (Break) isStatement() {}



type Call struct {

	Arguments []Value

	Function Callable

}


func (Call) isStatement() {}

func (Call) isValue() {}



type Conditional struct {

	Else Optional[Block]

	Ifs []If

}


func (Conditional) isStatement() {}



type ConstantDef struct {

	Name string

	Value ConstantValue

}


func (ConstantDef) isDefinition() {}



type Continue struct {

}


func (Continue) isStatement() {}



type Declare struct {

	Name string

	Value Value

}


func (Declare) isStatement() {}

func (Declare) isDefinition() {}



type EmptyList struct {

	Type Type

}


func (EmptyList) isConstantValue() {}

func (EmptyList) isValue() {}



type EqualOverride struct {

	Block Block

	OtherName string

}




type FieldDef struct {

	Name string

	Type Type

}


func (FieldDef) isDefinition() {}



type For struct {

	AfterEach Optional[Statement]

	Block Block

	Condition Value

	Initialization Optional[Statement]

}


func (For) isStatement() {}



type ForEach struct {

	Block Block

	ItemName string

	Iterable Value

}


func (ForEach) isDefinition() {}

func (ForEach) isStatement() {}



type FunctionDef struct {

	Arguments []ArgumentDef

	Block Block

	Name string

	ReturnTYpe Type

}


func (FunctionDef) isCallable() {}



type HashOverride struct {

	Block Block

}




type If struct {

	Block Block

	Condition Value

}




type Int64 struct {

}


func (Int64) isType() {}



type KeyValue struct {

	Key Value

	Value Value

}




type Length struct {

	Of Value

}


func (Length) isValue() {}



type List struct {

	Item Type

}


func (List) isType() {}



type LiteralBool struct {

	Value bool

}


func (LiteralBool) isConstantValue() {}

func (LiteralBool) isValue() {}



type LiteralInt struct {

	Value int64

}


func (LiteralInt) isConstantValue() {}

func (LiteralInt) isValue() {}



type LiteralList struct {

	Values []Value

}


func (LiteralList) isConstantValue() {}

func (LiteralList) isValue() {}



type LiteralMap struct {

	Keys []KeyValue

}


func (LiteralMap) isConstantValue() {}

func (LiteralMap) isValue() {}



type LiteralRune struct {

	Value rune

}


func (LiteralRune) isConstantValue() {}

func (LiteralRune) isValue() {}



type LiteralSet struct {

	Values []Value

}


func (LiteralSet) isConstantValue() {}

func (LiteralSet) isValue() {}



type LiteralString struct {

	Value string

}


func (LiteralString) isConstantValue() {}

func (LiteralString) isValue() {}



type Lookup struct {

	From Value

	Key Value

}


func (Lookup) isValue() {}



type Map struct {

	Key Type

	Value Type

}


func (Map) isType() {}



type Model struct {

	Name string

}


func (Model) isType() {}



type ModelDef struct {

	EqualOverride EqualOverride

	Fields []FieldDef

	HashOverride HashOverride

	Methods []FunctionDef

	Name string

}




type Module struct {

	Constants []ConstantDef

	Functions []FunctionDef

	Models []ModelDef

	Name string

}




type New struct {

	Model Model

}


func (New) isValue() {}



type Null struct {

	Type Type

}


func (Null) isConstantValue() {}

func (Null) isValue() {}



type Pop struct {

	List Value

}


func (Pop) isStatement() {}

func (Pop) isValue() {}



type Property struct {

	Name string

	Of Value

}


func (Property) isAssignable() {}

func (Property) isValue() {}



type Push struct {

	List Value

	Value Value

}


func (Push) isStatement() {}



type Return struct {

	Value Value

}


func (Return) isStatement() {}



type Rune struct {

}


func (Rune) isType() {}



type Self struct {

}


func (Self) isValue() {}



type Set struct {

	Item Type

}


func (Set) isType() {}



type SetContains struct {

	Set Set

	Value Value

}


func (SetContains) isValue() {}



type String struct {

}


func (String) isType() {}



type Variable struct {

	Name string

}


func (Variable) isAssignable() {}

func (Variable) isValue() {}



type Void struct {

}


func (Void) isType() {}


