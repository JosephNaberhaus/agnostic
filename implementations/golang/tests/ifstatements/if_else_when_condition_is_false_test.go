package ifstatements

import (
  "github.com/stretchr/testify/assert"
  "testing"
)

func TestGolang_IfElseWhenConditionIsFalse(t *testing.T) {
  input := false
  _ = input
  ifOutput := false
  _ = ifOutput
  elseOutput := false
  _ = elseOutput
  if input {
    ifOutput = true
  } else {
    elseOutput = true
  }
  assert.False(t, ifOutput)
  assert.True(t, elseOutput)
}
