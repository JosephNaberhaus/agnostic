package ast

// Unary Operation

type UnaryOperator int

const (
	Not UnaryOperator = iota + 1
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
	Add BinaryOperator = iota + 1
	Subtract
	Equals
)

type BinaryOperation struct {
	Left, Right Value
	Operator    BinaryOperator
}

func (b BinaryOperation) isValue() {}
