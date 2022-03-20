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
						Left:     agnostic.IntLiteralValue(10),
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
					Expected: agnostic.IntLiteralValue(1),
					Actual: agnostic.ComputedValue{
						Left:     agnostic.IntLiteralValue(10),
						Operator: agnostic.Modulo,
						Right:    agnostic.IntLiteralValue(3),
					},
				},
			},
		},
		{
			Name: "EqualInt",
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
			Name: "UnequalInt",
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
			Name: "EqualString",
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
			Name: "UnequalInt",
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
			Name: "GreaterThan",
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
			Name: "GreaterThanEqual",
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
			Name: "NotGreaterThan",
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
			Name: "GreaterThanOrEqualToEqual",
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
			Name: "GreaterThanOrEqualToGreatThan",
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
			Name: "NotGreaterThanOrEqual",
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
			Name: "LessThan",
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
			Name: "LessThanEqual",
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
			Name: "NotLessThan",
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
			Name: "LessThanOrEqualToEqual",
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
			Name: "LessThanOrEqualToGreatThan",
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
			Name: "NotLessThanOrEqual",
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
