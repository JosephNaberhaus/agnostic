package test

import "github.com/JosephNaberhaus/agnostic"

type Assertion interface {
	isAssertion()
}

type IsEqual struct {
	Expected, Actual agnostic.Value
}

// isAssertion implements the Assertion interface so that IsEqual can be used as an Assertion
func (i IsEqual) isAssertion() {}

type IsNotEqual struct {
	NotExpected, Actual agnostic.Value
}

// isAssertion implements the Assertion interface so that IsNotEqual can be used as an Assertion
func (i IsNotEqual) isAssertion() {}

type IsTrue struct {
	Actual agnostic.Value
}

// isAssertion implements the Assertion interface so that IsNotEqual can be used as an Assertion
func (i IsTrue) isAssertion() {}

type IsFalse struct {
	Actual agnostic.Value
}

// isAssertion implements the Assertion interface so that IsNotEqual can be used as an Assertion
func (i IsFalse) isAssertion() {}
