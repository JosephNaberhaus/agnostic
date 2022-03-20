package operator

import (
  "github.com/stretchr/testify/assert"
  "testing"
)

func TestGolang_NotGreaterThan(t *testing.T) {
  assert.False(t, 5 > 6)
}
