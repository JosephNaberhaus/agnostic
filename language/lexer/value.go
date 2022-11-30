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

type operatorType int

const (
	openParen operatorType = iota + 1
	argumentSeparator
	property
	not
	subtract
	add
	equal
	function
	lookup
	newModel
)

const negativeInf = math.MinInt

func (o operatorType) precedence() int {
	switch o {
	case openParen, argumentSeparator:
		return negativeInf
	case equal:
		return 1
	case add, subtract:
		return 2
	case not, newModel:
		return 3
	case property, function, lookup:
		return 4
	}

	panic("unreachable")
}

type semanticOperator operatorType

const (
	openParenOperator         = semanticOperator(openParen)
	argumentSeparatorOperator = semanticOperator(argumentSeparator)
)

func (s semanticOperator) operatorType() operatorType {
	return operatorType(s)
}

type binaryOperator operatorType

func (b binaryOperator) operatorType() operatorType {
	return operatorType(b)
}

const (
	subtractOperator = binaryOperator(subtract)
	addOperator      = binaryOperator(add)
	equalOperator    = binaryOperator(equal)
	propertyOperator = binaryOperator(property)
)

type unaryPrefixOperator operatorType

func (u unaryPrefixOperator) operatorType() operatorType {
	return operatorType(u)
}

const (
	newModelOperator = unaryPrefixOperator(newModel)
	notOperator      = unaryPrefixOperator(not)
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

		return ast.Property{
			Of:   left,
			Name: rightAsVariable.Name,
		}, nil
	}

	var astBinaryOperator ast.BinaryOperator
	switch op {
	case addOperator:
		astBinaryOperator = ast.Add
	case subtractOperator:
		astBinaryOperator = ast.Subtract
	case equalOperator:
		astBinaryOperator = ast.Equals
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
			case unaryPrefixOperator:
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
			handle(
				first(
					mapResultToConstant(inOrder(
						skip(stringConsumer("new")),
						allWhitespaceConsumer(),
					), newModelOperator),
					mapResultToConstant(stringConsumer("!"), notOperator),
				),
				func(op unaryPrefixOperator) error {
					return handleOperator(op)
				},
			),
			handle(
				first(
					mapResultToConstant(stringConsumer("+"), addOperator),
					mapResultToConstant(stringConsumer("-"), subtractOperator),
					mapResultToConstant(stringConsumer("."), propertyOperator),
					mapResultToConstant(stringConsumer("=="), equalOperator),
				),
				func(operator binaryOperator) error {
					return handleOperator(operator)
				},
			),
			handle(
				skip(stringConsumer(",")),
				func(_ void) error {
					return handleOperator(argumentSeparatorOperator)
				},
			),
			handle(
				skip(stringConsumer("(")),
				func(_ void) error {
					if lastHandled.isValue() {
						isConstructor := operatorStack.isNotEmpty() && operatorStack.peek().operatorType() == newModel
						if isConstructor {
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

						if operatorStack.peek().operatorType() == argumentSeparator {
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
					return handleOperator(new(lookupOperator))
				},
			),
			handle(
				skip(stringConsumer("]")),
				func(_ void) error {
					for true {
						if operatorStack.isEmpty() {
							return errors.New("no open bracket found")
						}

						if _, ok := operatorStack.peek().(*lookupOperator); ok {
							return addOperatorToValueStack(operatorStack.pop())

							break
						}

						err = addOperatorToValueStack(operatorStack.pop())
						if err != nil {
							return err
						}
					}

					return nil
				},
			),
			handleNoError(
				first(
					castToValue(intLiteralConsumer()),
					mapResult(
						alphaConsumer(),
						func(name string) (ast.Value, error) {
							return ast.Variable{
								Name: name,
							}, nil
						},
					),
				),
				handleValue,
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
