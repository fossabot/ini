package ini

// kinds of Tokens
type TokenType int

const (
	tokenSection    TokenType = iota + 1 // [section]
	tokenSubsection                      // [section "subsection"]
	tokenComment                         // ; comment
	tokenKey                             // key = value
	tokenValue                           // key = value
	tokenRaw                             // some content
	tokenNL                              // \r\n or \n
	tokenError
	tokenEOF
)

func (typ TokenType) String() string {

	switch typ {
	case tokenSection:
		return "SEC"
	case tokenSubsection:
		return "SUB"
	case tokenComment:
		return "COM"
	case tokenKey:
		return "KEY"
	case tokenValue:
		return "VAL"
	case tokenRaw:
		return "RAW"
	case tokenNL:
		return "NL"
	case tokenError:
		return "ERROR"
	case tokenEOF:
		return "EOF"
	default:
		return "UNDEFINED"
	}
}
