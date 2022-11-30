package code

type ArgumentDefMetadata struct{}

type FieldDefMetadata struct {
	Parent *ModelDef
}

type MethodDefMetadata struct {
	Parent *ModelDef
}

type ModelDefMetadata struct {
	FieldMap  map[string]*FieldDef
	MethodMap map[string]*MethodDef
}

type VariableMetadata struct {
	Definition Definition
	Type       Type
}

type PropertyMetadata struct {
	Type Type
}
