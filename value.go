package agnostic

type Value interface {
	// isValue is not used, but allows us to ensure that only allowed values are passed in as a Value
	isValue()
}

type BooleanLiteralValue bool

// isValue implements the Value interface so that BooleanLiteralValue can be used as a Value
func (b BooleanLiteralValue) isValue() {}

type IntLiteralValue int

// isValue implements the Value interface so that IntLiteralValue can be used as a Value
func (i IntLiteralValue) isValue() {}

type FloatLiteralValue float64

// isValue implements the Value interface so that FloatLiteralValue can be used as a Value
func (f FloatLiteralValue) isValue() {}

type StringLiteralValue string

// isValue implements the Value interface so that StringLiteralValue can be used as a Value
func (s StringLiteralValue) isValue() {}

type ArrayElementValue struct {
	Array Value
	Index Value
}

// isValue implements the Value interface so that ArrayElementValue can be used as a Value
func (a ArrayElementValue) isValue() {}

type MapElementValue struct {
	Map Value
	Key Value
}

// isValue implements the Value interface so that MapElementValue can be used as a Value
func (m MapElementValue) isValue() {}

type FieldValue struct {
	Model     Value
	FieldName string
}

// isValue implements the Value interface so that FieldValue can be used as a Value
func (f FieldValue) isValue() {}

type OwnField string

// isValue implements the Value interface so that OwnField can be used as a Value
func (o OwnField) isValue() {}

type VariableValue string

// isValue implements the Value interface so that VariableValue can be used as a Value
func (v VariableValue) isValue() {}

type Operator int

const (
	Add Operator = iota
	Subtract
	Multiply
	IntegerDivision
	Modulo
	Equal
	NotEqual
	GreaterThan
	GreaterThanOrEqualTo
	LessThan
	LessThanOrEqualTo
)

type ComputedValue struct {
	Left, Right Value
	Operator    Operator
}

// isValue implements the Value interface so that ComputedValue can be used as a Value
func (c ComputedValue) isValue() {}
