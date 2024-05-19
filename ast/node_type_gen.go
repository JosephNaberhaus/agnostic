package ast


type Assignable interface {
    isAssignable()
}

type Statement interface {
    isStatement()
}

type Definition interface {
    isDefinition()
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
