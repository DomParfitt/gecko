package automata

import (
	"fmt"
	"sort"
)

// Executer interface
type Executer interface {
	execute(input string) bool
}

// FiniteState struct
type FiniteState struct {
	nextState      int
	currentState   int
	terminalStates []int
	transitions    map[int]map[rune]int
}

// New creates a default instance of a FiniteState
// struct and returns a pointer to it
func New() *FiniteState {
	return &FiniteState{
		nextState:      1,
		currentState:   0,
		terminalStates: []int{0},
		transitions:    make(map[int]map[rune]int),
	}
}

// Create a new finite state with two states and a single set of transitions
// from the first to the second using the provided characters
func Create(chars []rune) *FiniteState {
	f := New()
	f.AddTransition(0, 1, chars)
	f.terminalStates = []int{1}
	return f
}

func (f *FiniteState) addTerminal(terminal int) {
	for _, val := range f.terminalStates {
		if val == terminal {
			return
		}
	}
	f.terminalStates = append(f.terminalStates, terminal)
}

// AddTransition from one state to another which consumes one of the
// provided characters. If there already exists a transition from the
// 'from' state using one of the provided charcters then it is overwritten.
func (f *FiniteState) AddTransition(from, to int, chars []rune) {

	//Update the next state indicator if necessary
	if from >= f.nextState {
		f.nextState = from + 1
	}

	if to >= f.nextState {
		f.nextState = to + 1
	}

	// If we have a transition set from this state already
	// then add/update
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

// Append the given automata onto the end of this one
func (f *FiniteState) Append(other *FiniteState) {
	offset := f.nextState
	f.nextState += other.nextState

	//Update transitions from the other initial
	for from, transition := range other.transitions {
		if from == 0 {
			for _, terminal := range f.terminalStates {
				for ch, to := range transition {
					f.AddTransition(terminal, to+offset, []rune{ch})
					// }
				}
			}
		} else {
			for ch, to := range transition {
				if to == 0 {
					for terminal := range f.terminalStates {
						f.AddTransition(from+offset, terminal, []rune{ch})
					}
				} else {
					f.AddTransition(from+offset, to+offset, []rune{ch})
				}
			}
		}
	}

	//Set new terminal
	newTerms := []int{}
	for _, term := range other.terminalStates {
		newTerms = append(newTerms, term+offset)
	}
	f.terminalStates = newTerms

}

//Union the given automata with this one
func (f *FiniteState) Union(other *FiniteState) {
	offset := f.nextState
	f.nextState += other.nextState

	//Copy transitions from other
	for from, transition := range other.transitions {

		isFromTerm := other.isTerminal(from)

		//Anything from the other's initial goes from f's initial
		//Anything else gets offset
		if from != 0 {
			from += offset
		}

		if isFromTerm {
			f.addTerminal(from)
		}

		for ch, to := range transition {

			isToTerm := other.isTerminal(to)

			if to != 0 {
				to += offset
			}

			if isToTerm {
				f.addTerminal(to)
			}

			f.AddTransition(from, to, []rune{ch})
		}
	}
}

// Loop this automata on itself
func (f *FiniteState) Loop() {
	for from, transition := range f.transitions {

		//If the transition comes from the initial state then we need
		//matching transitions from each terminal state
		if from == 0 {
			for _, state := range f.terminalStates {
				for ch, to := range transition {
					f.AddTransition(state, to, []rune{ch})
				}
			}
		}
	}

	f.addTerminal(0)
}

// Negate this automata, i.e. make it non-accepting on it's original pattern
func (f *FiniteState) Negate() {
	terminals := []int{}

	for from, transition := range f.transitions {
		if !f.isTerminal(from) {
			terminals = append(terminals, from)
		}

		for _, to := range transition {
			if !f.isTerminal(to) {
				terminals = append(terminals, to)
			}
		}
	}

	f.terminalStates = terminals
}

func (f *FiniteState) String() string {
	sort.Ints(f.terminalStates)
	str := fmt.Sprintf("Terminals: %v\n", f.terminalStates)
	for from, transition := range f.transitions {
		tran := ""
		for ch, to := range transition {
			tran += fmt.Sprintf("\n    %c => %d", ch, to)
		}
		str += fmt.Sprintf("%d: %s\n", from, tran)
	}
	return str
}

// Execute the provided input string on the automata that
// f represents.
// Returns true if the input string is accepted by the automata
// and false otherwise
func (f *FiniteState) Execute(input string) bool {
	f.currentState = 0
	for _, ch := range input {
		if !f.consume(ch) {
			return false
		}
	}
	return f.isTerminal(f.currentState)
}

// Consume a character and update the state of the
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

func (f *FiniteState) isTerminal(state int) bool {
	for _, val := range f.terminalStates {
		if state == val {
			return true
		}
	}
	return false
}
