package agnostic

import (
	"github.com/JosephNaberhaus/agnostic/ast"
	"github.com/JosephNaberhaus/agnostic/implementation/languages/java"
	"github.com/JosephNaberhaus/agnostic/implementation/test/constant"
	"github.com/JosephNaberhaus/agnostic/implementation/test/empty"
	"github.com/JosephNaberhaus/agnostic/implementation/test/erroring"
	"github.com/JosephNaberhaus/agnostic/implementation/text"
	"github.com/JosephNaberhaus/agnostic/language/parser"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"testing"
)

func TestMap(t *testing.T) {
	input := ast.Module{}
	expectedOutput := "test output"

	mapper := &constant.Mapper{Constant: text.Span(expectedOutput)}
	output, err := Map(input, mapper)
	require.NoError(t, err)
	assert.Equal(t, expectedOutput, output)
}

func TestMap_AstToCodeError(t *testing.T) {
	input := ast.Module{
		Models: []ast.ModelDef{
			{
				Name: "model with an invalid name",
			},
		},
	}

	mapper := &empty.Mapper{}
	_, err := Map(input, mapper)
	assert.Error(t, err)
}

func TestMap_MapperError(t *testing.T) {
	input := ast.Module{}

	mapper := &erroring.Mapper{}
	_, err := Map(input, mapper)
	require.ErrorIs(t, err, erroring.Error)
}

func TestTest(t *testing.T) {
	data, err := ioutil.ReadFile("test.agnostic")
	require.NoError(t, err)

	result, err := parser.ToAst(string(data))
	require.NoError(t, err)

	mapped, err := Map(result, java.Mapper{})
	require.NoError(t, err)
	println(mapped)
}
