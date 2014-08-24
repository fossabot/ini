package lexer

import (
	"asciigoat.org/ini/token"
)

func lineLexer(l *Lexer) LexerStateFn {
	for {
		r0, l0 := l.next()
		switch r0 {
		case EOF:
			l.emitLine(0, 0)
			l.emitEOF()
			break
		case '\n':
			l.emitLine(1, l0)
			l.emitEOL()
		case '\r':
			r1, l1 := l.next()
			if r1 == '\n' {
				l.emitLine(2, l0+l1)
				l.emitEOL()
			} else {
				l.back(1, l1)
			}
		}
	}
	return nil
}

func (l *Lexer) emitLine(runes, bytes uint) {
	if runes > 0 {
		l.back(runes, bytes)
		if !l.empty() {
			l.emit(token.TokenRaw)
		}
		l.forth(runes, bytes)
	} else if !l.empty() {
		l.emit(token.TokenRaw)
	}
}
