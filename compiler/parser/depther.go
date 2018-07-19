package parser

type Depther interface {
	Depth() int
}

func (r *RegExpr) Depth() int {
	depth := 1

	if r.union != nil {
		depth += r.union.Depth()
	} else if r.simple != nil {
		depth += r.simple.Depth()
	}

	return depth
}

func (u *Union) Depth() int {
	depth := 1
	regExDepth := u.regex.Depth()
	simpleDepth := u.simple.Depth()

	if regExDepth > simpleDepth {
		depth += regExDepth
	} else {
		depth += simpleDepth
	}

	return depth
}

func (s *SimpleExpr) Depth() int {
	depth := 1

	if s.concatenation != nil {
		depth += s.concatenation.Depth()
	} else if s.basic != nil {
		depth += s.basic.Depth()
	}

	return depth
}

func (c *Concatenation) Depth() int {
	depth := 1

	simpleDepth := c.simple.Depth()
	basicDepth := c.basic.Depth()

	if simpleDepth > basicDepth {
		depth += simpleDepth
	} else {
		depth += basicDepth
	}

	return depth
}

func (b *BasicExpr) Depth() int {
	depth := 1

	if b.star != nil {
		depth += b.star.Depth()
	} else if b.plus != nil {
		depth += b.plus.Depth()
	} else if b.element != nil {
		depth += b.element.Depth()
	}

	return depth
}

func (s *Star) Depth() int {
	return 1 + s.element.Depth()
}

func (p *Plus) Depth() int {
	return 1 + p.element.Depth()
}

func (e *Element) Depth() int {
	return 1
}
