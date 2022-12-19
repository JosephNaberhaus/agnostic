//go:generate go run ../tool/mapper_generator -exclude=optional.go

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

type ConstantValue interface {
	Value
	// isConstantValue is only a type-guard to limit what can be used as a ConstantValue.
	isConstantValue()
}

type Callable interface {
	Node
	// isCallable is only a type-guard to limit what can be used as a Callable
	isCallable()
}

type CallableDef interface {
	Node
	// isCallableDef is only a type-guard to limit what can be used as a CallableDef
	isCallableDef()
}
