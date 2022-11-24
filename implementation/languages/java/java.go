package java

import (
	"fmt"
	"github.com/JosephNaberhaus/agnostic/code"
	"github.com/JosephNaberhaus/agnostic/implementation"
	"github.com/JosephNaberhaus/agnostic/implementation/text"
)

type Mapper struct{}

var _ implementation.Mapper = (*Mapper)(nil)

func (m Mapper) Config() implementation.Config {
	return implementation.Config{
		Indent: "    ",
	}
}

func (m Mapper) MapLiteralInt32(original *code.LiteralInt32) (text.Node, error) {
	return text.Span(fmt.Sprintf("int(%d)", original.Value)), nil
}

func (m Mapper) MapLiteralString(original *code.LiteralString) (text.Node, error) {
	return text.Span(fmt.Sprintf("\"%s\"", original.Value)), nil
}

func (m Mapper) MapFieldDef(original *code.FieldDef) (text.Node, error) {
	typeNode, err := code.MapType[text.Node](original.Type, m)
	if err != nil {
		return nil, err
	}

	return text.Group{
		typeNode,
		text.Span(" "),
		text.Span(original.Name),
		text.Span(";"),
	}, nil
}

func (m Mapper) MapArgumentDef(original *code.ArgumentDef) (text.Node, error) {
	typeNode, err := code.MapType[text.Node](original.Type, m)
	if err != nil {
		return nil, err
	}

	return text.Group{
		typeNode,
		text.Span(" "),
		text.Span(original.Name),
	}, nil
}

func (m Mapper) MapMethodDef(original *code.MethodDef) (text.Node, error) {
	returnTypeNode, err := code.MapType[text.Node](original.ReturnType, m)
	if err != nil {
		return nil, err
	}

	var argumentNodes []text.Node
	for i, argument := range original.Arguments {
		argumentNode, err := m.MapArgumentDef(argument)
		if err != nil {
			return nil, err
		}

		argumentNodes = append(argumentNodes, argumentNode)

		if i != len(original.Arguments)-1 {
			argumentNodes = append(argumentNodes, text.Span(", "))
		}
	}

	var statementNodes []text.Node
	for _, statement := range original.Statements {
		statementNode, err := code.MapStatement[text.Node](statement, m)
		if err != nil {
			return nil, err
		}

		statementNodes = append(statementNodes, statementNode)
	}

	return text.Block{
		text.Group{
			text.Span("public "),
			returnTypeNode,
			text.Span(" "),
			text.Span(original.Name),
			text.Span("("),
			text.Group(argumentNodes),
			text.Span(") {"),
		},
		text.IndentedBlock(statementNodes),
		text.Span("}"),
	}, nil
}

func (m Mapper) MapModelDef(original *code.ModelDef) (text.Node, error) {
	var fieldNodes []text.Node
	for _, field := range original.Fields {
		fieldNode, err := m.MapFieldDef(field)
		if err != nil {
			return nil, err
		}

		fieldNodes = append(fieldNodes, fieldNode)
	}

	var methodNodes []text.Node
	for _, method := range original.Methods {
		methodNode, err := m.MapMethodDef(method)
		if err != nil {
			return nil, err
		}

		methodNodes = append(methodNodes, methodNode)
	}

	return text.Block{
		text.Group{
			text.Span("class "),
			text.Span(original.Name),
			text.Span(" {"),
		},
		text.IndentedBlock{
			text.Block(fieldNodes),
			text.Group(methodNodes),
		},
		text.Span("}"),
	}, nil
}

func (m Mapper) MapModule(original *code.Module) (text.Node, error) {
	var modelNodes []text.Node
	for _, model := range original.Models {
		modelNode, err := m.MapModelDef(model)
		if err != nil {
			return nil, err
		}

		modelNodes = append(modelNodes, modelNode)
	}

	return text.Group(modelNodes), nil
}

func (m Mapper) MapUnaryOperator(original code.UnaryOperator) (text.Node, error) {
	switch original {
	case code.Not:
		return text.Span("!"), nil
	case code.Negate:
		return text.Span("-"), nil
	}

	// TODO remove the need for this
	panic("oh no")
}

func (m Mapper) MapUnaryOperation(original *code.UnaryOperation) (text.Node, error) {
	operatorNode, err := m.MapUnaryOperator(original.Operator)
	if err != nil {
		return nil, err
	}

	valueNode, err := code.MapValue[text.Node](original.Value, m)
	if err != nil {
		return nil, err
	}

	return text.Group{
		operatorNode,
		valueNode,
	}, nil
}

func (m Mapper) MapBinaryOperator(original code.BinaryOperator) (text.Node, error) {
	// TODO remove this
	return text.Span("+"), nil
}

func (m Mapper) MapBinaryOperation(original *code.BinaryOperation) (text.Node, error) {
	leftNode, err := code.MapValue[text.Node](original.Left, m)
	if err != nil {
		return nil, err
	}

	operatorNode, err := m.MapBinaryOperator(original.Operator)
	if err != nil {
		return nil, err
	}

	rightNode, err := code.MapValue[text.Node](original.Right, m)
	if err != nil {
		return nil, err
	}

	return text.Group{
		leftNode,
		operatorNode,
		rightNode,
	}, nil
}

func (m Mapper) MapAssignment(original *code.Assignment) (text.Node, error) {
	toNode, err := code.MapAssignable[text.Node](original.To, m)
	if err != nil {
		return nil, err
	}

	fromNode, err := code.MapValue[text.Node](original.From, m)
	if err != nil {
		return nil, err
	}

	return text.Group{
		toNode,
		text.Span(" = "),
		fromNode,
		text.Span(";"),
	}, nil
}

func (m Mapper) MapModel(original *code.Model) (text.Node, error) {
	return text.Span(original.Name), nil
}

func (m Mapper) MapPrimitive(original code.Primitive) (text.Node, error) {
	switch original {
	case code.Boolean:
		return text.Span("bool"), nil
	case code.Int32:
		return text.Span("int"), nil
	case code.String:
		return text.Span("String"), nil
	case code.Void:
		return text.Span("void"), nil
	}

	// TODO remove the need for this
	panic("no!!!")
}

func (m Mapper) MapIf(original *code.If) (text.Node, error) {
	conditionNode, err := code.MapValue[text.Node](original.Condition, m)
	if err != nil {
		return nil, err
	}

	var statementNodes []text.Node
	for _, statement := range original.Statements {
		statementNode, err := code.MapStatement[text.Node](statement, m)
		if err != nil {
			return nil, err
		}

		statementNodes = append(statementNodes, statementNode)
	}

	return text.Block{
		text.Group{
			text.Span("if ("),
			conditionNode,
			text.Span(") {"),
		},
		text.IndentedBlock(statementNodes),
		// The closing "}" is added by MapConditional
	}, nil
}

func (m Mapper) MapElseIf(original *code.ElseIf) (text.Node, error) {
	conditionNode, err := code.MapValue[text.Node](original.Condition, m)
	if err != nil {
		return nil, err
	}

	var statementNodes []text.Node
	for _, statement := range original.Statements {
		statementNode, err := code.MapStatement[text.Node](statement, m)
		if err != nil {
			return nil, err
		}

		statementNodes = append(statementNodes, statementNode)
	}

	return text.Block{
		text.Group{
			text.Span("} else if ("),
			conditionNode,
			text.Span(") {"),
		},
		text.IndentedBlock(statementNodes),
		// The closing "}" is added by MapConditional
	}, nil
}

func (m Mapper) MapElse(original *code.Else) (text.Node, error) {
	var statementNodes []text.Node
	for _, statement := range original.Statements {
		statementNode, err := code.MapStatement[text.Node](statement, m)
		if err != nil {
			return nil, err
		}

		statementNodes = append(statementNodes, statementNode)
	}

	return text.Block{
		text.Span("} else {"),
		text.IndentedBlock(statementNodes),
		// The closing "}" is added by MapConditional
	}, nil
}

func (m Mapper) MapConditional(original *code.Conditional) (text.Node, error) {
	ifNode, err := m.MapIf(original.If)
	if err != nil {
		return nil, err
	}

	var elseIfNodes []text.Node
	for _, elseIf := range original.ElseIfs {
		elseIfNode, err := m.MapElseIf(elseIf)
		if err != nil {
			return nil, err
		}

		elseIfNodes = append(elseIfNodes, elseIfNode)
	}

	elseNode, err := m.MapElse(original.Else)
	if err != nil {
		return nil, err
	}

	return text.Block{
		ifNode,
		text.Group(elseIfNodes),
		elseNode,
		text.Span("}"),
	}, nil
}

func (m Mapper) MapProperty(original *code.Property) (text.Node, error) {
	valueNode, err := code.MapValue[text.Node](original.Of, m)
	if err != nil {
		return nil, err
	}

	return text.Group{
		valueNode,
		text.Span("."),
		text.Span(original.Name),
	}, nil
}

func (m Mapper) MapThis(original *code.This) (text.Node, error) {
	return text.Span("this"), nil
}
