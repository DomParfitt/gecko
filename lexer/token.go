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
	OpenBrace
	CloseBrace
	Star
	Plus
	Caret
	Escape
	Pipe
	None
)

//Token represents a lexed charcter. Contains the
// type of the token as well as its raw value
type Token struct {
	Type  Type
	Value rune
}

//Error token
func Error() Token {
	return Token{
		Type:  None,
		Value: ' ',
	}
}

func (t Token) String() string {
	return fmt.Sprintf("[%s]: %c", t.Type, t.Value)
}

// Match the provided character to a Type of Token.
func Match(ch rune) Type {
	if ch == '(' || ch == '[' || ch == '{' {
		return OpenBrace
	}

	if ch == ')' || ch == ']' || ch == '}' {
		return CloseBrace
	}

	if ch == '*' {
		return Star
	}

	if ch == '+' {
		return Plus
	}

	if ch == '^' {
		return Caret
	}

	if ch == '\\' {
		return Escape
	}

	if ch == '|' {
		return Pipe
	}

	return Character

}
