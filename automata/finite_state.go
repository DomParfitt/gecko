package automata

// Executer interface
type Executer interface {
	execute(input string) bool
}

// FiniteState struct
type FiniteState struct {
	currentState int
	states       []bool
	transitions  map[int]map[rune]int
}

// New creates a default instance of a FiniteState
// struct and returns a pointer to it
func New() *FiniteState {
	return &FiniteState{
		currentState: 0,
		states:       []bool{false},
		transitions:  make(map[int]map[rune]int),
	}
}

// AddState to the current set of states
func (f *FiniteState) AddState(isTerminal bool) {
	f.states = append(f.states, isTerminal)
}

// AddTransition from one state to another which consumes one of the
// provided characters. If there already exists a transition from the
// 'from' state using one of the provided charcters then it is overwritten.
func (f *FiniteState) AddTransition(from, to int, chars []rune) {
	// If we have a transition set from this state already
	// the add/update
	if transitionsFrom, ok := f.transitions[from]; ok {
		for _, ch := range chars {
			transitionsFrom[ch] = to
		}
	} else {
		transitionsFrom := make(map[rune]int)
		for _, ch := range chars {
			transitionsFrom[ch] = to
		}
		f.transitions[from] = transitionsFrom
	}
}

// Execute the provided input string on the automata that
// f represents.
// Returns true if the input string is accepted by the automata
// and false otherwise
func (f *FiniteState) Execute(input string) bool {
	for _, ch := range input {
		if !f.consume(ch) {
			return false
		}
	}
	return f.states[f.currentState]
}

// Consume a character and updates the state of the
// automata as required.
// Returns a bool indicating success or failure, where
// failure indicates that the given character could not
// be consumed from the current state of the automata
func (f *FiniteState) consume(ch rune) bool {
	return f.transition(f.currentState, ch)
}

// Transition from the 'from' state to the 'to' state,
// consuming the provided character in the process.
// Returns true if successful, false if there is no such
// transition between the two states using the given character
func (f *FiniteState) transition(from int, ch rune) bool {

	if to, ok := f.transitions[from][ch]; ok {
		f.currentState = to
		return true
	}

	return false
}
