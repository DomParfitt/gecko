package parser

import (
	"github.com/DomParfitt/gecko/core/lexer"
)

//RegExpr ::= Union | SimpleExpr
type RegExpr struct {
	union  *Union
	simple *SimpleExpr
}

//Union ::= RegExpr "|" SimpleExpr
type Union struct {
	regex  *RegExpr
	simple *SimpleExpr
}

//SimpleExpr ::= Concatenation | BasicExpr
type SimpleExpr struct {
	concatenation *Concatenation
	basic         *BasicExpr
}

//Concatenation ::= SimpleExpr BasicExpr
type Concatenation struct {
	simple *SimpleExpr
	basic  *BasicExpr
}

//BasicExpr ::= Star | Plus | Element
type BasicExpr struct {
	star     *Star
	plus     *Plus
	question *Question
	element  *Element
}

//Star ::= Element "*"
type Star struct {
	element *Element
}

//Plus ::= Element "+"
type Plus struct {
	element *Element
}

//Question ::= Element "?"
type Question struct {
	element *Element
}

//Repetition ::= Element "{x [,y]}"
type Repetition struct {
	min, max int
	element  *Element
}

//Element ::= Character | Group | Set
type Element struct {
	// Value     rune
	character *Character
	group     *Group
	set       *Set
	escape    *Escape
}

//Group ::= (RegExpr)
type Group struct {
	regExpr *RegExpr
}

//Escape ::= "\" Character
type Escape struct {
	// character *Character
	base *Base
}

//Set ::= PositiveSet | NegativeSet
type Set struct {
	positive *PositiveSet
	negative *NegativeSet
}

//PositiveSet ::= "[" SetItems "]"
type PositiveSet struct {
	items *SetItems
}

//NegativeSet ::= "[" "^" SetItems "]"
type NegativeSet struct {
	items *SetItems
}

//SetItems ::= SetItem SetItems
type SetItems struct {
	item  *SetItem
	items *SetItems
}

//SetItem ::= Range | Character
type SetItem struct {
	rnge      *Range
	character *Character
}

//Range ::= Character "-" Character
type Range struct {
	start *Character
	end   *Character
}

//Character ::= literal character
type Character struct {
	// Value rune
	base *Base
}

//Base ::= <literal character>
type Base struct {
	tokenType lexer.Type
	Value     rune
}
