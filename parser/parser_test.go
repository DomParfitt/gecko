package parser

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/DomParfitt/gecko/lexer"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want *Parser
	}{
		{"New", parserFrom("")},
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
		want    *RegExpr
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

func TestParser_reset(t *testing.T) {
	tests := []struct {
		name string
		p    *Parser
		want func()
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.reset(); !reflect.DeepEqual(got, tt.want) {
				// t.Errorf("Parser.reset() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParser_base(t *testing.T) {
	tests := []struct {
		name  string
		p     *Parser
		want  *Element
		want1 bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.p.base()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parser.base() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Parser.base() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestParser_star(t *testing.T) {
	tests := []struct {
		name  string
		p     *Parser
		want  *Star
		want1 bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.p.star()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parser.star() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Parser.star() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestParser_plus(t *testing.T) {
	tests := []struct {
		name  string
		p     *Parser
		want  *Plus
		want1 bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.p.plus()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parser.plus() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Parser.plus() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestParser_basicExpr(t *testing.T) {
	tests := []struct {
		name  string
		p     *Parser
		want  *BasicExpr
		want1 bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.p.basicExpr()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parser.basicExpr() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Parser.basicExpr() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestParser_concatenation(t *testing.T) {
	tests := []struct {
		name  string
		p     *Parser
		want  *Concatenation
		want1 bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.p.concatenation()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parser.concatenation() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Parser.concatenation() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestParser_simpleExpr(t *testing.T) {
	tests := []struct {
		name  string
		p     *Parser
		want  *SimpleExpr
		want1 bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.p.simpleExpr()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parser.simpleExpr() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Parser.simpleExpr() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestParser_union(t *testing.T) {
	tests := []struct {
		name  string
		p     *Parser
		want  *Union
		want1 bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.p.union()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parser.union() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Parser.union() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestParser_regExpr(t *testing.T) {
	tests := []struct {
		name  string
		p     *Parser
		want  *RegExpr
		want1 bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.p.regExpr()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parser.regExpr() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Parser.regExpr() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test(t *testing.T) {
	txt := "a|b"
	tokens := lexer.Tokenize(txt)

	for _, token := range tokens {
		fmt.Println(token)
	}
	parser := New()
	tree, err := parser.Parse(tokens)

	if err != nil {
		fmt.Println(err)
		return
	}

	exec := tree.Compile()
	fmt.Println(exec)
	result := exec.Execute("a")
	if !result {
		t.Errorf("failed compiling and executing")
	}
	fmt.Printf("%t", result)
}

func parserFrom(input string) *Parser {
	tokens := lexer.Tokenize(input)
	p := New()
	p.tokens = tokens
	return p
}
