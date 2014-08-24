package token

// kinds of Tokens
type TokenType int

const (
	TokenSection    TokenType = iota + 1 // [section]
	TokenSubsection                      // [section "subsection"]
	TokenComment                         // ; comment
	TokenKey                             // key = value
	TokenValue                           // key = value
	TokenRaw                             // some content
	TokenEOL                             // \r\n or \n
	TokenError
	TokenEOF
)

func (typ TokenType) String() string {

	switch typ {
	case TokenSection:
		return "SEC"
	case TokenSubsection:
		return "SUB"
	case TokenComment:
		return "COM"
	case TokenKey:
		return "KEY"
	case TokenValue:
		return "VAL"
	case TokenRaw:
		return "RAW"
	case TokenEOL:
		return "EOL"
	case TokenError:
		return "ERROR"
	case TokenEOF:
		return "EOF"
	default:
		return "UNDEFINED"
	}
}
