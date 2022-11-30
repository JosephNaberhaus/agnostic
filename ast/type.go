package ast

// Model

type Model struct {
	Name string
}

func (m Model) isType() {}

// Primitive

type Primitive int

const (
	Boolean Primitive = iota + 1
	Int
	String
	Void
)

func (p Primitive) isType() {}

// List

type List struct {
	Base Type
}

func (l List) isType() {}

// Map

type Map struct {
	Key, Value Type
}

func (m Map) isType() {}
