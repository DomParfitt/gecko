package stack

import (
	"github.com/DomParfitt/gecko/lexer"
)

//Stack struct
type Stack struct {
	stack []lexer.Token
}

//New empty stack
func New() *Stack {
	return &Stack{
		stack: []lexer.Token{},
	}
}

//Push a token onto the top of the stack
func (s *Stack) Push(token lexer.Token) {
	s.stack = append(s.stack, token)
}

//Pop the topmost token from the stack
func (s *Stack) Pop() (lexer.Token, bool) {
	if len(s.stack) > 0 {
		token := s.stack[len(s.stack)-1]
		s.stack = append(s.stack[:s.Size()-1])
		return token, true
	}

	return lexer.Error(), false
}

//Size of the stack
func (s *Stack) Size() int {
	return len(s.stack)
}
