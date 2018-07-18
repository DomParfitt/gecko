package parser

import (
	"reflect"
	"testing"

	"github.com/DomParfitt/gecko/automata"
)

func TestElement_Compile(t *testing.T) {
	tests := []struct {
		name string
		e    *Element
		want *automata.FiniteState
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Compile(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Element.Compile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPlus_Compile(t *testing.T) {
	tests := []struct {
		name string
		p    *Plus
		want *automata.FiniteState
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Compile(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Plus.Compile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStar_Compile(t *testing.T) {
	tests := []struct {
		name string
		s    *Star
		want *automata.FiniteState
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Compile(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Star.Compile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBasicExpr_Compile(t *testing.T) {
	tests := []struct {
		name string
		b    *BasicExpr
		want *automata.FiniteState
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.Compile(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BasicExpr.Compile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConcatenation_Compile(t *testing.T) {
	tests := []struct {
		name string
		c    *Concatenation
		want *automata.FiniteState
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Compile(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Concatenation.Compile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSimpleExpr_Compile(t *testing.T) {
	tests := []struct {
		name string
		s    *SimpleExpr
		want *automata.FiniteState
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Compile(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SimpleExpr.Compile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnion_Compile(t *testing.T) {
	tests := []struct {
		name string
		u    *Union
		want *automata.FiniteState
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.u.Compile(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Union.Compile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRegExpr_Compile(t *testing.T) {
	tests := []struct {
		name string
		r    *RegExpr
		want *automata.FiniteState
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.Compile(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RegExpr.Compile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_compile(t *testing.T) {
	type args struct {
		ch         chan<- *automata.FiniteState
		compilable Compiler
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			compile(tt.args.ch, tt.args.compilable)
		})
	}
}
