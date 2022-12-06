package ast

// Unary Operation

type UnaryOperator int

const (
	Not UnaryOperator = iota + 1
	Negate
	CastToInt
	CastToString
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
	Multiply
	Divide
	Equal
	NotEqual
	LessThan
	LessThanOrEqualTo
	GreaterThan
	GreaterThanOrEqualTo
	Or
	And
)

type BinaryOperation struct {
	Left, Right Value
	Operator    BinaryOperator
}

func (b BinaryOperation) isValue() {}
