package automata

import (
	"fmt"
	"reflect"
	"testing"
)

func TestExecute(t *testing.T) {
	f := New()
	if f.Execute("hello") {
		t.Errorf("Expected failure on executing default machine")
	}
}

func TestExecutePass(t *testing.T) {
	f := Create([]rune{'a'})
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
	f.terminalStates = []int{2}

	g := Create([]rune{'c'})

	f.Append(g)
	fmt.Println(f)

	if !f.Execute("abc") {
		t.Errorf("Error appending")
	}
}

func TestUnion(t *testing.T) {
	f := Create([]rune{'a'})

	g := Create([]rune{'b'})

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
	f.terminalStates = []int{2}
	f.Loop()

	fmt.Println(f)

	if !f.Execute("abcdabcd") {
		t.Errorf("Error looping")
	}
}

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want *FiniteState
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFiniteState_AddTransition(t *testing.T) {
	type args struct {
		from  int
		to    int
		chars []rune
	}
	tests := []struct {
		name string
		f    *FiniteState
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.f.AddTransition(tt.args.from, tt.args.to, tt.args.chars)
		})
	}
}

func TestFiniteState_Append(t *testing.T) {
	type args struct {
		other *FiniteState
		input string
	}
	tests := []struct {
		name string
		f    *FiniteState
		args args
		want bool
	}{
		{"Simple Append", Create([]rune{'a'}), args{other: Create([]rune{'b'}), input: "ab"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.f.Append(tt.args.other)
			if got := tt.f.Execute(tt.args.input); got != tt.want {
				t.Errorf("FiniteState.Append() on input %s = %v, want %v", tt.args.input, got, tt.want)
			}
		})
	}
}

func TestFiniteState_Union(t *testing.T) {
	type args struct {
		other *FiniteState
	}
	tests := []struct {
		name string
		f    *FiniteState
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.f.Union(tt.args.other)
		})
	}
}

func TestFiniteState_Loop(t *testing.T) {
	tests := []struct {
		name string
		f    *FiniteState
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.f.Loop()
		})
	}
}

func TestFiniteState_String(t *testing.T) {
	tests := []struct {
		name string
		f    *FiniteState
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.String(); got != tt.want {
				t.Errorf("FiniteState.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFiniteState_Execute(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		f    *FiniteState
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.Execute(tt.args.input); got != tt.want {
				t.Errorf("FiniteState.Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFiniteState_consume(t *testing.T) {
	type args struct {
		ch rune
	}
	tests := []struct {
		name string
		f    *FiniteState
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.consume(tt.args.ch); got != tt.want {
				t.Errorf("FiniteState.consume() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFiniteState_transition(t *testing.T) {
	type args struct {
		from int
		ch   rune
	}
	tests := []struct {
		name string
		f    *FiniteState
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.transition(tt.args.from, tt.args.ch); got != tt.want {
				t.Errorf("FiniteState.transition() = %v, want %v", got, tt.want)
			}
		})
	}
}
