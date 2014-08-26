package token

import (
	"fmt"
)

// a Token on a document
type Token struct {
	Typ  TokenType
	Val  string
	Name string
	Line uint
	Col  uint
}

func (t *Token) String() string {
	s := fmt.Sprintf("%s:%v:%v", t.Name, t.Line, t.Col)

	switch t.Typ {
	case TokenEOL, TokenEOF:
		return fmt.Sprintf("%s: %s", s, t.Typ)
	default:
		if len(t.Val) > 10 {
			return fmt.Sprintf("%s: %s:%.10q... (%v)", s, t.Typ, t.Val, len(t.Val))
		} else {
			return fmt.Sprintf("%s: %s:%q (%v)", s, t.Typ, t.Val, len(t.Val))
		}
	}
}
