package operator

import (
  "github.com/stretchr/testify/assert"
  "testing"
)

func TestGolang_IntegerDivisionRoundsDown(t *testing.T) {
  assert.Equal(t, 3, 10 / 3)
}
