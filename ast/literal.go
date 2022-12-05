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
