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
	var s string
	if len(t.Name) > 0 {
		s = fmt.Sprintf("%s:%v:%v", t.Name, t.Line, t.Col)
	} else {
		s = fmt.Sprintf("%v:%v", t.Line, t.Col)
	}

	switch t.Typ {
	case TokenEOL, TokenEOF:
		return fmt.Sprintf("%s:%s", s, t.Typ)
	default:
		if len(t.Val) > 10 {
			return fmt.Sprintf("%s:%s:%.10q... (%v)", s, t.Typ, t.Val, len(t.Val))
		} else {
			return fmt.Sprintf("%s:%s:%q (%v)", s, t.Typ, t.Val, len(t.Val))
		}
	}
}

// Loc modified the Line and Col fields of the Token and returns the Token again
// to be used from Tests
func (t *Token) Loc(l, c uint) *Token {
	t.Line, t.Col = l, c
	return t
}
