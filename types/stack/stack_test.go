package stack

import (
	"reflect"
	"testing"

	"github.com/DomParfitt/gecko/lexer"
)

func TestStack_Push(t *testing.T) {
	type args struct {
		token lexer.Token
	}
	tests := []struct {
		name string
		s    *Stack
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Push(tt.args.token)
		})
	}
}

func TestStack_Pop(t *testing.T) {
	tests := []struct {
		name  string
		s     *Stack
		want  lexer.Token
		want1 bool
		want2 int
	}{
		{name: "No Entries", s: &Stack{stack: []lexer.Token{}}, want: lexer.Token{Type: lexer.Digit, Value: ' '}, want1: false, want2: 0},
		{name: "One Entry", s: &Stack{stack: []lexer.Token{lexer.Token{Type: lexer.Digit, Value: '1'}}}, want: lexer.Token{Type: lexer.Digit, Value: '1'}, want1: true, want2: 0},
		{name: "Two Entries", s: &Stack{stack: []lexer.Token{lexer.Token{Type: lexer.Digit, Value: '1'}, lexer.Token{Type: lexer.Letter, Value: 'a'}}}, want: lexer.Token{Type: lexer.Letter, Value: 'a'}, want1: true, want2: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Pop()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Stack.Pop() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Stack.Pop() got1 = %v, want %v", got1, tt.want1)
			}
			if got := tt.s.Size(); got != tt.want2 {
				t.Errorf("Stack.Size() after Pop got = %v, want %v", got, tt.want2)
			}
		})
	}
}

func TestStack_Size(t *testing.T) {
	tests := []struct {
		name string
		s    *Stack
		want int
	}{
		{name: "No Entries", s: &Stack{stack: []lexer.Token{}}, want: 0},
		{name: "One Entry", s: &Stack{stack: []lexer.Token{lexer.Token{Type: lexer.Digit, Value: '1'}}}, want: 1},
		{name: "Two Entries", s: &Stack{stack: []lexer.Token{lexer.Token{Type: lexer.Digit, Value: '1'}, lexer.Token{Type: lexer.Letter, Value: 'a'}}}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Size(); got != tt.want {
				t.Errorf("Stack.Size() = %v, want %v", got, tt.want)
			}
		})
	}
}
