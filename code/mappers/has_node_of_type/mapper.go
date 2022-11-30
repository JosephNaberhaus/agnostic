package find_node_of_type

import "github.com/JosephNaberhaus/agnostic/code"

type Mapper[Target code.Node] struct{}

func (m Mapper[Target]) isTargetType(node code.Node) bool {
	if _, ok := node.(Target); ok {
		return true
	}

	return false
}

func (m Mapper[T]) MapFunction(original *code.Function) (bool, error) {
	if m.isTargetType(original) {
		return true, nil
	}

	return false, nil
}

func (m Mapper[T]) MapFunctionProperty(original *code.FunctionProperty) (bool, error) {
	if m.isTargetType(original) {
		return true, nil
	}

	return code.MapValue[bool](original.Of, m)
}

func (m Mapper[T]) MapLiteralInt(original *code.LiteralInt) (bool, error) {
	if m.isTargetType(original) {
		return true, nil
	}

	return false, nil
}

func (m Mapper[T]) MapLiteralString(original *code.LiteralString) (bool, error) {
	if m.isTargetType(original) {
		return true, nil
	}

	return false, nil
}

func (m Mapper[T]) MapFieldDef(original *code.FieldDef) (bool, error) {
	if m.isTargetType(original) {
		return true, nil
	}

	return
}

func (m Mapper[T]) MapArgumentDef(original *code.ArgumentDef) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (m Mapper[T]) MapMethodDef(original *code.MethodDef) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (m Mapper[T]) MapModelDef(original *code.ModelDef) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (m Mapper[T]) MapVariable(original *code.Variable) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (m Mapper[T]) MapProperty(original *code.Property) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (m Mapper[T]) MapModule(original *code.Module) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (m Mapper[T]) MapFunctionDef(original *code.FunctionDef) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (m Mapper[T]) MapUnaryOperator(original code.UnaryOperator) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (m Mapper[T]) MapUnaryOperation(original *code.UnaryOperation) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (m Mapper[T]) MapBinaryOperator(original code.BinaryOperator) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (m Mapper[T]) MapBinaryOperation(original *code.BinaryOperation) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (m Mapper[T]) MapAssignment(original *code.Assignment) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (m Mapper[T]) MapIf(original *code.If) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (m Mapper[T]) MapElseIf(original *code.ElseIf) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (m Mapper[T]) MapElse(original *code.Else) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (m Mapper[T]) MapConditional(original *code.Conditional) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (m Mapper[T]) MapReturn(original *code.Return) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (m Mapper[T]) MapDeclare(original *code.Declare) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (m Mapper[T]) MapModel(original *code.Model) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (m Mapper[T]) MapPrimitive(original code.Primitive) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (m Mapper[T]) MapList(original *code.List) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (m Mapper[T]) MapMap(original *code.Map) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (m Mapper[T]) MapCall(original *code.Call) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (m Mapper[T]) MapLookup(original *code.Lookup) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (m Mapper[T]) MapNew(original *code.New) (bool, error) {
	//TODO implement me
	panic("implement me")
}
