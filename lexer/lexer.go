package lexer

import (
	"github.com/DomParfitt/gecko/lexer/token"
)

func Tokenize(str string) []token.Token {
	tokens := []token.Token{}
	for _, ch := range str {
		tokenType, err := token.Match(ch)
		if err == nil {
			tokens = append(tokens, token.Token{Token: tokenType, Value: ch})
		}
	}
	return tokens
}
