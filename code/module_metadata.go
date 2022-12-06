package code

type ModuleMetadata struct {
	ModelMap     map[string]*ModelDef
	FunctionMap  map[string]*FunctionDef
	ConstantsMap map[string]*ConstantDef
}

type ConstantDefMetadata struct {
	Parent *Module
	Type   Type
}

type FunctionDefMetadata struct {
	IsMethod bool
}
