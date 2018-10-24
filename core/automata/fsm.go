package automata

import (
	"sort"
)

// FSM structure representing a Finite State Machine
type FSM struct {
	currentNode int
	Nodes       map[int]bool         //[]Node
	Edges       map[Edge]interface{} //[]Edge
}

// Edge structure representing a transition within an FSM
type Edge struct {
	From  int
	To    int
	Label rune
}

// Copy makes a copy of the given Edge
func (e Edge) copy() Edge {
	return e.copyWithOffset(0)
}

// copyWithOffset makes a copy of the given Edge with the
// From and To states offset by the given value
func (e Edge) copyWithOffset(offset int) Edge {
	return Edge{
		From:  e.From + offset,
		To:    e.To + offset,
		Label: e.Label,
	}
}

// NewFSM initialises a new FSM with a two states and a transition between them using
// the given character
func NewFSM(char rune) *FSM {
	fsm := emptyFSM()
	fsm.AddEdge(0, 1, char)
	fsm.Nodes[1] = true
	return fsm
}

// emptyFSM creates an empty FSM with no nodes or edges
func emptyFSM() *FSM {
	return &FSM{
		currentNode: 0,
		Nodes:       make(map[int]bool),
		Edges:       make(map[Edge]interface{}), //[]Edge{},
	}
}

func (f *FSM) Execute(input string) bool {
	f.currentNode = 0
	for _, ch := range input {
		edges := f.edgesFrom(f.currentNode)

		if len(edges) == 0 {
			return false
		}

		for _, edge := range edges {
			if edge.Label == ch {
				f.currentNode = edge.To
				break
			}
			return false
		}
	}
	return f.Nodes[f.currentNode]
}

// Append adds the other FSM to the end of the current one.
// I.e. if given A = [0]-a->[1*], and B = [0]-b->[1*] then appending
// B to A (A.Append(B)) should result in A = [0]-a->[1]-b->[2*]
func (f *FSM) Append(other *FSM) {
	offset := f.nextState()
	copy := other.copyWithOffset(offset)
	terminals := f.terminals()

	for edge := range copy.Edges {
		from := edge.From
		to := edge.To
		char := edge.Label
		if from-offset == 0 {
			for _, terminal := range terminals {
				f.AddEdge(terminal, to, char)
				f.Nodes[terminal] = false
			}
		} else if to-offset == 0 {
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

// Union adds the other FSM to the current one as a branch.
// I.e. if given A = [0]-a->[1*], and B = [0]-b->[1*] then unioning
// B to A (A.Union(B)) should result in A = [0]-(a,b)->[1*]
func (f *FSM) Union(other *FSM) {
	// offset := f.nextState()
	// copy := other.copyWithOffset(offset)

	// for edge := range copy.Edges {
	// 	from := edge.From
	// 	to := edge.To
	// 	char := edge.Label

	// 	if from-offset == 0 {
	// 		from = 0
	// 	}

	// }
}

// Loop sets the FSM to loop back on itself.
// I.e. if given A = [0]-a->[1*] then looping A on itself (A.Loop())
// should result in A = [0*]-a->[1*]-a->[1*]
func (f *FSM) Loop() {
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

// hasState returns whether a state is present in the FSM or not
func (f *FSM) hasState(state int) bool {
	_, ok := f.Nodes[state]
	return ok
}

// nextState returns the next unused state ID
func (f *FSM) nextState() int {
	next := 0
	for state := range f.Nodes {
		if state > next {
			next = state
		}
	}
	return next
}

// addState adds the given state if it doesn't exist
func (f *FSM) addState(state int, terminal bool) {
	if _, ok := f.Nodes[state]; !ok {
		f.Nodes[state] = terminal
	}
}

// allStates retrieves all the states in the FSM as a slice
func (f *FSM) allStates() []int {
	states := []int{}
	for state := range f.Nodes {
		states = append(states, state)
	}
	sort.Ints(states)
	return states
}

// terminals retrieves all the terminal states in the FSM as a slice
func (f *FSM) terminals() []int {
	terminals := []int{}
	for state, terminal := range f.Nodes {
		if terminal {
			terminals = append(terminals, state)
		}
	}
	sort.Ints(terminals)
	return terminals
}

// AddEdge adds a new Edge to the FSM if a matching Edge does not
// already exist, adding new states if required.
func (f *FSM) AddEdge(from, to int, char rune) {
	f.addState(from, false)
	f.addState(to, false)

	edge := Edge{From: from, To: to, Label: char}

	if _, exists := f.Edges[edge]; !exists {
		f.Edges[edge] = new(interface{})
	}

}

// edgesTo retrieves all the Edges going to a particular state
func (f *FSM) edgesTo(to int) []Edge {
	edges := []Edge{}
	for edge := range f.Edges {
		if edge.To == to {
			edges = append(edges, edge)
		}
	}
	return edges
}

// edgesFrom retrieves all the Edges coming from a particular state
func (f *FSM) edgesFrom(from int) []Edge {
	edges := []Edge{}
	for edge := range f.Edges {
		if edge.From == from {
			edges = append(edges, edge)
		}
	}
	return edges
}

// Copy the FSM, creating a new instance with identical values
func (f *FSM) copy() *FSM {
	return f.copyWithOffset(0)
}

// copyWithOffset copies the FSM applying the given offset to all states
func (f *FSM) copyWithOffset(offset int) *FSM {
	// Make an empty FSM as the start
	copy := emptyFSM()

	// Copy all the nodes
	for state, terminal := range f.Nodes {
		copy.Nodes[state+offset] = terminal
	}

	// Copy all the edges
	for edge := range f.Edges {
		edgeCopy := edge.copyWithOffset(offset)
		copy.Edges[edgeCopy] = new(interface{})
	}

	return copy
}
