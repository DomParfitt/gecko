package parser

import (
	"github.com/DomParfitt/gecko/server/api"
)

type Transformer interface {
	transform() api.AST
}
