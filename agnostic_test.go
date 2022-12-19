package agnostic

import (
	"fmt"
	"github.com/JosephNaberhaus/agnostic/code/mappers"
	"github.com/JosephNaberhaus/agnostic/implementation/languages/java"
	"github.com/JosephNaberhaus/agnostic/interpreter"
	"github.com/JosephNaberhaus/agnostic/language/lexer"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"testing"
)

func TestTest(t *testing.T) {
	data, err := ioutil.ReadFile("test.agnostic")
	require.NoError(t, err)

	result, err := lexer.Test(string(data))
	require.NoError(t, err)

	mapped, err := Map(result, &java.Mapper{})
	require.NoError(t, err)
	println(mapped)
}

func TestInterpreter(t *testing.T) {
	data, err := ioutil.ReadFile("test.agnostic")
	require.NoError(t, err)

	result, err := lexer.Test(string(data))
	require.NoError(t, err)

	codeModule, err := mappers.AstToCode(result)
	require.NoError(t, err)

	input := &interpreter.List{}
	input.Push(interpreter.Value{Value: "30373"})
	input.Push(interpreter.Value{Value: "25512"})
	input.Push(interpreter.Value{Value: "65332"})
	input.Push(interpreter.Value{Value: "33549"})
	input.Push(interpreter.Value{Value: "35390"})

	fmt.Printf("%v\n", interpreter.Run(codeModule, "partOne", interpreter.Value{Value: input}))
}
