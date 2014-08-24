package token

import (
	"fmt"
	"testing"
)

func TestTokenTypeToString(t *testing.T) {
	var foo TokenType

	for _, o := range []struct {
		typ TokenType
		str string
	}{
		{foo, "UNDEFINED"},
		{TokenSection, "SEC"},
		{TokenSubsection, "SUB"},
		{TokenComment, "COM"},
		{TokenKey, "KEY"},
		{TokenValue, "VAL"},
		{TokenRaw, "RAW"},
		{TokenEOL, "EOL"},
		{TokenError, "ERROR"},
		{TokenEOF, "EOF"},
		{1234, "UNDEFINED"},
	} {
		str := fmt.Sprintf("%s", o.typ)
		if str != o.str {
			t.Errorf("TokenType:%v stringified as %s instead of %s.", int(o.typ), str, o.str)
		}
	}
}
