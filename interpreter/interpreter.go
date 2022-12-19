package interpreter

import (
	"fmt"
	"github.com/JosephNaberhaus/agnostic/code"
	"github.com/JosephNaberhaus/agnostic/stack"
)

func Run(module *code.Module, entrypoint string, arguments ...Value) Value {
	interpreter := newInterpreter(module)

	return interpreter.execute(
		module.FunctionMap[entrypoint],
		newCallScope(nil, module.FunctionMap[entrypoint], arguments),
	)
}

type interpreter struct {
	constants map[string]Value
	functions map[string]*code.FunctionDef

	callStack stack.Stack[*callScope]
}

func (i interpreter) startBlock() {
	i.callStack.Peek().startBlock()
}

func (i interpreter) endBlock() {
	i.callStack.Peek().endBlock()
}

func (i interpreter) declare(name string, value Value) {
	i.callStack.Peek().declareVariable(name, value)
}

func (i interpreter) assign(name string, value Value) {
	i.callStack.Peek().assignVariable(name, value)
}

func (i interpreter) self() *Model {
	return i.callStack.Peek().model
}

func (i interpreter) variableByName(name string) Value {
	if variable, ok := i.callStack.Peek().getVariable(name); ok {
		return variable
	}

	if constant, ok := i.constants[name]; ok {
		return constant
	}

	panic(fmt.Sprintf("variable not found: %s", name))
}

func (i interpreter) functionByName(name string) (*Model, *code.FunctionDef) {
	if model, method, ok := i.callStack.Peek().getMethod(name); ok {
		return model, method
	}

	if function, ok := i.functions[name]; ok {
		return nil, function
	}

	panic(fmt.Sprintf("function not found: %s", name))
}

func (i interpreter) execute(function *code.FunctionDef, scope *callScope) Value {
	i.callStack.Push(scope)

	mapper := &statementMapper{runtime: i}
	return mapper.MapBlock(function.Block).value
}

func newInterpreter(module *code.Module) *interpreter {
	interpreter := &interpreter{
		constants: map[string]Value{},
		functions: map[string]*code.FunctionDef{},
		callStack: stack.Stack[*callScope]{},
	}

	for _, constantDef := range module.Constants {
		constantValue := code.MapValueNoError[Value](constantDef.Value, &valueMapper{runtime: interpreter})
		interpreter.constants[constantDef.Name] = constantValue
	}

	for _, functionDef := range module.Functions {
		interpreter.functions[functionDef.Name] = functionDef
	}

	return interpreter
}
