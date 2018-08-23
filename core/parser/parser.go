package parser

import (
	"fmt"
	"github.com/DomParfitt/gecko/core/lexer"
)

//Parser struct
type Parser struct {
	cursor int
	tokens []lexer.Token
}

//New parser
func New() *Parser {
	return &Parser{
		cursor: 0,
		tokens: []lexer.Token{},
	}
}

// Parse a list of tokens into an executable
func (p *Parser) Parse(tokens []lexer.Token) (*RegExpr, error) {
	p.cursor = 0
	p.tokens = tokens

	regExpr, ok := p.regExpr()

	if !ok {
		return nil, fmt.Errorf("the token stream could not be parsed")
	}

	return regExpr, nil
}

// Consume the next token in the list and increment the cursor
func (p *Parser) consume() (lexer.Token, bool) {

	// If cursor is incremented beyond final token then we
	// can't consume any further so return failure
	if p.cursor == len(p.tokens) {
		if len(p.tokens) == 0 {
			return lexer.Token{Type: lexer.Character, Value: ' '}, false
		}
		return p.tokens[p.cursor-1], false
	}

	token := p.tokens[p.cursor]
	p.cursor++

	return token, true
}

//Consume and match against a given token type, resetting if
// not matching
func (p *Parser) consumeAndMatch(expected lexer.Type) bool {

	reset := p.reset()

	token, ok := p.consume()

	if !ok {
		reset()
		return false
	}

	if token.Type != expected {
		reset()
		return false
	}

	return true
}

//Reset the cursor to a given value
func (p *Parser) reset() func() {
	cursor := p.cursor
	return func() {
		p.cursor = cursor
	}

}

func (p *Parser) regExpr() (*RegExpr, bool) {

	reset := p.reset()

	union, ok := p.union()
	if ok {
		return &RegExpr{union: union}, true
	}

	reset()

	simple, ok := p.simpleExpr()
	if ok {
		return &RegExpr{simple: simple}, true
	}

	reset()
	return nil, false
}

func (p *Parser) union() (*Union, bool) {

	reset := p.reset()

	simple, ok := p.simpleExpr()
	if !ok {
		reset()
		return nil, false
	}

	if !p.consumeAndMatch(lexer.Pipe) {
		return nil, false
	}

	regex, ok := p.regExpr()

	if !ok {
		reset()
		return nil, false
	}

	return &Union{regex, simple}, true
}

func (p *Parser) simpleExpr() (*SimpleExpr, bool) {
	reset := p.reset()
	concatenation, ok := p.concatenation()
	if ok {
		return &SimpleExpr{concatenation: concatenation}, true
	}

	reset()
	basic, ok := p.basicExpr()
	if ok {
		return &SimpleExpr{basic: basic}, true
	}

	reset()
	return nil, false
}

func (p *Parser) concatenation() (*Concatenation, bool) {
	reset := p.reset()

	basic, ok := p.basicExpr()

	if !ok {
		reset()
		return nil, false
	}

	simple, ok := p.simpleExpr()

	if !ok {
		reset()
		return nil, false
	}

	return &Concatenation{simple, basic}, true

}

func (p *Parser) basicExpr() (*BasicExpr, bool) {
	reset := p.reset()
	star, ok := p.star()
	if ok {
		return &BasicExpr{star: star}, true
	}

	reset()
	plus, ok := p.plus()
	if ok {
		return &BasicExpr{plus: plus}, true
	}

	reset()
	question, ok := p.question()
	if ok {
		return &BasicExpr{question: question}, true
	}

	reset()
	base, ok := p.element()
	if ok {
		return &BasicExpr{element: base}, true
	}

	return nil, false
}

func (p *Parser) star() (*Star, bool) {
	reset := p.reset()

	base, ok := p.element()
	if !ok {
		reset()
		return nil, false
	}

	if !p.consumeAndMatch(lexer.Star) {
		return nil, false
	}

	return &Star{base}, true
}

func (p *Parser) plus() (*Plus, bool) {
	reset := p.reset()

	base, ok := p.element()
	if !ok {
		reset()
		return nil, false
	}

	if !p.consumeAndMatch(lexer.Plus) {
		return nil, false
	}

	return &Plus{base}, true
}

func (p *Parser) question() (*Question, bool) {
	reset := p.reset()

	base, ok := p.element()
	if !ok {
		reset()
		return nil, false
	}

	if !p.consumeAndMatch(lexer.Question) {
		return nil, false
	}

	return &Question{base}, true
}

func (p *Parser) element() (*Element, bool) {
	reset := p.reset()

	group, ok := p.group()
	if ok {
		return &Element{group: group}, true
	}

	reset()

	set, ok := p.set()
	if ok {
		return &Element{set: set}, true
	}

	reset()

	character, ok := p.character()
	if ok {
		return &Element{character: character}, true
	}

	reset()

	escape, ok := p.escape()
	if ok {
		return &Element{escape: escape}, true
	}

	reset()
	return nil, false
}

func (p *Parser) group() (*Group, bool) {
	reset := p.reset()

	if !p.consumeAndMatch(lexer.OpenParen) {
		return nil, false
	}

	regex, ok := p.regExpr()

	if !ok {
		reset()
		return nil, false
	}

	if !p.consumeAndMatch(lexer.CloseParen) {
		return nil, false
	}

	return &Group{regex}, true
}

func (p *Parser) escape() (*Escape, bool) {
	reset := p.reset()

	if !p.consumeAndMatch(lexer.Escape) {
		return nil, false
	}

	base, ok := p.base()
	if !ok {
		reset()
		return nil, false
	}

	return &Escape{base}, true

}

func (p *Parser) set() (*Set, bool) {
	reset := p.reset()

	positive, ok := p.positiveSet()
	if ok {
		return &Set{positive: positive}, true
	}

	reset()

	negative, ok := p.negativeSet()
	if ok {
		return &Set{negative: negative}, true
	}

	reset()
	return nil, false

}

func (p *Parser) positiveSet() (*PositiveSet, bool) {
	reset := p.reset()

	if !p.consumeAndMatch(lexer.OpenBracket) {
		return nil, false
	}

	setItems, ok := p.setItems()

	if !ok {
		reset()
		return nil, false
	}

	if !p.consumeAndMatch(lexer.CloseBracket) {
		return nil, false
	}

	return &PositiveSet{setItems}, true

}

func (p *Parser) negativeSet() (*NegativeSet, bool) {
	reset := p.reset()

	if !p.consumeAndMatch(lexer.OpenBracket) {
		return nil, false
	}

	if !p.consumeAndMatch(lexer.Caret) {
		return nil, false
	}

	setItems, ok := p.setItems()

	if !ok {
		reset()
		return nil, false
	}

	if !p.consumeAndMatch(lexer.CloseBracket) {
		return nil, false
	}

	return &NegativeSet{setItems}, true
}

func (p *Parser) setItems() (*SetItems, bool) {
	reset := p.reset()

	item, ok := p.setItem()

	if !ok {
		reset()
		return nil, false
	}

	items, ok := p.setItems()

	return &SetItems{item: item, items: items}, true

}

func (p *Parser) setItem() (*SetItem, bool) {
	reset := p.reset()

	rnge, ok := p.rangeExpr()

	if ok {
		return &SetItem{rnge: rnge}, true
	}

	reset()

	character, ok := p.character()

	if ok {
		return &SetItem{character: character}, true
	}

	reset()
	return nil, false

}

func (p *Parser) rangeExpr() (*Range, bool) {
	reset := p.reset()
	start, ok := p.character()

	if !ok {
		reset()
		return nil, false
	}

	if !p.consumeAndMatch(lexer.Dash) {
		return nil, false
	}

	end, ok := p.character()

	if !ok {
		reset()
		return nil, false
	}

	return &Range{start, end}, true

}

func (p *Parser) character() (*Character, bool) {
	reset := p.reset()

	base, ok := p.base()

	if !ok {
		reset()
		return nil, false
	}

	if base.tokenType != lexer.Character {
		reset()
		return nil, false
	}

	return &Character{base}, true
}

func (p *Parser) base() (*Base, bool) {
	reset := p.reset()

	token, ok := p.consume()

	if !ok {
		reset()
		return nil, false
	}

	return &Base{Value: token.Value, tokenType: token.Type}, true
}
