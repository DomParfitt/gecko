package parser

import (
	"github.com/DomParfitt/gecko/automata"
)

//Compiler interface for compiling something into a Finite State Machine
type Compiler interface {
	Compile() *automata.FiniteState
}

//Compile an Element into a Finite State Machine
func (e *Element) Compile() *automata.FiniteState {
	if e.group != nil {
		return e.group.Compile()
	}

	return automata.Create([]rune{e.Value})
}

//Compile a Plus into a Finite State Machine
func (p *Plus) Compile() *automata.FiniteState {
	a := p.element.Compile()
	b := p.element.Compile()
	b.Loop()
	a.Append(b)
	return a
}

//Compile a Star into a Finite State Machine
func (s *Star) Compile() *automata.FiniteState {
	a := s.element.Compile()
	a.Loop()
	return a
}

//Compile a BasicExpr into a Finite State Machine
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

//Compile a Concatenation into a Finite State Machine
func (c *Concatenation) Compile() *automata.FiniteState {
	a := c.basic.Compile()
	b := c.simple.Compile()
	a.Append(b)
	return a
}

//Compile a SimpleExpr into a Finite State Machine
func (s *SimpleExpr) Compile() *automata.FiniteState {
	if s.concatenation != nil {
		return s.concatenation.Compile()
	}

	if s.basic != nil {
		return s.basic.Compile()
	}

	panic("invalid")
}

//Compile a Union into a Finite State Machine
func (u *Union) Compile() *automata.FiniteState {
	a := u.simple.Compile()
	b := u.regex.Compile()
	a.Union(b)
	return a
}

//Compile a Group into a Finite State Machine
func (g *Group) Compile() *automata.FiniteState {
	return g.regExpr.Compile()
}

//Compile a RegExpr into a Finite State Machine
func (r *RegExpr) Compile() *automata.FiniteState {
	if r.union != nil {
		return r.union.Compile()
	}

	if r.simple != nil {
		return r.simple.Compile()
	}

	panic("invalid")
}

// Compile something implementing the Compiler interface and return the result
// on the provided channel
func compile(ch chan<- *automata.FiniteState, compilable Compiler) {
	ch <- compilable.Compile()
}
