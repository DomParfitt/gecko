package main

import (
	"fmt"
	"github.com/DomParfitt/gecko/compiler"
	"os"
	"time"
)

func main() {
	start := time.Now()
	args := os.Args[1:]
	pattern := args[0]
	input := args[1]

	compiler := compiler.New()
	result, err := compiler.MatchPattern(pattern, input)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%t\n", result)
	end := time.Now()
	fmt.Printf("Compiled the pattern %s in %f seconds\n", pattern, end.Sub(start).Seconds())

}
