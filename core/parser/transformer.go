package parser

import (
	"github.com/DomParfitt/gecko/server/api"
)

type Transformer interface {
	Transform() api.AST
}

func (r *RegExpr) Transform() api.AST {
	children := []api.AST{}

	if r.union != nil {
		children = append(children, r.union.Transform())
	}

	if r.simple != nil {
		children = append(children, r.simple.Transform())
	}

	return api.AST{
		Label:    "RegExpr",
		Children: children,
	}
}

func (u *Union) Transform() api.AST {
	children := []api.AST{}
	children = append(children, u.regex.Transform())
	children = append(children, api.AST{Label: "|", Children: []api.AST{}})
	children = append(children, u.simple.Transform())

	return api.AST{
		Label:    "Union",
		Children: children,
	}
}

func (s *SimpleExpr) Transform() api.AST {
	children := []api.AST{}
	if s.concatenation != nil {
		children = append(children, s.concatenation.Transform())
	}

	if s.basic != nil {
		children = append(children, s.basic.Transform())
	}

	return api.AST{
		Label:    "SimpleExpr",
		Children: children,
	}
}

func (c *Concatenation) Transform() api.AST {
	children := []api.AST{}
	children = append(children, c.simple.Transform())
	children = append(children, c.basic.Transform())
	return api.AST{
		Label:    "Concatenation",
		Children: children,
	}
}

func (b *BasicExpr) Transform() api.AST {
	children := []api.AST{}

	if b.star != nil {
		children = append(children, b.star.Transform())
	}

	if b.plus != nil {
		children = append(children, b.plus.Transform())
	}

	if b.question != nil {
		children = append(children, b.question.Transform())
	}

	if b.element != nil {
		children = append(children, b.element.Transform())
	}

	return api.AST{
		Label:    "BasicExpr",
		Children: children,
	}
}

func (s *Star) Transform() api.AST {
	children := []api.AST{}
	children = append(children, s.element.Transform())
	children = append(children, api.AST{Label: "*", Children: []api.AST{}})

	return api.AST{
		Label:    "Star",
		Children: children,
	}
}

func (p *Plus) Transform() api.AST {
	children := []api.AST{}
	children = append(children, p.element.Transform())
	children = append(children, api.AST{Label: "+", Children: []api.AST{}})

	return api.AST{
		Label:    "Plus",
		Children: children,
	}
}

func (q *Question) Transform() api.AST {
	children := []api.AST{}
	children = append(children, q.element.Transform())
	children = append(children, api.AST{Label: "?", Children: []api.AST{}})

	return api.AST{
		Label:    "Question",
		Children: children,
	}
}

func (e *Element) Transform() api.AST {
	children := []api.AST{}

	if e.character != nil {
		children = append(children, e.character.Transform())
	}

	if e.group != nil {
		children = append(children, e.group.Transform())
	}

	if e.set != nil {
		children = append(children, e.set.Transform())
	}

	if e.escape != nil {
		children = append(children, e.escape.Transform())
	}

	return api.AST{
		Label:    "Element",
		Children: children,
	}
}

func (g *Group) Transform() api.AST {
	children := []api.AST{}
	children = append(children, api.AST{Label: "(", Children: []api.AST{}})
	children = append(children, g.regExpr.Transform())
	children = append(children, api.AST{Label: ")", Children: []api.AST{}})

	return api.AST{
		Label:    "Group",
		Children: children,
	}
}

func (e *Escape) Transform() api.AST {
	children := []api.AST{}
	children = append(children, api.AST{Label: "\\", Children: []api.AST{}})
	children = append(children, e.base.Transform())
	return api.AST{
		Label:    "Escape",
		Children: children,
	}
}

func (s *Set) Transform() api.AST {
	children := []api.AST{}

	if s.positive != nil {
		children = append(children, s.positive.Transform())
	}

	if s.negative != nil {
		children = append(children, s.negative.Transform())
	}

	return api.AST{
		Label:    "Set",
		Children: children,
	}
}

func (p *PositiveSet) Transform() api.AST {
	children := []api.AST{}

	children = append(children, api.AST{Label: "[", Children: []api.AST{}})
	children = append(children, p.items.Transform())
	children = append(children, api.AST{Label: "]", Children: []api.AST{}})

	return api.AST{
		Label:    "PositiveSet",
		Children: children,
	}
}

func (n *NegativeSet) Transform() api.AST {
	children := []api.AST{}
	children = append(children, api.AST{Label: "[", Children: []api.AST{}})
	children = append(children, api.AST{Label: "^", Children: []api.AST{}})
	children = append(children, n.items.Transform())
	children = append(children, api.AST{Label: "]", Children: []api.AST{}})

	return api.AST{
		Label:    "NegativeSet",
		Children: children,
	}
}

func (s *SetItems) Transform() api.AST {
	children := []api.AST{}
	children = append(children, s.item.Transform())

	if s.items != nil {
		children = append(children, s.items.Transform())
	}
	return api.AST{
		Label:    "SetItems",
		Children: children,
	}
}

func (s *SetItem) Transform() api.AST {
	children := []api.AST{}

	if s.rnge != nil {
		children = append(children, s.rnge.Transform())
	}

	if s.character != nil {
		children = append(children, s.character.Transform())
	}

	return api.AST{
		Label:    "SetItem",
		Children: children,
	}
}

func (r *Range) Transform() api.AST {
	children := []api.AST{}
	children = append(children, r.start.Transform())
	children = append(children, api.AST{Label: "-", Children: []api.AST{}})
	children = append(children, r.end.Transform())

	return api.AST{
		Label:    "Range",
		Children: children,
	}
}

func (c *Character) Transform() api.AST {
	children := []api.AST{}
	children = append(children, c.base.Transform())
	return api.AST{
		Label:    "Character",
		Children: children,
	}
}

func (b *Base) Transform() api.AST {
	children := []api.AST{}
	children = append(children, api.AST{Label: string(b.Value), Children: []api.AST{}})
	return api.AST{
		Label:    "Base",
		Children: children,
	}
}
