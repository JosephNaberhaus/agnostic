package operator

import (
  "github.com/stretchr/testify/assert"
  "testing"
)

func TestGolang_LessThan(t *testing.T) {
  assert.True(t, 9 < 10)
}
