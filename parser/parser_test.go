package parser

import (
	"fmt"
	"github.com/DomParfitt/gecko/lexer"
	"testing"
)

func TestParser(t *testing.T) {
	p := New()
	p.tokens = []lexer.Token{lexer.Token{Type: lexer.Letter, Value: 'a'}, lexer.Token{Type: lexer.Closure, Value: '*'}}
	if !p.regExpr() {
		fmt.Errorf("Expected to parse as regex")
	}
}
