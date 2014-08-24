package token

import (
	"fmt"
)

// a Token on a document
type Token struct {
	Typ  TokenType
	Val  string
	Line uint
	Col  uint
}

func (t *Token) String() string {
	switch t.Typ {
	case tokenNL, tokenEOF:
		return fmt.Sprintf("%s at %v:%v", t.Typ, t.Line, t.Col)
	default:
		if len(t.Val) > 10 {
			return fmt.Sprintf("%s:.10%q... at %v:%v", t.Typ, t.Val, t.Line, t.Col)
		} else {
			return fmt.Sprintf("%s:%q at %v:%v", t.Typ, t.Val, t.Line, t.Col)
		}
	}
}
