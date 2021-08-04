package calculator

type token struct {
}

type scanner struct {
	exp     string
	runes   []rune
	tokens  []*token
	start   int
	current int
}

func (s *scanner) scanExpression() []*token {
	return nil
}
