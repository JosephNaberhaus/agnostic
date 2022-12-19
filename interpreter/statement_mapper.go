package interpreter

import "github.com/JosephNaberhaus/agnostic/code"

type returnResult struct {
	value Value
}

type statementMapper struct {
	runtime runtime
}

func (s *statementMapper) MapBlock(original *code.Block) *returnResult {
	for _, statement := range original.Statements {
		result := code.MapStatementNoError[*returnResult](statement, s)
		if result != nil {
			return result
		}
	}

	return nil
}

func (s *statementMapper) MapAssignment(original *code.Assignment) *returnResult {
	fromValue := code.MapValueNoError[Value](original.From, &valueMapper{runtime: s.runtime})

	switch to := original.To.(type) {
	case *code.Variable:
		s.runtime.assign(to.Name, fromValue)
	case *code.Property:
		toModel := code.MapValueNoError[Value](original.To, &valueMapper{runtime: s.runtime}).Value.(*Model)
		toModel.Properties[to.Name] = fromValue
	case *code.Lookup:
		switch to.LookupType {
		case code.LookupTypeList:
			toList := code.MapValueNoError[Value](original.To, &valueMapper{runtime: s.runtime}).Value.(*List)
			index := code.MapValueNoError[Value](to.Key, &valueMapper{runtime: s.runtime}).Value.(int64)
			toList.Set(index, fromValue)
		case code.LookupTypeMap:
			toMap := code.MapValueNoError[Value](original.To, &valueMapper{runtime: s.runtime}).Value.(*Map)
			key := code.MapValueNoError[Value](to.Key, &valueMapper{runtime: s.runtime})
			toMap.Put(key, fromValue)
		}
	default:
		panic("unsupported assignment")
	}

	return nil
}

func (s *statementMapper) MapConditional(original *code.Conditional) *returnResult {
	ifCondition := code.MapValueNoError[Value](original.If.Condition, &valueMapper{runtime: s.runtime}).Value.(bool)
	if ifCondition {
		return s.MapBlock(original.If.Block)
	}

	for _, elseIf := range original.ElseIfs {
		elseIfCondition := code.MapValueNoError[Value](elseIf.Condition, &valueMapper{runtime: s.runtime}).Value.(bool)
		if elseIfCondition {
			return s.MapBlock(elseIf.Block)
		}
	}

	if original.Else != nil {
		return s.MapBlock(original.Else.Block)
	}

	return nil
}

func (s *statementMapper) MapReturn(original *code.Return) *returnResult {
	value := code.MapValueNoError[Value](original.Value, &valueMapper{runtime: s.runtime})

	return &returnResult{value: value}
}

func (s *statementMapper) MapDeclare(original *code.Declare) *returnResult {
	value := code.MapValueNoError[Value](original.Value, &valueMapper{runtime: s.runtime})
	s.runtime.declare(original.Name, value)

	return nil
}

func (s *statementMapper) MapFor(original *code.For) *returnResult {
	s.runtime.startBlock()

	if original.Initialization.IsSet() {
		code.MapStatementNoError[*returnResult](original.Initialization.Value(), s)
	}

	for true {
		condition := code.MapValueNoError[Value](original.Condition, &valueMapper{runtime: s.runtime}).Value.(bool)
		if !condition {
			break
		}

		result := s.MapBlock(original.Block)
		if result != nil {
			return result
		}

		if original.AfterEach.IsSet() {
			code.MapStatementNoError[*returnResult](original.AfterEach.Value(), s)
		}
	}

	s.runtime.endBlock()

	return nil
}

func (s *statementMapper) MapForIn(original *code.ForIn) *returnResult {
	switch original.ForInType {
	case code.ForInTypeList:
		list := code.MapValueNoError[Value](original.Iterable, &valueMapper{runtime: s.runtime}).Value.(*List)
		for _, item := range *list {
			s.runtime.startBlock()

			s.runtime.declare(original.ItemName, item)

			result := s.MapBlock(original.Block)
			if result != nil {
				return result
			}

			s.runtime.endBlock()
		}
	case code.ForInTypeSet:
		set := code.MapValueNoError[Value](original.Iterable, &valueMapper{runtime: s.runtime}).Value.(*Set)
		for item := range *set {
			s.runtime.startBlock()

			s.runtime.declare(original.ItemName, item)

			result := s.MapBlock(original.Block)
			if result != nil {
				return result
			}

			s.runtime.endBlock()
		}
	}

	return nil
}

func (s *statementMapper) MapAddToSet(original *code.AddToSet) *returnResult {
	toSet := code.MapValueNoError[Value](original.To, &valueMapper{runtime: s.runtime}).Value.(*Set)
	value := code.MapValueNoError[Value](original.Value, &valueMapper{runtime: s.runtime})
	toSet.Add(value)

	return nil
}

func (s *statementMapper) MapPush(original *code.Push) *returnResult {
	toList := code.MapValueNoError[Value](original.To, &valueMapper{runtime: s.runtime}).Value.(*List)
	value := code.MapValueNoError[Value](original.Value, &valueMapper{runtime: s.runtime})
	toList.Push(value)

	return nil
}

func (s *statementMapper) MapDeclareNull(original *code.DeclareNull) *returnResult {
	s.runtime.declare(original.Name, Value{})

	return nil
}

func (s *statementMapper) MapBreak(_ *code.Break) *returnResult {
	// We handle break inside MapFor
	panic("unreachable")
}

func (s *statementMapper) MapContinue(_ *code.Continue) *returnResult {
	// We handle continue inside MapFor
	panic("unreachable")
}

func (s *statementMapper) MapCall(original *code.Call) *returnResult {
	_ = code.MapValueNoError[Value](original, &valueMapper{runtime: s.runtime}).Value.(*List)

	return nil
}

func (s *statementMapper) MapPop(original *code.Pop) *returnResult {
	list := code.MapValueNoError[Value](original.Value, &valueMapper{runtime: s.runtime}).Value.(*List)
	list.Pop()

	return nil
}
