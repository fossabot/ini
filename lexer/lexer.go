package lexer

import (
	"asciigoat.org/ini/token"
	"unicode/utf8"
)

const (
	EOF = -1
)

type Lexer struct {
	name  string
	input string

	start     uint
	line, col uint

	pos, runes uint

	nextState LexerStateFn
	tokens    chan token.Token
}

type LexerStateFn func(*Lexer) LexerStateFn

func NewLexer(name string, input string) (*Lexer, chan token.Token) {
	l := &Lexer{
		name:  name,
		input: input,

		line: 1,
		col:  1,

		nextState: lineLexer,
		tokens:    make(chan token.Token),
	}
	return l, l.tokens
}

func (l *Lexer) Run() {
	for l.nextState != nil {
		l.nextState = l.nextState(l)
	}
	close(l.tokens)
}

// Emit Token
func (l *Lexer) emit(typ token.TokenType) {
	l.tokens <- token.Token{typ, l.input[l.start:l.pos],
		l.line, l.col}
	l.start = l.pos
	l.col += l.runes
}

func (l *Lexer) emitEOL() {
	l.emit(token.TokenEOL)
	l.line += 1
	l.col = 1
}

func (l *Lexer) emitEOF() {
	l.emit(token.TokenEOF)
}

// Helpers
func (l *Lexer) empty() bool {
	return l.start <= l.pos
}

func (l *Lexer) next() (rune, uint) {
	if l.pos >= uint(len(l.input)) {
		return EOF, 0
	}
	r, size := utf8.DecodeRuneInString(l.input[l.pos:])
	w := uint(size)

	l.pos += w   // byte offset
	l.runes += 1 // column step
	return r, w
}

func (l *Lexer) back(runes, bytes uint) {
	l.runes -= runes
	l.pos -= bytes
}

func (l *Lexer) forth(runes, bytes uint) {
	l.runes += runes
	l.pos += bytes
}