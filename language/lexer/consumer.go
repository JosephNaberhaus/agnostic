package lexer

type void *struct{}

type consumer[T any] func(text parserState) (parserState, T, error)

// first creates a consumer that tries each consumer it's given and returns the result of the first that succeeds.
// If none of the consumers succeed, then it will return the error of the consumer that made it the furthest.
func first[T any](consumers ...consumer[T]) consumer[T] {
	return func(inputText parserState) (parserState, T, error) {
		var furthestError error
		for _, consumer := range consumers {
			outputText, result, err := consumer(inputText)
			if err != nil {
				furthestError = takeFurthest(err, furthestError)
				continue
			}

			return outputText, result, nil
		}

		var zero T
		return parserState{}, zero, furthestError
	}
}

// inOrder creates a consumer that runs each consumer it is given in order.
// If any of the consumers fail, then the error is returned.
func inOrder(consumers ...consumer[void]) consumer[void] {
	return func(text parserState) (parserState, void, error) {
		for _, consumer := range consumers {
			newText, _, err := consumer(text)
			if err != nil {
				return parserState{}, nil, err
			}

			text = newText
		}

		return text, nil, nil
	}
}

// optional creates a consumer that tries to run the consumer that it is given.
// If the consumer fails, then the error is swallowed.
func optional(consumer consumer[void]) consumer[void] {
	return func(text parserState) (parserState, void, error) {
		newText, result, err := consumer(text)
		if err != nil {
			text.addError(err)
			return text, nil, nil
		}

		return newText, result, nil
	}
}

func repeat(consumer consumer[void]) consumer[void] {
	return func(text parserState) (parserState, void, error) {
		for text.isNotEmpty() {
			newText, _, err := consumer(text)
			if err != nil {
				text.addError(err)
				return text, nil, nil
			}

			// Check if we've stalled.
			if newText.numConsumed == text.numConsumed {
				return parserState{}, nil, createError(newText, "stalled")
			}

			text = newText
		}

		return text, nil, nil
	}
}

func skip[T any](consumer consumer[T]) consumer[void] {
	return func(r parserState) (parserState, void, error) {
		newRunes, _, err := consumer(r)
		if err != nil {
			return parserState{}, nil, err
		}

		return newRunes, nil, nil
	}
}

func handle[T any](consumer consumer[T], handler func(T) error) consumer[void] {
	return func(r parserState) (parserState, void, error) {
		newTokens, result, err := consumer(r)
		if err != nil {
			return parserState{}, nil, err
		}

		err = handler(result)
		if err != nil {
			return parserState{}, nil, err
		}

		return newTokens, nil, nil
	}
}

func handleNoError[T any](consumer consumer[T], handler func(T)) consumer[void] {
	return func(r parserState) (parserState, void, error) {
		newTokens, result, err := consumer(r)
		if err != nil {
			return parserState{}, nil, err
		}

		handler(result)

		return newTokens, nil, nil
	}
}

func attempt[T any](result *T, consumer consumer[void]) consumer[T] {
	return func(text parserState) (parserState, T, error) {
		var zero T
		*result = zero

		newText, _, err := consumer(text)
		if err != nil {
			return parserState{}, zero, err
		}

		return newText, *result, nil
	}
}

func mapResult[T, V any](consumer consumer[T], mapper func(T) (V, error)) consumer[V] {
	return func(text parserState) (parserState, V, error) {
		newText, result, err := consumer(text)
		if err != nil {
			var zero V
			return parserState{}, zero, err
		}

		mappedResult, err := mapper(result)
		if err != nil {
			var zero V
			return parserState{}, zero, err
		}

		return newText, mappedResult, nil
	}
}

func mapResultToConstant[T, V any](consumer consumer[T], result V) consumer[V] {
	return mapResult(
		consumer,
		func(_ T) (V, error) {
			return result, nil
		},
	)
}

func deferred[T any](consumerFactory func() consumer[T]) consumer[T] {
	return func(text parserState) (parserState, T, error) {
		newText, result, err := consumerFactory()(text)
		if err != nil {
			var zero T
			return parserState{}, zero, err
		}

		return newText, result, nil
	}
}

func lookahead(consumer consumer[void]) consumer[void] {
	return func(state parserState) (parserState, void, error) {
		_, _, err := consumer(state)
		if err != nil {
			return parserState{}, nil, err
		}

		return state, nil, nil
	}
}
