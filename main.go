package main

import (
	"fmt"
	"github.com/DomParfitt/gecko/lexer"
)

func main() {
	txt := "hello, gecko"
	tokens := lexer.Tokenize(txt)
	for _, token := range tokens {
		fmt.Println(token)
	}

}
