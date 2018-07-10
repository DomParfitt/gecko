package parser

type RegExpr struct {
	union  *UnionExpr
	simple *SimpleExpr
}

type UnionExpr struct {
	regex  *RegExpr
	simple *SimpleExpr
}

type SimpleExpr struct {
	concatenation *ConcatenationExpr
	basic         *BasicExpr
}

type ConcatenationExpr struct {
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
