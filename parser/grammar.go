package parser

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
	star    *Star
	plus    *Plus
	element *Element
}

//Star ::= Element "*"
type Star struct {
	element *Element
}

//Plus ::= Element "+"
type Plus struct {
	element *Element
}

//Element ::= Character
type Element struct {
	Value string
}
