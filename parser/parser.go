package parser

import (
	"github.com/DomParfitt/gecko/automata"
	"github.com/DomParfitt/gecko/lexer/token"
	"github.com/DomParfitt/gecko/tree"
)

//Parser struct
type Parser struct {
	cursor int
	tokens []token.Token
	tree   *tree.AbstractSyntax
}

//New parser
func New() *Parser {
	return &Parser{
		cursor: 0,
		tokens: []token.Token{},
		tree:   nil,
	}
}

// Parse a list of tokens into an executable
func (p *Parser) Parse(tokens []token.Token) (*automata.FiniteState, error) {
	p.cursor = 0
	p.tokens = tokens
	automata := automata.New()
	//TODO: Logic
	return automata, nil
}

// Consume the next token in the list and increment the cursor
func (p *Parser) consume() (token.Token, bool) {

	// If cursor is incremented beyond final token then we
	// can't consume any further so return failure
	if p.cursor == len(p.tokens) {
		return token.Token{Token: token.Digit, Value: ' '}, false
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

func literal() bool {

	return false
}
