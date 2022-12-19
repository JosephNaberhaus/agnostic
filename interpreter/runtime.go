package interpreter

import "github.com/JosephNaberhaus/agnostic/code"

type runtime interface {
	self() *Model

	variableByName(name string) Value
	functionByName(name string) (*Model, *code.FunctionDef)

	startBlock()
	endBlock()

	declare(name string, value Value)
	assign(name string, value Value)

	execute(function *code.FunctionDef, scope *callScope) Value
}
