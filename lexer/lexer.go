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
	tokens    chan *token.Token
}

type LexerStateFn func(*Lexer) LexerStateFn

func NewLexer(name string, input string, buf uint) (*Lexer, chan *token.Token) {
	l := &Lexer{
		name:  name,
		input: input,

		line: 1,
		col:  1,

		nextState: lineLexer,
		tokens:    make(chan *token.Token, buf),
	}
	return l, l.tokens
}

func (l *Lexer) Run() {
	for l.nextState != nil {
		l.nextState = l.nextState(l)
	}
}

func (l *Lexer) NextToken() *token.Token {
	for {
		select {
		case t := <-l.tokens:
			return t
		default:
			l.nextState = l.nextState(l)
		}
	}
}

// Emit Token
func (l *Lexer) emitNotEmpty(typ token.TokenType) {
	if !l.empty() {
		l.emit(typ)
	}
}

func (l *Lexer) emitBackNotEmpty(runes, bytes uint, typ token.TokenType) {
	if !l.emptyBack(runes, bytes) {
		l.emitBack(runes, bytes, typ)
	}
}
func (l *Lexer) emitBack(runes, bytes uint, typ token.TokenType) {
	t := &token.Token{typ, l.input[l.start : l.pos-bytes], l.name, l.line, l.col}
	l.tokens <- t

	l.start = l.pos - bytes
	l.col += l.runes - runes
}
func (l *Lexer) emit(typ token.TokenType) {
	t := &token.Token{typ, l.input[l.start:l.pos], l.name, l.line, l.col}
	l.tokens <- t

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
	close(l.tokens)
}

// Helpers
func (l *Lexer) empty() bool {
	return l.start <= l.pos
}
func (l *Lexer) emptyBack(_, bytes uint) bool {
	return (l.start + bytes) >= l.pos
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
