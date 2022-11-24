package ast

// Model

type Model struct {
	Name string
}

func (m Model) isType() {}

// Primitive

type Primitive int

const (
	Boolean Primitive = iota
	Int32
	String
	Void
)

func (p Primitive) isType() {}
