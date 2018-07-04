package lexer

import (
	"github.com/DomParfitt/gecko/lexer/token"
)

//Tokenize lexes a string into a list of
// tokens
func Tokenize(str string) []token.Token {
	var tokens []token.Token
	for _, ch := range str {
		tokenType, ok := token.Match(ch)
		if ok {
			tokens = append(tokens, token.Token{Token: tokenType, Value: ch})
		}
	}
	return tokens
}
