package test

import "github.com/JosephNaberhaus/agnostic"

type Suite struct {
	Name   string
	Models []agnostic.Model
	Tests  []Test
}

type Test struct {
	Name       string
	Before     []agnostic.Statement
	Assertions []Assertion
}
