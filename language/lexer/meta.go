package lexer

type TokenMeta struct {
	Start int
	End   int
	Kind  TokenKind
}

type TokenKind int

const (
	TokenKind_keyword TokenKind = iota
	TokenKind_type
	TokenKind_module
	TokenKind_function
	TokenKind_comment
)

func meta[T any](consumer consumer[T], kind TokenKind) consumer[T] {
	return func(state parserState) (parserState, T, error) {
		start := state.numConsumed
		state, val, err := consumer(state)
		end := state.numConsumed

		if start < end {
			state.tokens = append(state.tokens, TokenMeta{
				Start: start,
				End:   end,
				Kind:  kind,
			})
		}

		return state, val, err
	}
}
