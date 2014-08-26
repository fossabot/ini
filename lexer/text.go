package lexer

import (
	"asciigoat.org/ini/token"
)

func lineLexer(l *Lexer) LexerStateFn {
	for {
		r0, l0 := l.next()
		switch r0 {
		case EOF:
			l.emitNotEmpty(token.TokenText)
			l.emitEOF()
			return nil
		case '\n':
			l.emitBackNotEmpty(1, l0, token.TokenText)
			l.emitEOL()
			return lineLexer
		case '\r':
			r1, l1 := l.next()
			if r1 == '\n' {
				l.emitBackNotEmpty(2, l0+l1, token.TokenText)
				l.emitEOL()
				return lineLexer
			} else {
				l.back(1, l1)
			}
		}
	}
}
