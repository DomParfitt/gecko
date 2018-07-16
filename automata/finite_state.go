package automata

import (
	"fmt"
)

// Executer interface
type Executer interface {
	execute(input string) bool
}

// FiniteState struct
type FiniteState struct {
	nextState     int
	currentState  int
	terminalState int
	transitions   map[int]map[rune]int
}

// New creates a default instance of a FiniteState
// struct and returns a pointer to it
func New() *FiniteState {
	return &FiniteState{
		nextState:     1,
		currentState:  0,
		terminalState: 0,
		transitions:   make(map[int]map[rune]int),
	}
}

//SetTerminal state to provided
func (f *FiniteState) SetTerminal(terminal int) {
	f.terminalState = terminal
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

	//Update transitions to the original terminal state
	for _, transition := range f.transitions {
		for ch, to := range transition {
			if to == f.terminalState {
				transition[ch] = offset
			}
		}
	}

	//Copy transitions from other
	for from, transition := range other.transitions {
		for ch, to := range transition {
			f.AddTransition(from+offset, to+offset, []rune{ch})
		}
	}

	//Update transitions from the original terminal state
	if transition, ok := f.transitions[f.terminalState]; ok {
		for ch, to := range transition {
			f.AddTransition(other.terminalState+offset, to, []rune{ch})
		}
		delete(f.transitions, f.terminalState)
	}

	//Set new terminal
	f.terminalState = other.terminalState + offset

}

//Union the given automata with this one
func (f *FiniteState) Union(other *FiniteState) {
	offset := f.nextState
	f.nextState += other.nextState

	//Copy transitions from other
	for from, transition := range other.transitions {

		if from != 0 {
			from += offset
		}

		if from == other.terminalState {
			from = f.terminalState
		}

		for ch, to := range transition {

			if to == other.terminalState {
				to = f.terminalState
			} else if to != 0 {
				to += offset
			}

			f.AddTransition(from, to, []rune{ch})
		}
	}
}

// Loop this automata on itself
func (f *FiniteState) Loop() {
	for from, transition := range f.transitions {
		if from == f.terminalState {
			for ch, to := range transition {
				f.AddTransition(0, to, []rune{ch})
			}
			delete(f.transitions, from)
		} else {
			for ch, to := range transition {
				if to == f.terminalState {
					transition[ch] = 0
				}
			}
		}
	}

	f.terminalState = 0
}

func (f *FiniteState) String() string {
	str := fmt.Sprintf("Terminal: %d\n", f.terminalState)
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
	return f.currentState == f.terminalState
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
