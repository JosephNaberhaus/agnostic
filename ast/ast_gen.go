package ast


type ArgumentDef struct {

	Name string

	Type Type

}


func (_ ArgumentDef) isDefinition() {}



type FunctionDef struct {

	Arguments []ArgumentDef

	Block Block

	Name string

	ReturnTYpe Type

}


func (_ FunctionDef) isCallable() {}



type EqualOverride struct {

	Block Block

	OtherName string

}




type FieldDef struct {

	Name string

	Type Type

}


func (_ FieldDef) isDefinition() {}



type HashOverride struct {

	Block Block

}




type ModelDef struct {

	EqualOverride EqualOverride

	Fields []FieldDef

	HashOverride HashOverride

	Methods []FunctionDef

	Name string

}




type ConstantDef struct {

	Name string

	Value ConstantValue

}


func (_ ConstantDef) isDefinition() {}



type Module struct {

	Constants []ConstantDef

	Functions []FunctionDef

	Models []ModelDef

	Name string

}




type AddToSet struct {

	Set Value

	Value Value

}


func (_ AddToSet) isStatement() {}



type Assignment struct {

	From Value

	To Value

}


func (_ Assignment) isStatement() {}



type Block struct {

	Statements []Statement

}




type Break struct {

}


func (_ Break) isStatement() {}



type Conditional struct {

	Else Optional[Block]

	Ifs []If

}


func (_ Conditional) isStatement() {}



type Continue struct {

}


func (_ Continue) isStatement() {}



type Declare struct {

	Name string

	Value Value

}


func (_ Declare) isStatement() {}

func (_ Declare) isDefinition() {}



type For struct {

	AfterEach Optional[Statement]

	Block Block

	Condition Value

	Initialization Optional[Statement]

}


func (_ For) isStatement() {}



type ForEach struct {

	Block Block

	ItemName string

	Iterable Value

}


func (_ ForEach) isDefinition() {}

func (_ ForEach) isStatement() {}



type If struct {

	Block Block

	Condition Value

}




type Pop struct {

	List Value

}


func (_ Pop) isStatement() {}

func (_ Pop) isValue() {}



type Push struct {

	List Value

	Value Value

}


func (_ Push) isStatement() {}



type Return struct {

	Value Value

}


func (_ Return) isStatement() {}



type Bool struct {

}


func (_ Bool) isType() {}



type Int64 struct {

}


func (_ Int64) isType() {}



type List struct {

	Item Type

}


func (_ List) isType() {}



type Map struct {

	Key Type

	Value Type

}


func (_ Map) isType() {}



type Model struct {

	Name string

}


func (_ Model) isType() {}



type Rune struct {

}


func (_ Rune) isType() {}



type Set struct {

	Item Type

}


func (_ Set) isType() {}



type String struct {

}


func (_ String) isType() {}



type Void struct {

}


func (_ Void) isType() {}



type Call struct {

	Arguments []Value

	Function Callable

}


func (_ Call) isStatement() {}

func (_ Call) isValue() {}



type EmptyList struct {

	Type Type

}


func (_ EmptyList) isConstantValue() {}

func (_ EmptyList) isValue() {}



type KeyValue struct {

	Key Value

	Value Value

}




type Length struct {

	Of Value

}


func (_ Length) isValue() {}



type LiteralBool struct {

	Value bool

}


func (_ LiteralBool) isConstantValue() {}

func (_ LiteralBool) isValue() {}



type LiteralInt struct {

	Value int64

}


func (_ LiteralInt) isConstantValue() {}

func (_ LiteralInt) isValue() {}



type LiteralList struct {

	Values []Value

}


func (_ LiteralList) isConstantValue() {}

func (_ LiteralList) isValue() {}



type LiteralMap struct {

	Keys []KeyValue

}


func (_ LiteralMap) isConstantValue() {}

func (_ LiteralMap) isValue() {}



type LiteralRune struct {

	Value rune

}


func (_ LiteralRune) isConstantValue() {}

func (_ LiteralRune) isValue() {}



type LiteralSet struct {

	Values []Value

}


func (_ LiteralSet) isConstantValue() {}

func (_ LiteralSet) isValue() {}



type LiteralString struct {

	Value string

}


func (_ LiteralString) isConstantValue() {}

func (_ LiteralString) isValue() {}



type Lookup struct {

	From Value

	Key Value

}


func (_ Lookup) isValue() {}



type New struct {

	Model Model

}


func (_ New) isValue() {}



type Null struct {

	Type Type

}


func (_ Null) isConstantValue() {}

func (_ Null) isValue() {}



type Property struct {

	Name string

	Of Value

}


func (_ Property) isAssignable() {}

func (_ Property) isValue() {}



type Self struct {

}


func (_ Self) isValue() {}



type SetContains struct {

	Set Set

	Value Value

}


func (_ SetContains) isValue() {}



type Variable struct {

	Name string

}


func (_ Variable) isAssignable() {}

func (_ Variable) isValue() {}


