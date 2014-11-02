package lexer

import (
	"unicode"
)

func sectionLexer(l *Lexer) LexerStateFn {
	l.skip() // trash the [

loop:
	for {
		r0, _ := l.next()
		switch r0 {
		case EOF:
			l.emitError("Unexpected EOF")
			l.emitEOF()
			return nil
		case '\n', '\r':
			break loop
		default:
			if unicode.IsSpace(r0) {
				l.skip()
			} else if unicode.IsControl(r0) {
				break loop
			} else {
				return sectionNameLexer
			}
		}
	}

	l.emitError("Invalid Character")
	return nil
}

func sectionNameLexer(l *Lexer) LexerStateFn {
loop:
	for {
		r0, l0 := l.next()
		switch r0 {
		case EOF:
			l.emitBackNotEmpty(1, l0, token.TokenText)
			l.emitError("Unexpected EOF")
			l.emitEOF()
			return nil
		case '\n':
			l.emitBackNotEmpty(1, l0, token.TokenText)
			l.emitError("Unexpected EOL")
			l.emitEOL()
			return nil
		case '\r':

		case unicode.IsLetter(r0) || unicode.IsNumber(r0) || unicode.IsMark(r0):

		}
	}
}
