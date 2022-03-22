package operator

import (
  "github.com/stretchr/testify/assert"
  "testing"
)

func TestGolang_EqualWhenNotEqual(t *testing.T) {
  assert.False(t, 1 == 42)
}
