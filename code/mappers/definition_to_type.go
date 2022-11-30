package mappers

import (
	"errors"
	"github.com/JosephNaberhaus/agnostic/code"
)

// mapValueToType finds the base code.Type of a code.Definition.
func mapDefinitionToType(definition code.Definition) (code.Type, error) {
	return code.MapDefinition[code.Type](definition, &definitionToTypeMapper{})
}

func mapDefinitionToModelType(definition code.Definition) (*code.Model, error) {
	definitionType, err := mapDefinitionToType(definition)
	if err != nil {
		return nil, err
	}

	modelType, ok := definitionType.(*code.Model)
	if !ok {
		// TODO improve
		return nil, errors.New("definition is not a model")
	}

	return modelType, nil
}

type definitionToTypeMapper struct{}

var _ code.DefinitionMapper[code.Type] = (*definitionToTypeMapper)(nil)

func (d *definitionToTypeMapper) MapFieldDef(original *code.FieldDef) (code.Type, error) {
	return original.Type, nil
}

func (d *definitionToTypeMapper) MapArgumentDef(original *code.ArgumentDef) (code.Type, error) {
	return original.Type, nil
}

func (d *definitionToTypeMapper) MapDeclare(original *code.Declare) (code.Type, error) {
	return original.Type, nil
}
