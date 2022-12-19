package ast

// TODO move to value.go

// Literal Int32

type LiteralInt struct {
	Value int64
}

func (l LiteralInt) isValue() {}

func (l LiteralInt) isConstantValue() {}

// Literal Rune

type LiteralRune struct {
	Value rune
}

func (l LiteralRune) isValue() {}

func (l LiteralRune) isConstantValue() {}

// Literal String

type LiteralString struct {
	Value string
}

func (l LiteralString) isValue() {}

func (l LiteralString) isConstantValue() {}

// Literal Bool

type LiteralBool struct {
	Value bool
}

func (l LiteralBool) isValue() {}

func (l LiteralBool) isConstantValue() {}

// Literal List

type LiteralList struct {
	Items []Value
}

func (l LiteralList) isValue() {}

func (l LiteralList) isConstantValue() {}

// Literal Map

type KeyValue struct {
	Key, Value Value
}

type LiteralMap struct {
	Entries []KeyValue
}

func (l LiteralMap) isValue() {}

func (l LiteralMap) isConstantValue() {}

// Literal Set

type LiteralSet struct {
	Items []Value
}

func (l LiteralSet) isValue() {}

func (l LiteralSet) isConstantValue() {}

// Empty List

type EmptyList struct {
	Type Type
}

func (e EmptyList) isValue() {}

func (e EmptyList) isConstantValue() {}

// Empty Set

type EmptySet struct {
	Type Type
}

func (e EmptySet) isValue() {}

func (e EmptySet) isConstantValue() {}

// LiteralStruct

type LiteralProperty struct {
	Name  string
	Value Value
}

type LiteralStruct struct {
	Name       string
	Properties []LiteralProperty
}

func (l LiteralStruct) isValue() {}

func (l LiteralStruct) isConstantValue() {}
