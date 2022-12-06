package ast

// Call

type Call struct {
	Function  Callable
	Arguments []Value
}

func (c Call) isValue() {}

func (c Call) isStatement() {}

// Lookup

type Lookup struct {
	From, Key Value
}

func (l Lookup) isValue() {}

// New

type New struct {
	Model Model
}

func (n New) isValue() {}

// Length

type Length struct {
	Value Value
}

func (l Length) isValue() {}

// SetContains

type SetContains struct {
	Set   Value
	Value Value
}

func (s SetContains) isValue() {}

// Pop

type Pop struct {
	Value Value
}

func (p Pop) isValue() {}

func (p Pop) isStatement() {}
