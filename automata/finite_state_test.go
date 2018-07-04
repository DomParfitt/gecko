package automata

import (
	"testing"
)

func TestExecute(t *testing.T) {
	f := New()
	if f.Execute("hello") {
		t.Errorf("Expected failure on executing default machine")
	}
}

func TestExecutePass(t *testing.T) {
	f := New()
	f.AddState(true)
	f.AddTransition(0, 1, []rune{'a'})
	if !f.Execute("a") {
		t.Errorf("Expected success transitioning")
	}
}

func TestAddTransitionToNonExistantState(t *testing.T) {
	f := New()
	f.AddTransition(1, 2, []rune{'a'})
}
