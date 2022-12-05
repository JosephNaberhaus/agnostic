package ast

// Call

type Call struct {
	Function  Callable
	Arguments []Value
}

func (c Call) isValue() {}

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
