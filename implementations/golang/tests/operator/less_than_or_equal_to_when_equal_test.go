package operator

import (
  "github.com/stretchr/testify/assert"
  "testing"
)

func TestGolang_LessThanOrEqualToWhenEqual(t *testing.T) {
  assert.True(t, 10 <= 10)
}
