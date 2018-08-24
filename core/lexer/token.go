package lexer

import (
	"fmt"
)

//go:generate stringer -type=Type

// Type of Token
type Type int

// Accepted Types of Token
const (
	Character Type = iota
	OpenParen
	CloseParen
	OpenBracket
	CloseBracket
	OpenBrace
	CloseBrace
	Star
	Plus
	Question
	Caret
	Escape
	Pipe
	Dash
)

//Token represents a lexed character. Contains the
// type of the token as well as its raw value
type Token struct {
	Type  Type
	Value rune
}

func (t Token) String() string {
	return fmt.Sprintf("[%s]: '%c'", t.Type, t.Value)
}

// Match the provided character to a Type of Token.
func Match(ch rune) Type {
	switch ch {
	case '(':
		return OpenParen
	case ')':
		return CloseParen
	case '[':
		return OpenBracket
	case ']':
		return CloseBracket
	case '{':
		return OpenBrace
	case '}':
		return CloseBrace
	case '*':
		return Star
	case '+':
		return Plus
	case '?':
		return Question
	case '^':
		return Caret
	case '\\':
		return Escape
	case '|':
		return Pipe
	case '-':
		return Dash
	default:
		return Character
	}
}
