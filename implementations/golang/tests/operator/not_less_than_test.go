package operator

import (
  "github.com/stretchr/testify/assert"
  "testing"
)

func TestGolang_NotLessThan(t *testing.T) {
  assert.False(t, 6 < 5)
}
