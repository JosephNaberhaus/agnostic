package lexer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRunes_IsEmpty(t *testing.T) {
	var r runes
	assert.True(t, r.isEmpty())

	r.text = []rune("four")
	assert.False(t, r.isEmpty())
}

func TestRunes_IsNotEmpty(t *testing.T) {
	var r runes
	assert.False(t, r.isNotEmpty())

	r.text = []rune("four")
	assert.True(t, r.isNotEmpty())
}

func TestRunes_NumRemaining(t *testing.T) {
	var r runes
	assert.Zero(t, r.numRemaining())

	r.text = []rune("four")
	assert.Equal(t, 4, r.numRemaining())
}

func TestRunes_Peek(t *testing.T) {
	var r runes
	assert.Empty(t, r.peek(0))
	assert.Empty(t, r.peek(2))

	r.text = []rune("four")
	assert.Empty(t, r.peek(0))
	assert.Equal(t, []rune("fo"), r.peek(2))
	assert.Equal(t, []rune("four"), r.peek(4))
	assert.Equal(t, []rune("four"), r.peek(6))
}

func TestRunes_PeekOne(t *testing.T) {
	var r runes
	assert.Equal(t, nilRune, r.peekOne())

	r.text = []rune("four")
	assert.Equal(t, 'f', r.peekOne())
}

func TestRunes_PeekStr(t *testing.T) {
	var r runes
	assert.Equal(t, "", r.peekStr(0))
	assert.Equal(t, "", r.peekStr(2))

	r.text = []rune("four")
	assert.Equal(t, "", r.peekStr(0))
	assert.Equal(t, "fo", r.peekStr(2))
	assert.Equal(t, "four", r.peekStr(4))
	assert.Equal(t, "four", r.peekStr(6))
}

func TestRunes_Consume(t *testing.T) {
	var r runes
	assert.Empty(t, r.consume(0))
	assert.Empty(t, r.consume(2))

	r.text = []rune("fourfourfour")
	assert.Equal(t, []rune("four"), r.consume(4))
	assert.Equal(t, 4, r.numConsumed)
	assert.Equal(t, []rune("fourfour"), r.consume(8))
	assert.Equal(t, 12, r.numConsumed)
	assert.Empty(t, r.consume(4))
	assert.Equal(t, 12, r.numConsumed)
}

func TestRunes_ConsumeOne(t *testing.T) {
	var r runes
	assert.Equal(t, nilRune, r.consumeOne())

	r.text = []rune("four")
	assert.Equal(t, 'f', r.consumeOne())
	assert.Equal(t, 1, r.numConsumed)
	assert.Equal(t, 'o', r.consumeOne())
	assert.Equal(t, 2, r.numConsumed)
}

func TestRunes_ConsumeStr(t *testing.T) {
	var r runes
	assert.Equal(t, "", r.consumeStr(0))
	assert.Equal(t, "", r.consumeStr(2))

	r.text = []rune("fourfourfour")
	assert.Equal(t, "four", r.consumeStr(4))
	assert.Equal(t, 4, r.numConsumed)
	assert.Equal(t, "fourfour", r.consumeStr(8))
	assert.Equal(t, 12, r.numConsumed)
	assert.Equal(t, "", r.consumeStr(4))
	assert.Equal(t, 12, r.numConsumed)
}

func TestRunes_ConsumeWhile(t *testing.T) {
	alwaysTrue := func(r rune) bool {
		return true
	}
	alwaysFalse := func(r rune) bool {
		return false
	}
	trueNTimes := func(n int) func(r rune) bool {
		cur := 0
		return func(r rune) bool {
			cur++
			return cur <= n
		}
	}

	var r runes
	assert.Empty(t, r.consumeWhile(alwaysFalse))
	assert.Empty(t, r.consumeWhile(alwaysTrue))

	r.text = []rune("fourfourfour")
	assert.Empty(t, r.consumeWhile(alwaysFalse))
	assert.Equal(t, []rune("fourfour"), r.consumeWhile(trueNTimes(8)))
	assert.Equal(t, []rune("four"), r.consumeWhile(alwaysTrue))
}

func TestRunes_ConsumeWhileStr(t *testing.T) {
	alwaysTrue := func(r rune) bool {
		return true
	}
	alwaysFalse := func(r rune) bool {
		return false
	}
	trueNTimes := func(n int) func(r rune) bool {
		cur := 0
		return func(r rune) bool {
			cur++
			return cur <= n
		}
	}

	var r runes
	assert.Equal(t, "", r.consumeWhileStr(alwaysFalse))
	assert.Equal(t, "", r.consumeWhileStr(alwaysTrue))

	r.text = []rune("fourfourfour")
	assert.Equal(t, "", r.consumeWhileStr(alwaysFalse))
	assert.Equal(t, "fourfour", r.consumeWhileStr(trueNTimes(8)))
	assert.Equal(t, "four", r.consumeWhileStr(alwaysTrue))
}

func TestRunes_ConsumeUntil(t *testing.T) {
	alwaysTrue := func(r rune) bool {
		return true
	}
	alwaysFalse := func(r rune) bool {
		return false
	}
	falseNTimes := func(n int) func(r rune) bool {
		cur := 0
		return func(r rune) bool {
			cur++
			return cur > n
		}
	}

	var r runes
	assert.Empty(t, r.consumeUntil(alwaysTrue))
	assert.Empty(t, r.consumeUntil(alwaysFalse))

	r.text = []rune("fourfourfour")
	assert.Empty(t, r.consumeUntil(alwaysTrue))
	assert.Equal(t, []rune("fourfour"), r.consumeUntil(falseNTimes(8)))
	assert.Equal(t, []rune("four"), r.consumeUntil(alwaysFalse))
}

func TestRunes_ConsumeUntilStr(t *testing.T) {
	alwaysTrue := func(r rune) bool {
		return true
	}
	alwaysFalse := func(r rune) bool {
		return false
	}
	falseNTimes := func(n int) func(r rune) bool {
		cur := 0
		return func(r rune) bool {
			cur++
			return cur > n
		}
	}

	var r runes
	assert.Equal(t, "", r.consumeUntilStr(alwaysTrue))
	assert.Equal(t, "", r.consumeUntilStr(alwaysFalse))

	r.text = []rune("fourfourfour")
	assert.Equal(t, "", r.consumeUntilStr(alwaysTrue))
	assert.Equal(t, "fourfour", r.consumeUntilStr(falseNTimes(8)))
	assert.Equal(t, "four", r.consumeUntilStr(alwaysFalse))
}
