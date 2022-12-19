package lexer

import (
	"fmt"
	"testing"
)

func TestThisThing(t *testing.T) {
	state := newRunes("Point{\n                        x: p.x + moveDir.x,\n                        y: p.y + moveDir.y}\n                    }")
	endState, result, err := valueConsumer()(state)
	fmt.Printf("done! %v %v %v", endState, result, err)
}
