package parser

type RegExpr struct {
	union  *Union
	simple *SimpleExpr
}

type Union struct {
	regex  *RegExpr
	simple *SimpleExpr
}

type SimpleExpr struct {
	concatenation *Concatenation
	basic         *BasicExpr
}

type Concatenation struct {
	simple *SimpleExpr
	basic  *BasicExpr
}

type BasicExpr struct {
	star    *Star
	plus    *Plus
	element *Element
}

type Star struct {
	element *Element
}

type Plus struct {
	element *Element
}

type Element struct {
	Value string
}
