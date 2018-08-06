package core

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want *Compiler
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

func TestCompiler_Compile(t *testing.T) {
	type args struct {
		pattern string
	}
	tests := []struct {
		name    string
		c       *Compiler
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.Compile(tt.args.pattern); (err != nil) != tt.wantErr {
				t.Errorf("Compiler.Compile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCompiler_Match(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		c       *Compiler
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Match(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Compiler.Match() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Compiler.Match() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCompiler_MatchPattern(t *testing.T) {
	type args struct {
		pattern string
		input   string
	}
	tests := []struct {
		name    string
		c       *Compiler
		args    args
		want    bool
		wantErr bool
	}{
		{"", New(), args{pattern: "a|b", input: "a"}, true, false},
		{"", New(), args{pattern: "a|b", input: "b"}, true, false},
		{"", New(), args{pattern: "a|b", input: "ab"}, false, false},
		{"", New(), args{pattern: "a|b*", input: "a"}, true, false},
		{"", New(), args{pattern: "a|b*", input: ""}, true, false},
		{"", New(), args{pattern: "a|b*", input: "b"}, true, false},
		{"", New(), args{pattern: "a|b*", input: "bb"}, true, false},
		{"", New(), args{pattern: "a|b*", input: "ab"}, false, false},
		{"", New(), args{pattern: "a|ab", input: "ab"}, true, false},
		{"", New(), args{pattern: "a|ab", input: "a"}, true, false},
		{"", New(), args{pattern: "a|ab", input: "aab"}, false, false},
		{"", New(), args{pattern: "a|ab", input: ""}, false, false},
		{"", New(), args{pattern: "ab|a", input: "ab"}, true, false},
		{"", New(), args{pattern: "ab|a", input: "a"}, true, false},
		{"", New(), args{pattern: "ab|a", input: "aab"}, false, false},
		{"", New(), args{pattern: "ab|a", input: ""}, false, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.MatchPattern(tt.args.pattern, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Compiler.MatchPattern() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Compiler.MatchPattern() = %v, want %v", got, tt.want)
			}
		})
	}
}
