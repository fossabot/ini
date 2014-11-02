package lexer

import (
	"asciigoat.org/ini/token"
	"testing"
)

func newLexer(t *testing.T, str string) *Lexer {
	t.Logf("data = %q", str)
	l, _ := NewLexer("", str, 2)
	return l
}

func cmp(t *testing.T, lex *Lexer, tokens []*token.Token) {
	i := 0

	for {
		var x, y *token.Token

		x = lex.NextToken()
		if i < len(tokens) {
			y = tokens[i]
		}
		i = i + 1

		if x == nil || y == nil || *x != *y {
			t.Errorf("token[%v] failed: %v != %v", i, x, y)
		} else {
			t.Logf("token[%v] %v", i, x)
		}

		if x == nil || x.Typ == token.TokenEOF {
			x = nil
			for j := i; j < len(tokens); {
				y = tokens[j]
				j++
				t.Errorf("token[%v] failed: %v != %v", j, x, y)
			}
			return
		}
	}
}

// Really Empty
func TestEmpty1(t *testing.T) {
	l := newLexer(t, "")
	tokens := []*token.Token{
		l.Token(token.TokenEOF, "").Loc(1, 1),
	}

	cmp(t, l, tokens)
}

// Whitespace Empty
func TestEmpty2(t *testing.T) {
	l := newLexer(t, "\n")
	tokens := []*token.Token{
		l.Token(token.TokenEOL, "\n").Loc(1, 1),
		l.Token(token.TokenEOF, "").Loc(2, 1),
	}

	cmp(t, l, tokens)
}

// Whitespace Empty
func TestEmpty3(t *testing.T) {
	l := newLexer(t, "   \n")
	tokens := []*token.Token{
		l.Token(token.TokenEOL, "\n").Loc(1, 4),
		l.Token(token.TokenEOF, "").Loc(2, 1),
	}

	cmp(t, l, tokens)
}

// Whitespace Empty
func TestEmpty4(t *testing.T) {
	l := newLexer(t, "   \n\t ")
	tokens := []*token.Token{
		l.Token(token.TokenEOL, "\n").Loc(1, 4),
		l.Token(token.TokenEOF, "").Loc(2, 3),
	}

	cmp(t, l, tokens)
}

// Empty Preamble
func TestPreamble1(t *testing.T) {
	l := newLexer(t, "[section1]\nkey1 = value1\n")
	tokens := []*token.Token{
		l.Token(token.TokenSection, "section1").Loc(1, 1),
		l.Token(token.TokenEOL, "\n").Loc(1, 11),
		l.Token(token.TokenText, "key1 = value1").Loc(2, 1),
		l.Token(token.TokenEOL, "\n").Loc(2, 14),
		l.Token(token.TokenEOF, "").Loc(3, 1),
	}

	cmp(t, l, tokens)
}

// Whitespace Preamble
func TestPreamble2(t *testing.T) {
	l := newLexer(t, "\n\t\n[section1]\nkey1 = value1")
	tokens := []*token.Token{
		l.Token(token.TokenEOL, "\n").Loc(1, 1),
		l.Token(token.TokenEOL, "\n").Loc(2, 2),
		l.Token(token.TokenSection, "section1").Loc(3, 1),
		l.Token(token.TokenEOL, "\n").Loc(3, 11),
		l.Token(token.TokenText, "key1 = value1").Loc(4, 1),
		l.Token(token.TokenEOF, "").Loc(4, 14),
	}

	cmp(t, l, tokens)
}

// Comment Preamble
func TestPreamble3(t *testing.T) {
	l := newLexer(t, `
; comment 1
; comment 2

[section1]
key1 = value1
`)
	tokens := []*token.Token{
		l.Token(token.TokenEOL, "\n").Loc(1, 1),
		l.Token(token.TokenComment, "; comment 1").Loc(2, 1),
		l.Token(token.TokenEOL, "\n").Loc(2, 12),
		l.Token(token.TokenComment, "; comment 2").Loc(3, 1),
		l.Token(token.TokenEOL, "\n").Loc(3, 12),
		l.Token(token.TokenEOL, "\n").Loc(4, 1),
		l.Token(token.TokenSection, "section1").Loc(5, 1),
		l.Token(token.TokenEOL, "\n").Loc(5, 11),
		l.Token(token.TokenText, "key1 = value1").Loc(6, 1),
		l.Token(token.TokenEOL, "\n").Loc(6, 14),
		l.Token(token.TokenEOF, "").Loc(7, 1),
	}

	cmp(t, l, tokens)
}
