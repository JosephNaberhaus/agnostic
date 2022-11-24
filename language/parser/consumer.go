package parser

import "github.com/JosephNaberhaus/agnostic/language/token"

type void = *struct{}

type parser[T any] func(t tokenQueue) (tokenQueue, T, error)

func skipTypeConsumer(tokenType token.Type) parser[void] {
	return func(t tokenQueue) (tokenQueue, void, error) {
		consumed := t.consumeOne()
		if consumed.Type != tokenType {

		}
	}
}

func first[T any](parsers ...parser[T]) parser[T] {
	return func(t tokenQueue) (tokenQueue, T, error) {
		// TODO: take furthest
		var lastError error
		for _, parser := range parsers {
			newTokens, result, err := parser(t)
			if err != nil {
				lastError = err
				continue
			}

			return newTokens, result, nil
		}

		var zero T
		return tokenQueue{}, zero, lastError
	}
}

func inOrder(parsers ...parser[void]) parser[void] {
	return func(t tokenQueue) (tokenQueue, void, error) {
		for _, parser := range parsers {
			newTokens, _, err := parser(t)
			if err != nil {
				return tokenQueue{}, nil, err
			}

			t = newTokens
		}

		return t, nil, nil
	}
}

func handle[T any](parser parser[T], handler func(T)) parser[void] {
	return func(t tokenQueue) (tokenQueue, void, error) {
		newTokens, result, err := parser(t)
		if err != nil {
			return tokenQueue{}, nil, err
		}

		handler(result)

		return newTokens, nil, err
	}
}
