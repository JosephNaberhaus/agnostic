package agnostic

import (
	"github.com/JosephNaberhaus/agnostic/implementation/languages/java"
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

	mapped, err := Map(result, java.Mapper{})
	require.NoError(t, err)
	println(mapped)
}
