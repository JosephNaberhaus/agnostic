package ast

// Assignment

type Assignment struct {
	To   Assignable
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
