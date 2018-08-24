//Package lexer contains the defintions for Gecko's lexical tokens
//as well as functionality to tokenize a string into a list of tokens
package lexer

//Tokenize lexes a string into a list of
// tokens
func Tokenize(str string) []Token {
	var tokens []Token
	for _, ch := range str {
		tokenType := Match(ch)
		tokens = append(tokens, Token{Type: tokenType, Value: ch})
	}
	return tokens
}
