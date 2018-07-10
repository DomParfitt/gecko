package parser

import (
	"github.com/DomParfitt/gecko/automata"
)

type Compiler interface {
	Compile() *automata.FiniteState
}

func (e *Element) Compile() *automata.FiniteState {
	a := automata.New()
	a.AddTransition(0, 1, []rune{e.Value})
	return a
}

func (p *Plus) Compile() *automata.FiniteState {
	a := p.element.Compile()
	b := p.element.Compile()
	b.Loop()
	a.Append(b)
	return a
}

func (s *Star) Compile() *automata.FiniteState {
	a := s.element.Compile()
	a.Loop()
	return a
}

func (b *BasicExpr) Compile() *automata.FiniteState {
	if b.star != nil {
		return b.star.Compile()
	}

	if b.plus != nil {
		return b.plus.Compile()
	}

	if b.element != nil {
		return b.element.Compile()
	}

	panic("invalid")
}

func (c *Concatenation) Compile() *automata.FiniteState {
	a := c.simple.Compile()
	b := c.basic.Compile()
	a.Append(b)
	return a
}

func (s *SimpleExpr) Compile() *automata.FiniteState {
	if s.concatenation != nil {
		return s.concatenation.Compile()
	}

	if s.basic != nil {
		return s.basic.Compile()
	}

	panic("invalid")
}

func (u *Union) Compile() *automata.FiniteState {
	a := u.regex.Compile()
	b := u.simple.Compile()
	a.Union(b)
	return a
}

func (r *RegExpr) Compile() *automata.FiniteState {
	if r.union != nil {
		return r.union.Compile()
	}

	if r.simple != nil {
		return r.simple.Compile()
	}

	panic("invalid")
}
