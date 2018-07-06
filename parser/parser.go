package parser

import (
	"github.com/DomParfitt/gecko/lexer"
	"github.com/DomParfitt/gecko/stack"
	"github.com/DomParfitt/gecko/tree"
)

//Parser struct
type Parser struct {
	cursor int
	tokens []lexer.Token
	stack  *stack.Stack
	tree   *tree.AbstractSyntax
}

//New parser
func New() *Parser {
	return &Parser{
		cursor: 0,
		tokens: []lexer.Token{},
		stack:  stack.New(),
		tree:   nil,
	}
}

// Parse a list of tokens into an executable
func (p *Parser) Parse(tokens []lexer.Token) (*tree.AbstractSyntax, error) {
	p.cursor = 0
	p.tokens = tokens

	for token, ok := p.consume(); ok; {
		if p.literal(token) {
			tree := tree.New(token)
			tree.AddLeft(p.tree)
			p.tree = tree
		} else if p.operator(token) {

		}
	}

	return p.tree, nil
}

// Consume the next token in the list and increment the cursor
func (p *Parser) consume() (lexer.Token, bool) {

	// If cursor is incremented beyond final token then we
	// can't consume any further so return failure
	if p.cursor == len(p.tokens) {
		return lexer.Error(), false
	}

	token := p.tokens[p.cursor]
	p.cursor++

	return token, true
}

// Replace the previous token, decrementing the cursor
func (p *Parser) replace() {
	if p.cursor > 0 {
		p.cursor--
	}
}

func (p *Parser) literal(token lexer.Token) bool {
	return token.Type == lexer.Digit || token.Type == lexer.Letter
}

func (p *Parser) concatenation() {
	if token, ok := p.consume(); ok {
		if p.literal(token) {
			p.concatenation()
		}
	}
}

func (p *Parser) operator(token lexer.Token) bool {
	return token.Type == lexer.Pipe
}

func (p *Parser) group() {
	if token, ok := p.consume(); ok && token.Type == lexer.OpenBrace {
		p.stack.Push(token)
	}
}
