package ast

// Assignment

type Assignment struct {
	To   Value
	From Value
}

func (a Assignment) isStatement() {}

// If

type If struct {
	Condition  Value
	Statements []Statement
}

type ElseIf struct {
	Condition  Value
	Statements []Statement
}

type Else struct {
	Statements []Statement
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
