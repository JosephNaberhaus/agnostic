package agnostic

import (
	"fmt"
	"github.com/JosephNaberhaus/agnostic/code/mappers"
	"github.com/JosephNaberhaus/agnostic/implementation/languages/java"
	"github.com/JosephNaberhaus/agnostic/interpreter"
	"github.com/JosephNaberhaus/agnostic/language/lexer"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestTest(t *testing.T) {
	data, err := os.ReadFile("resources/test_parser.agn")
	require.NoError(t, err)

	result, _, err := lexer.Parse(string(data))
	require.NoError(t, err)

	mapped, err := Map(result, &java.Mapper{})
	require.NoError(t, err)
	println(mapped)
}

func TestInterpreter(t *testing.T) {
	data, err := os.ReadFile("resources/test_parser.agn")
	require.NoError(t, err)

	result, _, err := lexer.Parse(string(data))
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

func TestTokenMeta(t *testing.T) {
	data, err := os.ReadFile("resources/test_token_meta.agn")
	require.NoError(t, err)

	_, actualTokens, err := lexer.Parse(string(data))
	require.NoError(t, err)

	expectedTokens := []lexer.TokenMeta{
		{Start: 0, End: 6, Kind: lexer.TokenKind_keyword},
		{Start: 7, End: 11, Kind: lexer.TokenKind_module},
		{Start: 13, End: 21, Kind: lexer.TokenKind_function},
		{Start: 29, End: 35, Kind: lexer.TokenKind_keyword},
		{Start: 38, End: 41, Kind: lexer.TokenKind_keyword},
		{Start: 65, End: 68, Kind: lexer.TokenKind_keyword},
		{Start: 114, End: 116, Kind: lexer.TokenKind_keyword},
	}

	assert.Equal(t, expectedTokens, actualTokens)
}
