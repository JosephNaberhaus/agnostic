package interpreter

import (
	"fmt"
	"github.com/JosephNaberhaus/agnostic/code"
	"github.com/JosephNaberhaus/agnostic/stack"
)

type block struct {
	variables map[string]Value
}

func newBlock() block {
	return block{
		variables: map[string]Value{},
	}
}

type callScope struct {
	// The model that owns the function, or null if this is a function call.
	model      *Model
	arguments  map[string]Value
	blockStack stack.Stack[block]
}

func (c *callScope) startBlock() {
	c.blockStack.Push(newBlock())
}

func (c *callScope) endBlock() {
	c.blockStack.Pop()
}

func newCallScope(model *Model, function *code.FunctionDef, arguments []Value) *callScope {
	argumentsByName := map[string]Value{}
	for i, argument := range arguments {
		name := function.Arguments[i].Name
		argumentsByName[name] = argument
	}

	scope := &callScope{
		model:      model,
		arguments:  argumentsByName,
		blockStack: stack.Stack[block]{},
	}
	scope.startBlock()

	return scope
}

func (c *callScope) getVariable(name string) (Value, bool) {
	blockStack := c.blockStack.Copy()
	for blockStack.IsNotEmpty() {
		if variable, ok := blockStack.Pop().variables[name]; ok {
			return variable, true
		}
	}

	if argument, ok := c.arguments[name]; ok {
		return argument, true
	}

	if property, ok := c.model.Properties[name]; ok {
		return property, true
	}

	return Value{}, false
}

func (c *callScope) declareVariable(name string, value Value) {
	c.blockStack.Peek().variables[name] = value
}

func (c *callScope) assignVariable(name string, value Value) {
	blockStack := c.blockStack.Copy()
	for blockStack.IsNotEmpty() {
		curBlock := blockStack.Pop()
		if _, ok := curBlock.variables[name]; ok {
			curBlock.variables[name] = value
			return
		}
	}

	if _, ok := c.model.Properties[name]; ok {
		c.model.Properties[name] = value
		return
	}

	panic(fmt.Sprintf("unknown variable: %s", name))
}

func (c *callScope) getMethod(name string) (*Model, *code.FunctionDef, bool) {
	if c.model == nil {
		return nil, nil, false
	}

	if method, ok := c.model.Methods[name]; ok {
		return c.model, method, true
	}

	return nil, nil, false
}
