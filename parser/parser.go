package parser

import (
	"fmt"
	"github.com/DomParfitt/gecko/lexer"
	"github.com/DomParfitt/gecko/types/stack"
	"github.com/DomParfitt/gecko/types/tree"
)

//Parser struct
type Parser struct {
	cursor int
	tokens []lexer.Token
	stack  *stack.Stack
	tree   *tree.AbstractSyntax
	expr   *RegExpr
}

//New parser
func New() *Parser {
	return &Parser{
		cursor: 0,
		tokens: []lexer.Token{},
		stack:  stack.New(),
		tree:   nil,
		expr:   nil,
	}
}

// Parse a list of tokens into an executable
func (p *Parser) Parse(tokens []lexer.Token) (*tree.AbstractSyntax, error) {
	p.cursor = 0
	p.tokens = tokens

	if !p.regExpr() {
		return nil, fmt.Errorf("the token stream could not be parsed")
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

func (p *Parser) base() bool {
	token, ok := p.consume()

	if !ok {
		return false
	}

	if token.Type != lexer.Character {
		p.replace()
		return false
	}

	return true
}

func (p *Parser) star() bool {

	if !p.base() {
		return false
	}

	token, ok := p.consume()

	if !ok {
		p.replace()
		return false
	}

	if token.Type != lexer.Star {
		p.replace()
		p.replace()
		return false
	}

	return true
}

func (p *Parser) plus() bool {

	if !p.base() {
		return false
	}

	token, ok := p.consume()

	if !ok {
		p.replace()
		return false
	}

	if token.Type != lexer.Plus {
		p.replace()
		p.replace()
		return false
	}

	return true
}

func (p *Parser) basicExpr() bool {
	return p.star() || p.plus() || p.base()
}

func (p *Parser) concatenation() bool {
	return p.basicExpr() && p.simpleExpr()
}

func (p *Parser) simpleExpr() bool {
	return p.basicExpr() || p.concatenation()
}

func (p *Parser) union() bool {
	if !p.simpleExpr() {
		return false
	}

	token, ok := p.consume()

	if !ok {
		p.replace()
		return false
	}

	if token.Type != lexer.Pipe {
		p.replace()
		p.replace()
		return false
	}

	return p.regExpr()
}

func (p *Parser) regExpr() bool {
	return p.simpleExpr() || p.union()
}
