package tree

import (
	"github.com/DomParfitt/gecko/lexer/token"
)

//AbstractSyntax tree struct
type AbstractSyntax struct {
	token token.Token
	left  *AbstractSyntax
	right *AbstractSyntax
}

//New AbstractSyntax tree from a token
func New(token token.Token) *AbstractSyntax {
	return &AbstractSyntax{
		token: token,
		left:  nil,
		right: nil,
	}
}

func (t *AbstractSyntax) AddLeft(tree *AbstractSyntax) {
	if t.left == nil {
		t.left = tree
	} else {
		t.left.AddLeft(tree)
	}
}

func (t *AbstractSyntax) AddRight(tree *AbstractSyntax) {
	if t.right == nil {
		t.right = tree
	} else {
		t.right.AddRight(tree)
	}
}
