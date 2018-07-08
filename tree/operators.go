package tree

//go:generate stringer -type=Operator

type Operator int

const (
	Concatenation Operator = iota
	Union
	Star
	Plus
)
