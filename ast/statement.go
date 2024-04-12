package ast

// Block

type Block struct {
	Meta
	Statements []Statement
}

// Assignment

type Assignment struct {
	Meta
	To   Value
	From Value
}

func (a Assignment) isStatement() {}

// If

type If struct {
	Meta
	Condition Value
	Block     Block
}

type ElseIf struct {
	Meta
	Condition Value
	Block     Block
}

type Else struct {
	Meta
	Block Block
}

type Conditional struct {
	Meta
	If      If
	ElseIfs []ElseIf
	Else    Else
}

func (c Conditional) isStatement() {}

// Return

type Return struct {
	Meta
	Value Value
}

func (r Return) isStatement() {}

// Declare

type Declare struct {
	Meta
	Value Value
	Name  string
}

func (d Declare) isStatement() {}

func (d Declare) isDefinition() {}

// For

type For struct {
	Meta
	Initialization Optional[Statement]
	Condition      Value
	AfterEach      Optional[Statement]
	Block          Block
}

func (f For) isStatement() {}

// For In

type ForIn struct {
	Meta
	Iterable Value
	ItemName string
	Block    Block
}

func (f ForIn) isStatement() {}

func (f ForIn) isDefinition() {}

// AddToSet

// TODO this should probably be a value that returns void
type AddToSet struct {
	Meta
	To    Value
	Value Value
}

func (a AddToSet) isStatement() {}

// Push

type Push struct {
	Meta
	To    Value
	Value Value
}

func (p Push) isStatement() {}

// Declare

type DeclareNull struct {
	Meta
	Name string
	Type Type
}

func (d DeclareNull) isStatement() {}

func (d DeclareNull) isDefinition() {}

// Break

type Break struct {
	Meta
}

func (b Break) isStatement() {}

// Continue

type Continue struct {
	Meta
}

func (c Continue) isStatement() {}
