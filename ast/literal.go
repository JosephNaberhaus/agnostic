package ast

// TODO move to value.go

// Literal Int32

type LiteralInt struct {
	Meta
	Value int64
}

func (l LiteralInt) isValue() {}

func (l LiteralInt) isConstantValue() {}

// Literal Rune

type LiteralRune struct {
	Meta
	Value rune
}

func (l LiteralRune) isValue() {}

func (l LiteralRune) isConstantValue() {}

// Literal String

type LiteralString struct {
	Meta
	Value string
}

func (l LiteralString) isValue() {}

func (l LiteralString) isConstantValue() {}

// Literal Bool

type LiteralBool struct {
	Meta
	Value bool
}

func (l LiteralBool) isValue() {}

func (l LiteralBool) isConstantValue() {}

// Literal List

type LiteralList struct {
	Meta
	Items []Value
}

func (l LiteralList) isValue() {}

func (l LiteralList) isConstantValue() {}

// Literal Map

type KeyValue struct {
	Meta
	Key, Value Value
}

type LiteralMap struct {
	Meta
	Entries []KeyValue
}

func (l LiteralMap) isValue() {}

func (l LiteralMap) isConstantValue() {}

// Literal Set

type LiteralSet struct {
	Meta
	Items []Value
}

func (l LiteralSet) isValue() {}

func (l LiteralSet) isConstantValue() {}

// Empty List

type EmptyList struct {
	Meta
	Type Type
}

func (e EmptyList) isValue() {}

func (e EmptyList) isConstantValue() {}

// Empty Set

type EmptySet struct {
	Meta
	Type Type
}

func (e EmptySet) isValue() {}

func (e EmptySet) isConstantValue() {}

// LiteralStruct

type LiteralProperty struct {
	Meta
	Name  string
	Value Value
}

type LiteralStruct struct {
	Meta
	Name       string
	Properties []LiteralProperty
}

func (l LiteralStruct) isValue() {}

func (l LiteralStruct) isConstantValue() {}
