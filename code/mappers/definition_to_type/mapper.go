package definition_to_type

import "github.com/JosephNaberhaus/agnostic/code"

type Mapper struct{}

func (m Mapper) MapFieldDef(original *code.FieldDef) code.Type {
	return original.Type
}

func (m Mapper) MapArgumentDef(original *code.ArgumentDef) code.Type {
	return original.Type
}

func (m Mapper) MapDeclare(original *code.Declare) code.Type {
	return original.Type
}

func (m Mapper) MapForIn(original *code.ForIn) code.Type {
	return original.ItemType
}

func (m Mapper) MapConstantDef(original *code.ConstantDef) code.Type {
	return original.Type
}
