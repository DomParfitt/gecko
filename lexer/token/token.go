package token

import (
	"fmt"
)

type TokenType int

const (
	Digit TokenType = iota
	Letter
	OpenBrace
	CloseBrace
	Wildcard
)

type Token struct {
	Token TokenType
	Value rune
}

func (t TokenType) String() string {
	switch t {
	case Digit:
		return "Digit"
	case Letter:
		return "Letter"
	case OpenBrace:
		return "OpenBrace"
	case CloseBrace:
		return "CloseBrace"
	case Wildcard:
		return "Wildcard"
	default:
		return "UnknownToken"
	}
}

func (t Token) String() string {
	return fmt.Sprintf("[%s]: %c", t.Token, t.Value)
}

func Match(ch rune) (TokenType, error) {
	if ch >= '0' && ch <= '9' {
		return Digit, nil
	}

	if (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') {
		return Letter, nil
	}

	if ch == '(' || ch == '[' || ch == '{' {
		return OpenBrace, nil
	}

	if ch == ')' || ch == ']' || ch == '}' {
		return CloseBrace, nil
	}

	if ch == '*' || ch == '+' {
		return Wildcard, nil
	}

	return Digit, fmt.Errorf("No matching TokenType for character: %c", ch)

}
