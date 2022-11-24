package token

type Token struct {
	Type       Type
	Text       string
	Start, End int
}
