package code

type ArgumentDefMetadata struct{}

type FieldDefMetadata struct {
	Parent *ModelDef
}

type MethodDefMetadata struct {
	Parent *ModelDef
}

type ModelDefMetadata struct {
	FieldMap map[string]*FieldDef
}

type PropertyMetadata struct {
	Type Type
}

type ThisMetadata struct {
	This *ModelDef
}
