package operator

import (
  "github.com/stretchr/testify/assert"
  "testing"
)

func TestGolang_EqualWhenEqual(t *testing.T) {
  assert.True(t, 3 == 3)
}
