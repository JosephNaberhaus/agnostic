package operator

import (
  "github.com/stretchr/testify/assert"
  "testing"
)

func TestGolang_GreaterThanOrEqualToWhenLessThan(t *testing.T) {
  assert.False(t, 3 >= 5)
}
