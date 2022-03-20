package operator

import (
  "github.com/stretchr/testify/assert"
  "testing"
)

func TestGolang_EqualString(t *testing.T) {
  assert.True(t, "test" == "test")
}
