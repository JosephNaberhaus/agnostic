package mappers

import (
	"errors"
	"fmt"
	"github.com/JosephNaberhaus/agnostic/ast"
	"github.com/JosephNaberhaus/agnostic/code"
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

func (a *astToCodeMapper) MapLiteralInt(original ast.LiteralInt) (code.Node, error) {
	value := new(code.LiteralInt)
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

	err := validateVariableName(original.Name)
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

	err := validateVariableName(original.Name)
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

	var err error
	value.Function, err = mapAstNodeTo[*code.FunctionDef](original.Function, a)
	if err != nil {
		return nil, err
	}

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

	err := validateModelName(original.Name)
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

	value.MethodMap = map[string]*code.MethodDef{}
	for _, codeMethod := range value.Methods {
		if _, alreadyExists := value.MethodMap[codeMethod.Function.Name]; alreadyExists {
			// TODO improve
			return nil, errors.New("already exists")
		}

		value.MethodMap[codeMethod.Function.Name] = codeMethod
	}

	return value, nil
}

func (a *astToCodeMapper) MapModule(original ast.Module) (code.Node, error) {
	value := new(code.Module)
	a.stack.push(value)
	defer a.stack.pop()

	err := validateModuleName(original.Name)
	if err != nil {
		return nil, err
	}
	value.Name = original.Name

	for _, originalModel := range original.Models {
		model, err := mapAstNodeTo[*code.ModelDef](originalModel, a)
		if err != nil {
			return nil, err
		}
		value.Models = append(value.Models, model)
	}

	for _, originalFunction := range original.Functions {
		function, err := mapAstNodeTo[*code.FunctionDef](originalFunction, a)
		if err != nil {
			return nil, err
		}
		value.Functions = append(value.Functions, function)
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
	value.To, err = mapAstNodeTo[code.Value](original.To, a)
	if err != nil {
		return nil, err
	}

	// TODO validate that you can actually assign to `TO`

	value.From, err = mapAstNodeTo[code.Value](original.From, a)
	if err != nil {
		return nil, err
	}

	value.StatementMetadata, err = a.getStatementMetadata()
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
	value := new(code.If)
	a.stack.push(value)
	defer a.stack.pop()

	var err error
	value.Condition, err = mapAstNodeTo[code.Value](original.Condition, a)
	if err != nil {
		return nil, err
	}

	err = validateCondition(value.Condition)
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

	err = validateCondition(value.Condition)
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

	if len(original.Else.Statements) > 0 {
		value.Else, err = mapAstNodeTo[*code.Else](original.Else, a)
		if err != nil {
			return nil, err
		}
	}

	value.StatementMetadata, err = a.getStatementMetadata()
	if err != nil {
		return nil, err
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

func (a *astToCodeMapper) MapVariable(original ast.Variable) (code.Node, error) {
	value := new(code.Variable)
	a.stack.push(value)
	defer a.stack.pop()

	value.Name = original.Name

	var err error
	value.Definition, err = mapFindDefinition(original.Name, a.stack)
	if err != nil {
		return nil, err
	}

	value.Type, err = mapValueToType(value)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (a *astToCodeMapper) MapList(original ast.List) (code.Node, error) {
	value := new(code.List)
	a.stack.push(value)
	defer a.stack.pop()

	var err error
	value.Base, err = mapAstNodeTo[code.Type](original.Base, a)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (a *astToCodeMapper) MapMap(original ast.Map) (code.Node, error) {
	value := new(code.Map)
	a.stack.push(value)
	defer a.stack.pop()

	var err error
	value.Key, err = mapAstNodeTo[code.Type](original.Key, a)
	if err != nil {
		return nil, err
	}

	value.Value, err = mapAstNodeTo[code.Type](original.Value, a)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (a *astToCodeMapper) MapLookup(original ast.Lookup) (code.Node, error) {
	value := new(code.Lookup)
	a.stack.push(value)
	defer a.stack.pop()

	var err error
	value.From, err = mapAstNodeTo[code.Value](original.From, a)
	if err != nil {
		return nil, err
	}

	value.Key, err = mapAstNodeTo[code.Value](original.Key, a)
	if err != nil {
		return nil, err
	}

	value.OutputType, err = mapValueToType(value)
	if err != nil {
		return nil, err
	}

	fromType, err := mapValueToType(value.From)
	if err != nil {
		return nil, err
	}

	keyType, err := mapValueToType(value.Key)
	if err != nil {
		return nil, err
	}

	switch fromType := fromType.(type) {
	case *code.Map:
		value.LookupType = code.LookupTypeMap

		if fromType.Key != keyType {
			return nil, errors.New("mismatched key type for map")
		}
	case *code.List:
		value.LookupType = code.LookupTypeList

		if keyType != code.Int {
			return nil, errors.New("lists must be indexed by ints")
		}
	default:
		return nil, errors.New("invalid from type")
	}

	return value, nil
}

func (a *astToCodeMapper) MapFunctionDef(original ast.FunctionDef) (code.Node, error) {
	value := new(code.FunctionDef)
	a.stack.push(value)
	defer a.stack.pop()

	err := validateMethodName(original.Name)
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

	_, value.IsMethod = firstOfType[*code.MethodDef](a.stack)

	a.queueDefer(func() error {
		for _, originalStatement := range original.Statements {
			statement, err := mapAstNodeTo[code.Statement](originalStatement, a)
			if err != nil {
				return err
			}
			value.Statements = append(value.Statements, statement)
		}

		err = validateFunction(value)
		if err != nil {
			return err
		}

		return nil
	})

	return value, nil
}

func (a *astToCodeMapper) MapReturn(original ast.Return) (code.Node, error) {
	value := new(code.Return)
	a.stack.push(value)
	defer a.stack.pop()

	var err error
	value.Value, err = mapAstNodeTo[code.Value](original.Value, a)
	if err != nil {
		return nil, err
	}

	value.StatementMetadata, err = a.getStatementMetadata()
	if err != nil {
		return nil, err
	}

	err = validateReturn(value)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (a *astToCodeMapper) MapCall(original ast.Call) (code.Node, error) {
	value := new(code.Call)
	a.stack.push(value)
	defer a.stack.pop()

	var err error
	value.Function, err = mapAstNodeTo[code.Callable](original.Function, a)
	if err != nil {
		return nil, err
	}

	value.Arguments, err = mapAstNodesTo[code.Value](original.Arguments, a)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (a *astToCodeMapper) MapFunction(original ast.Function) (code.Node, error) {
	value := new(code.Function)
	a.stack.push(value)
	defer a.stack.pop()

	value.Name = original.Name

	return value, nil
}

func (a *astToCodeMapper) MapFunctionProperty(original ast.FunctionProperty) (code.Node, error) {
	value := new(code.FunctionProperty)
	a.stack.push(value)
	defer a.stack.pop()

	var err error
	value.Of, err = mapAstNodeTo[code.Value](original.Of, a)
	if err != nil {
		return nil, err
	}

	value.Name = original.Name

	return value, nil
}

func (a *astToCodeMapper) MapNew(original ast.New) (code.Node, error) {
	value := new(code.New)
	a.stack.push(value)
	defer a.stack.pop()

	var err error
	value.Model, err = mapAstNodeTo[*code.Model](original.Model, a)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (a *astToCodeMapper) MapDeclare(original ast.Declare) (code.Node, error) {
	value := new(code.Declare)
	a.stack.push(value)
	defer a.stack.pop()

	value.Name = original.Name

	var err error
	value.Value, err = mapAstNodeTo[code.Value](original.Value, a)
	if err != nil {
		return nil, err
	}

	value.StatementMetadata, err = a.getStatementMetadata()
	if err != nil {
		return nil, err
	}

	value.Type, err = mapValueToType(value.Value)
	if err != nil {
		return nil, err
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

func mapAstNodesTo[T code.Node, V ast.Node](originalNodes []V, mapper *astToCodeMapper) ([]T, error) {
	var mappedCodeNodes []T
	for _, original := range originalNodes {
		mappedCode, err := mapAstNodeTo[T](original, mapper)
		if err != nil {
			return nil, err
		}
		mappedCodeNodes = append(mappedCodeNodes, mappedCode)
	}

	return mappedCodeNodes, nil
}

func (a *astToCodeMapper) getStatementMetadata() (code.StatementMetadata, error) {
	var value code.StatementMetadata

	var ok bool
	value.Parent, ok = firstOfType[*code.FunctionDef](a.stack)
	if !ok {
		return code.StatementMetadata{}, errors.New("no function def found as parent of statement")
	}

	return value, nil
}
