package token

type Type int

const (
	// Literals

	NumberLit Type = iota
	StringLit
	BooleanLit

	// Keywords

	Model
	Func

	// Structure

	StartBlock
	EndBlock
	StartArguments
	EndArguments
	Semicolon

	// Operators

	OpenParen
	CloseParen
	Add
	Multiply
	Subtract
	Divide
	Assign

	// Names

	ModelName
	FieldName
	FuncName
	ArgName
	VariableName
	TypeName

	// Comments

	ModelComment
	FieldComment
	FuncComment
	StatementComment

	// Control

	EOF
)
