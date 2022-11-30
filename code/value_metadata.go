package code

type CallMetadata struct{}

type LookupFromType int

const (
	LookupTypeList LookupFromType = iota + 1
	LookupTypeMap
)

type LookupMetadata struct {
	LookupType LookupFromType
	OutputType Type
}

type NewMetadata struct{}
