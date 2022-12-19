package java

import (
	"fmt"
	"github.com/JosephNaberhaus/agnostic/code"
	"github.com/JosephNaberhaus/agnostic/code/mappers/has_node_of_type"
	"github.com/JosephNaberhaus/agnostic/implementation"
	"github.com/JosephNaberhaus/agnostic/implementation/text"
)

type Mapper struct{}

func (m *Mapper) Config() implementation.Config {
	return implementation.Config{
		Indent: "    ",
	}
}

func (m *Mapper) MapLiteralInt(original *code.LiteralInt) text.Node {
	return text.Span(fmt.Sprintf("%dL", original.Value))
}

func (m *Mapper) MapLiteralString(original *code.LiteralString) text.Node {
	return text.Span(fmt.Sprintf("\"%s\"", original.Value))
}

func (m *Mapper) MapLiteralRune(original *code.LiteralRune) text.Node {
	var preview string
	switch original.Value {
	case '\n':
		preview = "\\n"
	default:
		preview = string(original.Value)
	}

	return text.Span(fmt.Sprintf("%d /* '%s' */", int(original.Value), preview))
}

func (m *Mapper) MapFieldDef(original *code.FieldDef) text.Node {
	typeNode := code.MapTypeNoError[text.Node](original.Type, m)

	return text.Group{
		typeNode,
		text.Span(" "),
		text.Span(original.Name),
		text.Span(";"),
	}
}

func (m *Mapper) MapArgumentDef(original *code.ArgumentDef) text.Node {
	typeNode := code.MapTypeNoError[text.Node](original.Type, m)

	return text.Group{
		typeNode,
		text.Span(" "),
		text.Span(original.Name),
	}
}

func (m *Mapper) MapModelDef(original *code.ModelDef) text.Node {
	var constructorParams []text.Node
	var constructorAssignments []text.Node
	for _, field := range original.Fields {
		fieldType := code.MapTypeNoError[text.Node](field.Type, m)

		constructorParams = append(constructorParams, text.Group{
			fieldType,
			text.Span(" "),
			text.Span(field.Name),
		})

		constructorAssignments = append(constructorAssignments, text.Group{
			text.Span("this."),
			text.Span(field.Name),
			text.Span(" = "),
			text.Span(field.Name),
			text.Span(";"),
		})
	}

	fieldNodes := code.MapNodesNoError[text.Node](original.Fields, m)
	methodNodes := code.MapNodesNoError[text.Node](original.Methods, m)

	var equalOverride text.Node
	if original.EqualOverride.IsSet() {
		equalOverride = m.MapEqualOverride(original.EqualOverride.Value())
	}

	var hashOverride text.Node
	if original.HashOverride.IsSet() {
		hashOverride = m.MapHashOverride(original.HashOverride.Value())
	}

	return text.Block{
		text.Group{
			text.Span("class "),
			text.Span(original.Name),
			text.Span(" {"),
		},
		text.IndentedBlock{
			text.Group{
				text.Span(original.Name),
				text.Span("("),
				text.Join{
					Nodes: constructorParams,
					Sep:   ", ",
				},
				text.Span(") {"),
			},
			text.IndentedBlock(constructorAssignments),
			text.Span("}"),
			text.Span(""),
			text.Block(fieldNodes),
			text.Span(""),
			text.Block(methodNodes),
			equalOverride,
			hashOverride,
		},
		text.Span("}"),
	}
}

func (m *Mapper) MapModule(original *code.Module) text.Node {
	modelNodes := code.MapNodesNoError[text.Node](original.Models, m)
	functionNodes := code.MapNodesNoError[text.Node](original.Functions, m)
	constantNodes := code.MapNodesNoError[text.Node](original.Constants, m)

	// TODO: this could be simplified
	objectsImport := text.Span("import java.util.Objects;")

	var importArrays text.Node
	var importArrayListNode text.Node
	if code.MapNodeNoError[bool](original, has_node_of_type.Mapper[*code.LiteralList]{}) {
		importArrays = text.Span("import java.util.Arrays;")
		importArrayListNode = text.Span("import java.util.ArrayList;")
	} else if code.MapNodeNoError[bool](original, has_node_of_type.Mapper[*code.List]{}) || code.MapNodeNoError[bool](original, has_node_of_type.Mapper[*code.EmptyList]{}) {
		importArrayListNode = text.Span("import java.util.ArrayList;")
	}

	var importsMapNode text.Node
	var importHashmapNode text.Node
	if code.MapNodeNoError[bool](original, has_node_of_type.Mapper[*code.LiteralMap]{}) {
		importArrays = text.Span("import java.util.Map;")
		importHashmapNode = text.Span("import java.util.HashMap;")
	} else if code.MapNodeNoError[bool](original, has_node_of_type.Mapper[*code.Map]{}) {
		importHashmapNode = text.Span("import java.util.HashMap;")
	}

	var importSetNode text.Node
	var importHashSetNode text.Node
	if code.MapNodeNoError[bool](original, has_node_of_type.Mapper[*code.LiteralSet]{}) {
		importArrays = text.Span("import java.util.Set;")
		importHashmapNode = text.Span("import java.util.HashSet;")
	} else if code.MapNodeNoError[bool](original, has_node_of_type.Mapper[*code.Set]{}) || code.MapNodeNoError[bool](original, has_node_of_type.Mapper[*code.EmptySet]{}) {
		importHashmapNode = text.Span("import java.util.HashSet;")
	}

	return text.Block{
		importArrays,
		importArrayListNode,
		importsMapNode,
		importHashmapNode,
		importSetNode,
		importHashSetNode,
		objectsImport,
		text.Span(fmt.Sprintf("class %sFunctions {", original.Name)),
		text.IndentedBlock(functionNodes),
		text.Span("}"),
		text.Span(""),
		text.Span(fmt.Sprintf("class %sConstants {", original.Name)),
		text.IndentedBlock(constantNodes),
		text.Span("}"),
		text.Span(""),
		text.Block(modelNodes),
	}
}

func (m *Mapper) MapUnaryOperator(original code.UnaryOperator) text.Node {
	switch original {
	case code.Not:
		return text.Span("!")
	case code.Negate:
		return text.Span("-")
	}

	// TODO remove the need for this
	panic("oh no")
}

func (m *Mapper) MapUnaryOperation(original *code.UnaryOperation) text.Node {
	valueNode := code.MapValueNoError[text.Node](original.Value, m)

	switch original.Operator {
	case code.CastToInt:
		return text.Group{
			text.Span("(long)("),
			valueNode,
			text.Span(")"),
		}
	case code.CastToString:
		return text.Group{
			text.Span("new String(Character.toChars("),
			valueNode,
			text.Span("))"),
		}
	case code.Hash:
		return text.Group{
			text.Span("Objects.hashCode("),
			valueNode,
			text.Span(")"),
		}
	}

	operatorNode := m.MapUnaryOperator(original.Operator)

	return text.Group{
		operatorNode,
		valueNode,
	}
}

func (m *Mapper) MapBinaryOperator(original code.BinaryOperator) text.Node {
	switch original {
	case code.Multiply:
		return text.Span("*")
	case code.Divide:
		return text.Span("/")
	case code.Modulo:
		return text.Span("%")
	case code.Add:
		return text.Span("+")
	case code.Subtract:
		return text.Span("-")
	case code.NotEqual:
		return text.Span("!=")
	case code.LessThan:
		return text.Span("<")
	case code.LessThanOrEqualTo:
		return text.Span("<=")
	case code.GreaterThan:
		return text.Span(">")
	case code.GreaterThanOrEqualTo:
		return text.Span(">=")
	case code.Or:
		return text.Span("||")
	case code.And:
		return text.Span("&&")
	}

	// TODO remove the need for this
	panic("oh no")
}

func (m *Mapper) MapBinaryOperation(original *code.BinaryOperation) text.Node {
	leftNode := code.MapValueNoError[text.Node](original.Left, m)
	rightNode := code.MapValueNoError[text.Node](original.Right, m)

	if original.Operator == code.Equal {
		return text.Group{
			text.Span("Objects.equals("),
			leftNode,
			text.Span(", "),
			rightNode,
			text.Span(")"),
		}

	}

	operatorNode := m.MapBinaryOperator(original.Operator)

	return text.Group{
		text.Span("("),
		leftNode,
		text.Span(" "),
		operatorNode,
		text.Span(" "),
		rightNode,
		text.Span(")"),
	}
}

func (m *Mapper) MapAssignment(original *code.Assignment) text.Node {
	fromNode := code.MapValueNoError[text.Node](original.From, m)

	if lookup, ok := original.To.(*code.Lookup); ok {
		lookupFromNode := code.MapValueNoError[text.Node](lookup.From, m)
		lookupKeyNode := code.MapValueNoError[text.Node](lookup.Key, m)

		switch lookup.LookupType {
		case code.LookupTypeList:
			return text.Group{
				lookupFromNode,
				text.Span(".set(Long.valueOf("),
				lookupKeyNode,
				text.Span(").intValue(), "),
				fromNode,
				text.Span(")"),
			}
		case code.LookupTypeMap:
			return text.Group{
				lookupFromNode,
				text.Span(".put("),
				lookupKeyNode,
				text.Span(", "),
				fromNode,
				text.Span(")"),
			}
		}

		panic("unreachable")
	}

	toNode := code.MapValueNoError[text.Node](original.To, m)

	return text.Group{
		toNode,
		text.Span(" = "),
		fromNode,
	}
}

func (m *Mapper) MapModel(original *code.Model) text.Node {
	return text.Span(original.Name)
}

func (m *Mapper) MapPrimitive(original code.Primitive) text.Node {
	switch original {
	case code.Boolean:
		return text.Span("Boolean")
	case code.Int:
		return text.Span("Long")
	case code.Rune:
		return text.Span("Integer")
	case code.String:
		return text.Span("String")
	case code.Void:
		return text.Span("void")
	}

	// TODO remove the need for this
	panic("no!!!")
}

func (m *Mapper) MapIf(original *code.If) text.Node {
	conditionNode := code.MapValueNoError[text.Node](original.Condition, m)
	blockNode := m.MapBlock(original.Block)

	return text.Block{
		text.Group{
			text.Span("if ("),
			conditionNode,
			text.Span(") {"),
		},
		blockNode,
		// The closing "}" is added by MapConditional
	}
}

func (m *Mapper) MapElseIf(original *code.ElseIf) text.Node {
	conditionNode := code.MapValueNoError[text.Node](original.Condition, m)
	blockNode := m.MapBlock(original.Block)

	return text.Block{
		text.Group{
			text.Span("} else if ("),
			conditionNode,
			text.Span(") {"),
		},
		blockNode,
		// The closing "}" is added by MapConditional
	}
}

func (m *Mapper) MapElse(original *code.Else) text.Node {
	blockNode := m.MapBlock(original.Block)

	return text.Block{
		text.Span("} else {"),
		blockNode,
		// The closing "}" is added by MapConditional
	}
}

func (m *Mapper) MapConditional(original *code.Conditional) text.Node {
	ifNode := m.MapIf(original.If)

	var elseIfNode text.Node
	if len(original.ElseIfs) != 0 {
		elseIfNode = text.Group(code.MapNodesNoError[text.Node](original.ElseIfs, m))
	}

	var elseNode text.Node
	if original.Else != nil {
		elseNode = m.MapElse(original.Else)
	}

	return text.Block{
		ifNode,
		elseIfNode,
		elseNode,
		text.Span("}"),
	}
}

func (m *Mapper) MapProperty(original *code.Property) text.Node {
	ofNode := code.MapValueNoError[text.Node](original.Of, m)

	return text.Group{
		ofNode,
		text.Span("."),
		text.Span(original.Name),
	}
}

func (m *Mapper) MapVariable(original *code.Variable) text.Node {
	if original.IsConstant {
		return text.Group{
			text.Span(original.Definition.(*code.ConstantDef).ParentModule.Name),
			text.Span("Constants."),
			text.Span(original.Name),
		}
	}

	return text.Span(original.Name)
}

func (m *Mapper) MapList(original *code.List) text.Node {
	baseNode := code.MapTypeNoError[text.Node](original.Base, m)

	return text.Group{
		text.Span("ArrayList<"),
		baseNode,
		text.Span(">"),
	}
}

func (m *Mapper) MapMap(original *code.Map) text.Node {
	keyNode := code.MapTypeNoError[text.Node](original.Key, m)
	valueNode := code.MapTypeNoError[text.Node](original.Value, m)

	return text.Group{
		text.Span("HashMap<"),
		keyNode,
		text.Span(", "),
		valueNode,
		text.Span(">"),
	}
}

func (m *Mapper) MapLookup(original *code.Lookup) text.Node {
	fromNode := code.MapValueNoError[text.Node](original.From, m)
	keyNode := code.MapValueNoError[text.Node](original.Key, m)

	switch original.LookupType {
	case code.LookupTypeList:
		return text.Group{
			fromNode,
			text.Span(".get(Long.valueOf("),
			keyNode,
			text.Span(").intValue())"),
		}
	case code.LookupTypeMap:
		return text.Group{
			fromNode,
			text.Span(".get("),
			keyNode,
			text.Span(")"),
		}
	case code.LookupTypeString:
		return text.Group{
			fromNode,
			text.Span(".codePointAt("),
			fromNode,
			text.Span(".offsetByCodePoints(0, Long.valueOf("),
			keyNode,
			text.Span(").intValue()))"),
		}
	}

	panic("unreachable")
}

func (m *Mapper) MapFunctionDef(original *code.FunctionDef) text.Node {
	returnTypeNode := code.MapTypeNoError[text.Node](original.ReturnType, m)

	var modifierNode text.Node
	if !original.IsMethod {
		modifierNode = text.Span("static ")
	}

	argumentNodes := code.MapNodesNoError[text.Node](original.Arguments, m)
	blockNode := m.MapBlock(original.Block)

	return text.Block{
		text.Group{
			text.Span("public "),
			modifierNode,
			returnTypeNode,
			text.Span(" "),
			text.Span(original.Name),
			text.Span("("),
			text.Join{
				Nodes: argumentNodes,
				Sep:   ", ",
			},
			text.Span(") {"),
		},
		blockNode,
		text.Span("}"),
	}
}

func (m *Mapper) MapReturn(original *code.Return) text.Node {
	valueNode := code.MapValueNoError[text.Node](original.Value, m)

	return text.Group{
		text.Span("return "),
		valueNode,
	}
}

func (m *Mapper) MapCall(original *code.Call) text.Node {
	functionNode := code.MapCallableNoError[text.Node](original.Function, m)
	argumentNodes := code.MapNodesNoError[text.Node](original.Arguments, m)

	return text.Group{
		functionNode,
		text.Span("("),
		text.Join{
			Nodes: argumentNodes,
			Sep:   ", ",
		},
		text.Span(")"),
	}
}

func (m *Mapper) MapFunction(original *code.Function) text.Node {
	if !original.Definition.IsMethod {
		return text.Group{
			text.Span(original.ParentModule.Name),
			text.Span("Functions."),
			text.Span(original.Name),
		}
	}

	return text.Span(original.Name)
}

func (m *Mapper) MapFunctionProperty(original *code.FunctionProperty) text.Node {
	ofNode := code.MapValueNoError[text.Node](original.Of, m)

	return text.Group{
		ofNode,
		text.Span("."),
		text.Span(original.Name),
	}
}

func (m *Mapper) MapNew(original *code.New) text.Node {
	modelNode := m.MapModel(original.Model)

	return text.Group{
		text.Span("(new "),
		modelNode,
		text.Span("())"),
	}
}

func (m *Mapper) MapDeclare(original *code.Declare) text.Node {
	valueNode := code.MapValueNoError[text.Node](original.Value, m)
	typeNode := code.MapTypeNoError[text.Node](original.Type, m)

	return text.Group{
		typeNode,
		text.Span(" "),
		text.Span(original.Name),
		text.Span(" = "),
		valueNode,
	}
}

func (m *Mapper) MapFor(original *code.For) text.Node {
	var initializationNode text.Node
	if original.Initialization.IsSet() {
		initializationNode = code.MapStatementNoError[text.Node](original.Initialization.Value(), m)
	}

	conditionNode := code.MapValueNoError[text.Node](original.Condition, m)

	var afterEachNode text.Node
	if original.Initialization.IsSet() {
		afterEachNode = code.MapStatementNoError[text.Node](original.AfterEach.Value(), m)
	}

	blockNode := m.MapBlock(original.Block)

	return text.Block{
		text.Group{
			text.Span("for ("),
			initializationNode,
			text.Span("; "),
			conditionNode,
			text.Span("; "),
			afterEachNode,
			text.Span(") {"),
		},
		blockNode,
		text.Span("}"),
	}
}

func (m *Mapper) MapForIn(original *code.ForIn) text.Node {
	itemTypeNode := code.MapTypeNoError[text.Node](original.ItemType, m)
	iterableNode := code.MapValueNoError[text.Node](original.Iterable, m)
	blockNode := m.MapBlock(original.Block)

	return text.Block{
		text.Group{
			text.Span("for ("),
			itemTypeNode,
			text.Span(" "),
			text.Span(original.ItemName),
			text.Span(" : "),
			iterableNode,
			text.Span(") {"),
		},
		blockNode,
		text.Span("}"),
	}
}

func (m *Mapper) MapBlock(original *code.Block) text.Node {
	var statementNodes []text.Node
	for _, statement := range original.Statements {
		statementNode := code.MapStatementNoError[text.Node](statement, m)

		includeSemicolon := true
		switch statement.(type) {
		case *code.For, *code.ForIn, *code.Conditional:
			includeSemicolon = false
		}

		if includeSemicolon {
			statementNodes = append(statementNodes, text.Group{
				statementNode,
				text.Span(";"),
			})
		} else {
			statementNodes = append(statementNodes, statementNode)
		}
	}

	return text.IndentedBlock(statementNodes)
}

func (m *Mapper) MapLiteralList(original *code.LiteralList) text.Node {
	valueNodes := code.MapValuesNoError[text.Node](original.Items, m)

	return text.Group{
		text.Span("("),
		text.Span("new ArrayList(Arrays.asList("),
		text.Join{
			Nodes: valueNodes,
			Sep:   ", ",
		},
		text.Span(")))"),
	}
}

func (m *Mapper) MapLength(original *code.Length) text.Node {
	valueNode := code.MapValueNoError[text.Node](original.Value, m)

	switch original.LengthType {
	case code.LengthTypeString:
		return text.Group{
			text.Span("Long.valueOf("),
			valueNode,
			text.Span(".codePointCount(0, "),
			valueNode,
			text.Span(".length()))"),
		}
	case code.LengthTypeList, code.LengthTypeMap, code.LengthTypeSet:
		return text.Group{
			text.Span("Long.valueOf("),
			valueNode,
			text.Span(".size())"),
		}
	}

	panic("unreachable")
}

func (m *Mapper) MapConstantDef(original *code.ConstantDef) text.Node {
	typeNode := code.MapTypeNoError[text.Node](original.Type, m)
	valueNode := code.MapValueNoError[text.Node](original.Value, m)

	return text.Group{
		text.Span("public static "),
		typeNode,
		text.Span(" "),
		text.Span(original.Name),
		text.Span(" = "),
		valueNode,
		text.Span(";"),
	}
}

func (m *Mapper) MapKeyValue(original *code.KeyValue) text.Node {
	keyNode := code.MapValueNoError[text.Node](original.Key, m)
	valueNode := code.MapValueNoError[text.Node](original.Value, m)

	return text.Group{
		text.Span("Map.entry("),
		keyNode,
		text.Span(", "),
		valueNode,
		text.Span(")"),
	}
}

func (m *Mapper) MapLiteralMap(original *code.LiteralMap) text.Node {
	entryNodes := code.MapNodesNoError[text.Node](original.Entries, m)

	return text.Group{
		text.Span("new HashMap("),
		text.Span("Map.ofEntries("),
		text.Join{
			Nodes: entryNodes,
			Sep:   ",",
		},
		text.Span("))"),
	}
}

func (m *Mapper) MapLiteralSet(original *code.LiteralSet) text.Node {
	itemNodes := code.MapValuesNoError[text.Node](original.Items, m)

	return text.Group{
		text.Span("new HashSet(Set.of("),
		text.Join{
			Nodes: itemNodes,
			Sep:   ",",
		},
		text.Span("))"),
	}
}

func (m *Mapper) MapSet(original *code.Set) text.Node {
	baseNode := code.MapTypeNoError[text.Node](original.Base, m)

	return text.Group{
		text.Span("HashSet<"),
		baseNode,
		text.Span(">"),
	}
}

func (m *Mapper) MapEmptyList(original *code.EmptyList) text.Node {
	typeNode := code.MapTypeNoError[text.Node](original.Type, m)

	return text.Group{
		text.Span("(new ArrayList<"),
		typeNode,
		text.Span(">())"),
	}
}

func (m *Mapper) MapEmptySet(original *code.EmptySet) text.Node {
	typeNode := code.MapTypeNoError[text.Node](original.Type, m)

	return text.Group{
		text.Span("(new HashSet<"),
		typeNode,
		text.Span(">())"),
	}
}

func (m *Mapper) MapAddToSet(original *code.AddToSet) text.Node {
	toNode := code.MapValueNoError[text.Node](original.To, m)
	valueNode := code.MapValueNoError[text.Node](original.Value, m)

	return text.Group{
		toNode,
		text.Span(".add("),
		valueNode,
		text.Span(")"),
	}
}

func (m *Mapper) MapSetContains(original *code.SetContains) text.Node {
	setNode := code.MapValueNoError[text.Node](original.Set, m)

	valueNode := code.MapValueNoError[text.Node](original.Value, m)

	return text.Group{
		setNode,
		text.Span(".contains("),
		valueNode,
		text.Span(")"),
	}
}

func (m *Mapper) MapPush(original *code.Push) text.Node {
	toNode := code.MapValueNoError[text.Node](original.To, m)
	valueNode := code.MapValueNoError[text.Node](original.Value, m)

	return text.Group{
		toNode,
		text.Span(".add("),
		valueNode,
		text.Span(")"),
	}
}

func (m *Mapper) MapPop(original *code.Pop) text.Node {
	valueNode := code.MapValueNoError[text.Node](original.Value, m)

	return text.Group{
		valueNode,
		text.Span(".remove("),
		valueNode,
		text.Span(".size() - 1"),
		text.Span(")"),
	}
}

func (m *Mapper) MapLiteralBool(original *code.LiteralBool) text.Node {
	if original.Value {
		return text.Span("true")
	}

	return text.Span("false")
}

func (m *Mapper) MapNull(original *code.Null) text.Node {
	return text.Span("null")
}

func (m *Mapper) MapSelf(original *code.Self) text.Node {
	return text.Span("this")
}

func (m *Mapper) MapDeclareNull(original *code.DeclareNull) text.Node {
	typeNode := code.MapTypeNoError[text.Node](original.Type, m)

	return text.Group{
		typeNode,
		text.Span(" "),
		text.Span(original.Name),
		text.Span(" = null"),
	}
}

func (m *Mapper) MapBreak(original *code.Break) text.Node {
	return text.Span("break")
}

func (m *Mapper) MapContinue(original *code.Continue) text.Node {
	return text.Span("continue")
}

func (m *Mapper) MapEqualOverride(original *code.EqualOverride) text.Node {
	var blockNode = m.MapBlock(original.Block)

	return text.Block{
		text.Span("@Override"),
		text.Span("public boolean equals(Object otherObj) {"),
		text.IndentedBlock{
			text.Span("if (otherObj == null) {"),
			text.IndentedBlock{
				text.Span("return false;"),
			},
			text.Span("}"),
			text.Span(""),
			text.Span(fmt.Sprintf("if (!(otherObj instanceof %s)) {", original.ParentModel.Name)),
			text.IndentedBlock{
				text.Span("return false;"),
			},
			text.Span("}"),
			text.Span(""),
			text.Span(fmt.Sprintf("%s %s = (%s) otherObj;", original.ParentModel.Name, original.OtherName, original.ParentModel.Name)),
			text.Span(""),
		},
		blockNode,
		text.Span("}"),
	}
}

func (m *Mapper) MapHashOverride(original *code.HashOverride) text.Node {
	var blockNode = m.MapBlock(original.Block)

	return text.Block{
		text.Span("@Override"),
		text.Span("public int hashCode() {"),
		blockNode,
		text.Span("}"),
	}
}

func (m *Mapper) MapLiteralProperty(original *code.LiteralProperty) text.Node {
	panic("unreachable")
}

func (m *Mapper) MapLiteralStruct(original *code.LiteralStruct) text.Node {
	properties := map[string]text.Node{}
	for _, property := range original.Properties {
		properties[property.Name] = code.MapValueNoError[text.Node](property.Value, m)
	}

	var constructorValues []text.Node
	for _, field := range original.Definition.Fields {
		if value, ok := properties[field.Name]; ok {
			constructorValues = append(constructorValues, value)
		} else {
			constructorValues = append(constructorValues, text.Span("null"))
		}
	}

	return text.Group{
		text.Span("new "),
		text.Span(original.Name),
		text.Span("("),
		text.Join{
			Nodes: constructorValues,
			Sep:   ", ",
		},
		text.Span(")"),
	}
}
