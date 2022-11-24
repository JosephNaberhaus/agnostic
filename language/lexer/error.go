package lexer

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Error struct {
	pos     int
	message string
}

func (e Error) Error() string {
	return fmt.Sprintf("lexer error: %s", e.message)
}

func createError(r runes, message string) Error {
	return Error{
		pos:     r.numConsumed,
		message: message,
	}
}

// takeFurthest attempts to cast both errors to the Error type, and then returns the one that has the furthest position.
func takeFurthest(first, second error) error {
	var firstError, secondError Error

	if !errors.As(first, &firstError) {
		return second
	}

	if !errors.As(second, &secondError) {
		return first
	}

	if firstError.pos > secondError.pos {
		return first
	} else {
		return second
	}
}

type ContextualError struct {
	cause   Error
	rawText []rune
}

func lineColNumber(pos int, rawText []rune) (int, int) {
	lineNumber := 1
	colNumber := 0
	for i := 0; i < pos; i++ {
		if isNewline(rawText[i]) {
			colNumber = 0
			lineNumber++
		} else {
			colNumber++
		}
	}

	return lineNumber, colNumber
}

func (c ContextualError) Error() string {
	const contextSize = 10

	contextStart := c.cause.pos - contextSize
	if contextStart < 0 {
		contextStart = 0
	} else {
		// Move the context start so that it includes the entire line.
		for contextStart > 0 {
			if c.rawText[contextStart] == '\n' {
				contextStart++
				break
			}

			contextStart--
		}
	}

	contextEnd := c.cause.pos + contextSize
	if contextEnd >= len(c.rawText) {
		contextEnd = len(c.rawText) - 1
	} else {
		// Move the context end so that it includes the entire line.
		for contextEnd < len(c.rawText)-1 {
			if c.rawText[contextEnd] == '\n' {
				contextEnd--
				break
			}

			contextEnd++
		}
	}

	context := c.rawText[contextStart : contextEnd+1]

	startLineNumber, _ := lineColNumber(contextStart, c.rawText)
	errorLineNumber, errorColNumber := lineColNumber(c.cause.pos, c.rawText)
	endLineNumber, _ := lineColNumber(contextEnd, c.rawText)

	maxLineNumberLength := len(strconv.Itoa(endLineNumber))

	var output strings.Builder
	output.WriteString("Error when reading file:\n")
	output.WriteRune('\n')
	for lineOffset, line := range strings.Split(string(context), "\n") {
		lineNumber := startLineNumber + lineOffset
		lineNumberStr := strconv.Itoa(lineNumber)
		if len(lineNumberStr) < maxLineNumberLength {
			lineNumberStr = strings.Repeat(" ", maxLineNumberLength-len(lineNumberStr)) + lineNumberStr
		}

		output.WriteString(lineNumberStr)
		output.WriteString(": ")
		output.WriteString(line)
		output.WriteRune('\n')

		if lineNumber == errorLineNumber {
			output.WriteString(strings.Repeat(" ", maxLineNumberLength+1+errorColNumber))
			output.WriteRune('^')
			output.WriteRune('\n')
		}
	}

	output.WriteRune('\n')
	output.WriteString(c.cause.message)
	output.WriteRune('\n')

	return output.String()
}

func (c ContextualError) Unwrap() error {
	return c.cause
}

func contextualize(err error, rawText []rune) error {
	var lexerError Error
	if errors.As(err, &lexerError) {
		return ContextualError{
			cause:   lexerError,
			rawText: rawText,
		}
	}

	return err
}
