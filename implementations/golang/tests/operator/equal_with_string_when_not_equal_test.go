package operator

import (
  "github.com/stretchr/testify/assert"
  "testing"
)

func TestGolang_EqualWithStringWhenNotEqual(t *testing.T) {
  assert.False(t, "test" == "hello")
}
