package lexer

import (
	"asciigoat.org/ini/token"
	"unicode"
)

func preambleLexer(l *Lexer) LexerStateFn {
loop:
	for {
		r, l0 := l.next()
		switch r {
		case EOF:
			l.emitEOF()
			return nil
		case '\n':
			l.emitEOL()
		case '\r':
			r1, l1 := l.next()
			if r1 == '\n' {
				l.emitEOL()
				break
			}
			l.back(1, l1)
			break loop
		case '[':
			l.back(1, l0)
			return lineLexer
		case ';', '#':
			l.back(1, l0)
			return earlyCommentLexer
		default:
			if unicode.IsSpace(r) {
				l.skip()
			} else {
				break loop
			}
		}
	}

	l.emitError("Invalid Character")
	return nil
}

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
