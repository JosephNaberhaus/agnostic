// Code generated by tool/code_generator. DO NOT EDIT.

package code

type ArgumentDef struct {
	Name	string
	Type	Type
	ArgumentDefMetadata
}

func (a *ArgumentDef) isDefinition()	{}

type FunctionDef struct {
	Name		string
	Arguments	[]*ArgumentDef
	Block		*Block
	ReturnType	Type
	FunctionDefMetadata
}

func (f *FunctionDef) isCallableDef()	{}
