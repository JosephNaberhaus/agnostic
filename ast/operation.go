package ast

// Unary Operation

type UnaryOperator int

const (
	Not UnaryOperator = iota
	Negate
)

type UnaryOperation struct {
	Value    Value
	Operator UnaryOperator
}

func (u UnaryOperation) isValue() {}

// Binary Operation

type BinaryOperator int

const (
	Add BinaryOperator = iota
)

type BinaryOperation struct {
	Left, Right Value
	Operator    BinaryOperator
}

func (b BinaryOperation) isValue() {}
