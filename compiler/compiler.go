package compiler

import (
	"github.com/DomParfitt/gecko/automata"
	"github.com/DomParfitt/gecko/tree"
)

func Compile(tree *tree.AbstractSyntax) *automata.FiniteState {
	return automata.New()
}
