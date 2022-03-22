package operator

import (
  "github.com/stretchr/testify/assert"
  "testing"
)

func TestGolang_LessThanOrEqualToWithGreaterThan(t *testing.T) {
  assert.False(t, 5 <= 3)
}
