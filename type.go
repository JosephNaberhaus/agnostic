package agnostic

type Type interface {
	// isType is not used, but allows us to ensure that only allowed values are passed in as a Type
	isType()
}

type PrimitiveType int

const (
	BooleanType PrimitiveType = iota
	IntType
	FloatType
	StringType
)

// isType implements the Type interface so that PrimitiveType can be used as a Type
func (p PrimitiveType) isType() {}

type ArrayType struct {
	ElementType Type
}

// isType implements the Type interface so that ArrayType can be used as a Type
func (a ArrayType) isType() {}

type MapType struct {
	KeyType, ValueType Type
}

// isType implements the Type interface so that MapType can be used as a Type
func (m MapType) isType() {}

type ModelType struct {
	Model Model
}

// isType implements the Type interface so that ModelType can be used as a Type
func (m ModelType) isType() {}
