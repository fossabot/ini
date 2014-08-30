package lexer

import (
	"asciigoat.org/ini/token"
	"unicode"
)

func earlyCommentLexer(l *Lexer) LexerStateFn {
loop:
	for {
		r0, l0 := l.next()
		switch r0 {
		case EOF:
			l.emitNotEmpty(token.TokenComment)
			l.emitEOF()
			return nil
		case '\n':
			l.emitBackNotEmpty(1, l0, token.TokenComment)
			l.emitEOL()
			return preambleLexer
		case '\r':
			r1, l1 := l.next()
			if r1 == '\n' {
				l.emitBackNotEmpty(2, l0+l1, token.TokenComment)
				l.emitEOL()
				return preambleLexer
			} else {
				l.back(1, l1)
				l.emitBackNotEmpty(1, l0, token.TokenComment)
				break loop
			}
		default:
			if unicode.IsControl(r0) {
				l.emitBackNotEmpty(1, l0, token.TokenComment)
				break loop
			}
		}
	}

	l.emitError("Invalid Character")
	return nil
}
