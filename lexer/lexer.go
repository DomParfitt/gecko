package lexer

import (
	"fmt"
)

//Tokenize lexes a string into a list of
// tokens
func Tokenize(str string) ([]Token, error) {
	var tokens []Token
	for i, ch := range str {
		tokenType, ok := Match(ch)
		if !ok {
			return []Token{}, fmt.Errorf("unrecognised token: '%c' at position: %d", ch, i)
		}
		tokens = append(tokens, Token{Type: tokenType, Value: ch})
	}
	return tokens, nil
}
