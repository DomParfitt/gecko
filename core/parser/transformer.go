package parser

import (
	"github.com/DomParfitt/gecko/server/api"
)

type Transformer interface {
	transform() api.AST
}

func (r *RegExpr) transform() api.AST {
	children := []api.AST{}

	if r.union != nil {
		children = append(children, r.union.transform())
	}

	if r.simple != nil {
		children = append(children, r.simple.transform())
	}

	return api.AST{
		Label:    "RegExpr",
		Children: children,
	}
}

func (u *Union) transform() api.AST {
	children := []api.AST{}
	children = append(children, u.regex.transform())
	children = append(children, api.AST{Label: "|", Children: []api.AST{}})
	children = append(children, u.simple.transform())

	return api.AST{
		Label:    "Union",
		Children: children,
	}
}

func (s *SimpleExpr) transform() api.AST {
	children := []api.AST{}
	return api.AST{
		Label:    "Union",
		Children: children,
	}
}

func (c *Concatenation) transform() api.AST {
	children := []api.AST{}
	return api.AST{
		Label:    "Union",
		Children: children,
	}
}

func (b *BasicExpr) transform() api.AST {
	children := []api.AST{}
	return api.AST{
		Label:    "Union",
		Children: children,
	}
}

func (s *Star) transform() api.AST {
	children := []api.AST{}
	return api.AST{
		Label:    "Union",
		Children: children,
	}
}

func (p *Plus) transform() api.AST {
	children := []api.AST{}
	return api.AST{
		Label:    "Union",
		Children: children,
	}
}

func (q *Question) transform() api.AST {
	children := []api.AST{}
	return api.AST{
		Label:    "Union",
		Children: children,
	}
}

func (e *Element) transform() api.AST {
	children := []api.AST{}
	return api.AST{
		Label:    "Union",
		Children: children,
	}
}

func (g *Group) transform() api.AST {
	children := []api.AST{}
	return api.AST{
		Label:    "Union",
		Children: children,
	}
}

func (e *Escape) transform() api.AST {
	children := []api.AST{}
	return api.AST{
		Label:    "Union",
		Children: children,
	}
}

func (s *Set) transform() api.AST {
	children := []api.AST{}
	return api.AST{
		Label:    "Union",
		Children: children,
	}
}

func (p *PositiveSet) transform() api.AST {
	children := []api.AST{}
	return api.AST{
		Label:    "Union",
		Children: children,
	}
}

func (n *NegativeSet) transform() api.AST {
	children := []api.AST{}
	return api.AST{
		Label:    "Union",
		Children: children,
	}
}

func (s *SetItems) transform() api.AST {
	children := []api.AST{}
	return api.AST{
		Label:    "Union",
		Children: children,
	}
}

func (s *SetItem) transform() api.AST {
	children := []api.AST{}
	return api.AST{
		Label:    "Union",
		Children: children,
	}
}

func (r *Range) transform() api.AST {
	children := []api.AST{}
	return api.AST{
		Label:    "Union",
		Children: children,
	}
}

func (c *Character) transform() api.AST {
	children := []api.AST{}
	return api.AST{
		Label:    "Union",
		Children: children,
	}
}

func (b *Base) transform() api.AST {
	children := []api.AST{}
	return api.AST{
		Label:    "Union",
		Children: children,
	}
}
