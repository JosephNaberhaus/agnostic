package code

type StatementMetadata struct {
	Parent *MethodDef
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
