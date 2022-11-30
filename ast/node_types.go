//go:generate go run ../tool/mapper_generator/main.go

package ast

type Definition interface {
	Node
	// isDefinition is only a type-guard to limit what can be used as an Definition.
	isDefinition()
}

type Statement interface {
	Node
	// isStatement is only a type-guard to limit what can be used as a Statement.
	isStatement()
}

type Type interface {
	Node
	// isType is only a type guard to limit what can be used as a Type.
	isType()
}

type Value interface {
	Node
	// isValue is only a type-guard to limit what can be used as a Value.
	isValue()
}

type Callable interface {
	Node
	// isCallable is only a type-guard to limit what can be used as a callable
	isCallable()
}
