package mappers

import (
	"errors"
	"fmt"
	"github.com/JosephNaberhaus/agnostic/ast"
	"github.com/JosephNaberhaus/agnostic/code"
	"strings"
)

func AstToCode(original ast.Module) (*code.Module, error) {
	mapper := &astToCodeMapper{}

	result, err := mapAstNodeTo[*code.Module](original, mapper)
	if err != nil {
		return nil, err
	}

	for len(mapper.deferred) != 0 {
		deferredCall := mapper.dequeueDefer()
		mapper.stack = deferredCall.stackSnapshot
		err = deferredCall.deferredFunc()
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

type deferredCall struct {
	stackSnapshot codeStack
	deferredFunc  func() error
}

type astToCodeMapper struct {
	stack    codeStack
	deferred []deferredCall
}

var _ ast.NodeMapper[code.Node] = (*astToCodeMapper)(nil)

// queueDefer queues a function that will be called after all other mapping is done.
//
// This allows the mapping to act like a breadth first search at defined points. This is used when we need to know that
// the code models will be populated with all information from the AST (up to a certain depth).
func (a *astToCodeMapper) queueDefer(deferredFunc func() error) {
	a.deferred = append(a.deferred, deferredCall{
		stackSnapshot: a.stack.copy(),
		deferredFunc:  deferredFunc,
	})
}

func (a *astToCodeMapper) dequeueDefer() deferredCall {
	result := a.deferred[0]
	a.deferred = a.deferred[1:]
	return result
}

func (a *astToCodeMapper) MapLiteralInt32(original ast.LiteralInt32) (code.Node, error) {
	value := new(code.LiteralInt32)
	a.stack.push(value)
	defer a.stack.pop()

	value.Value = original.Value

	return value, nil
}

func (a *astToCodeMapper) MapLiteralString(original ast.LiteralString) (code.Node, error) {
	value := new(code.LiteralString)
	a.stack.push(value)
	defer a.stack.pop()

	value.Value = original.Value

	return value, nil
}

func (a *astToCodeMapper) MapFieldDef(original ast.FieldDef) (code.Node, error) {
	value := new(code.FieldDef)
	a.stack.push(value)
	defer a.stack.pop()

	err := validateName(original.Name)
	if err != nil {
		return nil, err
	}
	value.Name = original.Name

	value.Type, err = mapAstNodeTo[code.Type](original.Type, a)
	if err != nil {
		return nil, err
	}

	var ok bool

	value.Parent, ok = firstOfType[*code.ModelDef](a.stack)
	if !ok {
		// TODO improve
		return nil, errors.New("no model def found in the parent of a field")
	}

	return value, nil
}

func (a *astToCodeMapper) MapArgumentDef(original ast.ArgumentDef) (code.Node, error) {
	value := new(code.ArgumentDef)
	a.stack.push(value)
	defer a.stack.pop()

	err := validateName(original.Name)
	if err != nil {
		return nil, err
	}
	value.Name = original.Name

	value.Type, err = mapAstNodeTo[code.Type](original.Type, a)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (a *astToCodeMapper) MapMethodDef(original ast.MethodDef) (code.Node, error) {
	value := new(code.MethodDef)
	a.stack.push(value)
	defer a.stack.pop()

	err := validateName(original.Name)
	if err != nil {
		return nil, err
	}
	value.Name = original.Name

	for _, originalArguments := range original.Arguments {
		argument, err := mapAstNodeTo[*code.ArgumentDef](originalArguments, a)
		if err != nil {
			return nil, err
		}
		value.Arguments = append(value.Arguments, argument)
	}

	value.ReturnType, err = mapAstNodeTo[code.Type](original.ReturnType, a)
	if err != nil {
		return nil, err
	}

	a.queueDefer(func() error {
		for _, originalStatement := range original.Statements {
			statement, err := mapAstNodeTo[code.Statement](originalStatement, a)
			if err != nil {
				return err
			}
			value.Statements = append(value.Statements, statement)
		}

		return nil
	})

	var ok bool
	value.Parent, ok = a.stack.peekParent().(*code.ModelDef)
	if !ok {
		// TODO improve
		return nil, errors.New("parent is not a model definition")
	}

	return value, nil
}

func (a *astToCodeMapper) MapModelDef(original ast.ModelDef) (code.Node, error) {
	value := new(code.ModelDef)
	a.stack.push(value)
	defer a.stack.pop()

	err := validateName(original.Name)
	if err != nil {
		return nil, err
	}
	value.Name = original.Name

	for _, originalField := range original.Fields {
		field, err := mapAstNodeTo[*code.FieldDef](originalField, a)
		if err != nil {
			return nil, err
		}
		value.Fields = append(value.Fields, field)
	}

	for _, originalMethod := range original.Methods {
		method, err := mapAstNodeTo[*code.MethodDef](originalMethod, a)
		if err != nil {
			return nil, err
		}
		value.Methods = append(value.Methods, method)
	}

	value.FieldMap = map[string]*code.FieldDef{}
	for _, codeField := range value.Fields {
		if _, alreadyExists := value.FieldMap[codeField.Name]; alreadyExists {
			// TODO improve
			return nil, errors.New("already exists")
		}

		value.FieldMap[codeField.Name] = codeField
	}

	return value, nil
}

func (a *astToCodeMapper) MapModule(original ast.Module) (code.Node, error) {
	value := new(code.Module)
	a.stack.push(value)
	defer a.stack.pop()

	for _, originalModel := range original.Models {
		model, err := mapAstNodeTo[*code.ModelDef](originalModel, a)
		if err != nil {
			return nil, err
		}
		value.Models = append(value.Models, model)
	}

	return value, nil
}

func (a *astToCodeMapper) MapUnaryOperator(original ast.UnaryOperator) (code.Node, error) {
	return code.UnaryOperator(original), nil
}

func (a *astToCodeMapper) MapUnaryOperation(original ast.UnaryOperation) (code.Node, error) {
	value := new(code.UnaryOperation)
	a.stack.push(value)
	defer a.stack.pop()

	var err error
	value.Value, err = mapAstNodeTo[code.Value](original.Value, a)
	if err != nil {
		return nil, err
	}

	value.Operator, err = mapAstNodeTo[code.UnaryOperator](original.Operator, a)
	if err != nil {
		return nil, err
	}

	value.OutputType, err = mapValueToPrimitiveType(value)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (a *astToCodeMapper) MapBinaryOperator(original ast.BinaryOperator) (code.Node, error) {
	return code.BinaryOperator(original), nil
}

func (a *astToCodeMapper) MapBinaryOperation(original ast.BinaryOperation) (code.Node, error) {
	value := new(code.BinaryOperation)
	a.stack.push(value)
	defer a.stack.pop()

	var err error
	value.Left, err = mapAstNodeTo[code.Value](original.Left, a)
	if err != nil {
		return nil, err
	}

	value.Right, err = mapAstNodeTo[code.Value](original.Right, a)
	if err != nil {
		return nil, err
	}

	value.Operator, err = mapAstNodeTo[code.BinaryOperator](original.Operator, a)
	if err != nil {
		return nil, err
	}

	value.OutputType, err = mapValueToPrimitiveType(value)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (a *astToCodeMapper) MapAssignment(original ast.Assignment) (code.Node, error) {
	value := new(code.Assignment)
	a.stack.push(value)
	defer a.stack.pop()

	var err error
	value.To, err = mapAstNodeTo[code.Assignable](original.To, a)
	if err != nil {
		return nil, err
	}

	value.From, err = mapAstNodeTo[code.Value](original.From, a)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (a *astToCodeMapper) MapModel(original ast.Model) (code.Node, error) {
	value := new(code.Model)
	a.stack.push(value)
	defer a.stack.pop()

	value.Name = original.Name

	module, ok := firstOfType[*code.Module](a.stack)
	if !ok {
		// TODO improve
		return nil, errors.New("no module")
	}

	for _, model := range module.Models {
		if model.Name == original.Name {
			value.Definition = model
		}
	}
	if value.Definition == nil {
		// TODO improve
		return nil, errors.New("no definition found")
	}

	return value, nil
}

func (a *astToCodeMapper) MapPrimitive(original ast.Primitive) (code.Node, error) {
	return code.Primitive(original), nil
}

func (a *astToCodeMapper) MapIf(original ast.If) (code.Node, error) {
	var value *code.If
	a.stack.push(value)
	defer a.stack.pop()

	var err error
	value.Condition, err = mapAstNodeTo[code.Value](original.Condition, a)
	if err != nil {
		return nil, err
	}

	for _, originalStatement := range original.Statements {
		statement, err := mapAstNodeTo[code.Statement](originalStatement, a)
		if err != nil {
			return nil, err
		}
		value.Statements = append(value.Statements, statement)
	}

	return value, nil
}

func (a *astToCodeMapper) MapElseIf(original ast.ElseIf) (code.Node, error) {
	value := new(code.ElseIf)
	a.stack.push(value)
	defer a.stack.pop()

	var err error
	value.Condition, err = mapAstNodeTo[code.Value](original.Condition, a)
	if err != nil {
		return nil, err
	}

	for _, originalStatement := range original.Statements {
		statement, err := mapAstNodeTo[code.Statement](originalStatement, a)
		if err != nil {
			return nil, err
		}
		value.Statements = append(value.Statements, statement)
	}

	return value, nil
}

func (a *astToCodeMapper) MapElse(original ast.Else) (code.Node, error) {
	value := new(code.Else)
	a.stack.push(value)
	defer a.stack.pop()

	for _, originalStatement := range original.Statements {
		statement, err := mapAstNodeTo[code.Statement](originalStatement, a)
		if err != nil {
			return nil, err
		}
		value.Statements = append(value.Statements, statement)
	}

	return value, nil
}

func (a *astToCodeMapper) MapConditional(original ast.Conditional) (code.Node, error) {
	value := new(code.Conditional)
	a.stack.push(value)
	defer a.stack.pop()

	var err error
	value.If, err = mapAstNodeTo[*code.If](original.If, a)
	if err != nil {
		return nil, err
	}

	for _, originalElseIf := range original.ElseIfs {
		elseIf, err := mapAstNodeTo[*code.ElseIf](originalElseIf, a)
		if err != nil {
			return nil, err
		}

		value.ElseIfs = append(value.ElseIfs, elseIf)
	}

	value.Else, err = mapAstNodeTo[*code.Else](original.Else, a)
	if err != nil {
		return nil, err
	}

	var ok bool
	value.Parent, ok = firstOfType[*code.MethodDef](a.stack)
	if !ok {
		// TODO improve
		return nil, errors.New("no model def found in the parent of a conditional")
	}

	return value, nil
}

func (a *astToCodeMapper) MapProperty(original ast.Property) (code.Node, error) {
	value := new(code.Property)
	a.stack.push(value)
	defer a.stack.pop()

	var err error
	value.Of, err = mapAstNodeTo[code.Value](original.Of, a)
	if err != nil {
		return nil, err
	}

	value.Name = original.Name

	value.Type, err = mapValueToType(value)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (a *astToCodeMapper) MapThis(_ ast.This) (code.Node, error) {
	value := new(code.This)
	a.stack.push(value)
	defer a.stack.pop()

	var ok bool
	value.This, ok = firstOfType[*code.ModelDef](a.stack)
	if !ok {
		// TODO improve
		return nil, errors.New("no model definition found in the parent of a this")
	}

	return value, nil
}

func mapAstNodeTo[T code.Node](original ast.Node, mapper *astToCodeMapper) (T, error) {
	mappedCode, err := ast.MapNode[code.Node](original, mapper)
	if err != nil {
		var zero T
		return zero, err
	}

	codeType, ok := mappedCode.(T)
	if !ok {
		var zero T
		return zero, fmt.Errorf("expected %T but got %T", zero, mappedCode)
	}

	return codeType, nil
}

func validateName(name string) error {
	if strings.Contains(name, " ") {
		return errors.New("space characters")
	}

	if strings.ToUpper(name)[0] != name[0] {
		return fmt.Errorf("first letter of variable name must be capitalized: \"%s\"", name)
	}

	return nil
}
