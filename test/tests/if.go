package tests

import (
	"github.com/JosephNaberhaus/agnostic"
	"github.com/JosephNaberhaus/agnostic/test"
)

var ifSuite = test.Suite{
	Name: "IfStatements",
	Tests: []test.Test{
		{
			Name: "IfWhenConditionIsFalse",
			Before: []agnostic.Statement{
				agnostic.Declare{
					Name:  "input",
					Value: agnostic.BooleanLiteralValue(false),
				},
				agnostic.Declare{
					Name:  "output",
					Value: agnostic.BooleanLiteralValue(false),
				},
				agnostic.If{
					Condition: agnostic.VariableValue("input"),
					Statements: []agnostic.Statement{
						agnostic.AssignVar{
							Var:   agnostic.VariableValue("output"),
							Value: agnostic.BooleanLiteralValue(true),
						},
					},
				},
			},
			Assertions: []test.Assertion{
				test.IsFalse{
					Actual: agnostic.VariableValue("output"),
				},
			},
		},
		{
			Name: "IfWhenConditionIsTrue",
			Before: []agnostic.Statement{
				agnostic.Declare{
					Name:  "input",
					Value: agnostic.BooleanLiteralValue(true),
				},
				agnostic.Declare{
					Name:  "output",
					Value: agnostic.BooleanLiteralValue(false),
				},
				agnostic.If{
					Condition: agnostic.VariableValue("input"),
					Statements: []agnostic.Statement{
						agnostic.AssignVar{
							Var:   agnostic.VariableValue("output"),
							Value: agnostic.BooleanLiteralValue(true),
						},
					},
				},
			},
			Assertions: []test.Assertion{
				test.IsTrue{
					Actual: agnostic.VariableValue("output"),
				},
			},
		},
		{
			Name: "IfElseWhenConditionIsTrue",
			Before: []agnostic.Statement{
				agnostic.Declare{
					Name:  "input",
					Value: agnostic.BooleanLiteralValue(true),
				},
				agnostic.Declare{
					Name:  "ifOutput",
					Value: agnostic.BooleanLiteralValue(false),
				},
				agnostic.Declare{
					Name:  "elseOutput",
					Value: agnostic.BooleanLiteralValue(false),
				},
				agnostic.IfElse{
					Condition: agnostic.VariableValue("input"),
					TrueStatements: []agnostic.Statement{
						agnostic.AssignVar{
							Var:   agnostic.VariableValue("ifOutput"),
							Value: agnostic.BooleanLiteralValue(true),
						},
					},
					FalseStatements: []agnostic.Statement{
						agnostic.AssignVar{
							Var:   agnostic.VariableValue("elseOutput"),
							Value: agnostic.BooleanLiteralValue(true),
						},
					},
				},
			},
			Assertions: []test.Assertion{
				test.IsTrue{
					Actual: agnostic.VariableValue("ifOutput"),
				},
				test.IsFalse{
					Actual: agnostic.VariableValue("elseOutput"),
				},
			},
		},
		{
			Name: "IfElseWhenConditionIsFalse",
			Before: []agnostic.Statement{
				agnostic.Declare{
					Name:  "input",
					Value: agnostic.BooleanLiteralValue(false),
				},
				agnostic.Declare{
					Name:  "ifOutput",
					Value: agnostic.BooleanLiteralValue(false),
				},
				agnostic.Declare{
					Name:  "elseOutput",
					Value: agnostic.BooleanLiteralValue(false),
				},
				agnostic.IfElse{
					Condition: agnostic.VariableValue("input"),
					TrueStatements: []agnostic.Statement{
						agnostic.AssignVar{
							Var:   agnostic.VariableValue("ifOutput"),
							Value: agnostic.BooleanLiteralValue(true),
						},
					},
					FalseStatements: []agnostic.Statement{
						agnostic.AssignVar{
							Var:   agnostic.VariableValue("elseOutput"),
							Value: agnostic.BooleanLiteralValue(true),
						},
					},
				},
			},
			Assertions: []test.Assertion{
				test.IsFalse{
					Actual: agnostic.VariableValue("ifOutput"),
				},
				test.IsTrue{
					Actual: agnostic.VariableValue("elseOutput"),
				},
			},
		},
	},
}
