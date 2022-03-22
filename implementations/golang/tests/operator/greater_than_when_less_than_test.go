package operator

import (
  "github.com/stretchr/testify/assert"
  "testing"
)

func TestGolang_GreaterThanWhenLessThan(t *testing.T) {
  assert.False(t, 5 > 6)
}
