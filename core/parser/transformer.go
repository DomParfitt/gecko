package parser

import (
	"github.com/DomParfitt/gecko/server/api"
)

type Transformer interface {
	transform() api.AST
}

func (r *RegExpr) transform() api.AST {
	ast := api.AST{
		Label:    "RegExpr",
		Children: []api.AST{},
	}
	return ast
}

func (u *Union) transform() api.AST {

}

func (s *SimpleExpr) transform() api.AST {

}

func (c *Concatenation) transform() api.AST {

}

func (b *BasicExpr) transform() api.AST {

}

func (s *Star) transform() api.AST {

}

func (p *Plus) transform() api.AST {

}

func (q *Question) transform() api.AST {

}

func (e *Element) transform() api.AST {

}

func (g *Group) transform() api.AST {

}

func (e *Escape) transform() api.AST {

}

func (s *Set) transform() api.AST {

}

func (p *PositiveSet) transform() api.AST {

}

func (n *NegativeSet) transform() api.AST {

}

func (s *SetItems) transform() api.AST {

}

func (s *SetItem) transform() api.AST {

}

func (r *Range) transform() api.AST {

}

func (c *Character) transform() api.AST {

}

func (b *Base) transform() api.AST {

}
