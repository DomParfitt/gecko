package parser

import (
	"github.com/DomParfitt/gecko/automata"
	"github.com/DomParfitt/gecko/lexer/token"
)

func Parse(tokens []token.Token) *automata.FiniteState {
	return automata.New()
}
