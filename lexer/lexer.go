package lexer

//Tokenize lexes a string into a list of
// tokens
func Tokenize(str string) []Token {
	var tokens []Token
	for _, ch := range str {
		tokenType, ok := Match(ch)
		if ok {
			tokens = append(tokens, Token{Type: tokenType, Value: ch})
		}
	}
	return tokens
}
