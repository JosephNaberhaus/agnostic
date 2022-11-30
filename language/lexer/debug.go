package lexer

import (
	"fmt"
	"strings"
)

func debugPrint(message string) consumer[void] {
	return func(text parserState) (parserState, void, error) {
		println(message)
		return text, nil, nil
	}
}

func debugPrintRemaining() consumer[void] {
	return func(text parserState) (parserState, void, error) {
		var output strings.Builder

		output.WriteRune('[')
		for i, r := range text.remaining {
			output.WriteString(fmt.Sprintf("%#v", string(r)))
			if i != len(text.remaining)-1 {
				output.WriteString(", ")
			}
		}
		output.WriteRune(']')

		println(output.String())
		return text, nil, nil
	}
}
