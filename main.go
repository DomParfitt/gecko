package main

import (
	"fmt"
	"github.com/DomParfitt/gecko/lexer"
	"github.com/DomParfitt/gecko/parser"
)

func main() {
	txt := "a|b"
	tokens, err := lexer.Tokenize(txt)

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, token := range tokens {
		fmt.Println(token)
	}
	parser := parser.New()
	tree, err := parser.Parse(tokens)

	if err != nil {
		fmt.Println(err)
		return
	}

	exec := tree.Compile()
	result := exec.Execute("a")
	fmt.Printf("%t", result)

}
