package ast

// Literal Int32

type LiteralInt32 struct {
	Value int32
}

func (l LiteralInt32) isValue() {}

// Literal String

type LiteralString struct {
	Value string
}

func (l LiteralString) isValue() {}
