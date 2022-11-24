package lexer

// A meaningless rune value that is used when a method must return a rune, but there is no rune to return.
const nilRune rune = 0xF001

type runes struct {
	numConsumed int
	text        []rune
}

func newRunes(text string) runes {
	return runes{text: []rune(text)}
}

func (r *runes) isEmpty() bool {
	return len(r.text) == 0
}

func (r *runes) isNotEmpty() bool {
	return !r.isEmpty()
}

func (r *runes) numRemaining() int {
	return len(r.text)
}

func (r *runes) peek(n int) []rune {
	if n > len(r.text) {
		n = len(r.text)
	}

	return r.text[:n]
}

func (r *runes) peekOne() rune {
	if r.isEmpty() {
		return nilRune
	}

	return r.peek(1)[0]
}

func (r *runes) peekStr(n int) string {
	return string(r.peek(n))
}

func (r *runes) consume(n int) []rune {
	if n > len(r.text) {
		n = len(r.text)
	}

	consumed := r.text[:n]
	r.text = r.text[n:]
	r.numConsumed += n

	return consumed
}

func (r *runes) consumeOne() rune {
	if r.isEmpty() {
		return nilRune
	}

	return r.consume(1)[0]
}

func (r *runes) consumeStr(n int) string {
	return string(r.consume(n))
}

func (r *runes) consumeWhile(condition func(r rune) bool) []rune {
	var consumed []rune
	for r.isNotEmpty() && condition(r.peekOne()) {
		consumed = append(consumed, r.consumeOne())
	}

	return consumed
}

func (r *runes) consumeWhileStr(condition func(r rune) bool) string {
	return string(r.consumeWhile(condition))
}

func (r *runes) consumeUntil(condition func(r rune) bool) []rune {
	return r.consumeWhile(func(r rune) bool {
		return !condition(r)
	})
}

func (r *runes) consumeUntilStr(condition func(r rune) bool) string {
	return string(r.consumeUntil(condition))
}
