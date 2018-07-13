package main

import (
	"fmt"
	"github.com/DomParfitt/gecko/lexer"
	"github.com/DomParfitt/gecko/parser"
	"os"
	"time"
)

func main() {
	start := time.Now()
	args := os.Args[1:]
	pattern := args[0]
	input := args[1]
	tokens := lexer.Tokenize(pattern)

	parser := parser.New()
	regex, err := parser.Parse(tokens)

	if err != nil {
		fmt.Println(err)
		return
	}

	exec := regex.Compile()

	result := exec.Execute(input)
	// time.Sleep(5 * time.Second)
	fmt.Printf("%t\n", result)
	end := time.Now()
	fmt.Printf("Compiled the pattern %s in %f seconds\n", pattern, end.Sub(start).Seconds())

}
