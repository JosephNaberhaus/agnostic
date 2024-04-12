package ast

// Call

type Call struct {
	Meta
	Function  Callable
	Arguments []Value
}

func (c Call) isValue() {}

func (c Call) isStatement() {}

// Lookup

type Lookup struct {
	Meta
	From, Key Value
}

func (l Lookup) isValue() {}

// New

type New struct {
	Meta
	Model Model
}

func (n New) isValue() {}

// Length

type Length struct {
	Meta
	Value Value
}

func (l Length) isValue() {}

// SetContains

type SetContains struct {
	Meta
	Set   Value
	Value Value
}

func (s SetContains) isValue() {}

// Pop

type Pop struct {
	Meta
	Value Value
}

func (p Pop) isValue() {}

func (p Pop) isStatement() {}

// Null

type Null struct {
	Meta
}

func (n Null) isValue() {}

// Self

type Self struct {
	Meta
}

func (s Self) isValue() {}
