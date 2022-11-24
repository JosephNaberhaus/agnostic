package text

import "strings"

type Config struct {
	Indent string
}

type Node interface {
	String(config Config) string
}

// Group contains multiple text nodes. It will simply write the contents in sequence without newlines.
type Group []Node

func (g Group) String(config Config) string {
	sb := strings.Builder{}

	for _, span := range g {
		sb.WriteString(span.String(config))
	}

	return sb.String()
}

// IndentedBlock is a Block that adds another level of indent to its children.
type IndentedBlock []Node

func (i IndentedBlock) String(config Config) string {
	lines := strings.Split(Block(i).String(config), "\n")

	for i, line := range lines {
		lines[i] = config.Indent + line
	}

	return strings.Join(lines, "\n")
}

// Block is an ordered sequence of text nodes. Each Text node will be seperated by a newline in the text file.
type Block []Node

func (b Block) String(config Config) string {
	sb := strings.Builder{}
	for i, span := range b {
		sb.WriteString(span.String(config))

		if i != len(b)-1 {
			sb.WriteRune('\n')
		}
	}

	return sb.String()
}

// Span is a sequence of characters.
type Span string

func (s Span) String(config Config) string {
	return string(s)
}
