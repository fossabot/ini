package lexer

import (
	"asciigoat.org/ini/token"
	"fmt"
	"testing"
)

func dump(t *testing.T, name, data string) {
	lex, _ := NewLexer(name, data)

loop:
	for {
		x := lex.NextToken()
		fmt.Printf("token: %v\n", *x)
		if x.Typ == token.TokenEOF {
			break loop
		}
	}
}

// Really Empty
func TestEmpty1(t *testing.T) {
	dump(t, "Empty", ``)
}

// Whitespace Empty
func TestEmpty2(t *testing.T) {
	dump(t, "Empty", "\n")
}
func TestEmpty3(t *testing.T) {
	dump(t, "Empty", "   \n")
}
func TestEmpty4(t *testing.T) {
	dump(t, "Empty", "   \n\t")
}

// Empty Preamble
func TestPreamble1(t *testing.T) {
	dump(t, "Preamble1", "[section1]\nkey1 = value1\n")
}

// Whitespace Preamble
func TestPreamble2(t *testing.T) {
	dump(t, "Preamble2", "\n\t\n[section1]\nkey1 = value1")
}

// Comment Preamble
func TestPreamble3(t *testing.T) {
	dump(t, "Preamble1", `
; comment 1
; comment 2

[section1]
key1 = value1
`)
}
