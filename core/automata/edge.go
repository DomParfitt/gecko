package automata

// Edge structure representing a transition within an FiniteState
type Edge struct {
	From  int
	To    int
	Label rune
}

// copy makes a copy of the given Edge
func (e Edge) copy() Edge {
	return e.copyWithOffset(0)
}

// copyWithOffset makes a copy of the given Edge with the
// From and To states offset by the given value
func (e Edge) copyWithOffset(offset int) Edge {
	from := e.From
	to := e.To

	if from != 0 {
		from += offset
	}

	if to != 0 {
		to += offset
	}

	return Edge{
		From:  from,
		To:    to,
		Label: e.Label,
	}
}
