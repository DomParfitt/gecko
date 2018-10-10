// Package parser contains Gecko's recursive-descent parser, data structures representing
// the grammar productions and functions for compiling the AST to a Finite State Machine
package parser

import "github.com/DomParfitt/gecko/core/automata"

//Compiler interface for compiling something into a Finite State Machine
type Compiler interface {
	Compile() *automata.FiniteState
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

//Compile a Union into a Finite State Machine
func (u *Union) Compile() *automata.FiniteState {
	a := u.simple.Compile()
	b := u.regex.Compile()
	a.Union(b)
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

//Compile a Concatenation into a Finite State Machine
func (c *Concatenation) Compile() *automata.FiniteState {
	a := c.basic.Compile()
	b := c.simple.Compile()
	a.Append(b)
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

	if b.question != nil {
		return b.question.Compile()
	}

	if b.element != nil {
		return b.element.Compile()
	}

	panic("invalid")
}

//Compile a Star into a Finite State Machine
func (s *Star) Compile() *automata.FiniteState {
	a := s.element.Compile()
	a.Loop()
	return a
}

//Compile a Plus into a Finite State Machine
func (p *Plus) Compile() *automata.FiniteState {
	a := p.element.Compile()
	b := p.element.Compile()
	b.Loop()
	a.Append(b)
	return a
}

//Compile a Question into a Finite State Machine
func (q *Question) Compile() *automata.FiniteState {
	a := q.element.Compile()
	a.TerminalStates = append(a.TerminalStates, 0)
	return a
}

//Compile an Element into a Finite State Machine
func (e *Element) Compile() *automata.FiniteState {
	if e.group != nil {
		return e.group.Compile()
	}

	if e.set != nil {
		return e.set.Compile()
	}

	if e.character != nil {
		return e.character.Compile()
	}

	if e.escape != nil {
		return e.escape.Compile()
	}

	panic("invalid")
}

//Compile a Group into a Finite State Machine
func (g *Group) Compile() *automata.FiniteState {
	return g.regExpr.Compile()
}

//Compile an Escape into a Finite State Machine
func (e *Escape) Compile() *automata.FiniteState {
	// return e.character.Compile()
	return e.base.Compile()
}

//Compile a Set into a Finite State Machine
func (s *Set) Compile() *automata.FiniteState {
	if s.positive != nil {
		return s.positive.Compile()
	}

	if s.negative != nil {
		return s.negative.Compile()
	}

	panic("invalid")
}

//Compile a PositiveSet into a Finite State Machine
func (p *PositiveSet) Compile() *automata.FiniteState {
	return p.items.Compile()
}

//Compile a NegativeSet into a Finite State Machine
func (n *NegativeSet) Compile() *automata.FiniteState {
	a := n.items.Compile()
	a.Negate()
	return a
}

//Compile a SetItems into a Finite State Machine
func (s *SetItems) Compile() *automata.FiniteState {
	a := s.item.Compile()

	if s.items != nil {
		b := s.items.Compile()
		a.Union(b)
	}

	return a
}

//Compile a SetItem into a Finite State Machine
func (s *SetItem) Compile() *automata.FiniteState {
	if s.rnge != nil {
		return s.rnge.Compile()
	}

	if s.character != nil {
		return s.character.Compile()
	}

	panic("invalid")
}

//Compile a Range into a Finite State Machine
func (r *Range) Compile() *automata.FiniteState {
	chars := []rune{}
	for i := r.start.base.Value; i <= r.end.base.Value; i++ {
		chars = append(chars, i)
	}

	return automata.Create(chars)
}

//Compile a Character into a Finite State Machine
func (c *Character) Compile() *automata.FiniteState {
	// return automata.Create([]rune{c.Value})
	return c.base.Compile()
}

//Compile a Base into a Finite State Machine
func (b *Base) Compile() *automata.FiniteState {
	return automata.Create([]rune{b.Value})
}

// Compile something implementing the Compiler interface and return the result
// on the provided channel
func compile(ch chan<- *automata.FiniteState, compilable Compiler) {
	ch <- compilable.Compile()
}
