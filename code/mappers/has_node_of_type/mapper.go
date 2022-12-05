package has_node_of_type

import "github.com/JosephNaberhaus/agnostic/code"

type Mapper[Target code.Node] struct{}

func (m Mapper[Target]) isTargetType(node code.Node) bool {
	if _, ok := node.(Target); ok {
		return true
	}

	return false
}

func (m Mapper[T]) MapFunctionNoError(original *code.Function) bool {
	if m.isTargetType(original) {
		return true
	}

	return false
}

func (m Mapper[T]) MapFunctionPropertyNoError(original *code.FunctionProperty) bool {
	if m.isTargetType(original) {
		return true
	}

	if code.MapValueNoError[bool](original.Of, m) {
		return true
	}

	return false
}

func (m Mapper[T]) MapLiteralIntNoError(original *code.LiteralInt) bool {
	if m.isTargetType(original) {
		return true
	}

	return false
}

func (m Mapper[T]) MapLiteralStringNoError(original *code.LiteralString) bool {
	if m.isTargetType(original) {
		return true
	}

	return false
}

func (m Mapper[T]) MapLiteralRuneNoError(original *code.LiteralRune) bool {
	if m.isTargetType(original) {
		return true
	}

	return false
}

func (m Mapper[T]) MapFieldDefNoError(original *code.FieldDef) bool {
	if m.isTargetType(original) {
		return true
	}

	if code.MapTypeNoError[bool](original.Type, m) {
		return true
	}

	return false
}

func (m Mapper[T]) MapArgumentDefNoError(original *code.ArgumentDef) bool {
	if m.isTargetType(original) {
		return true
	}

	if code.MapTypeNoError[bool](original.Type, m) {
		return true
	}

	return false
}

func (m Mapper[T]) MapMethodDefNoError(original *code.MethodDef) bool {
	if m.isTargetType(original) {
		return true
	}

	if m.MapFunctionDefNoError(original.Function) {
		return true
	}

	return false
}

func (m Mapper[T]) MapModelDefNoError(original *code.ModelDef) bool {
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

func (m Mapper[T]) MapVariableNoError(original *code.Variable) bool {
	if m.isTargetType(original) {
		return true
	}

	return false
}

func (m Mapper[T]) MapPropertyNoError(original *code.Property) bool {
	if m.isTargetType(original) {
		return true
	}

	if code.MapValueNoError[bool](original.Of, m) {
		return true
	}

	return false
}

func (m Mapper[T]) MapModuleNoError(original *code.Module) bool {
	if m.isTargetType(original) {
		return true
	}

	if any(code.MapNodesNoError[bool](original.Models, m)) {
		return true
	}

	if any(code.MapNodesNoError[bool](original.Functions, m)) {
		return true
	}

	return false
}

func (m Mapper[T]) MapFunctionDefNoError(original *code.FunctionDef) bool {
	if m.isTargetType(original) {
		return true
	}

	if any(code.MapNodesNoError[bool](original.Arguments, m)) {
		return true
	}

	if m.MapBlockNoError(original.Block) {
		return true
	}

	if code.MapTypeNoError[bool](original.ReturnType, m) {
		return true
	}

	return false
}

func (m Mapper[T]) MapUnaryOperatorNoError(original code.UnaryOperator) bool {
	if m.isTargetType(original) {
		return true
	}

	return false
}

func (m Mapper[T]) MapUnaryOperationNoError(original *code.UnaryOperation) bool {
	if m.isTargetType(original) {
		return true
	}

	if code.MapValueNoError[bool](original.Value, m) {
		return true
	}

	if m.MapUnaryOperatorNoError(original.Operator) {
		return true
	}

	return false
}

func (m Mapper[T]) MapBinaryOperatorNoError(original code.BinaryOperator) bool {
	if m.isTargetType(original) {
		return true
	}

	return false
}

func (m Mapper[T]) MapBinaryOperationNoError(original *code.BinaryOperation) bool {
	if m.isTargetType(original) {
		return true
	}

	if code.MapValueNoError[bool](original.Left, m) {
		return true
	}

	if m.MapBinaryOperatorNoError(original.Operator) {
		return true
	}

	if code.MapValueNoError[bool](original.Right, m) {
		return true
	}

	return false
}

func (m Mapper[T]) MapAssignmentNoError(original *code.Assignment) bool {
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

func (m Mapper[T]) MapIfNoError(original *code.If) bool {
	if m.isTargetType(original) {
		return true
	}

	if code.MapValueNoError[bool](original.Condition, m) {
		return true
	}

	if m.MapBlockNoError(original.Block) {
		return true
	}

	return false
}

func (m Mapper[T]) MapElseIfNoError(original *code.ElseIf) bool {
	if m.isTargetType(original) {
		return true
	}

	if code.MapValueNoError[bool](original.Condition, m) {
		return true
	}

	if m.MapBlockNoError(original.Block) {
		return true
	}

	return false
}

func (m Mapper[T]) MapElseNoError(original *code.Else) bool {
	if m.isTargetType(original) {
		return true
	}

	if m.MapBlockNoError(original.Block) {
		return true
	}

	return false
}

func (m Mapper[T]) MapConditionalNoError(original *code.Conditional) bool {
	if m.isTargetType(original) {
		return true
	}

	if m.MapIfNoError(original.If) {
		return true
	}

	if any(code.MapNodesNoError[bool](original.ElseIfs, m)) {
		return true
	}

	if original.Else != nil && m.MapElseNoError(original.Else) {
		return true
	}

	return false
}

func (m Mapper[T]) MapReturnNoError(original *code.Return) bool {
	if m.isTargetType(original) {
		return true
	}

	if code.MapValueNoError[bool](original.Value, m) {
		return true
	}

	return false
}

func (m Mapper[T]) MapDeclareNoError(original *code.Declare) bool {
	if m.isTargetType(original) {
		return true
	}

	if code.MapValueNoError[bool](original.Value, m) {
		return true
	}

	return false
}

func (m Mapper[T]) MapModelNoError(original *code.Model) bool {
	if m.isTargetType(original) {
		return true
	}

	return false
}

func (m Mapper[T]) MapPrimitiveNoError(original code.Primitive) bool {
	if m.isTargetType(original) {
		return true
	}

	return false
}

func (m Mapper[T]) MapListNoError(original *code.List) bool {
	if m.isTargetType(original) {
		return true
	}

	if code.MapTypeNoError[bool](original.Base, m) {
		return true
	}

	return false
}

func (m Mapper[T]) MapMapNoError(original *code.Map) bool {
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

func (m Mapper[T]) MapCallNoError(original *code.Call) bool {
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

func (m Mapper[T]) MapLookupNoError(original *code.Lookup) bool {
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

func (m Mapper[T]) MapNewNoError(original *code.New) bool {
	if m.isTargetType(original) {
		return true
	}

	if code.MapTypeNoError[bool](original.Model, m) {
		return true
	}

	return false
}

func (m Mapper[T]) MapForNoError(original *code.For) bool {
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

	if m.MapBlockNoError(original.Block) {
		return true
	}

	return false
}

func (m Mapper[T]) MapForInNoError(original *code.ForIn) bool {
	if m.isTargetType(original) {
		return true
	}

	if code.MapValueNoError[bool](original.Iterable, m) {
		return true
	}

	if m.MapBlockNoError(original.Block) {
		return true
	}

	return false
}

func (m Mapper[T]) MapBlockNoError(original *code.Block) bool {
	if m.isTargetType(original) {
		return true
	}

	if any(code.MapStatementsNoError[bool](original.Statements, m)) {
		return true
	}

	return false
}

func (m Mapper[T]) MapLiteralListNoError(original *code.LiteralList) bool {
	if m.isTargetType(original) {
		return true
	}

	if any(code.MapValuesNoError[bool](original.Items, m)) {
		return true
	}

	return false
}

func (m Mapper[T]) MapLengthNoError(original *code.Length) bool {
	if m.isTargetType(original) {
		return true
	}

	if code.MapValueNoError[bool](original.Value, m) {
		return true
	}

	return false
}

func (m Mapper[T]) MapConstantDefNoError(original *code.ConstantDef) bool {
	if m.isTargetType(original) {
		return true
	}

	if code.MapValueNoError[bool](original.Value, m) {
		return true
	}

	return false
}

func (m Mapper[T]) MapKeyValueNoError(original *code.KeyValue) bool {
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

func (m Mapper[T]) MapLiteralMapNoError(original *code.LiteralMap) bool {
	if m.isTargetType(original) {
		return true
	}

	if any(code.MapNodesNoError[bool](original.Entries, m)) {
		return true
	}

	return false
}
