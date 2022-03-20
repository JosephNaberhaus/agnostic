package ifstatements

import (
  "github.com/stretchr/testify/assert"
  "testing"
)

func TestGolang_IfWhenConditionIsFalse(t *testing.T) {
  input := false
  _ = input
  output := false
  _ = output
  if input {
    output = true
  }
  assert.False(t, output)
}
