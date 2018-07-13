package lexer

import (
	"reflect"
	"testing"
)

func TestToken_String(t *testing.T) {
	tests := []struct {
		name string
		t    Token
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.t.String(); got != tt.want {
				t.Errorf("Token.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatch(t *testing.T) {
	type args struct {
		ch rune
	}
	tests := []struct {
		name string
		args args
		want Type
	}{
		{"Match Open Brace", args{'('}, OpenBrace},
		{"Match Open Brace", args{'['}, OpenBrace},
		{"Match Open Brace", args{'{'}, OpenBrace},
		{"Match Close Brace", args{')'}, CloseBrace},
		{"Match Close Brace", args{']'}, CloseBrace},
		{"Match Close Brace", args{'}'}, CloseBrace},
		{"Match Star", args{'*'}, Star},
		{"Match Plus", args{'+'}, Plus},
		{"Match Caret", args{'^'}, Caret},
		{"Match Escape", args{'\\'}, Escape},
		{"Match Pipe", args{'|'}, Pipe},
		{"Match Letter", args{'a'}, Character},
		{"Match Digit", args{'1'}, Character},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Match(tt.args.ch); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Match() = %v, want %v", got, tt.want)
			}
		})
	}
}
