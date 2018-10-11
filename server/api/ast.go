package api

// AST struct for sending JSON representation of
// the AST over the API
type AST struct {
	Label    string `json:"label"`
	Children []AST  `json:"children"`
}
