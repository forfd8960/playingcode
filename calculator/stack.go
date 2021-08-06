package calculator

// stack used to put operator and operand
type stack struct {
	tokens  []*token
	length  int
	current int
}

func (s *stack) push(t *token) {
	s.tokens = append(s.tokens, t)
	s.length++
}

func (s *stack) pop() *token {
	if s.isEmpty() {
		return nil
	}

	t := s.tokens[s.length-1]
	s.tokens = s.tokens[:s.length-1]
	s.length--
	return t
}

func (s *stack) peek() *token {
	if s.isEmpty() {
		return nil
	}
	return s.tokens[s.length-1]
}

func (s *stack) isEmpty() bool {
	return s.length == 0
}
