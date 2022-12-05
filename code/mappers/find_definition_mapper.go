package mappers

import (
	"fmt"
	"github.com/JosephNaberhaus/agnostic/code"
)

func mapFindDefinition(targetName string, stack code.Stack) (code.Definition, error) {
	mapper := &findDefinitionMapper{
		targetName: targetName,
	}

	for stack.IsNotEmpty() {
		definition, err := code.MapNode[code.Definition](stack.Pop(), mapper)
		if err != nil {
			return nil, err
		}

		if definition != nil {
			return definition, nil
		}
	}

	// TODO
	return nil, fmt.Errorf("couldn't find definition for %s", targetName)
}

type findDefinitionMapper struct {
	targetName string
}

var _ code.NodeMapper[code.Definition] = (*findDefinitionMapper)(nil)

func (f *findDefinitionMapper) MapLiteralInt(original *code.LiteralInt) (code.Definition, error) {
	return nil, nil
}

func (f *findDefinitionMapper) MapLiteralString(original *code.LiteralString) (code.Definition, error) {
	return nil, nil
}

func (f *findDefinitionMapper) MapLiteralRune(original *code.LiteralRune) (code.Definition, error) {
	return nil, nil
}

func (f *findDefinitionMapper) MapFieldDef(original *code.FieldDef) (code.Definition, error) {
	if original.Name == f.targetName {
		return original, nil
	}

	return nil, nil
}

func (f *findDefinitionMapper) MapArgumentDef(original *code.ArgumentDef) (code.Definition, error) {
	if original.Name == f.targetName {
		return original, nil
	}

	return nil, nil
}

func (f *findDefinitionMapper) MapMethodDef(original *code.MethodDef) (code.Definition, error) {
	return f.MapFunctionDef(original.Function)
}

func (f *findDefinitionMapper) MapModelDef(original *code.ModelDef) (code.Definition, error) {
	return findDefinitionInNodes(original.Fields, f)
}

func (f *findDefinitionMapper) MapVariable(original *code.Variable) (code.Definition, error) {
	return nil, nil
}

func (f *findDefinitionMapper) MapProperty(original *code.Property) (code.Definition, error) {
	return nil, nil
}

func (f *findDefinitionMapper) MapModule(original *code.Module) (code.Definition, error) {
	return findDefinitionInNodes(original.Constants, f)
}

func (f *findDefinitionMapper) MapUnaryOperator(original code.UnaryOperator) (code.Definition, error) {
	return nil, nil
}

func (f *findDefinitionMapper) MapUnaryOperation(original *code.UnaryOperation) (code.Definition, error) {
	return nil, nil
}

func (f *findDefinitionMapper) MapBinaryOperator(original code.BinaryOperator) (code.Definition, error) {
	return nil, nil
}

func (f *findDefinitionMapper) MapBinaryOperation(original *code.BinaryOperation) (code.Definition, error) {
	return nil, nil
}

func (f *findDefinitionMapper) MapAssignment(original *code.Assignment) (code.Definition, error) {
	return nil, nil
}

func (f *findDefinitionMapper) MapIf(original *code.If) (code.Definition, error) {
	return nil, nil
}

func (f *findDefinitionMapper) MapElseIf(original *code.ElseIf) (code.Definition, error) {
	return nil, nil
}

func (f *findDefinitionMapper) MapElse(original *code.Else) (code.Definition, error) {
	return nil, nil
}

func (f *findDefinitionMapper) MapConditional(original *code.Conditional) (code.Definition, error) {
	return nil, nil
}

func (f *findDefinitionMapper) MapModel(original *code.Model) (code.Definition, error) {
	return nil, nil
}

func (f *findDefinitionMapper) MapPrimitive(original code.Primitive) (code.Definition, error) {
	return nil, nil
}

func (f *findDefinitionMapper) MapList(original *code.List) (code.Definition, error) {
	return nil, nil
}

func (f *findDefinitionMapper) MapMap(original *code.Map) (code.Definition, error) {
	return nil, nil
}

func (f *findDefinitionMapper) MapLookup(original *code.Lookup) (code.Definition, error) {
	return nil, nil
}

func (f *findDefinitionMapper) MapFunctionDef(original *code.FunctionDef) (code.Definition, error) {
	return findDefinitionInNodes(original.Arguments, f)
}

func (f *findDefinitionMapper) MapReturn(original *code.Return) (code.Definition, error) {
	return nil, nil
}

func (f *findDefinitionMapper) MapCall(original *code.Call) (code.Definition, error) {
	return nil, nil
}

func (f *findDefinitionMapper) MapFunction(original *code.Function) (code.Definition, error) {
	return nil, nil
}

func (f *findDefinitionMapper) MapFunctionProperty(original *code.FunctionProperty) (code.Definition, error) {
	return nil, nil
}

func (f *findDefinitionMapper) MapNew(original *code.New) (code.Definition, error) {
	return nil, nil
}

func (f *findDefinitionMapper) MapDeclare(original *code.Declare) (code.Definition, error) {
	if original.Name == f.targetName {
		return original, nil
	}

	return nil, nil
}

func (f *findDefinitionMapper) MapFor(original *code.For) (code.Definition, error) {
	return code.MapStatement[code.Definition](original.Initialization, f)
}

func (f *findDefinitionMapper) MapForIn(original *code.ForIn) (code.Definition, error) {
	if original.ItemName == f.targetName {
		return original, nil
	}

	return nil, nil
}

func (f *findDefinitionMapper) MapBlock(original *code.Block) (code.Definition, error) {
	return findDefinitionInNodes(original.Statements, f)
}

func (f *findDefinitionMapper) MapLiteralList(original *code.LiteralList) (code.Definition, error) {
	return nil, nil
}

func (f *findDefinitionMapper) MapLength(original *code.Length) (code.Definition, error) {
	return nil, nil
}

func (f *findDefinitionMapper) MapConstantDef(original *code.ConstantDef) (code.Definition, error) {
	if original.Name == f.targetName {
		return original, nil
	}

	return nil, nil
}

func (f *findDefinitionMapper) MapKeyValue(original *code.KeyValue) (code.Definition, error) {
	return nil, nil
}

func (f *findDefinitionMapper) MapLiteralMap(original *code.LiteralMap) (code.Definition, error) {
	return nil, nil
}

func findDefinitionInNodes[T code.Node](nodes []T, mapper code.NodeMapper[code.Definition]) (code.Definition, error) {
	for _, node := range nodes {
		definition, err := code.MapNode[code.Definition](node, mapper)
		if err != nil {
			return nil, err
		}

		if definition != nil {
			return definition, nil
		}
	}

	return nil, nil
}
