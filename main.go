package main

import (
	"fmt"
	"github.com/DomParfitt/gecko/compiler"
	"github.com/DomParfitt/gecko/lexer"
	"github.com/DomParfitt/gecko/parser"
)

func main() {
	txt := "hello, gecko"
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

	exec := compiler.Compile(tree)
	result := exec.Execute("hello")
	fmt.Printf("%t", result)

}
