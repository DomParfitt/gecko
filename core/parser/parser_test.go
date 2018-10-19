//Package parser contains the implementation of Gecko's parser, data structures to
//represent the grammar as nodes of an AST and functionality to compile each of the
//AST nodes into an executable automata

package parser

import (
	"github.com/DomParfitt/gecko/core/lexer"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want *Parser
	}{
		{"New Parser", &Parser{cursor: 0, tokens: []lexer.Token{}}},
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
		wantErr bool
	}{
		{"Valid regex", New(), args{tokens: tokenize("ab")}, false},
		{"Partial match", New(), args{tokens: tokenize("ab[")}, true},
		{"Mismatched bracket", New(), args{tokens: tokenize("[")}, true},
		{"No tokens", New(), args{tokens: tokenize("")}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//Just want to check error condition, actual RegExpr structure checked by individual method tests
			_, err := tt.p.Parse(tt.args.tokens)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parser.Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
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
		{"Valid RegExpr from Simple", parserFrom("a"), regExprFromSimple('a'), true},
		{"Valid RegExpr from Union", parserFrom("a|b"), regExprFromUnion('a', 'b'), true},
		{"No remaining tokens", parserFrom(""), nil, false},
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

func TestParser_union(t *testing.T) {
	tests := []struct {
		name  string
		p     *Parser
		want  *Union
		want1 bool
	}{
		{"Valid Union", parserFrom("a|b"), union('a', 'b'), true},
		{"No first character", parserFrom("|b"), nil, false},
		{"No second character", parserFrom("a|"), nil, false},
		{"No pipe", parserFrom("ab"), nil, false},
		{"No remaining tokens", parserFrom(""), nil, false},
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

func TestParser_simpleExpr(t *testing.T) {
	tests := []struct {
		name  string
		p     *Parser
		want  *SimpleExpr
		want1 bool
	}{
		{"Valid SimpleExpr from BasicExpr", parserFrom("a"), simpleFromBasic('a'), true},
		{"Valid SimpleExpr from Concatenation", parserFrom("ab"), simpleFromConcat('a', 'b'), true},
		{"No remaining tokens", parserFrom(""), nil, false},
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

func TestParser_concatenation(t *testing.T) {
	tests := []struct {
		name  string
		p     *Parser
		want  *Concatenation
		want1 bool
	}{
		{"Valid Concatenation", parserFrom("ab"), &Concatenation{basic: basicFromCharElement('a'), simple: simpleFromBasic('b')}, true},
		{"No concatentation", parserFrom("a"), nil, false},
		{"No remaining tokens", parserFrom(""), nil, false},
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

func TestParser_basicExpr(t *testing.T) {
	tests := []struct {
		name  string
		p     *Parser
		want  *BasicExpr
		want1 bool
	}{
		{"Valid BasicExpr from Star", parserFrom("a*"), basicFromStar('a'), true},
		{"Valid BasicExpr from Plus", parserFrom("a+"), basicFromPlus('a'), true},
		{"Valid BasicExpr from Question", parserFrom("a?"), basicFromQuestion('a'), true},
		{"Valid BasicExpr from Element", parserFrom("a"), basicFromCharElement('a'), true},
		{"No remaining tokens", parserFrom(""), nil, false},
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

func TestParser_star(t *testing.T) {
	tests := []struct {
		name  string
		p     *Parser
		want  *Star
		want1 bool
	}{
		{"Valid Star", parserFrom("a*"), star('a'), true},
		{"No star symbol", parserFrom("a"), nil, false},
		{"No Element", parserFrom("*"), nil, false},
		{"No remaining tokens", parserFrom(""), nil, false},
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
		{"Valid Plus", parserFrom("a+"), plus('a'), true},
		{"No plus symbol", parserFrom("a"), nil, false},
		{"No Element", parserFrom("+"), nil, false},
		{"No remaining tokens", parserFrom(""), nil, false},
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

func TestParser_question(t *testing.T) {
	tests := []struct {
		name  string
		p     *Parser
		want  *Question
		want1 bool
	}{
		{"Valid Question", parserFrom("a?"), question('a'), true},
		{"No question mark", parserFrom("a"), nil, false},
		{"No Element", parserFrom("?"), nil, false},
		{"No remaining tokens", parserFrom(""), nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.p.question()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parser.question() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Parser.question() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestParser_element(t *testing.T) {
	tests := []struct {
		name  string
		p     *Parser
		want  *Element
		want1 bool
	}{
		// {"Valid Element from Group", parserFrom("(a)"), nil, true},
		{"Valid Element from Set", parserFrom("[a-z]"), elementFromSet('a', 'z'), true},
		{"Valid Element from Character", parserFrom("a"), elementFromChar('a'), true},
		{"Valid Element from Escape", parserFrom("\\*"), elementFromEscape(lexer.Star, '*'), true},
		{"No remaining tokens", parserFrom(""), nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.p.element()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parser.element() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Parser.element() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestParser_group(t *testing.T) {
	tests := []struct {
		name  string
		p     *Parser
		want  *Group
		want1 bool
	}{
		{"Valid Group", parserFrom("(a)"), group('a'), true},
		{"Missing open paren", parserFrom("a)"), nil, false},
		{"Missing closing paren", parserFrom("(a"), nil, false},
		{"Missing value", parserFrom("()"), nil, false},
		{"No remaining tokens", parserFrom(""), nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.p.group()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parser.group() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Parser.group() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestParser_escape(t *testing.T) {
	tests := []struct {
		name  string
		p     *Parser
		want  *Escape
		want1 bool
	}{
		{"Valid Escape", parserFrom("\\*"), escape(lexer.Star, '*'), true},
		{"Missing backslash", parserFrom("*"), nil, false},
		{"Missing escaped token", parserFrom("\\"), nil, false},
		{"No remaining tokens", parserFrom(""), nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.p.escape()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parser.escape() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Parser.escape() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestParser_set(t *testing.T) {
	tests := []struct {
		name  string
		p     *Parser
		want  *Set
		want1 bool
	}{
		{"Valid PositiveSet Range", parserFrom("[a-z]"), setFromPositiveRange('a', 'z'), true},
		{"Valid PositiveSet Character", parserFrom("[a]"), setFromPositiveChar('a'), true},
		{"Valid NegativeSet Range", parserFrom("[^a-z]"), setFromNegativeRange('a', 'z'), true},
		{"Valid NegativeSet Character", parserFrom("[^a]"), setFromNegativeChar('a'), true},
		{"Missing opening bracket", parserFrom("a-z]"), nil, false},
		{"Missing closing bracket", parserFrom("[a-z"), nil, false},
		{"Missing values", parserFrom("[]"), nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.p.set()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parser.set() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Parser.set() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestParser_positiveSet(t *testing.T) {
	tests := []struct {
		name  string
		p     *Parser
		want  *PositiveSet
		want1 bool
	}{
		{"Valid PositiveSet Range", parserFrom("[a-z]"), positiveSetFromRange('a', 'z'), true},
		{"Valid PositiveSet Character", parserFrom("[a]"), positiveSetFromChar('a'), true},
		{"Missing opening bracket", parserFrom("a-z]"), nil, false},
		{"Missing closing bracket", parserFrom("[a-z"), nil, false},
		{"Missing values", parserFrom("[]"), nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.p.positiveSet()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parser.positiveSet() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Parser.positiveSet() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestParser_negativeSet(t *testing.T) {
	tests := []struct {
		name  string
		p     *Parser
		want  *NegativeSet
		want1 bool
	}{
		{"Valid NegativeSet Range", parserFrom("[^a-z]"), negativeSetFromRange('a', 'z'), true},
		{"Valid NegativeSet Character", parserFrom("[^a]"), negativeSetFromChar('a'), true},
		{"Missing opening bracket", parserFrom("^a-z]"), nil, false},
		{"Missing closing bracket", parserFrom("[^a-z"), nil, false},
		{"Missing caret", parserFrom("[a-z]"), nil, false},
		{"Missing values", parserFrom("[^]"), nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.p.negativeSet()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parser.negativeSet() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Parser.negativeSet() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestParser_setItems(t *testing.T) {
	tests := []struct {
		name  string
		p     *Parser
		want  *SetItems
		want1 bool
	}{
		{"Valid SetItems with single SetItem", parserFrom("a-z"), setItemsRange('a', 'z'), true},
		{"Valid SetItems with multiple SetItem", parserFrom("a-zA"), &SetItems{item: setItemRange('a', 'z'), items: &SetItems{item: setItemCharacter('A')}}, true},
		{"Invalid token", parserFrom("*"), nil, false},
		{"No remaining tokens", parserFrom(""), nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.p.setItems()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parser.setItems() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Parser.setItems() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestParser_setItem(t *testing.T) {
	tests := []struct {
		name  string
		p     *Parser
		want  *SetItem
		want1 bool
	}{
		{"Valid SetItem with Range", parserFrom("a-z"), setItemRange('a', 'z'), true},
		{"Valid SetItem with Character", parserFrom("a"), setItemCharacter('a'), true},
		{"No remaining tokens", parserFrom(""), nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.p.setItem()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parser.setItem() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Parser.setItem() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestParser_rangeExpr(t *testing.T) {
	tests := []struct {
		name  string
		p     *Parser
		want  *Range
		want1 bool
	}{
		{"Valid Range", parserFrom("a-z"), rangeExpr('a', 'z'), true},
		{"First token not character", parserFrom("*-z"), nil, false},
		{"Final token not character", parserFrom("a-*"), nil, false},
		{"No dash", parserFrom("az"), nil, false},
		{"No remaining tokens", parserFrom(""), nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.p.rangeExpr()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parser.rangeExpr() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Parser.rangeExpr() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestParser_character(t *testing.T) {
	tests := []struct {
		name  string
		p     *Parser
		want  *Character
		want1 bool
	}{
		{"Valid Character", parserFrom("a"), character('a'), true},
		{"Non-character token", parserFrom("*"), nil, false},
		{"No remaining tokens", parserFrom(""), nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.p.character()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parser.character() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Parser.character() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestParser_base(t *testing.T) {
	tests := []struct {
		name  string
		p     *Parser
		want  *Base
		want1 bool
	}{
		{"Valid Base", parserFrom("a"), &Base{tokenType: lexer.Character, Value: 'a'}, true},
		{"No remaining tokens", parserFrom(""), nil, false},
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

func TestParser_consumeAndMatch(t *testing.T) {
	type args struct {
		expected lexer.Type
	}
	tests := []struct {
		name string
		p    *Parser
		args args
		want bool
	}{
		{"Token matches expected", parserFrom("a"), args{lexer.Character}, true},
		{"Token doesn't match expected", parserFrom("*"), args{lexer.Character}, false},
		{"No more tokens", parserFrom(""), args{lexer.Character}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.consumeAndMatch(tt.args.expected); got != tt.want {
				t.Errorf("Parser.consumeAndMatch() = %v, want %v", got, tt.want)
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

// HELPER FUNCTIONS

func tokenize(input string) []lexer.Token {
	return lexer.Tokenize(input)
}

func parserFrom(input string) *Parser {
	p := New()
	p.tokens = tokenize(input)
	return p
}

func regExprFromSimple(char rune) *RegExpr {
	return &RegExpr{simple: simpleFromBasic(char)}
}

func regExprFromUnion(first, second rune) *RegExpr {
	return &RegExpr{union: union(first, second)}
}

func union(first, second rune) *Union {
	return &Union{simple: simpleFromBasic(first), regex: regExprFromSimple(second)}
}

func concatenation(first, second rune) *Concatenation {
	return &Concatenation{basic: basicFromCharElement(first), simple: simpleFromBasic(second)}
}

func simpleFromBasic(char rune) *SimpleExpr {
	return &SimpleExpr{basic: basicFromCharElement(char)}
}

func simpleFromConcat(first, second rune) *SimpleExpr {
	return &SimpleExpr{concatenation: concatenation(first, second)}
}

func basicFromStar(char rune) *BasicExpr {
	return &BasicExpr{star: star(char)}
}

func basicFromPlus(char rune) *BasicExpr {
	return &BasicExpr{plus: plus(char)}
}

func basicFromQuestion(char rune) *BasicExpr {
	return &BasicExpr{question: question(char)}
}

func basicFromCharElement(char rune) *BasicExpr {
	return &BasicExpr{element: elementFromChar(char)}
}

func plus(char rune) *Plus {
	return &Plus{element: elementFromChar(char)}
}

func star(char rune) *Star {
	return &Star{element: elementFromChar(char)}
}

func question(char rune) *Question {
	return &Question{element: elementFromChar(char)}
}

func elementFromSet(first, last rune) *Element {
	return &Element{set: setFromPositiveRange(first, last)}
}

func elementFromChar(char rune) *Element {
	return &Element{character: character(char)}
}

func elementFromEscape(token lexer.Type, char rune) *Element {
	return &Element{escape: escape(token, char)}
}

func escape(token lexer.Type, char rune) *Escape {
	return &Escape{base: &Base{tokenType: token, Value: char}}
}

func group(char rune) *Group {
	return &Group{regExpr: regExprFromSimple(char)}
}

func setFromPositiveRange(first, last rune) *Set {
	return &Set{positive: positiveSetFromRange(first, last)}
}

func setFromPositiveChar(char rune) *Set {
	return &Set{positive: positiveSetFromChar(char)}
}

func setFromNegativeRange(first, last rune) *Set {
	return &Set{negative: negativeSetFromRange(first, last)}
}

func setFromNegativeChar(char rune) *Set {
	return &Set{negative: negativeSetFromChar(char)}
}

func positiveSetFromRange(first, last rune) *PositiveSet {
	return &PositiveSet{items: setItemsRange(first, last)}
}

func positiveSetFromChar(char rune) *PositiveSet {
	return &PositiveSet{items: setItemsCharacter(char)}
}

func negativeSetFromRange(first, last rune) *NegativeSet {
	return &NegativeSet{items: setItemsRange(first, last)}
}

func negativeSetFromChar(char rune) *NegativeSet {
	return &NegativeSet{items: setItemsCharacter(char)}
}

func setItemsCharacter(char rune) *SetItems {
	return &SetItems{item: setItemCharacter(char)}
}

func setItemsRange(first, last rune) *SetItems {
	return &SetItems{item: setItemRange(first, last)}
}

func setItemCharacter(char rune) *SetItem {
	return &SetItem{character: character(char)}
}

func setItemRange(first, last rune) *SetItem {
	return &SetItem{rnge: rangeExpr(first, last)}
}

// RangeExpr returns a valid Range struct with given first and
// last values for testing convenience
func rangeExpr(first, last rune) *Range {
	return &Range{start: character(first), end: character(last)}
}

// Character returns a valid Character struct with the given value
// for testing convenience
func character(char rune) *Character {
	return &Character{base: &Base{tokenType: lexer.Character, Value: char}}
}
