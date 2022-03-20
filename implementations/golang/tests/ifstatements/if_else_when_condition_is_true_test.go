package ifstatements

import (
  "github.com/stretchr/testify/assert"
  "testing"
)

func TestGolang_IfElseWhenConditionIsTrue(t *testing.T) {
  input := true
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
  assert.True(t, ifOutput)
  assert.False(t, elseOutput)
}
