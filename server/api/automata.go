package api

// Automata representation for JSON
type Automata struct {
	CurrentNode int    `json:"currentNode"`
	Nodes       []Node `json:"nodes"`
	Edges       []Edge `json:"edges"`
}

// Edge representation for JSON Automata
type Edge struct {
	From  int    `json:"from"`
	To    int    `json:"to"`
	Label string `json:"label"`
}

// Node representation for JSON Automata
type Node struct {
	ID         int  `json:"id"`
	IsTerminal bool `json:"isTerminal"`
}
