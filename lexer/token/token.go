package token

import (
	"fmt"
)

//go:generate stringer -type=Type

// Type of Token
type Type int

// Accepted Types of Token
const (
	Digit Type = iota
	Letter
	OpenBrace
	CloseBrace
	Wildcard
	Caret
	Escape
	Pipe
)

//Token represents a lexed charcter. Contains the
// type of the token as well as its raw value
type Token struct {
	Token Type
	Value rune
}

func (t Token) String() string {
	return fmt.Sprintf("[%s]: %c", t.Token, t.Value)
}

// Match the provided character to a Type of Token.
func Match(ch rune) (Type, bool) {
	if ch >= '0' && ch <= '9' {
		return Digit, true
	}

	if (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') {
		return Letter, true
	}

	if ch == '(' || ch == '[' || ch == '{' {
		return OpenBrace, true
	}

	if ch == ')' || ch == ']' || ch == '}' {
		return CloseBrace, true
	}

	if ch == '*' || ch == '+' {
		return Wildcard, true
	}

	if ch == '^' {
		return Caret, true
	}

	if ch == '\\' {
		return Escape, true
	}

	if ch == '|' {
		return Pipe, true
	}

	return Digit, false

}
