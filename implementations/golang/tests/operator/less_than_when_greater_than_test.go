package operator

import (
  "github.com/stretchr/testify/assert"
  "testing"
)

func TestGolang_LessThanWhenGreaterThan(t *testing.T) {
  assert.False(t, 6 < 5)
}
