package ast_to_code_mapper

import (
	"github.com/JosephNaberhaus/agnostic/code"
	"github.com/JosephNaberhaus/agnostic/internal/utils/stack"
)

type deferred struct {
	// The stack as it was when this deferred was created.
	stack stack.Stack[code.Node]
	// The function that was deferred.
	deferredFunc func() error
}

// queueDeferred queues a function that will be called after the remainder of the AST tree is exhuasted. This allows us
// to switch from depth first (the normal behavior of this mapper) to breadth first at arbitrary points.
//
// Once simple example of when this is needed is function calls. The AST does not require that a function be declared
// before it can be called. For example, function A can call function B even if function B appears after function A in
// the AST. This creates a problem with depth first search, because our mapper won't have parsed function B yet. To
// handle this, we simply call queueDeferred before we process the function body. Now the mapper will handle all of the
// function defs before it gets around to handling any function calls.
//
// There other uses for queueDeferred, but they all bare a resemblence to this example.
func (a *Mapper) queueDeferred(deferredFunc func() error) {
	a.deferred = append(a.deferred, deferred{
		stack:        a.stack.Copy(),
		deferredFunc: deferredFunc,
	})
}

// dequeueDeferred returns the oldest added deferred.
func (a *Mapper) dequeueDeferred() deferred {
	result := a.deferred[0]
	a.deferred = a.deferred[1:]
	return result
}
