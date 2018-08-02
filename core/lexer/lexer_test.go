package lexer

import (
	"testing"
)

func TestTokenize(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want []Token
	}{
		{"Empty String", args{""}, []Token{}},
		{"Single Token", args{"a"}, []Token{Token{Character, 'a'}}},
		{"Multiple Tokens", args{"a*|b"}, []Token{Token{Character, 'a'}, Token{Star, '*'}, Token{Pipe, '|'}, Token{Character, 'b'}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Tokenize(tt.args.str)
			for i, token := range got {
				if tt.want[i] != token {
					t.Errorf("Tokenize() = %v, want %v", token, tt.want[i])
				}
			}
		})
	}
}
