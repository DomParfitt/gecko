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
func (p *Parser) Parse(tokens []lexer.Token) (*RegExpr, error) {
	p.cursor = 0
	p.tokens = tokens

	regExpr, ok := p.regExpr()

	if !ok {
		return nil, fmt.Errorf("the token stream could not be parsed")
	}

	return regExpr, nil
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

func (p *Parser) base() (*Element, bool) {
	token, ok := p.consume()

	if !ok {
		return nil, false
	}

	if token.Type != lexer.Character {
		p.replace()
		return nil, false
	}

	return &Element{token.Value}, true
}

func (p *Parser) star() (*Star, bool) {

	base, ok := p.base()
	if !ok {
		return nil, false
	}

	token, ok := p.consume()

	if !ok {
		p.replace()
		return nil, false
	}

	if token.Type != lexer.Star {
		p.replace()
		p.replace()
		return nil, false
	}

	return &Star{base}, true
}

func (p *Parser) plus() (*Plus, bool) {

	base, ok := p.base()
	if !ok {
		return nil, false
	}

	token, ok := p.consume()

	if !ok {
		p.replace()
		return nil, false
	}

	if token.Type != lexer.Plus {
		p.replace()
		p.replace()
		return nil, false
	}

	return &Plus{base}, true
}

func (p *Parser) basicExpr() (*BasicExpr, bool) {
	star, ok := p.star()
	if ok {
		return &BasicExpr{star: star}, true
	}

	plus, ok := p.plus()
	if ok {
		return &BasicExpr{plus: plus}, true
	}

	base, ok := p.base()
	if ok {
		return &BasicExpr{element: base}, true
	}

	return nil, false
	// return p.star() || p.plus() || p.base()
}

func (p *Parser) concatenation() (*Concatenation, bool) {
	basic, ok := p.basicExpr()
	if ok {
		simple, ok := p.simpleExpr()
		if ok {
			return &Concatenation{simple, basic}, true
		}
	}
	return nil, false
}

func (p *Parser) simpleExpr() (*SimpleExpr, bool) {
	basic, ok := p.basicExpr()
	if ok {
		return &SimpleExpr{basic: basic}, true
	}

	concatenation, ok := p.concatenation()
	if ok {
		return &SimpleExpr{concatenation: concatenation}, true
	}

	return nil, false
	// return p.basicExpr() || p.concatenation()
}

func (p *Parser) union() (*Union, bool) {

	simple, ok := p.simpleExpr()
	if !ok {
		return nil, false
	}

	token, ok := p.consume()

	if !ok {
		p.replace()
		return nil, false
	}

	if token.Type != lexer.Pipe {
		p.replace()
		p.replace()
		return nil, false
	}

	regex, ok := p.regExpr()

	if !ok {
		return nil, false
	}

	return &Union{regex, simple}, true

	// return p.regExpr()
}

func (p *Parser) regExpr() (*RegExpr, bool) {
	union, ok := p.union()
	if ok {
		return &RegExpr{union: union}, true
	}

	simple, ok := p.simpleExpr()
	if ok {
		return &RegExpr{simple: simple}, true
	}

	return nil, false
	// return p.simpleExpr() || p.union()
}
