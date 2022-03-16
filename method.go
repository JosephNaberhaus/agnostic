package agnostic

type Statement interface {
	// isStatement is not used, but allows us to ensure that only allowed values are passed in as a Statement
	isStatement()
}

type Declare struct {
	Name  string
	Value Value
}

// isStatement implements the Statement interface so that Declare can be used as a Statement
func (d Declare) isStatement() {}

type AssignVar struct {
	Var, Value Value
}

// isStatement implements the Statement interface so that AssignVar can be used as a Statement
func (a AssignVar) isStatement() {}

type AssignField struct {
	Model     Value
	FieldName string
	Value     Value
}

// isStatement implements the Statement interface so that AssignField can be used as a Statement
func (a AssignField) isStatement() {}

type AssignOwnField struct {
	FieldName string
	Value     Value
}

// isStatement implements the Statement interface so that AssignOwnField can be used as a Statement
func (a AssignOwnField) isStatement() {}

type AppendValue struct {
	Array    Value
	ToAppend Value
}

// isStatement implements the Statement interface so that AppendValue can be used as a Statement
func (a AppendValue) isStatement() {}

type AppendArray struct {
	Array    Value
	ToAppend Value
}

// isStatement implements the Statement interface so that AppendArray can be used as a Statement
func (a AppendArray) isStatement() {}

type RemoveValue struct {
	Array Value
	Index Value
}

// isStatement implements the Statement interface so that RemoveValue can be used as a Statement
func (r RemoveValue) isStatement() {}

type MapPut struct {
	Map   Value
	Key   Value
	Value Value
}

// isStatement implements the Statement interface so that MapPut can be used as a Statement
func (m MapPut) isStatement() {}

type MapDelete struct {
	Map Value
	Key Value
}

// isStatement implements the Statement interface so that MapDelete can be used as a Statement
func (m MapDelete) isStatement() {}

type ForEach struct {
	Array                                Value
	IndexVariableName, ValueVariableName string
	Statements                           []Statement
}

// isStatement implements the Statement interface so that ForEach can be used as a Statement
func (f ForEach) isStatement() {}

type If struct {
	Condition  Value
	Statements []Statement
}

// isStatement implements the Statement interface so that If can be used as a Statement
func (i If) isStatement() {}

type IfElse struct {
	Condition       Value
	TrueStatements  []Statement
	FalseStatements []Statement
}

// isStatement implements the Statement interface so that IfElse can be used as a Statement
func (i IfElse) isStatement() {}

type Return struct {
	ToReturn Value
}

// isStatement implements the Statement interface so that Return can be used as a Statement
func (r Return) isStatement() {}

type Method struct {
	Name       string
	Parameters []Field
	Returns    Type
	Statements []Statement
}
