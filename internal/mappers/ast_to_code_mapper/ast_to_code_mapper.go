package ast_to_code_mapper

import (
	"github.com/JosephNaberhaus/agnostic/ast"
	"github.com/JosephNaberhaus/agnostic/code"
	"github.com/JosephNaberhaus/agnostic/internal/utils/stack"
)

type Mapper struct {
	// The current stack. This is the path from the root node of the AST to the node currently being processed.
	stack stack.Stack[code.Node]
	// List of diferred calls. These will be hanlded once the rest of the AST tree is exhausted.
	deferred []deferred
}

// MapRoot is the only valid entry-point into this Mapper.
func (m *Mapper) MapRoot(original ast.Root) (code.Node, error) {
	value := &code.Root{}
	m.stack.Push(value)
	defer m.stack.Pop()

	var err error
	value.Modules, err = mapAstNodesTo[*code.Module](original.Modules, m)
	if err != nil {
		return nil, err
	}

	// Handle any deferred functions.
	curStack := m.stack
	for len(m.deferred) > 0 {
		deferred := m.dequeueDeferred()

		m.stack = deferred.stack
		err = deferred.deferredFunc()
		if err != nil {
			return nil, err
		}
	}
	m.stack = curStack

	return value, nil
}

func (m *Mapper) MapAddToSet(original ast.AddToSet) (code.Node, error) {
	value := &code.AddToSet{}
	m.stack.Push(value)
	defer m.stack.Pop()

	var err error
	value.Set, err = mapAstNodeTo[code.Value](original.Set, m)
	if err != nil {
		return nil, err
	}

	value.Value, err = mapAstNodeTo[code.Value](original.Value, m)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (m *Mapper) MapArgumentDef(original ast.ArgumentDef) (code.Node, error) {
	value := &code.ArgumentDef{}
	m.stack.Push(value)
	defer m.stack.Pop()

	value.Name = original.Name

	var err error
	value.Type, err = mapAstNodeTo[code.Type](original.Type, m)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (m *Mapper) MapAssignment(original ast.Assignment) (code.Node, error) {
	value := &code.Assignment{}
	m.stack.Push(value)
	defer m.stack.Pop()

	var err error
	value.From, err = mapAstNodeTo[code.Value](original.From, m)
	if err != nil {
		return nil, err
	}

	value.To, err = mapAstNodeTo[code.Value](original.To, m)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (m *Mapper) MapBlock(original ast.Block) (code.Node, error) {
	value := &code.Block{}
	m.stack.Push(value)
	defer m.stack.Pop()

	var err error
	value.Statements, err = mapAstNodesTo[code.Statement](original.Statements, m)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (m *Mapper) MapBool(original ast.Bool) (code.Node, error) {
	value := &code.Bool{}
	m.stack.Push(value)
	defer m.stack.Pop()

	return value, nil
}

func (m *Mapper) MapBreak(original ast.Break) (code.Node, error) {
	value := &code.Break{}
	m.stack.Push(value)
	defer m.stack.Pop()

	return value, nil
}

func (m *Mapper) MapCall(original ast.Call) (code.Node, error) {
	value := &code.Call{}
	m.stack.Push(value)
	defer m.stack.Pop()

	var err error
	value.Arguments, err = mapAstNodesTo[code.Value](original.Arguments, m)
	if err != nil {
		return nil, err
	}

	value.Function, err = mapAstNodeTo[code.Callable](original.Function, m)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (m *Mapper) MapConditional(original ast.Conditional) (code.Node, error) {
	value := &code.Conditional{}
	m.stack.Push(value)
	defer m.stack.Pop()

	var err error
	if original.Else.IsSet() {
		value.Else, err = mapAstNodeTo[*code.Block](original.Else.Value(), m)
		if err != nil {
			return nil, err
		}
	}

	value.Ifs, err = mapAstNodesTo[*code.If](original.Ifs, m)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (m *Mapper) MapConstantDef(original ast.ConstantDef) (code.Node, error) {
	value := &code.ConstantDef{}
	m.stack.Push(value)
	defer m.stack.Pop()

	value.Name = original.Name

	var err error
	value.Value, err = mapAstNodeTo[code.ConstantValue](original.Value, m)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (m *Mapper) MapContinue(original ast.Continue) (code.Node, error) {
	value := &code.Continue{}
	m.stack.Push(value)
	defer m.stack.Pop()

	return value, nil
}

func (m *Mapper) MapDeclare(original ast.Declare) (code.Node, error) {
	value := &code.Declare{}
	m.stack.Push(value)
	defer m.stack.Pop()

	value.Name = original.Name

	var err error
	value.Value, err = mapAstNodeTo[code.Value](original.Value, m)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (m *Mapper) MapEmptyList(original ast.EmptyList) (code.Node, error) {
	value := &code.EmptyList{}
	m.stack.Push(value)
	defer m.stack.Pop()

	var err error
	value.Type, err = mapAstNodeTo[code.Type](original.Type, m)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (m *Mapper) MapEqualOverride(original ast.EqualOverride) (code.Node, error) {
	value := &code.EqualOverride{}
	m.stack.Push(value)
	defer m.stack.Pop()

	var err error
	value.Block, err = mapAstNodeTo[*code.Block](original.Block, m)
	if err != nil {
		return nil, err
	}

	value.OtherName = original.OtherName

	return value, nil
}

func (m *Mapper) MapFieldDef(original ast.FieldDef) (code.Node, error) {
	value := &code.FieldDef{}
	m.stack.Push(value)
	defer m.stack.Pop()

	value.Name = original.Name

	var err error
	value.Type, err = mapAstNodeTo[code.Type](original.Type, m)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (m *Mapper) MapFor(original ast.For) (code.Node, error) {
	value := &code.For{}
	m.stack.Push(value)
	defer m.stack.Pop()

	var err error
	if original.AfterEach.IsSet() {
		value.AfterEach, err = mapAstNodeTo[code.Statement](original.AfterEach.Value(), m)
		if err != nil {
			return nil, err
		}
	}

	value.Block, err = mapAstNodeTo[*code.Block](original.Block, m)
	if err != nil {
		return nil, err
	}

	value.Condition, err = mapAstNodeTo[code.Value](original.Condition, m)
	if err != nil {
		return nil, err
	}

	if original.Initialization.IsSet() {
		value.Initialization, err = mapAstNodeTo[code.Statement](original.Initialization.Value(), m)
		if err != nil {
			return nil, err
		}
	}

	return value, nil
}

func (m *Mapper) MapForEach(original ast.ForEach) (code.Node, error) {
	value := &code.ForEach{}
	m.stack.Push(value)
	defer m.stack.Pop()

	var err error
	value.Block, err = mapAstNodeTo[*code.Block](original.Block, m)
	if err != nil {
		return nil, err
	}

	value.ItemName = original.ItemName

	value.Iterable, err = mapAstNodeTo[code.Value](original.Iterable, m)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (m *Mapper) MapFunctionDef(original ast.FunctionDef) (code.Node, error) {
	value := &code.FunctionDef{}
	m.stack.Push(value)
	defer m.stack.Pop()

	var err error
	value.Arguments, err = mapAstNodesTo[*code.ArgumentDef](original.Arguments, m)
	if err != nil {
		return nil, err
	}

	value.Name = original.Name

	value.ReturnType, err = mapAstNodeTo[code.Type](original.ReturnType, m)
	if err != nil {
		return nil, err
	}

	// Defer because code inside this function might call a function that hasn't been processed yet.
	m.queueDeferred(func() error {
		value.Block, err = mapAstNodeTo[*code.Block](original.Block, m)
		if err != nil {
			return err
		}

		return nil
	})

	return value, nil
}

func (m *Mapper) MapHashOverride(original ast.HashOverride) (code.Node, error) {
	value := &code.HashOverride{}
	m.stack.Push(value)
	defer m.stack.Pop()

	var err error
	value.Block, err = mapAstNodeTo[*code.Block](original.Block, m)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (m *Mapper) MapIf(original ast.If) (code.Node, error) {
	value := &code.If{}
	m.stack.Push(value)
	defer m.stack.Pop()

	var err error
	value.Block, err = mapAstNodeTo[*code.Block](original.Block, m)
	if err != nil {
		return nil, err
	}

	value.Condition, err = mapAstNodeTo[code.Value](original.Condition, m)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (m *Mapper) MapInt64(original ast.Int64) (code.Node, error) {
	value := &code.Int64{}
	m.stack.Push(value)
	defer m.stack.Pop()

	return value, nil
}

func (m *Mapper) MapKeyValue(original ast.KeyValue) (code.Node, error) {
	value := &code.KeyValue{}
	m.stack.Push(value)
	defer m.stack.Pop()

	var err error
	value.Key, err = mapAstNodeTo[code.Value](original.Key, m)
	if err != nil {
		return nil, err
	}

	value.Value, err = mapAstNodeTo[code.Value](original.Value, m)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (m *Mapper) MapLength(original ast.Length) (code.Node, error) {
	value := &code.Length{}
	m.stack.Push(value)
	defer m.stack.Pop()

	var err error
	value.Of, err = mapAstNodeTo[code.Value](original.Of, m)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (m *Mapper) MapList(original ast.List) (code.Node, error) {
	value := &code.List{}
	m.stack.Push(value)
	defer m.stack.Pop()

	var err error
	value.Item, err = mapAstNodeTo[code.Type](original.Item, m)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (m *Mapper) MapLiteralBool(original ast.LiteralBool) (code.Node, error) {
	value := &code.LiteralBool{}
	m.stack.Push(value)
	defer m.stack.Pop()

	value.Value = original.Value

	return value, nil
}

func (m *Mapper) MapLiteralInt64(original ast.LiteralInt64) (code.Node, error) {
	value := &code.LiteralInt64{}
	m.stack.Push(value)
	defer m.stack.Pop()

	value.Value = original.Value

	return value, nil
}

func (m *Mapper) MapLiteralList(original ast.LiteralList) (code.Node, error) {
	value := &code.LiteralList{}
	m.stack.Push(value)
	defer m.stack.Pop()

	var err error
	value.Values, err = mapAstNodesTo[code.Value](original.Values, m)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (m *Mapper) MapLiteralMap(original ast.LiteralMap) (code.Node, error) {
	value := &code.LiteralMap{}
	m.stack.Push(value)
	defer m.stack.Pop()

	var err error
	value.Values, err = mapAstNodesTo[*code.KeyValue](original.Values, m)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (m *Mapper) MapLiteralRune(original ast.LiteralRune) (code.Node, error) {
	value := &code.LiteralRune{}
	m.stack.Push(value)
	defer m.stack.Pop()

	value.Value = original.Value

	return value, nil
}

func (m *Mapper) MapLiteralSet(original ast.LiteralSet) (code.Node, error) {
	value := &code.LiteralSet{}
	m.stack.Push(value)
	defer m.stack.Pop()

	var err error
	value.Values, err = mapAstNodesTo[code.Value](original.Values, m)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (m *Mapper) MapLiteralString(original ast.LiteralString) (code.Node, error) {
	value := &code.LiteralString{}
	m.stack.Push(value)
	defer m.stack.Pop()

	value.Value = original.Value

	return value, nil
}

func (m *Mapper) MapLookup(original ast.Lookup) (code.Node, error) {
	value := &code.Lookup{}
	m.stack.Push(value)
	defer m.stack.Pop()

	var err error
	value.From, err = mapAstNodeTo[code.Value](original.From, m)
	if err != nil {
		return nil, err
	}

	value.Key, err = mapAstNodeTo[code.Value](original.Key, m)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (m *Mapper) MapMap(original ast.Map) (code.Node, error) {
	value := &code.Map{}
	m.stack.Push(value)
	defer m.stack.Pop()

	var err error
	value.Key, err = mapAstNodeTo[code.Type](original.Key, m)
	if err != nil {
		return nil, err
	}

	value.Value, err = mapAstNodeTo[code.Type](original.Value, m)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (m *Mapper) MapModel(original ast.Model) (code.Node, error) {
	value := &code.Model{}
	m.stack.Push(value)
	defer m.stack.Pop()

	value.Name = original.Name

	return value, nil
}

func (m *Mapper) MapModelDef(original ast.ModelDef) (code.Node, error) {
	value := &code.ModelDef{}
	m.stack.Push(value)
	defer m.stack.Pop()

	value.Name = original.Name

	// Defer because something inside this model might refer to a mode that hasn't been processed yet.
	m.queueDeferred(func() error {
		var err error
		value.EqualOverride, err = mapAstNodeTo[*code.EqualOverride](original.EqualOverride, m)
		if err != nil {
			return err
		}

		value.Fields, err = mapAstNodesTo[*code.FieldDef](original.Fields, m)
		if err != nil {
			return err
		}

		value.HashOverride, err = mapAstNodeTo[*code.HashOverride](original.HashOverride, m)
		if err != nil {
			return err
		}

		value.Methods, err = mapAstNodesTo[*code.FunctionDef](original.Methods, m)
		if err != nil {
			return err
		}

		return nil
	})

	return value, nil
}

func (m *Mapper) MapModule(original ast.Module) (code.Node, error) {
	value := &code.Module{}
	m.stack.Push(value)
	defer m.stack.Pop()

	var err error
	value.Constants, err = mapAstNodesTo[*code.ConstantDef](original.Constants, m)
	if err != nil {
		return nil, err
	}

	value.Functions, err = mapAstNodesTo[*code.FunctionDef](original.Functions, m)
	if err != nil {
		return nil, err
	}

	value.Models, err = mapAstNodesTo[*code.ModelDef](original.Models, m)
	if err != nil {
		return nil, err
	}

	value.Name = original.Name

	return value, nil
}

func (m *Mapper) MapNew(original ast.New) (code.Node, error) {
	value := &code.New{}
	m.stack.Push(value)
	defer m.stack.Pop()

	var err error
	value.Model, err = mapAstNodeTo[*code.Model](original.Model, m)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (m *Mapper) MapNil(original ast.Nil) (code.Node, error) {
	value := &code.Nil{}
	m.stack.Push(value)
	defer m.stack.Pop()

	var err error
	value.Type, err = mapAstNodeTo[code.Type](original.Type, m)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (m *Mapper) MapPop(original ast.Pop) (code.Node, error) {
	value := &code.Pop{}
	m.stack.Push(value)
	defer m.stack.Pop()

	var err error
	value.List, err = mapAstNodeTo[code.Value](original.List, m)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (m *Mapper) MapProperty(original ast.Property) (code.Node, error) {
	value := &code.Property{}
	m.stack.Push(value)
	defer m.stack.Pop()

	value.Name = original.Name

	var err error
	value.Of, err = mapAstNodeTo[code.Value](original.Of, m)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (m *Mapper) MapPush(original ast.Push) (code.Node, error) {
	value := &code.Push{}
	m.stack.Push(value)
	defer m.stack.Pop()

	var err error
	value.List, err = mapAstNodeTo[code.Value](original.List, m)
	if err != nil {
		return nil, err
	}

	value.Value, err = mapAstNodeTo[code.Value](original.Value, m)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (m *Mapper) MapReturn(original ast.Return) (code.Node, error) {
	value := &code.Return{}
	m.stack.Push(value)
	defer m.stack.Pop()

	var err error
	value.Value, err = mapAstNodeTo[code.Value](original.Value, m)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (m *Mapper) MapRune(original ast.Rune) (code.Node, error) {
	value := &code.Rune{}
	m.stack.Push(value)
	defer m.stack.Pop()

	return value, nil
}

func (m *Mapper) MapSelf(original ast.Self) (code.Node, error) {
	value := &code.Self{}
	m.stack.Push(value)
	defer m.stack.Pop()

	return value, nil
}

func (m *Mapper) MapSet(original ast.Set) (code.Node, error) {
	value := &code.Set{}
	m.stack.Push(value)
	defer m.stack.Pop()

	var err error
	value.Item, err = mapAstNodeTo[code.Type](original.Item, m)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (m *Mapper) MapSetContains(original ast.SetContains) (code.Node, error) {
	value := &code.SetContains{}
	m.stack.Push(value)
	defer m.stack.Pop()

	var err error
	value.Set, err = mapAstNodeTo[code.Value](original.Set, m)
	if err != nil {
		return nil, err
	}

	value.Value, err = mapAstNodeTo[code.Value](original.Value, m)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (m *Mapper) MapString(original ast.String) (code.Node, error) {
	value := &code.String{}
	m.stack.Push(value)
	defer m.stack.Pop()

	return value, nil
}

func (m *Mapper) MapVariable(original ast.Variable) (code.Node, error) {
	value := &code.Variable{}
	m.stack.Push(value)
	defer m.stack.Pop()

	value.Name = original.Name

	return value, nil
}

func (m *Mapper) MapVoid(original ast.Void) (code.Node, error) {
	value := &code.Void{}
	m.stack.Push(value)
	defer m.stack.Pop()

	return value, nil
}
