package code

type ModuleMetadata struct {
	FunctionMap  map[string]*FunctionDef
	ConstantsMap map[string]*ConstantDef
}

type ConstantDefMetadata struct {
	Type Type
}

type FunctionDefMetadata struct {
	IsMethod bool
}
