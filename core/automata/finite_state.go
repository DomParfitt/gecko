// Package automata contains definitions and functionality relating to representation of
// regular expresssions as executable automata
package automata

import (
	"sort"
)

// FiniteState structure representing a Finite State Machine
type FiniteState struct {
	currentNode int
	Nodes       map[int]bool
	Edges       map[Edge]interface{}
}

// NewFiniteState initialises a new FiniteState with a two states and a transition between them using
// the given character
func NewFiniteState(char rune) *FiniteState {
	FiniteState := emptyFiniteState()
	FiniteState.AddEdge(0, 1, char)
	FiniteState.Nodes[1] = true
	return FiniteState
}

// emptyFiniteState creates an empty FiniteState with no nodes or edges
func emptyFiniteState() *FiniteState {
	return &FiniteState{
		currentNode: 0,
		Nodes:       make(map[int]bool),
		Edges:       make(map[Edge]interface{}),
	}
}

// Execute the input string against the automata.
// Returns true if the string matches, false otherwise.
func (f *FiniteState) Execute(input string) bool {
	f.currentNode = 0
	for _, ch := range input {
		if !f.transition(ch) {
			return false
		}
	}
	return f.Nodes[f.currentNode]
}

// transition attempts to transition from the current state using the provided character.
// Returns true if the transition is possible, false otherwise. Updates the current state
// if the transition is possible
func (f *FiniteState) transition(char rune) bool {
	edges := f.edgesFrom(f.currentNode)
	for _, edge := range edges {
		if edge.Label == char {
			f.currentNode = edge.To
			return true
		}
	}
	return false
}

// Append adds the other FiniteState to the end of the current one.
// I.e. if given A = [0]-a->[1*], and B = [0]-b->[1*] then appending
// B to A (A.Append(B)) should result in A = [0]-a->[1]-b->[2*]
func (f *FiniteState) Append(other *FiniteState) {
	offset := f.nextState()
	copy := other.copyWithOffset(offset)
	terminals := f.terminals()

	for edge := range copy.Edges {
		from := edge.From
		to := edge.To
		char := edge.Label
		if from == 0 {
			for _, terminal := range terminals {
				f.AddEdge(terminal, to, char)
				f.Nodes[terminal] = false
			}
		} else if to == 0 {
			for _, terminal := range terminals {
				f.AddEdge(from, terminal, char)
				f.Nodes[terminal] = false
			}
		} else {
			f.AddEdge(from, to, char)
		}
	}

	for _, terminal := range copy.terminals() {
		f.Nodes[terminal] = true
	}
}

// Union adds the other FiniteState to the current one as a branch.
// I.e. if given A = [0]-a->[1*], and B = [0]-b->[1*] then unioning
// B to A (A.Union(B)) should result in A = [0]-(a,b)->[1*]
func (f *FiniteState) Union(other *FiniteState) {
	offset := f.nextState()
	copy := other.copyWithOffset(offset)

	for edge := range copy.Edges {
		from := edge.From
		to := edge.To
		char := edge.Label

		f.AddEdge(from, to, char)
		f.Nodes[from] = copy.Nodes[edge.From]
		f.Nodes[to] = copy.Nodes[edge.To]
	}
}

// Loop sets the FiniteState to loop back on itself.
// I.e. if given A = [0]-a->[1*] then looping A on itself (A.Loop())
// should result in A = [0*]-a->[1*]-a->[1*]
func (f *FiniteState) Loop() {
	edgesFromInitial := f.edgesFrom(0)

	for _, edge := range edgesFromInitial {
		terminals := f.terminals()
		for _, terminal := range terminals {
			to := edge.To
			from := terminal
			char := edge.Label
			f.AddEdge(from, to, char)
		}
	}

	f.Nodes[0] = true
}

// AddEdge adds a new Edge to the FiniteState if a matching Edge does not
// already exist, adding new states if required.
func (f *FiniteState) AddEdge(from, to int, char rune) {
	f.addState(from, false)
	f.addState(to, false)

	edge := Edge{From: from, To: to, Label: char}

	if _, exists := f.Edges[edge]; !exists {
		f.Edges[edge] = new(interface{})
	}
}

// edgesTo retrieves all the Edges going to a particular state
func (f *FiniteState) edgesTo(to int) []Edge {
	return f.matchingEdges(func(edge Edge) bool {
		return edge.To == to
	})
}

// edgesFrom retrieves all the Edges coming from a particular state
func (f *FiniteState) edgesFrom(from int) []Edge {
	return f.matchingEdges(func(edge Edge) bool {
		return edge.From == from
	})
}

// matchingEdges retrieves all edges which are valid according to the given
// function, matcher. I.e. returns all edges where matcher(edge) == true
func (f *FiniteState) matchingEdges(matcher func(edge Edge) bool) []Edge {
	edges := []Edge{}
	for edge := range f.Edges {
		if matcher(edge) {
			edges = append(edges, edge)
		}
	}
	return edges
}

// hasState returns whether a state is present in the FiniteState or not
func (f *FiniteState) hasState(state int) bool {
	_, ok := f.Nodes[state]
	return ok
}

// nextState returns the next unused state ID
func (f *FiniteState) nextState() int {
	next := 0
	for state := range f.Nodes {
		if state > next {
			next = state
		}
	}
	return next
}

// addState adds the given state if it doesn't exist
func (f *FiniteState) addState(state int, terminal bool) {
	if _, ok := f.Nodes[state]; !ok {
		f.Nodes[state] = terminal
	}
}

// allStates retrieves all the states in the FiniteState as a slice
func (f *FiniteState) allStates() []int {
	states := []int{}
	for state := range f.Nodes {
		states = append(states, state)
	}
	sort.Ints(states)
	return states
}

// terminals retrieves all the terminal states in the FiniteState as a slice
func (f *FiniteState) terminals() []int {
	terminals := []int{}
	for state, terminal := range f.Nodes {
		if terminal {
			terminals = append(terminals, state)
		}
	}
	sort.Ints(terminals)
	return terminals
}

// Copy the FiniteState, creating a new instance with identical values
func (f *FiniteState) copy() *FiniteState {
	return f.copyWithOffset(0)
}

// copyWithOffset copies the FiniteState applying the given offset to all states
func (f *FiniteState) copyWithOffset(offset int) *FiniteState {
	// Make an empty FiniteState as the start
	copy := emptyFiniteState()

	// Copy all the nodes
	for state, terminal := range f.Nodes {
		if state == 0 {
			copy.Nodes[state] = terminal
		} else {
			copy.Nodes[state+offset] = terminal
		}
	}

	// Copy all the edges
	for edge := range f.Edges {
		edgeCopy := edge.copyWithOffset(offset)
		copy.Edges[edgeCopy] = new(interface{})
	}

	return copy
}
