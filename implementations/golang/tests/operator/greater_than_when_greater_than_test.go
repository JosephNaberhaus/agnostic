package operator

import (
  "github.com/stretchr/testify/assert"
  "testing"
)

func TestGolang_GreaterThanWhenGreaterThan(t *testing.T) {
  assert.True(t, 10 > 9)
}
