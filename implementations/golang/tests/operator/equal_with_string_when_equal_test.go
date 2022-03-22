package operator

import (
  "github.com/stretchr/testify/assert"
  "testing"
)

func TestGolang_EqualWithStringWhenEqual(t *testing.T) {
  assert.True(t, "test" == "test")
}
