package ast

type NodeMapper[T any] interface {
	MapAddToSet(value AddToSet) (T, error)

	MapArgumentDef(value ArgumentDef) (T, error)

	MapAssignment(value Assignment) (T, error)

	MapBlock(value Block) (T, error)

	MapBool(value Bool) (T, error)

	MapBreak(value Break) (T, error)

	MapCall(value Call) (T, error)

	MapConditional(value Conditional) (T, error)

	MapConstantDef(value ConstantDef) (T, error)

	MapContinue(value Continue) (T, error)

	MapDeclare(value Declare) (T, error)

	MapEmptyList(value EmptyList) (T, error)

	MapEqualOverride(value EqualOverride) (T, error)

	MapFieldDef(value FieldDef) (T, error)

	MapFor(value For) (T, error)

	MapForEach(value ForEach) (T, error)

	MapFunctionDef(value FunctionDef) (T, error)

	MapHashOverride(value HashOverride) (T, error)

	MapIf(value If) (T, error)

	MapInt64(value Int64) (T, error)

	MapKeyValue(value KeyValue) (T, error)

	MapLength(value Length) (T, error)

	MapList(value List) (T, error)

	MapLiteralBool(value LiteralBool) (T, error)

	MapLiteralInt(value LiteralInt) (T, error)

	MapLiteralList(value LiteralList) (T, error)

	MapLiteralMap(value LiteralMap) (T, error)

	MapLiteralRune(value LiteralRune) (T, error)

	MapLiteralSet(value LiteralSet) (T, error)

	MapLiteralString(value LiteralString) (T, error)

	MapLookup(value Lookup) (T, error)

	MapMap(value Map) (T, error)

	MapModel(value Model) (T, error)

	MapModelDef(value ModelDef) (T, error)

	MapModule(value Module) (T, error)

	MapNew(value New) (T, error)

	MapNull(value Null) (T, error)

	MapPop(value Pop) (T, error)

	MapProperty(value Property) (T, error)

	MapPush(value Push) (T, error)

	MapReturn(value Return) (T, error)

	MapRune(value Rune) (T, error)

	MapSelf(value Self) (T, error)

	MapSet(value Set) (T, error)

	MapSetContains(value SetContains) (T, error)

	MapString(value String) (T, error)

	MapVariable(value Variable) (T, error)

	MapVoid(value Void) (T, error)
}

func MapNode[T any](node Node, mapper NodeMapper[T]) (T, error) {
	switch value := node.(type) {

	case AddToSet:
		return mapper.MapAddToSet(value)

	case ArgumentDef:
		return mapper.MapArgumentDef(value)

	case Assignment:
		return mapper.MapAssignment(value)

	case Block:
		return mapper.MapBlock(value)

	case Bool:
		return mapper.MapBool(value)

	case Break:
		return mapper.MapBreak(value)

	case Call:
		return mapper.MapCall(value)

	case Conditional:
		return mapper.MapConditional(value)

	case ConstantDef:
		return mapper.MapConstantDef(value)

	case Continue:
		return mapper.MapContinue(value)

	case Declare:
		return mapper.MapDeclare(value)

	case EmptyList:
		return mapper.MapEmptyList(value)

	case EqualOverride:
		return mapper.MapEqualOverride(value)

	case FieldDef:
		return mapper.MapFieldDef(value)

	case For:
		return mapper.MapFor(value)

	case ForEach:
		return mapper.MapForEach(value)

	case FunctionDef:
		return mapper.MapFunctionDef(value)

	case HashOverride:
		return mapper.MapHashOverride(value)

	case If:
		return mapper.MapIf(value)

	case Int64:
		return mapper.MapInt64(value)

	case KeyValue:
		return mapper.MapKeyValue(value)

	case Length:
		return mapper.MapLength(value)

	case List:
		return mapper.MapList(value)

	case LiteralBool:
		return mapper.MapLiteralBool(value)

	case LiteralInt:
		return mapper.MapLiteralInt(value)

	case LiteralList:
		return mapper.MapLiteralList(value)

	case LiteralMap:
		return mapper.MapLiteralMap(value)

	case LiteralRune:
		return mapper.MapLiteralRune(value)

	case LiteralSet:
		return mapper.MapLiteralSet(value)

	case LiteralString:
		return mapper.MapLiteralString(value)

	case Lookup:
		return mapper.MapLookup(value)

	case Map:
		return mapper.MapMap(value)

	case Model:
		return mapper.MapModel(value)

	case ModelDef:
		return mapper.MapModelDef(value)

	case Module:
		return mapper.MapModule(value)

	case New:
		return mapper.MapNew(value)

	case Null:
		return mapper.MapNull(value)

	case Pop:
		return mapper.MapPop(value)

	case Property:
		return mapper.MapProperty(value)

	case Push:
		return mapper.MapPush(value)

	case Return:
		return mapper.MapReturn(value)

	case Rune:
		return mapper.MapRune(value)

	case Self:
		return mapper.MapSelf(value)

	case Set:
		return mapper.MapSet(value)

	case SetContains:
		return mapper.MapSetContains(value)

	case String:
		return mapper.MapString(value)

	case Variable:
		return mapper.MapVariable(value)

	case Void:
		return mapper.MapVoid(value)

	default:
		panic("unreachable")
	}
}

func MapEachNode[T any](nodes []Node, mapper NodeMapper[T]) ([]T, error) {
	results := make([]T, 0, len(nodes))
	for _, node := range nodes {
		result, err := MapNode(node, mapper)
		if err != nil {
			return nil, err
		}

		results = append(results, result)
	}

	return results, nil
}

type NodeMapperNoError[T any] interface {
	MapAddToSet(value AddToSet) T

	MapArgumentDef(value ArgumentDef) T

	MapAssignment(value Assignment) T

	MapBlock(value Block) T

	MapBool(value Bool) T

	MapBreak(value Break) T

	MapCall(value Call) T

	MapConditional(value Conditional) T

	MapConstantDef(value ConstantDef) T

	MapContinue(value Continue) T

	MapDeclare(value Declare) T

	MapEmptyList(value EmptyList) T

	MapEqualOverride(value EqualOverride) T

	MapFieldDef(value FieldDef) T

	MapFor(value For) T

	MapForEach(value ForEach) T

	MapFunctionDef(value FunctionDef) T

	MapHashOverride(value HashOverride) T

	MapIf(value If) T

	MapInt64(value Int64) T

	MapKeyValue(value KeyValue) T

	MapLength(value Length) T

	MapList(value List) T

	MapLiteralBool(value LiteralBool) T

	MapLiteralInt(value LiteralInt) T

	MapLiteralList(value LiteralList) T

	MapLiteralMap(value LiteralMap) T

	MapLiteralRune(value LiteralRune) T

	MapLiteralSet(value LiteralSet) T

	MapLiteralString(value LiteralString) T

	MapLookup(value Lookup) T

	MapMap(value Map) T

	MapModel(value Model) T

	MapModelDef(value ModelDef) T

	MapModule(value Module) T

	MapNew(value New) T

	MapNull(value Null) T

	MapPop(value Pop) T

	MapProperty(value Property) T

	MapPush(value Push) T

	MapReturn(value Return) T

	MapRune(value Rune) T

	MapSelf(value Self) T

	MapSet(value Set) T

	MapSetContains(value SetContains) T

	MapString(value String) T

	MapVariable(value Variable) T

	MapVoid(value Void) T
}

func MapNodeNoError[T any](node Node, mapper NodeMapperNoError[T]) T {
	switch value := node.(type) {

	case AddToSet:
		return mapper.MapAddToSet(value)

	case ArgumentDef:
		return mapper.MapArgumentDef(value)

	case Assignment:
		return mapper.MapAssignment(value)

	case Block:
		return mapper.MapBlock(value)

	case Bool:
		return mapper.MapBool(value)

	case Break:
		return mapper.MapBreak(value)

	case Call:
		return mapper.MapCall(value)

	case Conditional:
		return mapper.MapConditional(value)

	case ConstantDef:
		return mapper.MapConstantDef(value)

	case Continue:
		return mapper.MapContinue(value)

	case Declare:
		return mapper.MapDeclare(value)

	case EmptyList:
		return mapper.MapEmptyList(value)

	case EqualOverride:
		return mapper.MapEqualOverride(value)

	case FieldDef:
		return mapper.MapFieldDef(value)

	case For:
		return mapper.MapFor(value)

	case ForEach:
		return mapper.MapForEach(value)

	case FunctionDef:
		return mapper.MapFunctionDef(value)

	case HashOverride:
		return mapper.MapHashOverride(value)

	case If:
		return mapper.MapIf(value)

	case Int64:
		return mapper.MapInt64(value)

	case KeyValue:
		return mapper.MapKeyValue(value)

	case Length:
		return mapper.MapLength(value)

	case List:
		return mapper.MapList(value)

	case LiteralBool:
		return mapper.MapLiteralBool(value)

	case LiteralInt:
		return mapper.MapLiteralInt(value)

	case LiteralList:
		return mapper.MapLiteralList(value)

	case LiteralMap:
		return mapper.MapLiteralMap(value)

	case LiteralRune:
		return mapper.MapLiteralRune(value)

	case LiteralSet:
		return mapper.MapLiteralSet(value)

	case LiteralString:
		return mapper.MapLiteralString(value)

	case Lookup:
		return mapper.MapLookup(value)

	case Map:
		return mapper.MapMap(value)

	case Model:
		return mapper.MapModel(value)

	case ModelDef:
		return mapper.MapModelDef(value)

	case Module:
		return mapper.MapModule(value)

	case New:
		return mapper.MapNew(value)

	case Null:
		return mapper.MapNull(value)

	case Pop:
		return mapper.MapPop(value)

	case Property:
		return mapper.MapProperty(value)

	case Push:
		return mapper.MapPush(value)

	case Return:
		return mapper.MapReturn(value)

	case Rune:
		return mapper.MapRune(value)

	case Self:
		return mapper.MapSelf(value)

	case Set:
		return mapper.MapSet(value)

	case SetContains:
		return mapper.MapSetContains(value)

	case String:
		return mapper.MapString(value)

	case Variable:
		return mapper.MapVariable(value)

	case Void:
		return mapper.MapVoid(value)

	default:
		panic("unreachable")
	}
}

func MapEachNodeNoError[T any](nodes []Node, mapper NodeMapperNoError[T]) []T {
	results := make([]T, 0, len(nodes))
	for _, node := range nodes {
		result := MapNodeNoError(node, mapper)
		results = append(results, result)
	}

	return results
}

type AssignableMapper[T any] interface {
	MapProperty(value Property) (T, error)

	MapVariable(value Variable) (T, error)
}

func MapAssignable[T any](node Assignable, mapper AssignableMapper[T]) (T, error) {
	switch value := node.(type) {

	case Property:
		return mapper.MapProperty(value)

	case Variable:
		return mapper.MapVariable(value)

	default:
		panic("unreachable")
	}
}

func MapEachAssignable[T any](nodes []Assignable, mapper AssignableMapper[T]) ([]T, error) {
	results := make([]T, 0, len(nodes))
	for _, node := range nodes {
		result, err := MapAssignable(node, mapper)
		if err != nil {
			return nil, err
		}

		results = append(results, result)
	}

	return results, nil
}

type AssignableMapperNoError[T any] interface {
	MapProperty(value Property) T

	MapVariable(value Variable) T
}

func MapAssignableNoError[T any](node Assignable, mapper AssignableMapperNoError[T]) T {
	switch value := node.(type) {

	case Property:
		return mapper.MapProperty(value)

	case Variable:
		return mapper.MapVariable(value)

	default:
		panic("unreachable")
	}
}

func MapEachAssignableNoError[T any](nodes []Assignable, mapper AssignableMapperNoError[T]) []T {
	results := make([]T, 0, len(nodes))
	for _, node := range nodes {
		result := MapAssignableNoError(node, mapper)
		results = append(results, result)
	}

	return results
}

type CallableMapper[T any] interface {
	MapFunctionDef(value FunctionDef) (T, error)
}

func MapCallable[T any](node Callable, mapper CallableMapper[T]) (T, error) {
	switch value := node.(type) {

	case FunctionDef:
		return mapper.MapFunctionDef(value)

	default:
		panic("unreachable")
	}
}

func MapEachCallable[T any](nodes []Callable, mapper CallableMapper[T]) ([]T, error) {
	results := make([]T, 0, len(nodes))
	for _, node := range nodes {
		result, err := MapCallable(node, mapper)
		if err != nil {
			return nil, err
		}

		results = append(results, result)
	}

	return results, nil
}

type CallableMapperNoError[T any] interface {
	MapFunctionDef(value FunctionDef) T
}

func MapCallableNoError[T any](node Callable, mapper CallableMapperNoError[T]) T {
	switch value := node.(type) {

	case FunctionDef:
		return mapper.MapFunctionDef(value)

	default:
		panic("unreachable")
	}
}

func MapEachCallableNoError[T any](nodes []Callable, mapper CallableMapperNoError[T]) []T {
	results := make([]T, 0, len(nodes))
	for _, node := range nodes {
		result := MapCallableNoError(node, mapper)
		results = append(results, result)
	}

	return results
}

type ConstantValueMapper[T any] interface {
	MapEmptyList(value EmptyList) (T, error)

	MapLiteralBool(value LiteralBool) (T, error)

	MapLiteralInt(value LiteralInt) (T, error)

	MapLiteralList(value LiteralList) (T, error)

	MapLiteralMap(value LiteralMap) (T, error)

	MapLiteralRune(value LiteralRune) (T, error)

	MapLiteralSet(value LiteralSet) (T, error)

	MapLiteralString(value LiteralString) (T, error)

	MapNull(value Null) (T, error)
}

func MapConstantValue[T any](node ConstantValue, mapper ConstantValueMapper[T]) (T, error) {
	switch value := node.(type) {

	case EmptyList:
		return mapper.MapEmptyList(value)

	case LiteralBool:
		return mapper.MapLiteralBool(value)

	case LiteralInt:
		return mapper.MapLiteralInt(value)

	case LiteralList:
		return mapper.MapLiteralList(value)

	case LiteralMap:
		return mapper.MapLiteralMap(value)

	case LiteralRune:
		return mapper.MapLiteralRune(value)

	case LiteralSet:
		return mapper.MapLiteralSet(value)

	case LiteralString:
		return mapper.MapLiteralString(value)

	case Null:
		return mapper.MapNull(value)

	default:
		panic("unreachable")
	}
}

func MapEachConstantValue[T any](nodes []ConstantValue, mapper ConstantValueMapper[T]) ([]T, error) {
	results := make([]T, 0, len(nodes))
	for _, node := range nodes {
		result, err := MapConstantValue(node, mapper)
		if err != nil {
			return nil, err
		}

		results = append(results, result)
	}

	return results, nil
}

type ConstantValueMapperNoError[T any] interface {
	MapEmptyList(value EmptyList) T

	MapLiteralBool(value LiteralBool) T

	MapLiteralInt(value LiteralInt) T

	MapLiteralList(value LiteralList) T

	MapLiteralMap(value LiteralMap) T

	MapLiteralRune(value LiteralRune) T

	MapLiteralSet(value LiteralSet) T

	MapLiteralString(value LiteralString) T

	MapNull(value Null) T
}

func MapConstantValueNoError[T any](node ConstantValue, mapper ConstantValueMapperNoError[T]) T {
	switch value := node.(type) {

	case EmptyList:
		return mapper.MapEmptyList(value)

	case LiteralBool:
		return mapper.MapLiteralBool(value)

	case LiteralInt:
		return mapper.MapLiteralInt(value)

	case LiteralList:
		return mapper.MapLiteralList(value)

	case LiteralMap:
		return mapper.MapLiteralMap(value)

	case LiteralRune:
		return mapper.MapLiteralRune(value)

	case LiteralSet:
		return mapper.MapLiteralSet(value)

	case LiteralString:
		return mapper.MapLiteralString(value)

	case Null:
		return mapper.MapNull(value)

	default:
		panic("unreachable")
	}
}

func MapEachConstantValueNoError[T any](nodes []ConstantValue, mapper ConstantValueMapperNoError[T]) []T {
	results := make([]T, 0, len(nodes))
	for _, node := range nodes {
		result := MapConstantValueNoError(node, mapper)
		results = append(results, result)
	}

	return results
}

type DefinitionMapper[T any] interface {
	MapArgumentDef(value ArgumentDef) (T, error)

	MapConstantDef(value ConstantDef) (T, error)

	MapDeclare(value Declare) (T, error)

	MapFieldDef(value FieldDef) (T, error)

	MapForEach(value ForEach) (T, error)
}

func MapDefinition[T any](node Definition, mapper DefinitionMapper[T]) (T, error) {
	switch value := node.(type) {

	case ArgumentDef:
		return mapper.MapArgumentDef(value)

	case ConstantDef:
		return mapper.MapConstantDef(value)

	case Declare:
		return mapper.MapDeclare(value)

	case FieldDef:
		return mapper.MapFieldDef(value)

	case ForEach:
		return mapper.MapForEach(value)

	default:
		panic("unreachable")
	}
}

func MapEachDefinition[T any](nodes []Definition, mapper DefinitionMapper[T]) ([]T, error) {
	results := make([]T, 0, len(nodes))
	for _, node := range nodes {
		result, err := MapDefinition(node, mapper)
		if err != nil {
			return nil, err
		}

		results = append(results, result)
	}

	return results, nil
}

type DefinitionMapperNoError[T any] interface {
	MapArgumentDef(value ArgumentDef) T

	MapConstantDef(value ConstantDef) T

	MapDeclare(value Declare) T

	MapFieldDef(value FieldDef) T

	MapForEach(value ForEach) T
}

func MapDefinitionNoError[T any](node Definition, mapper DefinitionMapperNoError[T]) T {
	switch value := node.(type) {

	case ArgumentDef:
		return mapper.MapArgumentDef(value)

	case ConstantDef:
		return mapper.MapConstantDef(value)

	case Declare:
		return mapper.MapDeclare(value)

	case FieldDef:
		return mapper.MapFieldDef(value)

	case ForEach:
		return mapper.MapForEach(value)

	default:
		panic("unreachable")
	}
}

func MapEachDefinitionNoError[T any](nodes []Definition, mapper DefinitionMapperNoError[T]) []T {
	results := make([]T, 0, len(nodes))
	for _, node := range nodes {
		result := MapDefinitionNoError(node, mapper)
		results = append(results, result)
	}

	return results
}

type StatementMapper[T any] interface {
	MapAddToSet(value AddToSet) (T, error)

	MapAssignment(value Assignment) (T, error)

	MapBreak(value Break) (T, error)

	MapCall(value Call) (T, error)

	MapConditional(value Conditional) (T, error)

	MapContinue(value Continue) (T, error)

	MapDeclare(value Declare) (T, error)

	MapFor(value For) (T, error)

	MapForEach(value ForEach) (T, error)

	MapPop(value Pop) (T, error)

	MapPush(value Push) (T, error)

	MapReturn(value Return) (T, error)
}

func MapStatement[T any](node Statement, mapper StatementMapper[T]) (T, error) {
	switch value := node.(type) {

	case AddToSet:
		return mapper.MapAddToSet(value)

	case Assignment:
		return mapper.MapAssignment(value)

	case Break:
		return mapper.MapBreak(value)

	case Call:
		return mapper.MapCall(value)

	case Conditional:
		return mapper.MapConditional(value)

	case Continue:
		return mapper.MapContinue(value)

	case Declare:
		return mapper.MapDeclare(value)

	case For:
		return mapper.MapFor(value)

	case ForEach:
		return mapper.MapForEach(value)

	case Pop:
		return mapper.MapPop(value)

	case Push:
		return mapper.MapPush(value)

	case Return:
		return mapper.MapReturn(value)

	default:
		panic("unreachable")
	}
}

func MapEachStatement[T any](nodes []Statement, mapper StatementMapper[T]) ([]T, error) {
	results := make([]T, 0, len(nodes))
	for _, node := range nodes {
		result, err := MapStatement(node, mapper)
		if err != nil {
			return nil, err
		}

		results = append(results, result)
	}

	return results, nil
}

type StatementMapperNoError[T any] interface {
	MapAddToSet(value AddToSet) T

	MapAssignment(value Assignment) T

	MapBreak(value Break) T

	MapCall(value Call) T

	MapConditional(value Conditional) T

	MapContinue(value Continue) T

	MapDeclare(value Declare) T

	MapFor(value For) T

	MapForEach(value ForEach) T

	MapPop(value Pop) T

	MapPush(value Push) T

	MapReturn(value Return) T
}

func MapStatementNoError[T any](node Statement, mapper StatementMapperNoError[T]) T {
	switch value := node.(type) {

	case AddToSet:
		return mapper.MapAddToSet(value)

	case Assignment:
		return mapper.MapAssignment(value)

	case Break:
		return mapper.MapBreak(value)

	case Call:
		return mapper.MapCall(value)

	case Conditional:
		return mapper.MapConditional(value)

	case Continue:
		return mapper.MapContinue(value)

	case Declare:
		return mapper.MapDeclare(value)

	case For:
		return mapper.MapFor(value)

	case ForEach:
		return mapper.MapForEach(value)

	case Pop:
		return mapper.MapPop(value)

	case Push:
		return mapper.MapPush(value)

	case Return:
		return mapper.MapReturn(value)

	default:
		panic("unreachable")
	}
}

func MapEachStatementNoError[T any](nodes []Statement, mapper StatementMapperNoError[T]) []T {
	results := make([]T, 0, len(nodes))
	for _, node := range nodes {
		result := MapStatementNoError(node, mapper)
		results = append(results, result)
	}

	return results
}

type TypeMapper[T any] interface {
	MapBool(value Bool) (T, error)

	MapInt64(value Int64) (T, error)

	MapList(value List) (T, error)

	MapMap(value Map) (T, error)

	MapModel(value Model) (T, error)

	MapRune(value Rune) (T, error)

	MapSet(value Set) (T, error)

	MapString(value String) (T, error)

	MapVoid(value Void) (T, error)
}

func MapType[T any](node Type, mapper TypeMapper[T]) (T, error) {
	switch value := node.(type) {

	case Bool:
		return mapper.MapBool(value)

	case Int64:
		return mapper.MapInt64(value)

	case List:
		return mapper.MapList(value)

	case Map:
		return mapper.MapMap(value)

	case Model:
		return mapper.MapModel(value)

	case Rune:
		return mapper.MapRune(value)

	case Set:
		return mapper.MapSet(value)

	case String:
		return mapper.MapString(value)

	case Void:
		return mapper.MapVoid(value)

	default:
		panic("unreachable")
	}
}

func MapEachType[T any](nodes []Type, mapper TypeMapper[T]) ([]T, error) {
	results := make([]T, 0, len(nodes))
	for _, node := range nodes {
		result, err := MapType(node, mapper)
		if err != nil {
			return nil, err
		}

		results = append(results, result)
	}

	return results, nil
}

type TypeMapperNoError[T any] interface {
	MapBool(value Bool) T

	MapInt64(value Int64) T

	MapList(value List) T

	MapMap(value Map) T

	MapModel(value Model) T

	MapRune(value Rune) T

	MapSet(value Set) T

	MapString(value String) T

	MapVoid(value Void) T
}

func MapTypeNoError[T any](node Type, mapper TypeMapperNoError[T]) T {
	switch value := node.(type) {

	case Bool:
		return mapper.MapBool(value)

	case Int64:
		return mapper.MapInt64(value)

	case List:
		return mapper.MapList(value)

	case Map:
		return mapper.MapMap(value)

	case Model:
		return mapper.MapModel(value)

	case Rune:
		return mapper.MapRune(value)

	case Set:
		return mapper.MapSet(value)

	case String:
		return mapper.MapString(value)

	case Void:
		return mapper.MapVoid(value)

	default:
		panic("unreachable")
	}
}

func MapEachTypeNoError[T any](nodes []Type, mapper TypeMapperNoError[T]) []T {
	results := make([]T, 0, len(nodes))
	for _, node := range nodes {
		result := MapTypeNoError(node, mapper)
		results = append(results, result)
	}

	return results
}

type ValueMapper[T any] interface {
	MapCall(value Call) (T, error)

	MapEmptyList(value EmptyList) (T, error)

	MapLength(value Length) (T, error)

	MapLiteralBool(value LiteralBool) (T, error)

	MapLiteralInt(value LiteralInt) (T, error)

	MapLiteralList(value LiteralList) (T, error)

	MapLiteralMap(value LiteralMap) (T, error)

	MapLiteralRune(value LiteralRune) (T, error)

	MapLiteralSet(value LiteralSet) (T, error)

	MapLiteralString(value LiteralString) (T, error)

	MapLookup(value Lookup) (T, error)

	MapNew(value New) (T, error)

	MapNull(value Null) (T, error)

	MapPop(value Pop) (T, error)

	MapProperty(value Property) (T, error)

	MapSelf(value Self) (T, error)

	MapSetContains(value SetContains) (T, error)

	MapVariable(value Variable) (T, error)
}

func MapValue[T any](node Value, mapper ValueMapper[T]) (T, error) {
	switch value := node.(type) {

	case Call:
		return mapper.MapCall(value)

	case EmptyList:
		return mapper.MapEmptyList(value)

	case Length:
		return mapper.MapLength(value)

	case LiteralBool:
		return mapper.MapLiteralBool(value)

	case LiteralInt:
		return mapper.MapLiteralInt(value)

	case LiteralList:
		return mapper.MapLiteralList(value)

	case LiteralMap:
		return mapper.MapLiteralMap(value)

	case LiteralRune:
		return mapper.MapLiteralRune(value)

	case LiteralSet:
		return mapper.MapLiteralSet(value)

	case LiteralString:
		return mapper.MapLiteralString(value)

	case Lookup:
		return mapper.MapLookup(value)

	case New:
		return mapper.MapNew(value)

	case Null:
		return mapper.MapNull(value)

	case Pop:
		return mapper.MapPop(value)

	case Property:
		return mapper.MapProperty(value)

	case Self:
		return mapper.MapSelf(value)

	case SetContains:
		return mapper.MapSetContains(value)

	case Variable:
		return mapper.MapVariable(value)

	default:
		panic("unreachable")
	}
}

func MapEachValue[T any](nodes []Value, mapper ValueMapper[T]) ([]T, error) {
	results := make([]T, 0, len(nodes))
	for _, node := range nodes {
		result, err := MapValue(node, mapper)
		if err != nil {
			return nil, err
		}

		results = append(results, result)
	}

	return results, nil
}

type ValueMapperNoError[T any] interface {
	MapCall(value Call) T

	MapEmptyList(value EmptyList) T

	MapLength(value Length) T

	MapLiteralBool(value LiteralBool) T

	MapLiteralInt(value LiteralInt) T

	MapLiteralList(value LiteralList) T

	MapLiteralMap(value LiteralMap) T

	MapLiteralRune(value LiteralRune) T

	MapLiteralSet(value LiteralSet) T

	MapLiteralString(value LiteralString) T

	MapLookup(value Lookup) T

	MapNew(value New) T

	MapNull(value Null) T

	MapPop(value Pop) T

	MapProperty(value Property) T

	MapSelf(value Self) T

	MapSetContains(value SetContains) T

	MapVariable(value Variable) T
}

func MapValueNoError[T any](node Value, mapper ValueMapperNoError[T]) T {
	switch value := node.(type) {

	case Call:
		return mapper.MapCall(value)

	case EmptyList:
		return mapper.MapEmptyList(value)

	case Length:
		return mapper.MapLength(value)

	case LiteralBool:
		return mapper.MapLiteralBool(value)

	case LiteralInt:
		return mapper.MapLiteralInt(value)

	case LiteralList:
		return mapper.MapLiteralList(value)

	case LiteralMap:
		return mapper.MapLiteralMap(value)

	case LiteralRune:
		return mapper.MapLiteralRune(value)

	case LiteralSet:
		return mapper.MapLiteralSet(value)

	case LiteralString:
		return mapper.MapLiteralString(value)

	case Lookup:
		return mapper.MapLookup(value)

	case New:
		return mapper.MapNew(value)

	case Null:
		return mapper.MapNull(value)

	case Pop:
		return mapper.MapPop(value)

	case Property:
		return mapper.MapProperty(value)

	case Self:
		return mapper.MapSelf(value)

	case SetContains:
		return mapper.MapSetContains(value)

	case Variable:
		return mapper.MapVariable(value)

	default:
		panic("unreachable")
	}
}

func MapEachValueNoError[T any](nodes []Value, mapper ValueMapperNoError[T]) []T {
	results := make([]T, 0, len(nodes))
	for _, node := range nodes {
		result := MapValueNoError(node, mapper)
		results = append(results, result)
	}

	return results
}
