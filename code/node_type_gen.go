package code

type Node interface {
	// isNode is just a interface guard to restrict what can be used as a Node.
	isNode()
}

type Type interface {
	isType()
}

type Value interface {
	isValue()
}

type ConstantValue interface {
	isConstantValue()
}

type Callable interface {
	isCallable()
}

type Assignable interface {
	isAssignable()
}

type Statement interface {
	isStatement()
}

type Definition interface {
	isDefinition()
}
