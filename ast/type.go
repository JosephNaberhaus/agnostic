package ast

// Model

type Model struct {
	Meta
	Name string
}

func (m Model) isType() {}

// Primitive

// TODO: Primitive should be broken up into structs

type Primitive int

const (
	Boolean Primitive = iota + 1
	Int
	Rune
	String
	Void
)

func (p Primitive) isType() {}

// List

type List struct {
	Meta
	Base Type
}

func (l List) isType() {}

// Map

type Map struct {
	Meta
	Key, Value Type
}

func (m Map) isType() {}

// Set

type Set struct {
	Meta
	Base Type
}

func (s Set) isType() {}
