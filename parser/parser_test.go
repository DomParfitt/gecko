package parser

import (
	"reflect"
	"testing"

	"github.com/DomParfitt/gecko/lexer"
	"github.com/DomParfitt/gecko/tree"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want *Parser
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

func TestParser_Parse(t *testing.T) {
	type args struct {
		tokens []lexer.Token
	}
	tests := []struct {
		name    string
		p       *Parser
		args    args
		want    *tree.AbstractSyntax
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.p.Parse(tt.args.tokens)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parser.Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parser.Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParser_consume(t *testing.T) {
	tests := []struct {
		name  string
		p     *Parser
		want  lexer.Token
		want1 bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.p.consume()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parser.consume() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Parser.consume() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestParser_lookBack(t *testing.T) {
	tests := []struct {
		name  string
		p     *Parser
		want  lexer.Token
		want1 bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.p.lookBack()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parser.lookBack() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Parser.lookBack() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestParser_replace(t *testing.T) {
	tests := []struct {
		name string
		p    *Parser
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.p.replace()
		})
	}
}

func TestParser_base(t *testing.T) {
	tests := []struct {
		name string
		p    *Parser
		want bool
	}{
		// TODO: Add test cases.
		{"One", parserFor("a"), true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.base(); got != tt.want {
				t.Errorf("Parser.base() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParser_star(t *testing.T) {
	tests := []struct {
		name string
		p    *Parser
		want bool
	}{
		// TODO: Add test cases.
		{"One", parserFor("a*"), true},
		{"Fail", parserFor("*a"), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.star(); got != tt.want {
				t.Errorf("Parser.star() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParser_basicExpr(t *testing.T) {
	tests := []struct {
		name string
		p    *Parser
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.basicExpr(); got != tt.want {
				t.Errorf("Parser.basicExpr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParser_concatenation(t *testing.T) {
	tests := []struct {
		name string
		p    *Parser
		want bool
	}{
		// TODO: Add test cases.
		{"Simple", parserFor("abc"), true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.concatenation(); got != tt.want {
				t.Errorf("Parser.concatenation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParser_simpleExpr(t *testing.T) {
	tests := []struct {
		name string
		p    *Parser
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.simpleExpr(); got != tt.want {
				t.Errorf("Parser.simpleExpr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParser_union(t *testing.T) {
	tests := []struct {
		name string
		p    *Parser
		want bool
	}{
		// TODO: Add test cases.
		{"Simple", parserFor("a|b"), true},
		{"Simple", parserFor("a|bc*|d"), true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.union(); got != tt.want {
				t.Errorf("Parser.union() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParser_regExpr(t *testing.T) {
	tests := []struct {
		name string
		p    *Parser
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.regExpr(); got != tt.want {
				t.Errorf("Parser.regExpr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func parserFor(input string) *Parser {
	tokens, err := lexer.Tokenize(input)
	if err != nil {
		return nil
	}

	parser := New()
	parser.tokens = tokens

	return parser
}

func TestParser_plus(t *testing.T) {
	tests := []struct {
		name string
		p    *Parser
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.plus(); got != tt.want {
				t.Errorf("Parser.plus() = %v, want %v", got, tt.want)
			}
		})
	}
}
