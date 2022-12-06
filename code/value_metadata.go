package code

type CallMetadata struct {
	Definition *FunctionDef
}

type LookupFromType int

const (
	LookupTypeList LookupFromType = iota + 1
	LookupTypeMap
	LookupTypeString
)

type LookupMetadata struct {
	LookupType LookupFromType
	OutputType Type
}

type NewMetadata struct{}

type LengthType int

const (
	LengthTypeString LengthType = iota + 1
	LengthTypeList
	LengthTypeMap
)

type LengthMetadata struct {
	LengthType LengthType
}

type SetContainsMetadata struct{}

type PopMetadata struct{}
