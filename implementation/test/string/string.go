package string

import (
	"fmt"
	"github.com/JosephNaberhaus/agnostic/code"
	"github.com/JosephNaberhaus/agnostic/implementation"
	"github.com/JosephNaberhaus/agnostic/implementation/text"
)

// Mapper maps a code module to a simple string representation of it. Useful for testing and debugging.
type Mapper struct{}

var _ implementation.Mapper = (*Mapper)(nil)

func (m *Mapper) Config() implementation.Config {
	return implementation.Config{}
}

func (m *Mapper) MapLiteralInt32(original *code.LiteralInt32) (text.Node, error) {
	return text.Span(fmt.Sprintf("LiteralInt32(Value: %v)", original.Value)), nil
}

func (m *Mapper) MapLiteralString(original *code.LiteralString) (text.Node, error) {
	return text.Span(fmt.Sprintf("LiteralString(Value:%v)", original.Value)), nil
}

func (m *Mapper) MapFieldDef(original *code.FieldDef) (text.Node, error) {
	return text.Span(fmt.Sprintf("FieldDef(Name:%v,Type:%v)", original.Name, original.Type)), nil
}

func (m *Mapper) MapArgumentDef(original *code.ArgumentDef) (text.Node, error) {
	return text.Span(fmt.Sprintf("ArgumentDef(Name:%v,Type:%v)", original.Name, original.Type)), nil
}

func (m *Mapper) MapMethodDef(original *code.MethodDef) (text.Node, error) {
	return text.Span("MethodDef"), nil
}

func (m *Mapper) MapModelDef(original *code.ModelDef) (text.Node, error) {
	return text.Span("ModelDef"), nil
}

func (m *Mapper) MapProperty(original *code.Property) (text.Node, error) {
	return text.Span("Property"), nil
}

func (m *Mapper) MapThis(original *code.This) (text.Node, error) {
	return text.Span("This"), nil
}

func (m *Mapper) MapModule(original *code.Module) (text.Node, error) {
	return text.Span("Module"), nil
}

func (m *Mapper) MapUnaryOperator(original code.UnaryOperator) (text.Node, error) {
	return text.Span("UnaryOperator"), nil
}

func (m *Mapper) MapUnaryOperation(original *code.UnaryOperation) (text.Node, error) {
	return text.Span("UnaryOperation"), nil
}

func (m *Mapper) MapBinaryOperator(original code.BinaryOperator) (text.Node, error) {
	return text.Span("BinaryOperator"), nil
}

func (m *Mapper) MapBinaryOperation(original *code.BinaryOperation) (text.Node, error) {
	return text.Span("BinaryOperation"), nil
}

func (m *Mapper) MapAssignment(original *code.Assignment) (text.Node, error) {
	return text.Span("Assignment"), nil
}

func (m *Mapper) MapIf(original *code.If) (text.Node, error) {
	return text.Span("If"), nil
}

func (m *Mapper) MapElseIf(original *code.ElseIf) (text.Node, error) {
	return text.Span("ElseIf"), nil
}

func (m *Mapper) MapElse(original *code.Else) (text.Node, error) {
	return text.Span("Else"), nil
}

func (m *Mapper) MapConditional(original *code.Conditional) (text.Node, error) {
	return text.Span("Conditional"), nil
}

func (m *Mapper) MapModel(original *code.Model) (text.Node, error) {
	return text.Span("Model"), nil
}

func (m *Mapper) MapPrimitive(original code.Primitive) (text.Node, error) {
	return text.Span("Primitive"), nil
}
