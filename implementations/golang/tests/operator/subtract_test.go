package operator

import (
  "github.com/stretchr/testify/assert"
  "testing"
)

func TestGolang_Subtract(t *testing.T) {
  assert.Equal(t, 42, 50 - 8)
}
