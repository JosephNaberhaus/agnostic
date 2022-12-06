package code

type BlockMetadata struct{}

// TODO we don't actually need this for every statement. Just return.
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

type ForMetadata struct {
	StatementMetadata
}

type ForInType int

const (
	ForInTypeList ForInType = iota + 1
	ForInTypeSet
)

type ForInMetadata struct {
	StatementMetadata
	ForInType ForInType
	ItemType  Type
}

type CallFunctionMetadata struct {
	StatementMetadata
}

type AddToSetMetadata struct {
	StatementMetadata
}

type PushMetadata struct{}
