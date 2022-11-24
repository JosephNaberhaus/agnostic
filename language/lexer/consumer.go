package lexer

import (
	"github.com/JosephNaberhaus/agnostic/language/token"
)

type consumer func(text runes) (runes, []token.Token, error)

// first creates a consumer that tries each consumer it's given and returns the result of the first that succeeds.
// If none of the consumers succeed, then it will return the error of the consumer that made it the furthest.
func first(consumers ...consumer) consumer {
	return func(inputText runes) (runes, []token.Token, error) {
		var furthestError error
		for _, consumer := range consumers {
			outputText, tokens, err := consumer(inputText)
			if err != nil {
				furthestError = takeFurthest(err, furthestError)
				continue
			}

			return outputText, tokens, nil
		}

		return runes{}, nil, furthestError
	}
}

// inOrder creates a consumer that runs each consumer it is given in order.
// If any of the consumers fail, then the error is returned.
func inOrder(consumers ...consumer) consumer {
	return func(text runes) (runes, []token.Token, error) {
		var tokens []token.Token
		for _, consumer := range consumers {
			newText, newTokens, err := consumer(text)
			if err != nil {
				return runes{}, nil, err
			}

			text = newText
			tokens = append(tokens, newTokens...)
		}

		return text, tokens, nil
	}
}

// optional creates a consumer that tries to run the consumer that it is given.
// If the consumer fails, then the error is swallowed.
func optional(consumer consumer) consumer {
	return func(text runes) (runes, []token.Token, error) {
		newText, newTokens, err := consumer(text)
		if err != nil {
			return text, nil, nil
		}

		return newText, newTokens, nil
	}
}

// repeatAndThen creates a consumer that repeatedly runs the first consumer until it fails, then it runs the second consumer.
// If the second consumer fails then the furthest error will be returned.
//
// repeatAndThen would just take a single consumer as a parameter, but it needs to be possible to detect when it should
// return an error from the first consumer. That is why it takes the next consumer as its second parameter.
func repeatAndThen(first, second consumer) consumer {
	return func(text runes) (runes, []token.Token, error) {
		var tokens []token.Token

		isDone := false
		for !isDone {
			newText, newTokens, firstErr := first(text)
			if firstErr != nil {
				var secondErr error
				// TODO make better
				newText, newTokens, secondErr = second(text)
				if secondErr != nil {
					return runes{}, nil, takeFurthest(firstErr, secondErr)
				}

				isDone = true
			}

			text = newText
			tokens = append(tokens, newTokens...)
		}

		return text, tokens, nil
	}
}
