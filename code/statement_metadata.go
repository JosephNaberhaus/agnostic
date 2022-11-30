package code

type StatementMetadata struct {
	Parent *FunctionDef
}

type AssignmentMetadata struct {
	StatementMetadata
}

type IfMetadata struct{}

type ElseIfMetadata struct{}

type ElseMetadata struct{}

type ConditionalMetadata struct {
	StatementMetadata
}

type ReturnMetadata struct {
	StatementMetadata
}

type DeclareMetadata struct {
	StatementMetadata
	Type Type
}
