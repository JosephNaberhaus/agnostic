package operator

import (
  "github.com/stretchr/testify/assert"
  "testing"
)

func TestGolang_NotGreaterThanOrEqual(t *testing.T) {
  assert.False(t, 3 >= 5)
}
