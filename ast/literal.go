package ast

// TODO move to value.go

// Literal Int32

type LiteralInt struct {
	Value int64
}

func (l LiteralInt) isValue() {}

// Literal String

type LiteralString struct {
	Value string
}

func (l LiteralString) isValue() {}
