package populate_metadata_mapper

import (
	"github.com/JosephNaberhaus/agnostic/code"
	"github.com/JosephNaberhaus/agnostic/internal/utils/stack"
)

var _ code.NodeMapperOnlyError = Mapper{}

type Mapper struct {
	Stack stack.Stack[code.Node]
}

func (m Mapper) MapAddToSet(value code.AddToSet) error {
	return nil
}

func (m Mapper) MapArgumentDef(value code.ArgumentDef) error {
	return nil
}

func (m Mapper) MapAssignment(value code.Assignment) error {
	return nil
}

func (m Mapper) MapBlock(value code.Block) error {
	return nil
}

func (m Mapper) MapBool(value code.Bool) error {
	return nil
}

func (m Mapper) MapBreak(value code.Break) error {
	return nil
}

func (m Mapper) MapCall(value code.Call) error {
	return nil
}

func (m Mapper) MapConditional(value code.Conditional) error {
	return nil
}

func (m Mapper) MapConstantDef(value code.ConstantDef) error {
	return nil
}

func (m Mapper) MapContinue(value code.Continue) error {
	return nil
}

func (m Mapper) MapDeclare(value code.Declare) error {
	return nil
}

func (m Mapper) MapEmptyList(value code.EmptyList) error {
	return nil
}

func (m Mapper) MapEqualOverride(value code.EqualOverride) error {
	return nil
}

func (m Mapper) MapFieldDef(value code.FieldDef) error {
	return nil
}

func (m Mapper) MapFor(value code.For) error {
	return nil
}

func (m Mapper) MapForEach(value code.ForEach) error {
	return nil
}

func (m Mapper) MapFunctionDef(value code.FunctionDef) error {
	return nil
}

func (m Mapper) MapHashOverride(value code.HashOverride) error {
	return nil
}

func (m Mapper) MapIf(value code.If) error {
	return nil
}

func (m Mapper) MapInt64(value code.Int64) error {
	return nil
}

func (m Mapper) MapKeyValue(value code.KeyValue) error {
	return nil
}

func (m Mapper) MapLength(value code.Length) error {
	return nil
}

func (m Mapper) MapList(value code.List) error {
	return nil
}

func (m Mapper) MapLiteralBool(value code.LiteralBool) error {
	return nil
}

func (m Mapper) MapLiteralInt64(value code.LiteralInt64) error {
	return nil
}

func (m Mapper) MapLiteralList(value code.LiteralList) error {
	return nil
}

func (m Mapper) MapLiteralMap(value code.LiteralMap) error {
	return nil
}

func (m Mapper) MapLiteralRune(value code.LiteralRune) error {
	return nil
}

func (m Mapper) MapLiteralSet(value code.LiteralSet) error {
	return nil
}

func (m Mapper) MapLiteralString(value code.LiteralString) error {
	return nil
}

func (m Mapper) MapLookup(value code.Lookup) error {
	return nil
}

func (m Mapper) MapMap(value code.Map) error {
	return nil
}

func (m Mapper) MapModel(value code.Model) error {
	return nil
}

func (m Mapper) MapModelDef(value code.ModelDef) error {
	return nil
}

func (m Mapper) MapModule(value code.Module) error {
	return nil
}

func (m Mapper) MapNew(value code.New) error {
	return nil
}

func (m Mapper) MapNil(value code.Nil) error {
	return nil
}

func (m Mapper) MapPop(value code.Pop) error {
	return nil
}

func (m Mapper) MapProperty(value code.Property) error {
	return nil
}

func (m Mapper) MapPush(value code.Push) error {
	return nil
}

func (m Mapper) MapReturn(value code.Return) error {
	return nil
}

func (m Mapper) MapRoot(value code.Root) error {
	return nil
}

func (m Mapper) MapRune(value code.Rune) error {
	return nil
}

func (m Mapper) MapSelf(value code.Self) error {
	return nil
}

func (m Mapper) MapSet(value code.Set) error {
	return nil
}

func (m Mapper) MapSetContains(value code.SetContains) error {
	return nil
}

func (m Mapper) MapString(value code.String) error {
	return nil
}

func (m Mapper) MapVariable(value code.Variable) error {
	return nil
}

func (m Mapper) MapVoid(value code.Void) error {
	return nil
}
