package automata

import (
	"fmt"
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
	f.AddTransition(0, 1, []rune{'a'})
	f.SetTerminal(1)
	if !f.Execute("a") {
		t.Errorf("Expected success transitioning")
	}
}

func TestAddTransitionOverwrite(t *testing.T) {
	f := New()
	f.AddTransition(0, 1, []rune{'a'})
	if transition, ok := f.transitions[0]; ok {
		if to, ok := transition['a']; !ok || to != 1 {
			t.Errorf("Expected transition did not exist")
		}
	} else {
		t.Errorf("Expected transition did not exist")
	}

	f.AddTransition(0, 2, []rune{'a'})
	if transition, ok := f.transitions[0]; ok {
		if to, ok := transition['a']; !ok || to != 2 {
			t.Errorf("Expected transition did not exist")
		}
	} else {
		t.Errorf("Expected transition did not exist")
	}
}

func TestAddTransitionToNonExistantState(t *testing.T) {
	f := New()
	f.AddTransition(1, 2, []rune{'a'})
}

func TestAppend(t *testing.T) {
	f := New()
	f.AddTransition(0, 1, []rune{'a'})
	f.AddTransition(1, 2, []rune{'b'})
	f.AddTransition(2, 1, []rune{'x'})
	f.SetTerminal(2)

	g := New()
	g.AddTransition(0, 1, []rune{'c'})
	g.SetTerminal(1)

	f.Append(g)
	fmt.Println(f)

	if !f.Execute("abc") {
		t.Errorf("Error appending")
	}
}

func TestUnion(t *testing.T) {
	f := New()
	f.AddTransition(0, 1, []rune{'a'})
	f.SetTerminal(1)

	g := New()
	g.AddTransition(0, 1, []rune{'b'})
	g.SetTerminal(1)

	f.Union(g)

	if !f.Execute("a") {
		t.Errorf("Error unioning")
	}

	if !f.Execute("b") {
		t.Errorf("Error unioning")
	}

}

func TestLoop(t *testing.T) {
	f := New()
	f.AddTransition(0, 1, []rune{'a'})
	f.AddTransition(1, 2, []rune{'b'})
	f.AddTransition(2, 3, []rune{'c'})
	f.AddTransition(3, 2, []rune{'d'})
	f.SetTerminal(2)
	f.Loop()

	fmt.Println(f)

	if !f.Execute("abcdabcd") {
		t.Errorf("Error looping")
	}
}
