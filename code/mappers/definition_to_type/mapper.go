package definition_to_type

import "github.com/JosephNaberhaus/agnostic/code"

type Mapper struct{}

func (m Mapper) MapFieldDefNoError(original *code.FieldDef) code.Type {
	return original.Type
}

func (m Mapper) MapArgumentDefNoError(original *code.ArgumentDef) code.Type {
	return original.Type
}

func (m Mapper) MapDeclareNoError(original *code.Declare) code.Type {
	return original.Type
}

func (m Mapper) MapForInNoError(original *code.ForIn) code.Type {
	return original.ItemType
}

func (m Mapper) MapConstantDefNoError(original *code.ConstantDef) code.Type {
	return original.Type
}
