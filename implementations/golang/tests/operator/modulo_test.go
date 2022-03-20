package operator

import (
  "github.com/stretchr/testify/assert"
  "testing"
)

func TestGolang_Modulo(t *testing.T) {
  assert.Equal(t, 1, 10 % 3)
}
