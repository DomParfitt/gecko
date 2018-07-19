package compiler

import (
	"fmt"
	"github.com/DomParfitt/gecko/automata"
	"github.com/DomParfitt/gecko/lexer"
	"github.com/DomParfitt/gecko/parser"
)

//Compiler wraps the entire process for compiling and matching
// a regex
type Compiler struct {
	parser *parser.Parser
	exe    *automata.FiniteState
}

//New Compiler
func New() *Compiler {
	return &Compiler{
		parser: parser.New(),
	}
}

//Compile the provided pattern into an executable
func (c *Compiler) Compile(pattern string) error {
	tokens := lexer.Tokenize(pattern)
	ast, err := c.parser.Parse(tokens)

	if err != nil {
		return err
	}

	c.exe = ast.Compile()

	return nil
}

//Match the input string against the compiled executable.
//Returns an error if no pattern has been provided
func (c *Compiler) Match(input string) (bool, error) {
	if c.exe == nil {
		return false, fmt.Errorf("unable to match input: %s - no pattern initialised", input)
	}

	return c.exe.Execute(input), nil
}

//MatchPattern takes a pattern and an input, compiles the pattern and
//matches the input against it
func (c *Compiler) MatchPattern(pattern, input string) (bool, error) {
	err := c.Compile(pattern)

	if err != nil {
		return false, err
	}

	return c.Match(input)
}
