package lexer

import (
	"errors"
	"fmt"
	"github.com/JosephNaberhaus/agnostic/ast"
	"math"
	"math/big"
)

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

type operatorType int

const (
	openParen operatorType = iota + 1
	separator
	colon
	function
	listLiteral
	mapLiteral
	property
	not
	multiply
	divide
	subtract
	add
	equal
	lookup
	newModel
	lessThan
	greaterThan
	and
	or
	castToInt
)

const negativeInf = math.MinInt

func (o operatorType) precedence() int {
	switch o {
	case openParen, separator, colon, lookup, function, listLiteral, mapLiteral:
		return negativeInf
	case or:
		return 1
	case and:
		return 2
	case equal:
		return 3
	case lessThan, greaterThan:
		return 4
	case add, subtract:
		return 5
	case multiply, divide:
		return 6
	case not, newModel, castToInt:
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
	multiplyOperator    = binaryOperator(multiply)
	divideOperator      = binaryOperator(divide)
	subtractOperator    = binaryOperator(subtract)
	addOperator         = binaryOperator(add)
	equalOperator       = binaryOperator(equal)
	propertyOperator    = binaryOperator(property)
	lessThanOperator    = binaryOperator(lessThan)
	greaterThanOperator = binaryOperator(greaterThan)
	andOperator         = binaryOperator(and)
	orOperator          = binaryOperator(or)
)

type unaryOperator operatorType

func (u unaryOperator) operatorType() operatorType {
	return operatorType(u)
}

const (
	newModelOperator  = unaryOperator(newModel)
	notOperator       = unaryOperator(not)
	castToIntOperator = unaryOperator(castToInt)
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

type listLiteralOperator struct {
	numItems int
}

func (l *listLiteralOperator) operatorType() operatorType {
	return listLiteral
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
	case addOperator:
		astBinaryOperator = ast.Add
	case subtractOperator:
		astBinaryOperator = ast.Subtract
	case equalOperator:
		astBinaryOperator = ast.Equals
	case lessThanOperator:
		astBinaryOperator = ast.LessThan
	case greaterThanOperator:
		astBinaryOperator = ast.GreaterThan
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
	value ast.Value
	op    operator
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

		addOperatorToValueStack := func(op operator) error {
			switch op := op.(type) {
			case unaryOperator:
				result, err := unaryOperatorToAst(valueStack.pop(), op)
				if err != nil {
					return err
				}

				valueStack.push(result)
				return nil
			case binaryOperator:
				right := valueStack.pop()
				left := valueStack.pop()

				result, err := binaryOperationToAst(left, op, right)
				if err != nil {
					return err
				}

				valueStack.push(result)
				return nil
			case *functionOperator:
				arguments := make([]ast.Value, op.numArgs)
				for i := 0; i < op.numArgs; i++ {
					arguments[op.numArgs-i-1] = valueStack.pop()
				}

				// Convert the variable type into a callable type
				var function ast.Callable
				switch value := valueStack.pop().(type) {
				case ast.Property:
					function = ast.FunctionProperty{
						Of:   value.Of,
						Name: value.Name,
					}
				case ast.Variable:
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
				return nil
			case *listLiteralOperator:
				items := make([]ast.Value, op.numItems)
				for i := 0; i < op.numItems; i++ {
					items[op.numItems-i-1] = valueStack.pop()
				}

				valueStack.push(ast.LiteralList{Items: items})
				return nil
			case *mapLiteralOperator:
				entries := make([]ast.KeyValue, op.numEntries)
				for i := 0; i < op.numEntries; i++ {
					entries[op.numEntries-i-1] = ast.KeyValue{
						Value: valueStack.pop(),
						Key:   valueStack.pop(),
					}
				}

				valueStack.push(ast.LiteralMap{Entries: entries})
				return nil
			case *lookupOperator:
				valueStack.push(ast.Lookup{
					Key:  valueStack.pop(),
					From: valueStack.pop(),
				})
				return nil
			}

			return errors.New("invalid operator")
		}

		var lastHandled token

		handleValue := func(value ast.Value) {
			lastHandled = token{value: value}

			valueStack.push(value)
		}

		handleOperator := func(op operator) error {
			lastHandled = token{op: op}

			if op.operatorType().precedence() != negativeInf {
				for operatorStack.isNotEmpty() && op.operatorType().precedence() <= operatorStack.peek().operatorType().precedence() {

					err := addOperatorToValueStack(operatorStack.pop())
					if err != nil {
						return err
					}
				}
			}

			operatorStack.push(op)
			return nil
		}

		var err error
		state, _, err = repeat(first(
			allWhitespaceConsumer(),
			// TODO this should probably only be allowed if we're consuming a map or list literal and the last thing was a comma
			newlineConsumer(),
			handle(
				first(
					castToValue(intLiteralConsumer()),
					castToValue(runeLiteralConsumer()),
				),
				func(value ast.Value) error {
					handleValue(value)
					return nil
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
				),
				func(op unaryOperator) error {
					return handleOperator(op)
				},
			),
			handle(
				first(
					mapResultToConstant(stringConsumer("*"), multiplyOperator),
					mapResultToConstant(stringConsumer("/"), divideOperator),
					mapResultToConstant(stringConsumer("+"), addOperator),
					mapResultToConstant(stringConsumer("-"), subtractOperator),
					mapResultToConstant(stringConsumer("."), propertyOperator),
					mapResultToConstant(stringConsumer("=="), equalOperator),
					mapResultToConstant(stringConsumer("<"), lessThanOperator),
					mapResultToConstant(stringConsumer(">"), greaterThanOperator),
					mapResultToConstant(stringConsumer("&&"), andOperator),
					mapResultToConstant(stringConsumer("||"), orOperator),
				),
				func(operator binaryOperator) error {
					return handleOperator(operator)
				},
			),
			handle(
				skip(stringConsumer(",")),
				func(_ void) error {
					return handleOperator(separatorOperator)
				},
			),
			handle(
				skip(stringConsumer(":")),
				func(_ void) error {
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
						} else {
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
					if lastHandled.isValue() {
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
					return handleOperator(new(mapLiteralOperator))
				},
			),
			handle(
				skip(stringConsumer("}")),
				func(_ void) error {
					entryCount := 0
					if !lastHandled.isOperatorOfType(mapLiteral) {
						if operatorStack.pop() != colonOperator {
							return errors.New("expected colon")
						}

						entryCount = 1
					}

					for true {
						if op, ok := operatorStack.peek().(*mapLiteralOperator); ok {
							op.numEntries = entryCount
							return addOperatorToValueStack(operatorStack.pop())
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
		))(state)
		if err != nil {
			return parserState{}, nil, err
		}

		for operatorStack.isNotEmpty() {
			err = addOperatorToValueStack(operatorStack.pop())
			if err != nil {
				return parserState{}, nil, err
			}
		}

		if len(*valueStack) != 1 {
			return parserState{}, nil, errors.New("testing")
		}

		return state, valueStack.pop(), nil
	}
}
