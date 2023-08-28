package gid

const (
	ClassOther TokenClass = iota
	ClassLowerCase
	ClassUpperCase
	ClassDigit
)

type (
	TokenClass int
	Token      struct {
		Class TokenClass
		Runes []rune
	}
)

func NewToken(class TokenClass, r ...rune) *Token {
	return &Token{
		Class: class,
		Runes: r,
	}
}

func (tok *Token) Append(r rune) {
	tok.Runes = append(tok.Runes, r)
}

func (tok *Token) Pop() rune {
	result := tok.Runes[len(tok.Runes)-1]

	tok.Runes = tok.Runes[:len(tok.Runes)-1]

	return result
}

func (tok *Token) Valid() bool {
	return tok.Class != ClassOther && len(tok.Runes) > 0
}
