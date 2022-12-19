package mappers

import (
	"errors"
	"fmt"
	"github.com/JosephNaberhaus/agnostic/ast"
	"github.com/JosephNaberhaus/agnostic/code"
	"github.com/JosephNaberhaus/agnostic/code/mappers/callable_to_function_definition"
	"github.com/JosephNaberhaus/agnostic/code/mappers/find_function_definition"
	"github.com/JosephNaberhaus/agnostic/code/mappers/value_to_type"
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
	stackSnapshot code.Stack
	deferredFunc  func() error
}

type astToCodeMapper struct {
	stack    code.Stack
	deferred []deferredCall
}

var _ ast.NodeMapper[code.Node] = (*astToCodeMapper)(nil)

// queueDefer queues a function that will be called after all other mapping is done.
//
// This allows the mapping to act like a breadth first search at defined points. This is used when we need to know that
// the code models will be populated with all information from the AST (up to a certain depth).
func (a *astToCodeMapper) queueDefer(deferredFunc func() error) {
	a.deferred = append(a.deferred, deferredCall{
		stackSnapshot: a.stack.Copy(),
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
	a.stack.Push(value)
	defer a.stack.Pop()

	value.Value = original.Value

	return value, nil
}

func (a *astToCodeMapper) MapLiteralString(original ast.LiteralString) (code.Node, error) {
	value := new(code.LiteralString)
	a.stack.Push(value)
	defer a.stack.Pop()

	value.Value = original.Value

	return value, nil
}

func (a *astToCodeMapper) MapLiteralRune(original ast.LiteralRune) (code.Node, error) {
	value := new(code.LiteralRune)
	a.stack.Push(value)
	defer a.stack.Pop()

	value.Value = original.Value

	return value, nil
}

func (a *astToCodeMapper) MapFieldDef(original ast.FieldDef) (code.Node, error) {
	value := new(code.FieldDef)
	a.stack.Push(value)
	defer a.stack.Pop()

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

	value.ParentModel, ok = code.FirstOfType[*code.ModelDef](a.stack)
	if !ok {
		// TODO improve
		return nil, errors.New("no model def found in the parent of a field")
	}

	return value, nil
}

func (a *astToCodeMapper) MapArgumentDef(original ast.ArgumentDef) (code.Node, error) {
	value := new(code.ArgumentDef)
	a.stack.Push(value)
	defer a.stack.Pop()

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

func (a *astToCodeMapper) MapModelDef(original ast.ModelDef) (code.Node, error) {
	value := new(code.ModelDef)
	a.stack.Push(value)
	defer a.stack.Pop()

	err := validateModelName(original.Name)
	if err != nil {
		return nil, err
	}
	value.Name = original.Name

	a.queueDefer(func() error {
		for _, originalField := range original.Fields {
			field, err := mapAstNodeTo[*code.FieldDef](originalField, a)
			if err != nil {
				return err
			}
			value.Fields = append(value.Fields, field)
		}

		for _, originalMethod := range original.Methods {
			method, err := mapAstNodeTo[*code.FunctionDef](originalMethod, a)
			if err != nil {
				return err
			}
			value.Methods = append(value.Methods, method)
		}

		value.FieldMap = map[string]*code.FieldDef{}
		for _, codeField := range value.Fields {
			if _, alreadyExists := value.FieldMap[codeField.Name]; alreadyExists {
				// TODO improve
				return errors.New("already exists")
			}

			value.FieldMap[codeField.Name] = codeField
		}

		value.MethodMap = map[string]*code.FunctionDef{}
		for _, codeMethod := range value.Methods {
			if _, alreadyExists := value.MethodMap[codeMethod.Name]; alreadyExists {
				// TODO improve
				return errors.New("already exists")
			}

			value.MethodMap[codeMethod.Name] = codeMethod
		}

		if original.EqualOverride.IsSet() {
			equalOverride, err := mapAstNodeTo[*code.EqualOverride](original.EqualOverride.Value(), a)
			if err != nil {
				return err
			}
			value.EqualOverride.Set(equalOverride)
		}

		if original.HashOverride.IsSet() {
			hashOverride, err := mapAstNodeTo[*code.HashOverride](original.HashOverride.Value(), a)
			if err != nil {
				return err
			}
			value.HashOverride.Set(hashOverride)
		}

		return nil
	})

	return value, nil
}

func (a *astToCodeMapper) MapModule(original ast.Module) (code.Node, error) {
	value := new(code.Module)
	a.stack.Push(value)
	defer a.stack.Pop()

	err := validateModuleName(original.Name)
	if err != nil {
		return nil, err
	}
	value.Name = original.Name

	value.ConstantsMap = map[string]*code.ConstantDef{}
	for _, originalConstant := range original.Constants {
		constant, err := mapAstNodeTo[*code.ConstantDef](originalConstant, a)
		if err != nil {
			return nil, err
		}
		value.Constants = append(value.Constants, constant)
		value.ConstantsMap[constant.Name] = constant
	}

	value.ModelMap = map[string]*code.ModelDef{}
	for _, originalModel := range original.Models {
		model, err := mapAstNodeTo[*code.ModelDef](originalModel, a)
		if err != nil {
			return nil, err
		}
		value.Models = append(value.Models, model)
		value.ModelMap[model.Name] = model
	}

	value.FunctionMap = map[string]*code.FunctionDef{}
	for _, originalFunction := range original.Functions {
		function, err := mapAstNodeTo[*code.FunctionDef](originalFunction, a)
		if err != nil {
			return nil, err
		}
		value.Functions = append(value.Functions, function)
		value.FunctionMap[function.Name] = function
	}

	return value, nil
}

func (a *astToCodeMapper) MapUnaryOperator(original ast.UnaryOperator) (code.Node, error) {
	return code.UnaryOperator(original), nil
}

func (a *astToCodeMapper) MapUnaryOperation(original ast.UnaryOperation) (code.Node, error) {
	value := new(code.UnaryOperation)
	a.stack.Push(value)
	defer a.stack.Pop()

	var err error
	value.Value, err = mapAstNodeTo[code.Value](original.Value, a)
	if err != nil {
		return nil, err
	}

	value.Operator, err = mapAstNodeTo[code.UnaryOperator](original.Operator, a)
	if err != nil {
		return nil, err
	}

	value.OutputType, err = value_to_type.MapValueToPrimitiveType(value)
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
	a.stack.Push(value)
	defer a.stack.Pop()

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

	value.OutputType, err = value_to_type.MapValueToPrimitiveType(value)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (a *astToCodeMapper) MapAssignment(original ast.Assignment) (code.Node, error) {
	value := new(code.Assignment)
	a.stack.Push(value)
	defer a.stack.Pop()

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

	return value, nil
}

func (a *astToCodeMapper) MapModel(original ast.Model) (code.Node, error) {
	value := new(code.Model)
	a.stack.Push(value)
	defer a.stack.Pop()

	value.Name = original.Name

	module, ok := code.FirstOfType[*code.Module](a.stack)
	if !ok {
		// TODO improve
		return nil, errors.New("no module")
	}

	value.Definition, ok = module.ModelMap[value.Name]
	if !ok {
		return nil, errors.New("no definition found for " + value.Name)
	}

	return value, nil
}

func (a *astToCodeMapper) MapPrimitive(original ast.Primitive) (code.Node, error) {
	return code.Primitive(original), nil
}

func (a *astToCodeMapper) MapIf(original ast.If) (code.Node, error) {
	value := new(code.If)
	a.stack.Push(value)
	defer a.stack.Pop()

	var err error
	value.Condition, err = mapAstNodeTo[code.Value](original.Condition, a)
	if err != nil {
		return nil, err
	}

	err = validateCondition(value.Condition)
	if err != nil {
		return nil, err
	}

	value.Block, err = mapAstNodeTo[*code.Block](original.Block, a)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (a *astToCodeMapper) MapElseIf(original ast.ElseIf) (code.Node, error) {
	value := new(code.ElseIf)
	a.stack.Push(value)
	defer a.stack.Pop()

	var err error
	value.Condition, err = mapAstNodeTo[code.Value](original.Condition, a)
	if err != nil {
		return nil, err
	}

	err = validateCondition(value.Condition)
	if err != nil {
		return nil, err
	}

	value.Block, err = mapAstNodeTo[*code.Block](original.Block, a)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (a *astToCodeMapper) MapElse(original ast.Else) (code.Node, error) {
	value := new(code.Else)
	a.stack.Push(value)
	defer a.stack.Pop()

	var err error
	value.Block, err = mapAstNodeTo[*code.Block](original.Block, a)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (a *astToCodeMapper) MapConditional(original ast.Conditional) (code.Node, error) {
	value := new(code.Conditional)
	a.stack.Push(value)
	defer a.stack.Pop()

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

	if len(original.Else.Block.Statements) > 0 {
		value.Else, err = mapAstNodeTo[*code.Else](original.Else, a)
		if err != nil {
			return nil, err
		}
	}

	return value, nil
}

func (a *astToCodeMapper) MapProperty(original ast.Property) (code.Node, error) {
	value := new(code.Property)
	a.stack.Push(value)
	defer a.stack.Pop()

	var err error
	value.Of, err = mapAstNodeTo[code.Value](original.Of, a)
	if err != nil {
		return nil, err
	}

	value.Name = original.Name

	value.Type, err = code.MapValue[code.Type](value, value_to_type.Mapper{})
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (a *astToCodeMapper) MapVariable(original ast.Variable) (code.Node, error) {
	value := new(code.Variable)
	a.stack.Push(value)
	defer a.stack.Pop()

	value.Name = original.Name

	var err error
	value.Definition, err = mapFindDefinition(original.Name, a.stack)
	if err != nil {
		return nil, err
	}

	_, value.IsConstant = value.Definition.(*code.ConstantDef)

	value.Type, err = code.MapValue[code.Type](value, value_to_type.Mapper{})
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (a *astToCodeMapper) MapList(original ast.List) (code.Node, error) {
	value := new(code.List)
	a.stack.Push(value)
	defer a.stack.Pop()

	var err error
	value.Base, err = mapAstNodeTo[code.Type](original.Base, a)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (a *astToCodeMapper) MapMap(original ast.Map) (code.Node, error) {
	value := new(code.Map)
	a.stack.Push(value)
	defer a.stack.Pop()

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
	a.stack.Push(value)
	defer a.stack.Pop()

	var err error
	value.From, err = mapAstNodeTo[code.Value](original.From, a)
	if err != nil {
		return nil, err
	}

	value.Key, err = mapAstNodeTo[code.Value](original.Key, a)
	if err != nil {
		return nil, err
	}

	fromType, err := code.MapValue[code.Type](value.From, value_to_type.Mapper{})
	if err != nil {
		return nil, err
	}

	keyType, err := code.MapValue[code.Type](value.Key, value_to_type.Mapper{})
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
	case code.Primitive:
		// TODO be better
		if fromType != code.String {
			panic("be better")
		}

		value.LookupType = code.LookupTypeString
	default:
		return nil, errors.New("invalid from type")
	}

	value.OutputType, err = code.MapValue[code.Type](value, value_to_type.Mapper{})
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (a *astToCodeMapper) MapFunctionDef(original ast.FunctionDef) (code.Node, error) {
	value := new(code.FunctionDef)
	a.stack.Push(value)
	defer a.stack.Pop()

	err := validateMethodName(original.Name)
	if err != nil {
		return nil, err
	}
	value.Name = original.Name

	a.queueDefer(func() error {
		for _, originalArguments := range original.Arguments {
			argument, err := mapAstNodeTo[*code.ArgumentDef](originalArguments, a)
			if err != nil {
				return err
			}
			value.Arguments = append(value.Arguments, argument)
		}

		value.ReturnType, err = mapAstNodeTo[code.Type](original.ReturnType, a)
		if err != nil {
			return err
		}

		a.queueDefer(func() error {
			value.Block, err = mapAstNodeTo[*code.Block](original.Block, a)
			if err != nil {
				return err
			}

			err = validateFunction(value)
			if err != nil {
				return err
			}

			return nil
		})

		return nil
	})

	_, value.IsMethod = code.FirstOfType[*code.ModelDef](a.stack)

	return value, nil
}

func (a *astToCodeMapper) MapReturn(original ast.Return) (code.Node, error) {
	value := new(code.Return)
	a.stack.Push(value)
	defer a.stack.Pop()

	var err error
	value.Value, err = mapAstNodeTo[code.Value](original.Value, a)
	if err != nil {
		return nil, err
	}

	var ok bool
	value.CallableDef, ok = code.FirstOfType[code.CallableDef](a.stack)
	if !ok {
		return nil, errors.New("no function def found as parent of statement")
	}

	err = validateReturn(value)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (a *astToCodeMapper) MapCall(original ast.Call) (code.Node, error) {
	value := new(code.Call)
	a.stack.Push(value)
	defer a.stack.Pop()

	// Check whether the call is to one of the built-in functions.
	if functionProperty, ok := original.Function.(ast.FunctionProperty); ok {
		functionPropertyOf, err := mapAstNodeTo[code.Value](functionProperty.Of, a)
		if err != nil {
			return nil, err
		}

		functionPropertyOfType, err := code.MapValue[code.Type](functionPropertyOf, value_to_type.Mapper{})
		if err != nil {
			return nil, err
		}

		if _, isSet := functionPropertyOfType.(*code.Set); isSet {
			switch functionProperty.Name {
			case "add":
				if len(original.Arguments) != 1 {
					return nil, errors.New("set add requires one argument")
				}

				return a.MapAddToSet(ast.AddToSet{
					To:    functionProperty.Of,
					Value: original.Arguments[0],
				})
			case "contains":
				if len(original.Arguments) != 1 {
					return nil, errors.New("set contains requires one argument")
				}

				return a.MapSetContains(ast.SetContains{
					Set:   functionProperty.Of,
					Value: original.Arguments[0],
				})
			}
		} else if _, isList := functionPropertyOfType.(*code.List); isList {
			switch functionProperty.Name {
			case "push":
				if len(original.Arguments) != 1 {
					return nil, errors.New("list push requires one argument")
				}

				return a.MapPush(ast.Push{
					To:    functionProperty.Of,
					Value: original.Arguments[0],
				})
			case "pop":
				if len(original.Arguments) != 0 {
					return nil, errors.New("list pop requires no arguments")
				}

				return a.MapPop(ast.Pop{
					Value: functionProperty.Of,
				})
			}
		}
	}

	var err error
	value.Function, err = mapAstNodeTo[code.Callable](original.Function, a)
	if err != nil {
		return nil, err
	}

	value.Arguments, err = mapAstNodesTo[code.Value](original.Arguments, a)
	if err != nil {
		return nil, err
	}

	value.Definition, err = code.MapCallable[*code.FunctionDef](value.Function, callable_to_function_definition.Mapper{})
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (a *astToCodeMapper) MapFunction(original ast.Function) (code.Node, error) {
	value := new(code.Function)
	a.stack.Push(value)
	defer a.stack.Pop()

	value.Name = original.Name

	var err error
	value.Definition, err = find_function_definition.InStack(value.Name, a.stack)
	if err != nil {
		return nil, err
	}

	var ok bool
	value.ParentModule, ok = code.FirstOfType[*code.Module](a.stack)
	if !ok {
		return nil, errors.New("expected module parent")
	}

	return value, nil
}

func (a *astToCodeMapper) MapFunctionProperty(original ast.FunctionProperty) (code.Node, error) {
	value := new(code.FunctionProperty)
	a.stack.Push(value)
	defer a.stack.Pop()

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
	a.stack.Push(value)
	defer a.stack.Pop()

	var err error
	value.Model, err = mapAstNodeTo[*code.Model](original.Model, a)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (a *astToCodeMapper) MapDeclare(original ast.Declare) (code.Node, error) {
	value := new(code.Declare)
	a.stack.Push(value)
	defer a.stack.Pop()

	value.Name = original.Name

	var err error
	value.Value, err = mapAstNodeTo[code.Value](original.Value, a)
	if err != nil {
		return nil, err
	}

	value.Type, err = code.MapValue[code.Type](value.Value, value_to_type.Mapper{})
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (a *astToCodeMapper) MapFor(original ast.For) (code.Node, error) {
	value := new(code.For)
	a.stack.Push(value)
	defer a.stack.Pop()

	if original.Initialization.IsSet() {
		initialization, err := mapAstNodeTo[code.Statement](original.Initialization.Value(), a)
		if err != nil {
			return nil, err
		}
		value.Initialization.Set(initialization)
	}

	var err error
	value.Condition, err = mapAstNodeTo[code.Value](original.Condition, a)
	if err != nil {
		return nil, err
	}

	if original.AfterEach.IsSet() {
		afterEach, err := mapAstNodeTo[code.Statement](original.AfterEach.Value(), a)
		if err != nil {
			return nil, err
		}
		value.AfterEach.Set(afterEach)
	}

	value.Block, err = mapAstNodeTo[*code.Block](original.Block, a)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (a *astToCodeMapper) MapForIn(original ast.ForIn) (code.Node, error) {
	value := new(code.ForIn)
	a.stack.Push(value)
	defer a.stack.Pop()

	var err error
	value.Iterable, err = mapAstNodeTo[code.Value](original.Iterable, a)
	if err != nil {
		return nil, err
	}

	value.ItemName = original.ItemName

	iterableType, err := code.MapValue[code.Type](value.Iterable, value_to_type.Mapper{})
	if err != nil {
		return nil, err
	}

	switch iterableType := iterableType.(type) {
	case *code.List:
		value.ForInType = code.ForInTypeList
		value.ItemType = iterableType.Base
	case *code.Set:
		value.ForInType = code.ForInTypeSet
		value.ItemType = iterableType.Base
	default:
		return nil, errors.New("for-in can only be used with a list or set")
	}

	value.Block, err = mapAstNodeTo[*code.Block](original.Block, a)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (a *astToCodeMapper) MapBlock(original ast.Block) (code.Node, error) {
	value := new(code.Block)
	a.stack.Push(value)
	defer a.stack.Pop()

	for _, originalStatement := range original.Statements {
		statement, err := mapAstNodeTo[code.Statement](originalStatement, a)
		if err != nil {
			return nil, err
		}
		value.Statements = append(value.Statements, statement)
	}

	return value, nil
}

func (a *astToCodeMapper) MapLiteralList(original ast.LiteralList) (code.Node, error) {
	value := new(code.LiteralList)
	a.stack.Push(value)
	defer a.stack.Pop()

	var err error
	value.Items, err = mapAstNodesTo[code.Value](original.Items, a)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (a *astToCodeMapper) MapLength(original ast.Length) (code.Node, error) {
	value := new(code.Length)
	a.stack.Push(value)
	defer a.stack.Pop()

	var err error
	value.Value, err = mapAstNodeTo[code.Value](original.Value, a)
	if err != nil {
		return nil, err
	}

	valueType, err := code.MapValue[code.Type](value.Value, value_to_type.Mapper{})
	if err != nil {
		return nil, err
	}

	// TODO Unify errors
	switch valueType := valueType.(type) {
	case code.Primitive:
		if valueType != code.String {
			return nil, errors.New("unexpected type for length")
		}

		value.LengthType = code.LengthTypeString
	case *code.List:
		value.LengthType = code.LengthTypeList
	case *code.Map:
		value.LengthType = code.LengthTypeMap
	case *code.Set:
		value.LengthType = code.LengthTypeSet
	default:
		return nil, errors.New("unexpected type for length")
	}

	return value, nil
}

func (a *astToCodeMapper) MapConstantDef(original ast.ConstantDef) (code.Node, error) {
	value := new(code.ConstantDef)
	a.stack.Push(value)
	defer a.stack.Pop()

	value.Name = original.Name

	var ok bool
	value.ParentModule, ok = code.FirstOfType[*code.Module](a.stack)
	if !ok {
		return nil, errors.New("expected module to be parent of constant")
	}

	a.queueDefer(func() error {
		var err error
		value.Value, err = mapAstNodeTo[code.ConstantValue](original.Value, a)
		if err != nil {
			return err
		}

		value.Type, err = code.MapConstantValue[code.Type](value.Value, value_to_type.Mapper{})
		if err != nil {
			return err
		}

		return nil
	})

	return value, nil
}

func (a *astToCodeMapper) MapKeyValue(original ast.KeyValue) (code.Node, error) {
	value := new(code.KeyValue)
	a.stack.Push(value)
	defer a.stack.Pop()

	var err error
	value.Key, err = mapAstNodeTo[code.Value](original.Key, a)
	if err != nil {
		return nil, err
	}

	value.Value, err = mapAstNodeTo[code.Value](original.Value, a)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (a *astToCodeMapper) MapLiteralMap(original ast.LiteralMap) (code.Node, error) {
	value := new(code.LiteralMap)
	a.stack.Push(value)
	defer a.stack.Pop()

	var err error
	value.Entries, err = mapAstNodesTo[*code.KeyValue](original.Entries, a)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (a *astToCodeMapper) MapLiteralSet(original ast.LiteralSet) (code.Node, error) {
	value := new(code.LiteralSet)
	a.stack.Push(value)
	defer a.stack.Pop()

	var err error
	value.Items, err = mapAstNodesTo[code.Value](original.Items, a)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (a *astToCodeMapper) MapSet(original ast.Set) (code.Node, error) {
	value := new(code.Set)
	a.stack.Push(value)
	defer a.stack.Pop()

	var err error
	value.Base, err = mapAstNodeTo[code.Type](original.Base, a)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (a *astToCodeMapper) MapEmptyList(original ast.EmptyList) (code.Node, error) {
	value := new(code.EmptyList)
	a.stack.Push(value)
	defer a.stack.Pop()

	var err error
	value.Type, err = mapAstNodeTo[code.Type](original.Type, a)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (a *astToCodeMapper) MapEmptySet(original ast.EmptySet) (code.Node, error) {
	value := new(code.EmptySet)
	a.stack.Push(value)
	defer a.stack.Pop()

	var err error
	value.Type, err = mapAstNodeTo[code.Type](original.Type, a)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (a *astToCodeMapper) MapAddToSet(original ast.AddToSet) (code.Node, error) {
	value := new(code.AddToSet)
	a.stack.Push(value)
	defer a.stack.Pop()

	var err error
	value.To, err = mapAstNodeTo[code.Value](original.To, a)
	if err != nil {
		return nil, err
	}

	value.Value, err = mapAstNodeTo[code.Value](original.Value, a)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (a *astToCodeMapper) MapSetContains(original ast.SetContains) (code.Node, error) {
	value := new(code.SetContains)
	a.stack.Push(value)
	defer a.stack.Pop()

	var err error
	value.Set, err = mapAstNodeTo[code.Value](original.Set, a)
	if err != nil {
		return nil, err
	}

	value.Value, err = mapAstNodeTo[code.Value](original.Value, a)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (a *astToCodeMapper) MapPush(original ast.Push) (code.Node, error) {
	value := new(code.Push)
	a.stack.Push(value)
	defer a.stack.Pop()

	var err error
	value.To, err = mapAstNodeTo[code.Value](original.To, a)
	if err != nil {
		return nil, err
	}

	value.Value, err = mapAstNodeTo[code.Value](original.Value, a)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (a *astToCodeMapper) MapPop(original ast.Pop) (code.Node, error) {
	value := new(code.Pop)
	a.stack.Push(value)
	defer a.stack.Pop()

	var err error
	value.Value, err = mapAstNodeTo[code.Value](original.Value, a)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (a *astToCodeMapper) MapLiteralBool(original ast.LiteralBool) (code.Node, error) {
	value := new(code.LiteralBool)
	a.stack.Push(value)
	defer a.stack.Pop()

	value.Value = original.Value

	return value, nil
}

func (a *astToCodeMapper) MapNull(_ ast.Null) (code.Node, error) {
	value := new(code.Null)
	a.stack.Push(value)
	defer a.stack.Pop()

	value.Parent = a.stack.PeekParent()

	return value, nil
}

func (a *astToCodeMapper) MapSelf(_ ast.Self) (code.Node, error) {
	value := new(code.Self)
	a.stack.Push(value)
	defer a.stack.Pop()

	var ok bool
	value.ParentModel, ok = code.FirstOfType[*code.ModelDef](a.stack)
	if !ok {
		return nil, errors.New("self can only be used as a child of a model")
	}

	return value, nil
}

func (a *astToCodeMapper) MapDeclareNull(original ast.DeclareNull) (code.Node, error) {
	value := new(code.DeclareNull)
	a.stack.Push(value)
	defer a.stack.Pop()

	value.Name = original.Name

	var err error
	value.Type, err = mapAstNodeTo[code.Type](original.Type, a)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (a *astToCodeMapper) MapBreak(_ ast.Break) (code.Node, error) {
	value := new(code.Break)
	a.stack.Push(value)
	defer a.stack.Pop()

	return value, nil
}

func (a *astToCodeMapper) MapContinue(_ ast.Continue) (code.Node, error) {
	value := new(code.Continue)
	a.stack.Push(value)
	defer a.stack.Pop()

	return value, nil
}

func (a *astToCodeMapper) MapEqualOverride(original ast.EqualOverride) (code.Node, error) {
	value := new(code.EqualOverride)
	a.stack.Push(value)
	defer a.stack.Pop()

	value.OtherName = original.OtherName

	var ok bool
	value.ParentModel, ok = code.FirstOfType[*code.ModelDef](a.stack)
	if !ok {
		return nil, errors.New("parent of equal override must be a model")
	}

	var err error
	value.Block, err = mapAstNodeTo[*code.Block](original.Block, a)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (a *astToCodeMapper) MapHashOverride(original ast.HashOverride) (code.Node, error) {
	value := new(code.HashOverride)
	a.stack.Push(value)
	defer a.stack.Pop()

	var err error
	value.Block, err = mapAstNodeTo[*code.Block](original.Block, a)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (a *astToCodeMapper) MapLiteralProperty(original ast.LiteralProperty) (code.Node, error) {
	value := new(code.LiteralProperty)
	a.stack.Push(value)
	defer a.stack.Pop()

	value.Name = original.Name

	var err error
	value.Value, err = mapAstNodeTo[code.Value](original.Value, a)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (a *astToCodeMapper) MapLiteralStruct(original ast.LiteralStruct) (code.Node, error) {
	value := new(code.LiteralStruct)
	a.stack.Push(value)
	defer a.stack.Pop()

	value.Name = original.Name

	var err error
	value.Properties, err = mapAstNodesTo[*code.LiteralProperty](original.Properties, a)
	if err != nil {
		return nil, err
	}

	module, ok := code.FirstOfType[*code.Module](a.stack)
	if !ok {
		// TODO improve
		return nil, errors.New("no module")
	}

	value.Definition, ok = module.ModelMap[value.Name]
	if !ok {
		return nil, errors.New("no definition found for " + value.Name)
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
