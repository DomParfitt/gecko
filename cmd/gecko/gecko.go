package main

import (
	"fmt"
	"github.com/DomParfitt/gecko/core"
	"os"
)

func main() {
	args := os.Args[1:]
	pattern := args[0]
	input := args[1]

	compiler := core.New()
	result, err := compiler.MatchPattern(pattern, input)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%t\n", result)
}
