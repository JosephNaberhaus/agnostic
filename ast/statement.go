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
	Initialization Statement
	Condition      Value
	AfterEach      Statement
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
