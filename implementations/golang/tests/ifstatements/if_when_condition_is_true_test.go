package ifstatements

import (
  "github.com/stretchr/testify/assert"
  "testing"
)

func TestGolang_IfWhenConditionIsTrue(t *testing.T) {
  input := true
  _ = input
  output := false
  _ = output
  if input {
    output = true
  }
  assert.True(t, output)
}
