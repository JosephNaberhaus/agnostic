package agnostic

type Field struct {
	Name string
	Type Type
}

type Model struct {
	Package *Package
	Name    string
	Fields  []Field
	Methods []Method
}
