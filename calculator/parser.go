package calculator

type node interface{}

type expr interface {
	visit(n node) expr
}

type parser struct {
	tokens  []*token
	current int
}

func (p *parser) parse() []expr {
	return nil
}
