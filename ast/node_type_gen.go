package ast


type Statement interface {
    isStatement()
}

type Value interface {
    isValue()
}

type Type interface {
    isType()
}

type ConstantValue interface {
    isConstantValue()
}

type Assignable interface {
    isAssignable()
}

type Definition interface {
    isDefinition()
}

type Callable interface {
    isCallable()
}
