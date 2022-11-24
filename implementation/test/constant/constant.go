package constant

import (
	"github.com/JosephNaberhaus/agnostic/code"
	"github.com/JosephNaberhaus/agnostic/implementation"
	"github.com/JosephNaberhaus/agnostic/implementation/text"
)

type Mapper struct {
	Constant text.Node
}

var _ implementation.Mapper = (*Mapper)(nil)

func (m *Mapper) Config() implementation.Config {
	return implementation.Config{}
}

func (m *Mapper) MapLiteralInt32(_ *code.LiteralInt32) (text.Node, error) {
	return m.Constant, nil
}

func (m *Mapper) MapLiteralString(_ *code.LiteralString) (text.Node, error) {
	return m.Constant, nil
}

func (m *Mapper) MapFieldDef(_ *code.FieldDef) (text.Node, error) {
	return m.Constant, nil
}

func (m *Mapper) MapArgumentDef(_ *code.ArgumentDef) (text.Node, error) {
	return m.Constant, nil
}

func (m *Mapper) MapMethodDef(_ *code.MethodDef) (text.Node, error) {
	return m.Constant, nil
}

func (m *Mapper) MapModelDef(_ *code.ModelDef) (text.Node, error) {
	return m.Constant, nil
}

func (m *Mapper) MapProperty(_ *code.Property) (text.Node, error) {
	return m.Constant, nil
}

func (m *Mapper) MapThis(_ *code.This) (text.Node, error) {
	return m.Constant, nil
}

func (m *Mapper) MapModule(_ *code.Module) (text.Node, error) {
	return m.Constant, nil
}

func (m *Mapper) MapUnaryOperator(_ code.UnaryOperator) (text.Node, error) {
	return m.Constant, nil
}

func (m *Mapper) MapUnaryOperation(_ *code.UnaryOperation) (text.Node, error) {
	return m.Constant, nil
}

func (m *Mapper) MapBinaryOperator(_ code.BinaryOperator) (text.Node, error) {
	return m.Constant, nil
}

func (m *Mapper) MapBinaryOperation(_ *code.BinaryOperation) (text.Node, error) {
	return m.Constant, nil
}

func (m *Mapper) MapAssignment(_ *code.Assignment) (text.Node, error) {
	return m.Constant, nil
}

func (m *Mapper) MapIf(_ *code.If) (text.Node, error) {
	return m.Constant, nil
}

func (m *Mapper) MapElseIf(_ *code.ElseIf) (text.Node, error) {
	return m.Constant, nil
}

func (m *Mapper) MapElse(_ *code.Else) (text.Node, error) {
	return m.Constant, nil
}

func (m *Mapper) MapConditional(_ *code.Conditional) (text.Node, error) {
	return m.Constant, nil
}

func (m *Mapper) MapModel(_ *code.Model) (text.Node, error) {
	return m.Constant, nil
}

func (m *Mapper) MapPrimitive(_ code.Primitive) (text.Node, error) {
	return m.Constant, nil
}
