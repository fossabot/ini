package lexer

import (
	"testing"
)

// Really Empty
func TestEmpty1(t *testing.T) {
	_, _ = NewLexer("Empty", ``)
}

// Whitespace Empty
func TestEmpty2(t *testing.T) {
	_, _ = NewLexer("Empty", `\n`)
}
func TestEmpty3(t *testing.T) {
	_, _ = NewLexer("Empty", `   \n`)
}
func TestEmpty4(t *testing.T) {
	_, _ = NewLexer("Empty", `   \n\t`)
}

// Empty Preamble
func TestPreamble1(t *testing.T) {
	_, _ = NewLexer("Preamble1", `[section1]\nkey1 = value1\n`)
}

// Whitespace Preamble
func TestPreamble2(t *testing.T) {
	_, _ = NewLexer("Preamble2", `\n\t\n[section1]\nkey1 = value1`)
}

// Comment Preamble
func TestPreamble3(t *testing.T) {
	_, _ = NewLexer("Preamble1", `
; comment 1
; comment 2

[section1]
key1 = value1
`)
}
