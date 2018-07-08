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

	for p.cursor < len(tokens) {
		reset := p.cursor
		if p.expression() {
			continue
		} else {
			p.cursor = reset
		}
	}

	return p.tree, nil
}

// Consume the next token in the list and increment the cursor
func (p *Parser) consume() (lexer.Token, bool) {

	// If cursor is incremented beyond final token then we
	// can't consume any further so return failure
	if p.cursor == len(p.tokens) {
		return p.tokens[p.cursor-1], false
	}

	token := p.tokens[p.cursor]
	p.cursor++

	return token, true
}

//LookBack and get the previous token
func (p *Parser) lookBack() (lexer.Token, bool) {

	if p.cursor == 0 {
		return p.tokens[p.cursor], false
	}

	token := p.tokens[p.cursor-1]

	return token, true
}

// Replace the previous token, decrementing the cursor
func (p *Parser) replace() {
	if p.cursor > 0 {
		p.cursor--
	}
}

// Literal is a digit or a letter
func (p *Parser) literal(token lexer.Token) bool {
	return token.Type == lexer.Digit || token.Type == lexer.Letter
}

// Expression is a literal or a literal followed by a wildcard
func (p *Parser) expression() bool {
	//Get a token
	token, ok := p.consume()

	// If not ok then we're at the end of the tokens
	// Wildcard must be final token
	// Either way look at previous token should be literal
	if !ok || token.Type == lexer.Closure {
		token, ok := p.lookBack()
		return ok && p.literal(token)
	}

	//Current token is literal, so look ahead
	if p.literal(token) {
		next, ok := p.consume()

		//No next token but current is valid
		if !ok {
			return true
		}

		//Next token is a wildcard which is valid and indicates end
		if token.Type == lexer.Closure {
			return true
		}

		//Next is literal so recurse
		if p.literal(next) {
			return p.expression()
		}
	}
	return false
}

func (p *Parser) term() bool {
	// token, ok := p.consume()
	return false
}
