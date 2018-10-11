package api

// MatchResponse struct representing the response object
// returned by the Match gateway
type MatchResponse struct {
	Pattern string `json:"pattern"`
	Input   string `json:"input"`
	Result  bool   `json:"result"`
	Steps   []int  `json:"steps"`
}
