package parser

import (
	"github.com/DomParfitt/gecko/server/api"
)

// Transformer interface for types that can be transformed into the
// API's AST representation
type Transformer interface {
	Transform() api.AST
}

// Transform a RegExpr into the API's AST representation
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

// Transform a Union into the API's AST representation
func (u *Union) Transform() api.AST {
	children := []api.AST{}
	children = append(children, u.simple.Transform())
	children = append(children, api.AST{Label: "|", Children: []api.AST{}})
	children = append(children, u.regex.Transform())

	return api.AST{
		Label:    "Union",
		Children: children,
	}
}

// Transform a SimpleExpr into the API's AST representation
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

// Transform a Concatenation into the API's AST representation
func (c *Concatenation) Transform() api.AST {
	children := []api.AST{}
	children = append(children, c.basic.Transform())
	children = append(children, c.simple.Transform())
	return api.AST{
		Label:    "Concatenation",
		Children: children,
	}
}

// Transform a BasicExpr into the API's AST representation
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

// Transform a Star into the API's AST representation
func (s *Star) Transform() api.AST {
	children := []api.AST{}
	children = append(children, s.element.Transform())
	children = append(children, api.AST{Label: "*", Children: []api.AST{}})

	return api.AST{
		Label:    "Star",
		Children: children,
	}
}

// Transform a Plus into the API's AST representation
func (p *Plus) Transform() api.AST {
	children := []api.AST{}
	children = append(children, p.element.Transform())
	children = append(children, api.AST{Label: "+", Children: []api.AST{}})

	return api.AST{
		Label:    "Plus",
		Children: children,
	}
}

// Transform a Question into the API's AST representation
func (q *Question) Transform() api.AST {
	children := []api.AST{}
	children = append(children, q.element.Transform())
	children = append(children, api.AST{Label: "?", Children: []api.AST{}})

	return api.AST{
		Label:    "Question",
		Children: children,
	}
}

// Transform an Element into the API's AST representation
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

// Transform a Group into the API's AST representation
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

// Transform an Escape into the API's AST representation
func (e *Escape) Transform() api.AST {
	children := []api.AST{}
	children = append(children, api.AST{Label: "\\", Children: []api.AST{}})
	children = append(children, e.base.Transform())
	return api.AST{
		Label:    "Escape",
		Children: children,
	}
}

// Transform a Set into the API's AST representation
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

// Transform a PositiveSet into the API's AST representation
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

// Transform a NegativeSet into the API's AST representation
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

// Transform a SetItems into the API's AST representation
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

// Transform a SetItem into the API's AST representation
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

// Transform a Range into the API's AST representation
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

// Transform a Character into the API's AST representation
func (c *Character) Transform() api.AST {
	children := []api.AST{}
	children = append(children, c.base.Transform())
	return api.AST{
		Label:    "Character",
		Children: children,
	}
}

// Transform a Base into the API's AST representation
func (b *Base) Transform() api.AST {
	children := []api.AST{}
	children = append(children, api.AST{Label: string(b.Value), Children: []api.AST{}})
	return api.AST{
		Label:    "Base",
		Children: children,
	}
}
