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

func (m Mapper) MapLiteralInt(original *code.LiteralInt) (text.Node, error) {
	return text.Span(fmt.Sprintf("%dL", original.Value)), nil
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
	return m.MapFunctionDef(original.Function)
}

func (m Mapper) MapModelDef(original *code.ModelDef) (text.Node, error) {
	fieldNodes, err := code.MapNodes[text.Node](original.Fields, m)
	if err != nil {
		return nil, err
	}

	methodNodes, err := code.MapNodes[text.Node](original.Methods, m)
	if err != nil {
		return nil, err
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
	modelNodes, err := code.MapNodes[text.Node](original.Models, m)
	if err != nil {
		return nil, err
	}

	functionNodes, err := code.MapNodes[text.Node](original.Functions, m)
	if err != nil {
		return nil, err
	}

	return text.Block{
		text.Span(fmt.Sprintf("class %sFunctions {", original.Name)),
		text.IndentedBlock(functionNodes),
		text.Span("}"),
		text.Span(""),
		text.Block(modelNodes),
	}, nil
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
	switch original {
	case code.Add:
		return text.Span("+"), nil
	case code.Equals:
		return text.Span("=="), nil
	}

	// TODO remove the need for this
	panic("oh no")
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
		text.Span("("),
		leftNode,
		text.Span(" "),
		operatorNode,
		text.Span(" "),
		rightNode,
		text.Span(")"),
	}, nil
}

func (m Mapper) MapAssignment(original *code.Assignment) (text.Node, error) {
	toNode, err := code.MapValue[text.Node](original.To, m)
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
	case code.Int:
		return text.Span("Long"), nil
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

	statementNodes, err := code.MapNodes[text.Node](original.Statements, m)
	if err != nil {
		return nil, err
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

	statementNodes, err := code.MapNodes[text.Node](original.Statements, m)
	if err != nil {
		return nil, err
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
	statementNodes, err := code.MapNodes[text.Node](original.Statements, m)
	if err != nil {
		return nil, err
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

	elseIfNodes, err := code.MapNodes[text.Node](original.ElseIfs, m)
	if err != nil {
		return nil, err
	}

	var elseNode text.Node
	if original.Else != nil {
		elseNode, err = m.MapElse(original.Else)
		if err != nil {
			return nil, err
		}
	}

	return text.Block{
		ifNode,
		text.Group(elseIfNodes),
		elseNode,
		text.Span("}"),
	}, nil
}

func (m Mapper) MapProperty(original *code.Property) (text.Node, error) {
	ofNode, err := code.MapValue[text.Node](original.Of, m)
	if err != nil {
		return nil, err
	}

	return text.Group{
		ofNode,
		text.Span("."),
		text.Span(original.Name),
	}, nil
}

func (m Mapper) MapVariable(original *code.Variable) (text.Node, error) {
	return text.Span(original.Name), nil
}

func (m Mapper) MapList(original *code.List) (text.Node, error) {
	baseNode, err := code.MapType[text.Node](original.Base, m)
	if err != nil {
		return nil, err
	}

	return text.Group{
		text.Span("List<"),
		baseNode,
		text.Span(">"),
	}, nil
}

func (m Mapper) MapMap(original *code.Map) (text.Node, error) {
	keyNode, err := code.MapType[text.Node](original.Key, m)
	if err != nil {
		return nil, err
	}

	valueNode, err := code.MapType[text.Node](original.Value, m)
	if err != nil {
		return nil, err
	}

	return text.Group{
		text.Span("Map<"),
		keyNode,
		text.Span(", "),
		valueNode,
		text.Span(">"),
	}, nil
}

func (m Mapper) MapLookup(original *code.Lookup) (text.Node, error) {
	fromNode, err := code.MapValue[text.Node](original.From, m)
	if err != nil {
		return nil, err
	}

	keyNode, err := code.MapValue[text.Node](original.Key, m)
	if err != nil {
		return nil, err
	}

	return text.Group{
		fromNode,
		text.Span(".get("),
		keyNode,
		text.Span(")"),
	}, nil
}

func (m Mapper) MapFunctionDef(original *code.FunctionDef) (text.Node, error) {
	returnTypeNode, err := code.MapType[text.Node](original.ReturnType, m)
	if err != nil {
		return nil, err
	}

	var modifierNode text.Node
	if !original.IsMethod {
		modifierNode = text.Span("static ")
	}

	argumentNodes, err := code.MapNodes[text.Node](original.Arguments, m)
	if err != nil {
		return nil, err
	}

	statementNodes, err := code.MapNodes[text.Node](original.Statements, m)
	if err != nil {
		return nil, err
	}

	return text.Block{
		text.Group{
			text.Span("public "),
			modifierNode,
			returnTypeNode,
			text.Span(" "),
			text.Span(original.Name),
			text.Span("("),
			text.Join{
				Nodes: argumentNodes,
				Sep:   ", ",
			},
			text.Span(") {"),
		},
		text.IndentedBlock(statementNodes),
		text.Span("}"),
	}, nil
}

func (m Mapper) MapReturn(original *code.Return) (text.Node, error) {
	valueNode, err := code.MapValue[text.Node](original.Value, m)
	if err != nil {
		return nil, err
	}

	return text.Group{
		text.Span("return "),
		valueNode,
		text.Span(";"),
	}, nil
}

func (m Mapper) MapCall(original *code.Call) (text.Node, error) {
	functionNode, err := code.MapCallable[text.Node](original.Function, m)
	if err != nil {
		return nil, err
	}

	argumentNodes, err := code.MapNodes[text.Node](original.Arguments, m)
	if err != nil {
		return nil, err
	}

	return text.Group{
		functionNode,
		text.Span("("),
		text.Join{
			Nodes: argumentNodes,
			Sep:   ", ",
		},
		text.Span(")"),
	}, nil
}

func (m Mapper) MapFunction(original *code.Function) (text.Node, error) {
	return text.Span(original.Name), nil
}

func (m Mapper) MapFunctionProperty(original *code.FunctionProperty) (text.Node, error) {
	ofNode, err := code.MapValue[text.Node](original.Of, m)
	if err != nil {
		return nil, err
	}

	return text.Group{
		ofNode,
		text.Span("."),
		text.Span(original.Name),
	}, nil
}

func (m Mapper) MapNew(original *code.New) (text.Node, error) {
	modelNode, err := m.MapModel(original.Model)
	if err != nil {
		return nil, err
	}

	return text.Group{
		text.Span("(new "),
		modelNode,
		text.Span("())"),
	}, nil
}

func (m Mapper) MapDeclare(original *code.Declare) (text.Node, error) {
	valueNode, err := code.MapValue[text.Node](original.Value, m)
	if err != nil {
		return nil, err
	}

	typeNode, err := code.MapType[text.Node](original.Type, m)
	if err != nil {
		return nil, err
	}

	return text.Group{
		typeNode,
		text.Span(" "),
		text.Span(original.Name),
		text.Span(" = "),
		valueNode,
		text.Span(";"),
	}, nil
}
