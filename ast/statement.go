package ast

// Block

type Block struct {
	Statements []Statement
}

// Assignment

type Assignment struct {
	To   Value
	From Value
}

func (a Assignment) isStatement() {}

// If

type If struct {
	Condition Value
	Block     Block
}

type ElseIf struct {
	Condition Value
	Block     Block
}

type Else struct {
	Block Block
}

type Conditional struct {
	If      If
	ElseIfs []ElseIf
	Else    Else
}

func (c Conditional) isStatement() {}

// Return

type Return struct {
	Value Value
}

func (r Return) isStatement() {}

// Declare

type Declare struct {
	Value Value
	Name  string
}

func (d Declare) isStatement() {}

func (d Declare) isDefinition() {}

// For

type For struct {
	Initialization Optional[Statement]
	Condition      Value
	AfterEach      Optional[Statement]
	Block          Block
}

func (f For) isStatement() {}

// For In

type ForIn struct {
	Iterable Value
	ItemName string
	Block    Block
}

func (f ForIn) isStatement() {}

func (f ForIn) isDefinition() {}

// AddToSet

// TODO this should probably be a value that returns void
type AddToSet struct {
	To    Value
	Value Value
}

func (a AddToSet) isStatement() {}

// Push

type Push struct {
	To    Value
	Value Value
}

func (p Push) isStatement() {}

// Declare

type DeclareNull struct {
	Name string
	Type Type
}

func (d DeclareNull) isStatement() {}

func (d DeclareNull) isDefinition() {}

// Break

type Break struct{}

func (b Break) isStatement() {}

// Continue

type Continue struct{}

func (c Continue) isStatement() {}
