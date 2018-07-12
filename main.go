package main

import (
	"fmt"
	"github.com/DomParfitt/gecko/lexer"
	"github.com/DomParfitt/gecko/parser"
	"os"
)

func main() {
	args := os.Args[1:]
	pattern := args[0]
	input := args[1]
	tokens := lexer.Tokenize(pattern)

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
	result := exec.Execute(input)
	fmt.Printf("%t", result)

}
