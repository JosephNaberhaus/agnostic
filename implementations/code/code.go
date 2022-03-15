package writer

import (
	"fmt"
	"strings"
)

type Code interface {
	isCode()
}

type Line string

// isCode implements the Code interface so that Line can be used as Code
func (l Line) isCode() {}

type Group []Code

// isCode implements the Code interface so that Group can be used as Code
func (g Group) isCode() {}

type Block []Code

// isCode implements the Code interface so that Block can be used as Code
func (b Block) isCode() {}

type Nil struct{}

// isCode implements the Code interface so that Nil can be used as Code
func (n Nil) isCode() {}

func CodeString(code Code, indent int) string {
	switch c := code.(type) {
	case Line:
		return strings.Repeat(" ", indent) + string(c) + "\n"
	case Group:
		sb := strings.Builder{}
		for _, child := range c {
			sb.WriteString(CodeString(child, indent))
		}
		return sb.String()
	case Block:
		sb := strings.Builder{}
		for _, child := range c {
			sb.WriteString(CodeString(child, indent+2))
		}
		return sb.String()
	case Nil:
		return ""
	default:
		panic(fmt.Errorf("unkown code \"%v\"", code))
	}
}
