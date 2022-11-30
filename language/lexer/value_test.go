package lexer

import (
	"fmt"
	"testing"
)

func TestThisThing(t *testing.T) {
	state := newRunes("(new test()).test")
	endState, result, err := valueConsumer()(state)
	fmt.Printf("done! %v %v %v", endState, result, err)
}
