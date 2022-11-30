package lexer

// A meaningless rune value that is used when a method must return a rune, but there is no rune to return.
const nilRune rune = 0xF001

type parserState struct {
	numConsumed int

	remaining []rune

	furthestIgnorableError error
}

func newRunes(text string) parserState {
	return parserState{
		remaining: []rune(text),
	}
}

func (p *parserState) addError(err error) {
	p.furthestIgnorableError = takeFurthest(err, p.furthestIgnorableError)
}

func (p *parserState) isEmpty() bool {
	return len(p.remaining) == 0
}

func (p *parserState) isNotEmpty() bool {
	return !p.isEmpty()
}

func (p *parserState) numRemaining() int {
	return len(p.remaining)
}

func (p *parserState) peek(n int) []rune {
	if n > len(p.remaining) {
		n = len(p.remaining)
	}

	return p.remaining[:n]
}

func (p *parserState) peekOne() rune {
	if p.isEmpty() {
		return nilRune
	}

	return p.peek(1)[0]
}

func (p *parserState) peekStr(n int) string {
	return string(p.peek(n))
}

func (p *parserState) consume(n int) []rune {
	if n > len(p.remaining) {
		n = len(p.remaining)
	}

	consumed := p.remaining[:n]
	p.remaining = p.remaining[n:]

	p.numConsumed += n

	return consumed
}

func (p *parserState) consumeOne() rune {
	if p.isEmpty() {
		return nilRune
	}

	return p.consume(1)[0]
}

func (p *parserState) consumeStr(n int) string {
	return string(p.consume(n))
}

func (p *parserState) consumeWhile(condition func(r rune) bool) []rune {
	var consumed []rune
	for p.isNotEmpty() && condition(p.peekOne()) {
		consumed = append(consumed, p.consumeOne())
	}

	return consumed
}

func (p *parserState) consumeWhileStr(condition func(r rune) bool) string {
	return string(p.consumeWhile(condition))
}

func (p *parserState) consumeUntil(condition func(r rune) bool) []rune {
	return p.consumeWhile(func(r rune) bool {
		return !condition(r)
	})
}

func (p *parserState) consumeUntilStr(condition func(r rune) bool) string {
	return string(p.consumeUntil(condition))
}
