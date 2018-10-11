package api

// PatternResponse represents the Response JSON
// returned by calls to the Pattern API
type PatternResponse struct {
	AST      AST      `json:"ast"`
	Automata Automata `json:"automata"`
}
