package lexer

import (
	"errors"
	"fmt"
	"github.com/JosephNaberhaus/agnostic/ast"
	"math"
	"math/big"
)

func selfConsumer() consumer[ast.Self] {
	return mapResultToConstant(stringConsumer("self"), ast.Self{})
}

func nullLiteralConsumer() consumer[ast.Null] {
	return mapResultToConstant(stringConsumer("null"), ast.Null{})
}

func boolLiteralConsumer() consumer[ast.LiteralBool] {
	return first(
		mapResultToConstant(stringConsumer("false"), ast.LiteralBool{Value: false}),
		mapResultToConstant(stringConsumer("true"), ast.LiteralBool{Value: true}),
	)
}

func intLiteralConsumer() consumer[ast.LiteralInt] {
	return mapResult(
		intConsumer(),
		func(number *big.Int) (ast.LiteralInt, error) {
			if big.NewInt(math.MaxInt64).Cmp(number) < 0 {
				return ast.LiteralInt{}, fmt.Errorf("number %v is greater than the maximum allowed int %v", number.String(), math.MaxInt64)
			}

			if big.NewInt(math.MinInt64).Cmp(number) > 0 {
				return ast.LiteralInt{}, fmt.Errorf("number %v is less than the minimum allowed int %v", number.String(), math.MaxInt64)
			}

			return ast.LiteralInt{
				Value: number.Int64(),
			}, nil
		},
	)
}

func stringLiteralConsumer() consumer[ast.LiteralString] {
	var result ast.LiteralString
	return attempt(
		&result,
		inOrder(
			skip(stringConsumer("\"")),
			handleNoError(
				func(state parserState) (parserState, string, error) {
					value := state.consumeUntilStr(func(r rune) bool {
						return r == '"'
					})
					return state, value, nil
				},
				func(value string) {
					result.Value = value
				},
			),
			skip(stringConsumer("\"")),
		),
	)
}

func runeLiteralConsumer() consumer[ast.LiteralRune] {
	var result ast.LiteralRune
	return attempt(
		&result,
		inOrder(
			skip(stringConsumer("'")),
			handleNoError(
				func(state parserState) (parserState, rune, error) {
					if state.isEmpty() {
						return parserState{}, 0, createError(state, "expected any rune")
					}

					if state.peekOne() == '\\' {
						state.consumeOne()

						switch state.consumeOne() {
						case 'n':
							return state, '\n', nil
						case 't':
							return state, '\t', nil
						}

						return parserState{}, 0, createError(state, "unsupported escape sequence")
					}

					result := state.consumeOne()
					return state, result, nil
				},
				func(value rune) {
					result.Value = value
				},
			),
			skip(stringConsumer("'")),
		),
	)
}

func emptyListConsumer() consumer[ast.EmptyList] {
	var result ast.EmptyList
	return attempt(
		&result,
		inOrder(
			skip(stringConsumer("make(")),
			anyWhitespaceConsumer(),
			handleNoError(
				listConsumer(),
				func(list ast.List) {
					result.Type = list.Base
				},
			),
			anyWhitespaceConsumer(),
			skip(stringConsumer(")")),
		),
	)
}

func emptySetConsumer() consumer[ast.EmptySet] {
	var result ast.EmptySet
	return attempt(
		&result,
		inOrder(
			skip(stringConsumer("make(")),
			anyWhitespaceConsumer(),
			handleNoError(
				setConsumer(),
				func(set ast.Set) {
					result.Type = set.Base
				},
			),
			anyWhitespaceConsumer(),
			skip(stringConsumer(")")),
		),
	)
}

type operatorType int

const (
	openParen operatorType = iota + 1
	separator
	colon
	function
	structLiteral
	listLiteral
	mapLiteral
	setLiteral
	property
	not
	multiply
	divide
	modulo
	subtract
	add
	equal
	notEqual
	lookup
	newModel
	lessThan
	lessThanOrEqualTo
	greaterThan
	greaterThanOrEqualTo
	and
	or
	castToInt
	castToString
)

const negativeInf = math.MinInt

func (o operatorType) precedence() int {
	switch o {
	case openParen, separator, colon, listLiteral, mapLiteral, setLiteral, function, structLiteral, lookup:
		return negativeInf
	case or:
		return 1
	case and:
		return 2
	case equal, notEqual:
		return 3
	case lessThan, lessThanOrEqualTo, greaterThan, greaterThanOrEqualTo:
		return 4
	case add, subtract:
		return 5
	case multiply, divide, modulo:
		return 6
	case not, newModel, castToInt, castToString:
		return 7
	case property:
		return 8
	}

	panic("unreachable")
}

type semanticOperator operatorType

const (
	openParenOperator = semanticOperator(openParen)
	separatorOperator = semanticOperator(separator)
	colonOperator     = semanticOperator(colon)
)

func (s semanticOperator) operatorType() operatorType {
	return operatorType(s)
}

type binaryOperator operatorType

func (b binaryOperator) operatorType() operatorType {
	return operatorType(b)
}

const (
	multiplyOperator             = binaryOperator(multiply)
	divideOperator               = binaryOperator(divide)
	moduloOperator               = binaryOperator(modulo)
	subtractOperator             = binaryOperator(subtract)
	addOperator                  = binaryOperator(add)
	equalOperator                = binaryOperator(equal)
	notEqualOperator             = binaryOperator(notEqual)
	propertyOperator             = binaryOperator(property)
	lessThanOperator             = binaryOperator(lessThan)
	lessThanOrEqualToOperator    = binaryOperator(lessThanOrEqualTo)
	greaterThanOperator          = binaryOperator(greaterThan)
	greaterThanOrEqualToOperator = binaryOperator(greaterThanOrEqualTo)
	andOperator                  = binaryOperator(and)
	orOperator                   = binaryOperator(or)
)

type unaryOperator operatorType

func (u unaryOperator) operatorType() operatorType {
	return operatorType(u)
}

const (
	newModelOperator     = unaryOperator(newModel)
	notOperator          = unaryOperator(not)
	castToIntOperator    = unaryOperator(castToInt)
	castToStringOperator = unaryOperator(castToString)
)

type operator interface {
	operatorType() operatorType
}

type functionOperator struct {
	numArgs int
}

func (f *functionOperator) operatorType() operatorType {
	return function
}

type structLiteralOperator struct {
	numProperties int
}

func (s *structLiteralOperator) operatorType() operatorType {
	return structLiteral
}

type listLiteralOperator struct {
	numItems int
}

func (l *listLiteralOperator) operatorType() operatorType {
	return listLiteral
}

type setLiteralOperator struct {
	numItems int
}

func (l *setLiteralOperator) operatorType() operatorType {
	return setLiteral
}

type mapLiteralOperator struct {
	numEntries int
}

func (l *mapLiteralOperator) operatorType() operatorType {
	return mapLiteral
}

type lookupOperator struct{}

func (l *lookupOperator) operatorType() operatorType {
	return lookup
}

type stack[T any] []T

func (s *stack[T]) push(value T) {
	*s = append(*s, value)
}

func (s *stack[T]) pop() T {
	if s.isEmpty() {
		var zero T
		return zero
	}

	value := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return value
}

func (s *stack[T]) peek() T {
	if s.isEmpty() {
		var zero T
		return zero
	}

	return (*s)[len(*s)-1]
}

func (s *stack[T]) isEmpty() bool {
	return len(*s) == 0
}

func (s *stack[T]) isNotEmpty() bool {
	return len(*s) != 0
}

func (s *stack[T]) length() int {
	return len(*s)
}

func (s *stack[T]) any(condition func(T) bool) bool {
	for _, value := range *s {
		if condition(value) {
			return true
		}
	}

	return false
}

func (s *stack[T]) copy() *stack[T] {
	stackCopy := make(stack[T], len(*s))
	copy(stackCopy, *s)
	return &stackCopy
}

func unaryOperatorToAst(value ast.Value, op operator) (ast.Value, error) {
	if op.operatorType() == newModel {
		valueAsVariable, ok := value.(ast.Variable)
		if !ok {
			// TODO
			return nil, errors.New("can only instantiate variable")
		}

		return ast.New{
			Model: ast.Model{
				Name: valueAsVariable.Name,
			},
		}, nil
	}

	var astUnaryOperator ast.UnaryOperator
	switch op.operatorType() {
	case not:
		astUnaryOperator = ast.Not
	case castToInt:
		astUnaryOperator = ast.CastToInt
	case castToString:
		astUnaryOperator = ast.CastToString
	}

	return ast.UnaryOperation{
		Value:    value,
		Operator: astUnaryOperator,
	}, nil
}

func binaryOperationToAst(left ast.Value, op operator, right ast.Value) (ast.Value, error) {
	if op.operatorType() == property {
		rightAsVariable, ok := right.(ast.Variable)
		if !ok {
			// TOOD
			return nil, errors.New("right-hand side of a property access must be a variable")
		}

		// Built-in properties
		switch rightAsVariable.Name {
		case "length":
			return ast.Length{Value: left}, nil
		}

		return ast.Property{
			Of:   left,
			Name: rightAsVariable.Name,
		}, nil
	}

	var astBinaryOperator ast.BinaryOperator
	switch op {
	case multiplyOperator:
		astBinaryOperator = ast.Multiply
	case divideOperator:
		astBinaryOperator = ast.Divide
	case moduloOperator:
		astBinaryOperator = ast.Modulo
	case addOperator:
		astBinaryOperator = ast.Add
	case subtractOperator:
		astBinaryOperator = ast.Subtract
	case equalOperator:
		astBinaryOperator = ast.Equal
	case notEqualOperator:
		astBinaryOperator = ast.NotEqual
	case lessThanOperator:
		astBinaryOperator = ast.LessThan
	case lessThanOrEqualToOperator:
		astBinaryOperator = ast.LessThanOrEqualTo
	case greaterThanOperator:
		astBinaryOperator = ast.GreaterThan
	case greaterThanOrEqualToOperator:
		astBinaryOperator = ast.GreaterThanOrEqualTo
	case andOperator:
		astBinaryOperator = ast.And
	case orOperator:
		astBinaryOperator = ast.Or
	default:
		panic("unreachable")
	}

	return ast.BinaryOperation{
		Left:     left,
		Operator: astBinaryOperator,
		Right:    right,
	}, nil
}

type token struct {
	value            ast.Value
	wasValueProducer bool
	op               operator
}

func (t token) isValue() bool {
	return t.value != nil
}

func (t token) isOperatorOfType(operatorType operatorType) bool {
	return t.op != nil && t.op.operatorType() == operatorType
}

func valueConsumer() consumer[ast.Value] {
	return func(state parserState) (parserState, ast.Value, error) {
		valueStack := new(stack[ast.Value])
		operatorStack := new(stack[operator])

		// TODO this is a hack
		var checkValid func()

		addOperatorToValueStackHandler := func(op operator, valueStack *stack[ast.Value]) error {
			switch op := op.(type) {
			case unaryOperator:
				if valueStack.length() < 1 {
					return errors.New("unary operation requires one value")
				}

				result, err := unaryOperatorToAst(valueStack.pop(), op)
				if err != nil {
					return err
				}

				valueStack.push(result)
			case binaryOperator:
				if valueStack.length() < 2 {
					return errors.New("binary operation requires two values")
				}

				right := valueStack.pop()
				left := valueStack.pop()

				result, err := binaryOperationToAst(left, op, right)
				if err != nil {
					return err
				}

				valueStack.push(result)
			case *functionOperator:
				arguments := make([]ast.Value, op.numArgs)
				for i := 0; i < op.numArgs; i++ {
					arguments[op.numArgs-i-1] = valueStack.pop()
				}

				// Convert the variable type into a callable type.
				var function ast.Callable
				switch value := valueStack.pop().(type) {
				case ast.Property:
					function = ast.FunctionProperty{
						Of:   value.Of,
						Name: value.Name,
					}
				case ast.Variable:
					// Check if it's one of the built-in functions.
					switch value.Name {
					case "hash":
						if len(arguments) != 1 {
							return errors.New("hash can only be called with exactly one argument")
						}

						valueStack.push(ast.UnaryOperation{
							Value:    arguments[0],
							Operator: ast.Hash,
						})
						return nil
					}

					function = ast.Function{
						Name: value.Name,
					}
				default:
					return errors.New("invalid type to call")
				}

				valueStack.push(ast.Call{
					Function:  function,
					Arguments: arguments,
				})
			case *structLiteralOperator:
				properties := make([]ast.LiteralProperty, op.numProperties)
				for i := 0; i < op.numProperties; i++ {
					valueValue := valueStack.pop()
					keyValue := valueStack.pop()

					if keyVariable, ok := keyValue.(ast.Variable); ok {
						properties[op.numProperties-i-1] = ast.LiteralProperty{
							Name:  keyVariable.Name,
							Value: valueValue,
						}
					} else {
						return errors.New("literal property must have variables as keys")
					}
				}

				if structNameVariable, ok := valueStack.pop().(ast.Variable); ok {
					valueStack.push(ast.LiteralStruct{
						Name:       structNameVariable.Name,
						Properties: properties,
					})
				} else {
					return errors.New("invalid type to make a struct literal of")
				}
			case *listLiteralOperator:
				items := make([]ast.Value, op.numItems)
				for i := 0; i < op.numItems; i++ {
					items[op.numItems-i-1] = valueStack.pop()
				}

				valueStack.push(ast.LiteralList{Items: items})
			case *mapLiteralOperator:
				entries := make([]ast.KeyValue, op.numEntries)
				for i := 0; i < op.numEntries; i++ {
					entries[op.numEntries-i-1] = ast.KeyValue{
						Value: valueStack.pop(),
						Key:   valueStack.pop(),
					}
				}

				valueStack.push(ast.LiteralMap{Entries: entries})
			case *setLiteralOperator:
				items := make([]ast.Value, op.numItems)
				for i := 0; i < op.numItems; i++ {
					items[op.numItems-i-1] = valueStack.pop()
				}

				valueStack.push(ast.LiteralSet{Items: items})
			case *lookupOperator:
				valueStack.push(ast.Lookup{
					Key:  valueStack.pop(),
					From: valueStack.pop(),
				})
			default:
				return errors.New("invalid operator")
			}

			return nil
		}

		addOperatorToValueStack := func(op operator) error {
			err := addOperatorToValueStackHandler(op, valueStack)
			if err != nil {
				return err
			}

			return nil
		}

		lastHandled := token{wasValueProducer: false}

		handleValue := func(value ast.Value) {
			lastHandled = token{value: value}

			valueStack.push(value)
		}

		handlePrecedence := func(precedence int) error {
			if precedence != negativeInf {
				for operatorStack.isNotEmpty() && precedence <= operatorStack.peek().operatorType().precedence() {

					err := addOperatorToValueStack(operatorStack.pop())
					if err != nil {
						return err
					}
				}
			}

			return nil
		}

		handleOperator := func(op operator) error {
			lastHandled = token{op: op}

			err := handlePrecedence(op.operatorType().precedence())
			if err != nil {
				return err
			}

			operatorStack.push(op)
			return nil
		}

		emptyStacks := func(valueStack *stack[ast.Value], operatorStack *stack[operator]) error {
			for operatorStack.isNotEmpty() {
				err := addOperatorToValueStackHandler(operatorStack.pop(), valueStack)
				if err != nil {
					return err
				}
			}

			return nil
		}

		var lastValidOutput ast.Value
		var lastValidOutputState parserState

		checkValid = func() {
			valueStackCopy := valueStack.copy()
			operatorStackCopy := operatorStack.copy()

			err := emptyStacks(valueStackCopy, operatorStackCopy)
			if err != nil {
				return
			}

			if valueStackCopy.length() == 1 && operatorStackCopy.isEmpty() {
				lastValidOutput = valueStackCopy.peek()
				lastValidOutputState = state
			}
		}

		var err error
		for state.isNotEmpty() {
			var newState parserState
			newState, _, err = first(
				allWhitespaceConsumer(),
				// TODO this should probably only be allowed if we're consuming a map or list literal and the last thing was a comma
				newlineConsumer(),
				handle(
					first(
						castToValue(nullLiteralConsumer()),
						castToValue(selfConsumer()),
						castToValue(boolLiteralConsumer()),
						castToValue(intLiteralConsumer()),
						castToValue(runeLiteralConsumer()),
						castToValue(stringLiteralConsumer()),
						castToValue(emptyListConsumer()),
						castToValue(emptySetConsumer()),
					),
					func(value ast.Value) error {
						handleValue(value)
						return nil
					},
				),
				handle(
					skip(stringConsumer(",")),
					func(_ void) error {
						for true {
							if operatorStack.isEmpty() {
								return errors.New("unexpected comma")
							}

							nextOpType := operatorStack.peek().operatorType()
							if nextOpType == mapLiteral || nextOpType == structLiteral || nextOpType == listLiteral || nextOpType == setLiteral || nextOpType == function || nextOpType == separator || nextOpType == colon {
								break
							}

							err = addOperatorToValueStack(operatorStack.pop())
							if err != nil {
								return err
							}
						}

						return handleOperator(separatorOperator)
					},
				),
				handle(
					skip(stringConsumer(":")),
					func(_ void) error {
						for true {
							if operatorStack.isEmpty() {
								return errors.New("unexpected colon")
							}

							nextOpType := operatorStack.peek().operatorType()
							if nextOpType == mapLiteral || nextOpType == structLiteral || nextOpType == separator {
								break
							}

							err = addOperatorToValueStack(operatorStack.pop())
							if err != nil {
								return err
							}
						}

						return handleOperator(colonOperator)
					},
				),
				handle(
					skip(stringConsumer("(")),
					func(_ void) error {
						if lastHandled.isValue() {
							if operatorStack.isNotEmpty() && operatorStack.peek().operatorType() == newModel {
								// TODO
							} else if operatorStack.isNotEmpty() && operatorStack.peek().operatorType() == castToInt {
								// TODO
							} else if operatorStack.isNotEmpty() && operatorStack.peek().operatorType() == castToString {
								// TODO
							} else {
								err = handlePrecedence(property.precedence())
								if err != nil {
									return err
								}
								return handleOperator(new(functionOperator))
							}
						}

						return handleOperator(openParenOperator)
					},
				),
				handle(
					skip(stringConsumer(")")),
					func(_ void) error {
						argCount := 0
						if !lastHandled.isOperatorOfType(function) {
							argCount = 1
						}

						for true {
							if operatorStack.isEmpty() {
								return errors.New("no open parentheses found")
							}

							if operatorStack.peek().operatorType() == openParen {
								operatorStack.pop()

								break
							}

							if operatorStack.peek().operatorType() == separator {
								operatorStack.pop()

								argCount++
								continue
							}

							if op, ok := operatorStack.peek().(*functionOperator); ok {
								op.numArgs = argCount

								lastHandled = token{wasValueProducer: true}

								return addOperatorToValueStack(operatorStack.pop())
							}

							err = addOperatorToValueStack(operatorStack.pop())
							if err != nil {
								return err
							}
						}

						return nil
					},
				),
				handle(
					skip(stringConsumer("[")),
					func(_ void) error {
						if lastHandled.isValue() || lastHandled.wasValueProducer {
							err = handlePrecedence(property.precedence())
							if err != nil {
								return err
							}
							return handleOperator(new(lookupOperator))
						}

						return handleOperator(new(listLiteralOperator))
					},
				),
				handle(
					skip(stringConsumer("]")),
					func(_ void) error {
						itemCount := 0
						if !lastHandled.isOperatorOfType(listLiteral) {
							itemCount = 1
						}

						for true {
							if _, ok := operatorStack.peek().(*lookupOperator); ok {
								return addOperatorToValueStack(operatorStack.pop())
							}

							if op, ok := operatorStack.peek().(*listLiteralOperator); ok {
								op.numItems = itemCount
								return addOperatorToValueStack(operatorStack.pop())
							}

							if operatorStack.peek() == separatorOperator {
								operatorStack.pop()

								itemCount++
								continue
							}

							err = addOperatorToValueStack(operatorStack.pop())
							if err != nil {
								return err
							}
						}

						return nil
					},
				),
				handle(
					skip(stringConsumer("{")),
					func(_ void) error {
						if lastHandled.isValue() {
							return handleOperator(new(structLiteralOperator))
						}

						return handleOperator(new(mapLiteralOperator))
					},
				),
				handle(
					skip(stringConsumer("}")),
					func(_ void) error {
						entryCount := 0
						for true {
							if op, ok := operatorStack.peek().(*mapLiteralOperator); ok {
								op.numEntries = entryCount
								return addOperatorToValueStack(operatorStack.pop())
							}

							if op, ok := operatorStack.peek().(*structLiteralOperator); ok {
								op.numProperties = entryCount
								return addOperatorToValueStack(operatorStack.pop())
							}

							if operatorStack.peek() == colonOperator {
								operatorStack.pop()
								entryCount++
								continue
							}

							if operatorStack.peek() == separatorOperator {
								operatorStack.pop()
								if operatorStack.pop() != colonOperator {
									return errors.New("expected colon")
								}

								entryCount++
								continue
							}

							err = addOperatorToValueStack(operatorStack.pop())
							if err != nil {
								return err
							}
						}

						return nil
					},
				),
				handle(
					skip(stringConsumer("<")),
					func(_ void) error {
						if lastHandled.isValue() || lastHandled.wasValueProducer {
							// TODO this is literally garbage
							return errors.New("set literal cannot be defined after a value")
						}

						return handleOperator(new(setLiteralOperator))
					},
				),
				handle(
					skip(stringConsumer(">")),
					func(_ void) error {
						numItems := 0
						if !lastHandled.isOperatorOfType(setLiteral) {
							numItems = 1
						}

						if !operatorStack.any(func(value operator) bool {
							_, isSetLiteral := value.(*setLiteralOperator)
							return isSetLiteral
						}) {
							return errors.New("expected set literal opeartor")
						}

						for true {
							if op, ok := operatorStack.peek().(*setLiteralOperator); ok {
								op.numItems = numItems
								return addOperatorToValueStack(operatorStack.pop())
							}

							if operatorStack.peek() == separatorOperator {
								operatorStack.pop()

								numItems++
								continue
							}

							err = addOperatorToValueStack(operatorStack.pop())
							if err != nil {
								return err
							}
						}

						return nil
					},
				),
				handle(
					first(
						mapResultToConstant(stringConsumer("*"), multiplyOperator),
						mapResultToConstant(stringConsumer("/"), divideOperator),
						mapResultToConstant(stringConsumer("%"), moduloOperator),
						mapResultToConstant(stringConsumer("+"), addOperator),
						mapResultToConstant(stringConsumer("-"), subtractOperator),
						mapResultToConstant(stringConsumer("."), propertyOperator),
						mapResultToConstant(stringConsumer("=="), equalOperator),
						mapResultToConstant(stringConsumer("!="), notEqualOperator),
						mapResultToConstant(stringConsumer("<="), lessThanOrEqualToOperator),
						mapResultToConstant(stringConsumer("<"), lessThanOperator),
						mapResultToConstant(stringConsumer(">="), greaterThanOrEqualToOperator),
						mapResultToConstant(stringConsumer(">"), greaterThanOperator),
						mapResultToConstant(stringConsumer("&&"), andOperator),
						mapResultToConstant(stringConsumer("||"), orOperator),
					),
					func(operator binaryOperator) error {
						return handleOperator(operator)
					},
				),
				handle(
					first(
						mapResultToConstant(inOrder(
							skip(stringConsumer("new")),
							allWhitespaceConsumer(),
							// TODO make sure that the next character is a (
						), newModelOperator),
						mapResultToConstant(stringConsumer("!"), notOperator),
						mapResultToConstant(inOrder(
							skip(stringConsumer("int")),
							// TODO make sure that the next character is a (
						), castToIntOperator),
						mapResultToConstant(inOrder(
							skip(stringConsumer("string")),
							// TODO make sure that the next character is a (
						), castToStringOperator),
					),
					func(op unaryOperator) error {
						return handleOperator(op)
					},
				),
				handle(
					mapResult(
						alphaConsumer(),
						func(name string) (ast.Variable, error) {
							return ast.Variable{
								Name: name,
							}, nil
						},
					),
					func(value ast.Variable) error {
						handleValue(value)
						return nil
					},
				),
			)(state)
			if err != nil {
				break
			}

			state = newState
			checkValid()
		}
		if err != nil {
			if lastValidOutput != nil {
				lastValidOutputState.addError(err)
				return lastValidOutputState, lastValidOutput, nil
			}

			return parserState{}, nil, err
		}

		if lastValidOutput != nil {
			return lastValidOutputState, lastValidOutput, err
		}

		return parserState{}, nil, errors.New("expected more operators")
	}
}
