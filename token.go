package ini

import (
	"fmt"
)

// a Token on a document
type Token struct {
	typ  TokenType
	val  string
	line uint
	col  uint
}

func (t *Token) String() string {
	switch t.typ {
	case tokenNL, tokenEOF:
		return fmt.Sprintf("%s at %v:%v", t.typ, t.line, t.col)
	default:
		if len(t.val) > 10 {
			return fmt.Sprintf("%s:.10%q... at %v:%v", t.typ, t.val, t.line, t.col)
		} else {
			return fmt.Sprintf("%s:%q at %v:%v", t.typ, t.val, t.line, t.col)
		}
	}
}
