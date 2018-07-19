package main

import (
	"github.com/DomParfitt/gecko/server"
)

func main() {
	// args := os.Args[1:]
	// pattern := args[0]
	// input := args[1]

	// compiler := compiler.New()
	// result, err := compiler.MatchPattern(pattern, input)

	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// fmt.Printf("%t\n", result)
	server.Run()
}
