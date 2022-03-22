package tests

import (
	"github.com/JosephNaberhaus/agnostic"
	"github.com/JosephNaberhaus/agnostic/test"
)

var operatorSuite = test.Suite{
	Name: "Operator",
	Tests: []test.Test{
		{
			Name: "Add",
			Assertions: []test.Assertion{
				test.IsEqual{
					Expected: agnostic.IntLiteralValue(10),
					Actual: agnostic.ComputedValue{
						Left:     agnostic.IntLiteralValue(8),
						Operator: agnostic.Add,
						Right:    agnostic.IntLiteralValue(2),
					},
				},
			},
		},
		{
			Name: "Subtract",
			Assertions: []test.Assertion{
				test.IsEqual{
					Expected: agnostic.IntLiteralValue(42),
					Actual: agnostic.ComputedValue{
						Left:     agnostic.IntLiteralValue(50),
						Operator: agnostic.Subtract,
						Right:    agnostic.IntLiteralValue(8),
					},
				},
			},
		},
		{
			Name: "Multiply",
			Assertions: []test.Assertion{
				test.IsEqual{
					Expected: agnostic.IntLiteralValue(48),
					Actual: agnostic.ComputedValue{
						Left:     agnostic.IntLiteralValue(4),
						Operator: agnostic.Multiply,
						Right:    agnostic.IntLiteralValue(12),
					},
				},
			},
		},
		{
			Name: "IntegerDivision",
			Assertions: []test.Assertion{
				test.IsEqual{
					Expected: agnostic.IntLiteralValue(12),
					Actual: agnostic.ComputedValue{
						Left:     agnostic.IntLiteralValue(120),
						Operator: agnostic.IntegerDivision,
						Right:    agnostic.IntLiteralValue(10),
					},
				},
			},
		},
		{
			Name: "IntegerDivisionRoundsDown",
			Assertions: []test.Assertion{
				test.IsEqual{
					Expected: agnostic.IntLiteralValue(3),
					Actual: agnostic.ComputedValue{
						Left:     agnostic.IntLiteralValue(11),
						Operator: agnostic.IntegerDivision,
						Right:    agnostic.IntLiteralValue(3),
					},
				},
			},
		},
		{
			Name: "Modulo",
			Assertions: []test.Assertion{
				test.IsEqual{
					Expected: agnostic.IntLiteralValue(2),
					Actual: agnostic.ComputedValue{
						Left:     agnostic.IntLiteralValue(11),
						Operator: agnostic.Modulo,
						Right:    agnostic.IntLiteralValue(3),
					},
				},
			},
		},
		{
			Name: "EqualWhenEqual",
			Assertions: []test.Assertion{
				test.IsTrue{
					Actual: agnostic.ComputedValue{
						Left:     agnostic.IntLiteralValue(3),
						Operator: agnostic.Equal,
						Right:    agnostic.IntLiteralValue(3),
					},
				},
			},
		},
		{
			Name: "EqualWhenNotEqual",
			Assertions: []test.Assertion{
				test.IsFalse{
					Actual: agnostic.ComputedValue{
						Left:     agnostic.IntLiteralValue(1),
						Operator: agnostic.Equal,
						Right:    agnostic.IntLiteralValue(42),
					},
				},
			},
		},
		{
			Name: "EqualWithStringWhenEqual",
			Assertions: []test.Assertion{
				test.IsTrue{
					Actual: agnostic.ComputedValue{
						Left:     agnostic.StringLiteralValue("test"),
						Operator: agnostic.Equal,
						Right:    agnostic.StringLiteralValue("test"),
					},
				},
			},
		},
		{
			Name: "EqualWithStringWhenNotEqual",
			Assertions: []test.Assertion{
				test.IsFalse{
					Actual: agnostic.ComputedValue{
						Left:     agnostic.StringLiteralValue("test"),
						Operator: agnostic.Equal,
						Right:    agnostic.StringLiteralValue("hello"),
					},
				},
			},
		},
		{
			Name: "GreaterThanWhenGreaterThan",
			Assertions: []test.Assertion{
				test.IsTrue{
					Actual: agnostic.ComputedValue{
						Left:     agnostic.IntLiteralValue(10),
						Operator: agnostic.GreaterThan,
						Right:    agnostic.IntLiteralValue(9),
					},
				},
			},
		},
		{
			Name: "GreaterThanWhenEqual",
			Assertions: []test.Assertion{
				test.IsFalse{
					Actual: agnostic.ComputedValue{
						Left:     agnostic.IntLiteralValue(10),
						Operator: agnostic.GreaterThan,
						Right:    agnostic.IntLiteralValue(10),
					},
				},
			},
		},
		{
			Name: "GreaterThanWhenLessThan",
			Assertions: []test.Assertion{
				test.IsFalse{
					Actual: agnostic.ComputedValue{
						Left:     agnostic.IntLiteralValue(5),
						Operator: agnostic.GreaterThan,
						Right:    agnostic.IntLiteralValue(6),
					},
				},
			},
		},
		{
			Name: "GreaterThanOrEqualToWhenGreaterThan",
			Assertions: []test.Assertion{
				test.IsTrue{
					Actual: agnostic.ComputedValue{
						Left:     agnostic.IntLiteralValue(6),
						Operator: agnostic.GreaterThanOrEqualTo,
						Right:    agnostic.IntLiteralValue(3),
					},
				},
			},
		},
		{
			Name: "GreaterThanOrEqualToWhenEqual",
			Assertions: []test.Assertion{
				test.IsTrue{
					Actual: agnostic.ComputedValue{
						Left:     agnostic.IntLiteralValue(10),
						Operator: agnostic.GreaterThanOrEqualTo,
						Right:    agnostic.IntLiteralValue(10),
					},
				},
			},
		},
		{
			Name: "GreaterThanOrEqualToWhenLessThan",
			Assertions: []test.Assertion{
				test.IsFalse{
					Actual: agnostic.ComputedValue{
						Left:     agnostic.IntLiteralValue(3),
						Operator: agnostic.GreaterThanOrEqualTo,
						Right:    agnostic.IntLiteralValue(5),
					},
				},
			},
		},
		{
			Name: "LessThanWhenLessThan",
			Assertions: []test.Assertion{
				test.IsTrue{
					Actual: agnostic.ComputedValue{
						Left:     agnostic.IntLiteralValue(9),
						Operator: agnostic.LessThan,
						Right:    agnostic.IntLiteralValue(10),
					},
				},
			},
		},
		{
			Name: "LessThanWhenEqual",
			Assertions: []test.Assertion{
				test.IsFalse{
					Actual: agnostic.ComputedValue{
						Left:     agnostic.IntLiteralValue(10),
						Operator: agnostic.LessThan,
						Right:    agnostic.IntLiteralValue(10),
					},
				},
			},
		},
		{
			Name: "LessThanWhenGreaterThan",
			Assertions: []test.Assertion{
				test.IsFalse{
					Actual: agnostic.ComputedValue{
						Left:     agnostic.IntLiteralValue(6),
						Operator: agnostic.LessThan,
						Right:    agnostic.IntLiteralValue(5),
					},
				},
			},
		},
		{
			Name: "LessThanOrEqualToWithLessThan",
			Assertions: []test.Assertion{
				test.IsTrue{
					Actual: agnostic.ComputedValue{
						Left:     agnostic.IntLiteralValue(3),
						Operator: agnostic.LessThanOrEqualTo,
						Right:    agnostic.IntLiteralValue(6),
					},
				},
			},
		},
		{
			Name: "LessThanOrEqualToWhenEqual",
			Assertions: []test.Assertion{
				test.IsTrue{
					Actual: agnostic.ComputedValue{
						Left:     agnostic.IntLiteralValue(10),
						Operator: agnostic.LessThanOrEqualTo,
						Right:    agnostic.IntLiteralValue(10),
					},
				},
			},
		},
		{
			Name: "LessThanOrEqualToWithGreaterThan",
			Assertions: []test.Assertion{
				test.IsFalse{
					Actual: agnostic.ComputedValue{
						Left:     agnostic.IntLiteralValue(5),
						Operator: agnostic.LessThanOrEqualTo,
						Right:    agnostic.IntLiteralValue(3),
					},
				},
			},
		},
	},
}
