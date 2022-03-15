package agnostic

type Access int

const (
	Private Access = iota
	PrivateWithGetter
	Public
)

type Field struct {
	Name   string
	Access Access
	Type   Type
}

type Model struct {
	Name    string
	Path    string
	Fields  []Field
	Methods []Method
}
