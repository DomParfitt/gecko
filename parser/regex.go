package parser

type Regex struct {
	union  *Union
	simple *SimpleRE
}

type Union struct {
	regex  *Regex
	simple *SimpleRE
}

type SimpleRE struct {
	concatenation *Concatenation
	basic         *BasicRE
}

type Concatenation struct {
	simple *SimpleRE
	basic  *BasicRE
}

type BasicRE struct {
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
