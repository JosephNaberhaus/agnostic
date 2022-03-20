package operator

import (
  "github.com/stretchr/testify/assert"
  "testing"
)

func TestGolang_GreaterThanOrEqualToGreatThan(t *testing.T) {
  assert.True(t, 6 >= 3)
}
