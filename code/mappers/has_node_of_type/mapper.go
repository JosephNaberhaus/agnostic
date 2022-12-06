package has_node_of_type

import "github.com/JosephNaberhaus/agnostic/code"

type Mapper[Target code.Node] struct{}

func (m Mapper[Target]) isTargetType(node code.Node) bool {
	if _, ok := node.(Target); ok {
		return true
	}

	return false
}

func (m Mapper[T]) MapFunction(original *code.Function) bool {
	if m.isTargetType(original) {
		return true
	}

	return false
}

func (m Mapper[T]) MapFunctionProperty(original *code.FunctionProperty) bool {
	if m.isTargetType(original) {
		return true
	}

	if code.MapValueNoError[bool](original.Of, m) {
		return true
	}

	return false
}

func (m Mapper[T]) MapLiteralInt(original *code.LiteralInt) bool {
	if m.isTargetType(original) {
		return true
	}

	return false
}

func (m Mapper[T]) MapLiteralString(original *code.LiteralString) bool {
	if m.isTargetType(original) {
		return true
	}

	return false
}

func (m Mapper[T]) MapLiteralRune(original *code.LiteralRune) bool {
	if m.isTargetType(original) {
		return true
	}

	return false
}

func (m Mapper[T]) MapFieldDef(original *code.FieldDef) bool {
	if m.isTargetType(original) {
		return true
	}

	if code.MapTypeNoError[bool](original.Type, m) {
		return true
	}

	return false
}

func (m Mapper[T]) MapArgumentDef(original *code.ArgumentDef) bool {
	if m.isTargetType(original) {
		return true
	}

	if code.MapTypeNoError[bool](original.Type, m) {
		return true
	}

	return false
}

func (m Mapper[T]) MapMethodDef(original *code.MethodDef) bool {
	if m.isTargetType(original) {
		return true
	}

	if m.MapFunctionDef(original.Function) {
		return true
	}

	return false
}

func (m Mapper[T]) MapModelDef(original *code.ModelDef) bool {
	if m.isTargetType(original) {
		return true
	}

	if any(code.MapNodesNoError[bool](original.Fields, m)) {
		return true
	}

	if any(code.MapNodesNoError[bool](original.Methods, m)) {
		return true
	}

	return false
}

func (m Mapper[T]) MapVariable(original *code.Variable) bool {
	if m.isTargetType(original) {
		return true
	}

	return false
}

func (m Mapper[T]) MapProperty(original *code.Property) bool {
	if m.isTargetType(original) {
		return true
	}

	if code.MapValueNoError[bool](original.Of, m) {
		return true
	}

	return false
}

func (m Mapper[T]) MapModule(original *code.Module) bool {
	if m.isTargetType(original) {
		return true
	}

	if any(code.MapNodesNoError[bool](original.Models, m)) {
		return true
	}

	if any(code.MapNodesNoError[bool](original.Functions, m)) {
		return true
	}

	if any(code.MapNodesNoError[bool](original.Constants, m)) {
		return true
	}

	return false
}

func (m Mapper[T]) MapFunctionDef(original *code.FunctionDef) bool {
	if m.isTargetType(original) {
		return true
	}

	if any(code.MapNodesNoError[bool](original.Arguments, m)) {
		return true
	}

	if m.MapBlock(original.Block) {
		return true
	}

	if code.MapTypeNoError[bool](original.ReturnType, m) {
		return true
	}

	return false
}

func (m Mapper[T]) MapUnaryOperator(original code.UnaryOperator) bool {
	if m.isTargetType(original) {
		return true
	}

	return false
}

func (m Mapper[T]) MapUnaryOperation(original *code.UnaryOperation) bool {
	if m.isTargetType(original) {
		return true
	}

	if code.MapValueNoError[bool](original.Value, m) {
		return true
	}

	if m.MapUnaryOperator(original.Operator) {
		return true
	}

	return false
}

func (m Mapper[T]) MapBinaryOperator(original code.BinaryOperator) bool {
	if m.isTargetType(original) {
		return true
	}

	return false
}

func (m Mapper[T]) MapBinaryOperation(original *code.BinaryOperation) bool {
	if m.isTargetType(original) {
		return true
	}

	if code.MapValueNoError[bool](original.Left, m) {
		return true
	}

	if m.MapBinaryOperator(original.Operator) {
		return true
	}

	if code.MapValueNoError[bool](original.Right, m) {
		return true
	}

	return false
}

func (m Mapper[T]) MapAssignment(original *code.Assignment) bool {
	if m.isTargetType(original) {
		return true
	}

	if code.MapValueNoError[bool](original.To, m) {
		return true
	}

	if code.MapValueNoError[bool](original.From, m) {
		return true
	}

	return false
}

func (m Mapper[T]) MapIf(original *code.If) bool {
	if m.isTargetType(original) {
		return true
	}

	if code.MapValueNoError[bool](original.Condition, m) {
		return true
	}

	if m.MapBlock(original.Block) {
		return true
	}

	return false
}

func (m Mapper[T]) MapElseIf(original *code.ElseIf) bool {
	if m.isTargetType(original) {
		return true
	}

	if code.MapValueNoError[bool](original.Condition, m) {
		return true
	}

	if m.MapBlock(original.Block) {
		return true
	}

	return false
}

func (m Mapper[T]) MapElse(original *code.Else) bool {
	if m.isTargetType(original) {
		return true
	}

	if m.MapBlock(original.Block) {
		return true
	}

	return false
}

func (m Mapper[T]) MapConditional(original *code.Conditional) bool {
	if m.isTargetType(original) {
		return true
	}

	if m.MapIf(original.If) {
		return true
	}

	if any(code.MapNodesNoError[bool](original.ElseIfs, m)) {
		return true
	}

	if original.Else != nil && m.MapElse(original.Else) {
		return true
	}

	return false
}

func (m Mapper[T]) MapReturn(original *code.Return) bool {
	if m.isTargetType(original) {
		return true
	}

	if code.MapValueNoError[bool](original.Value, m) {
		return true
	}

	return false
}

func (m Mapper[T]) MapDeclare(original *code.Declare) bool {
	if m.isTargetType(original) {
		return true
	}

	if code.MapValueNoError[bool](original.Value, m) {
		return true
	}

	return false
}

func (m Mapper[T]) MapModel(original *code.Model) bool {
	if m.isTargetType(original) {
		return true
	}

	return false
}

func (m Mapper[T]) MapPrimitive(original code.Primitive) bool {
	if m.isTargetType(original) {
		return true
	}

	return false
}

func (m Mapper[T]) MapList(original *code.List) bool {
	if m.isTargetType(original) {
		return true
	}

	if code.MapTypeNoError[bool](original.Base, m) {
		return true
	}

	return false
}

func (m Mapper[T]) MapMap(original *code.Map) bool {
	if m.isTargetType(original) {
		return true
	}

	if code.MapTypeNoError[bool](original.Key, m) {
		return true
	}

	if code.MapTypeNoError[bool](original.Value, m) {
		return true
	}

	return false
}

func (m Mapper[T]) MapCall(original *code.Call) bool {
	if m.isTargetType(original) {
		return true
	}

	if code.MapCallableNoError[bool](original.Function, m) {
		return true
	}

	if any(code.MapNodesNoError[bool](original.Arguments, m)) {
		return true
	}

	return false
}

func (m Mapper[T]) MapLookup(original *code.Lookup) bool {
	if m.isTargetType(original) {
		return true
	}

	if code.MapValueNoError[bool](original.From, m) {
		return true
	}

	if code.MapValueNoError[bool](original.Key, m) {
		return true
	}

	return false
}

func (m Mapper[T]) MapNew(original *code.New) bool {
	if m.isTargetType(original) {
		return true
	}

	if code.MapTypeNoError[bool](original.Model, m) {
		return true
	}

	return false
}

func (m Mapper[T]) MapFor(original *code.For) bool {
	if m.isTargetType(original) {
		return true
	}

	if code.MapStatementNoError[bool](original.Initialization, m) {
		return true
	}

	if code.MapValueNoError[bool](original.Condition, m) {
		return true
	}

	if code.MapStatementNoError[bool](original.AfterEach, m) {
		return true
	}

	if m.MapBlock(original.Block) {
		return true
	}

	return false
}

func (m Mapper[T]) MapForIn(original *code.ForIn) bool {
	if m.isTargetType(original) {
		return true
	}

	if code.MapValueNoError[bool](original.Iterable, m) {
		return true
	}

	if m.MapBlock(original.Block) {
		return true
	}

	return false
}

func (m Mapper[T]) MapBlock(original *code.Block) bool {
	if m.isTargetType(original) {
		return true
	}

	if any(code.MapStatementsNoError[bool](original.Statements, m)) {
		return true
	}

	return false
}

func (m Mapper[T]) MapLiteralList(original *code.LiteralList) bool {
	if m.isTargetType(original) {
		return true
	}

	if any(code.MapValuesNoError[bool](original.Items, m)) {
		return true
	}

	return false
}

func (m Mapper[T]) MapLength(original *code.Length) bool {
	if m.isTargetType(original) {
		return true
	}

	if code.MapValueNoError[bool](original.Value, m) {
		return true
	}

	return false
}

func (m Mapper[T]) MapConstantDef(original *code.ConstantDef) bool {
	if m.isTargetType(original) {
		return true
	}

	if code.MapValueNoError[bool](original.Value, m) {
		return true
	}

	return false
}

func (m Mapper[T]) MapKeyValue(original *code.KeyValue) bool {
	if m.isTargetType(original) {
		return true
	}

	if code.MapValueNoError[bool](original.Key, m) {
		return true
	}

	if code.MapValueNoError[bool](original.Value, m) {
		return true
	}

	return false
}

func (m Mapper[T]) MapLiteralMap(original *code.LiteralMap) bool {
	if m.isTargetType(original) {
		return true
	}

	if any(code.MapNodesNoError[bool](original.Entries, m)) {
		return true
	}

	return false
}

func (m Mapper[T]) MapLiteralSet(original *code.LiteralSet) bool {
	if m.isTargetType(original) {
		return true
	}

	if any(code.MapNodesNoError[bool](original.Items, m)) {
		return true
	}

	return false
}

func (m Mapper[T]) MapSet(original *code.Set) bool {
	if m.isTargetType(original) {
		return true
	}

	if code.MapNodeNoError[bool](original.Base, m) {
		return true
	}

	return false
}

func (m Mapper[T]) MapEmptyList(original *code.EmptyList) bool {
	if m.isTargetType(original) {
		return true
	}

	if code.MapTypeNoError[bool](original.Type, m) {
		return true
	}

	return false
}

func (m Mapper[T]) MapEmptySet(original *code.EmptySet) bool {
	if m.isTargetType(original) {
		return true
	}

	if code.MapTypeNoError[bool](original.Type, m) {
		return true
	}

	return false
}

func (m Mapper[T]) MapAddToSet(original *code.AddToSet) bool {
	if m.isTargetType(original) {
		return true
	}

	if code.MapValueNoError[bool](original.To, m) {
		return true
	}

	if code.MapValueNoError[bool](original.Value, m) {
		return true
	}

	return false
}

func (m Mapper[T]) MapSetContains(original *code.SetContains) bool {
	if m.isTargetType(original) {
		return true
	}

	if code.MapValueNoError[bool](original.Set, m) {
		return true
	}

	if code.MapValueNoError[bool](original.Value, m) {
		return true
	}

	return false
}

func (m Mapper[T]) MapPush(original *code.Push) bool {
	if m.isTargetType(original) {
		return true
	}

	if code.MapValueNoError[bool](original.To, m) {
		return true
	}

	if code.MapValueNoError[bool](original.Value, m) {
		return true
	}

	return false
}

func (m Mapper[T]) MapPop(original *code.Pop) bool {
	if m.isTargetType(original) {
		return true
	}

	if code.MapValueNoError[bool](original.Value, m) {
		return true
	}

	return false
}

func (m Mapper[T]) MapLiteralBool(original *code.LiteralBool) bool {
	if m.isTargetType(original) {
		return true
	}

	return false
}
