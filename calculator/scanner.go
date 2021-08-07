package calculator

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

var (
	errUnsupportToken = func(r rune) error {
		return fmt.Errorf("unsupported token: %s", string(r))
	}
)

type tokenType int

const (
	leftParent  tokenType = iota // (
	rightParent                  // )
	minus                        // -
	plus                         // +
	slash                        // "/"
	star                         // '*'
	doubleStar                   // x ** y -> math.pow
	number                       // 123, 9, 10
	eof                          // end token type
)

type token struct {
	TkType  tokenType
	Lexeme  string
	Literal interface{}
}

type scanner struct {
	exp     string
	runes   []rune
	tokens  []*token
	start   int
	current int
}

func newScanner(exp string) *scanner {
	return &scanner{
		exp:    exp,
		runes:  []rune(exp),
		tokens: make([]*token, 0, 10),
	}
}

func (s *scanner) String() string {
	sb := &strings.Builder{}
	for _, t := range s.tokens {
		sb.WriteString(fmt.Sprintf("type: %d, lexeme: %s\n", t.TkType, t.Lexeme))
	}
	return sb.String()
}

// scanExpression read the expression string and split it into tokens
func (s *scanner) scanExpression() error {
	for !s.isEnd() {
		s.start = s.current
		if err := s.scanToken(); err != nil {
			return err
		}
	}
	s.tokens = append(s.tokens, &token{
		TkType:  eof,
		Lexeme:  "",
		Literal: nil,
	})
	return nil
}

func (s *scanner) scanToken() error {
	r := s.advance()
	switch r {
	case '(':
		s.addToken(leftParent, nil)
	case ')':
		s.addToken(rightParent, nil)
	case '+':
		s.addToken(plus, nil)
	case '-':
		s.addToken(minus, nil)
	case '*':
		if s.match('*') { // check '**' token
			s.addToken(doubleStar, nil)
		} else {
			s.addToken(star, nil)
		}
	case '/':
		s.addToken(slash, nil)
	case ' ', '\r', '\t', '\n':
	default:
		if s.isDigit(r) {
			s.getNumber()
		} else {
			return errUnsupportToken(r)
		}
	}
	return nil
}

func (s *scanner) isEnd() bool {
	return s.current >= len(s.runes)
}

func (s *scanner) advance() rune {
	s.current++
	return s.runes[s.current-1]
}

// match check if current rune is equals r
// if euqals then current move to next rune
func (s *scanner) match(r rune) bool {
	if s.isEnd() {
		return false
	}

	if s.runes[s.current] != r {
		return false
	}

	s.current++
	return true
}

func (s *scanner) addToken(tokenType tokenType, literal interface{}) {
	lexeme := string(s.runes[s.start:s.current])
	s.tokens = append(s.tokens, &token{
		TkType:  tokenType,
		Lexeme:  lexeme,
		Literal: literal,
	})
}

func (s *scanner) getNumber() {
	for s.isDigit(s.peek()) {
		s.advance()
	}

	// float number
	if s.peek() == '.' && s.isDigit(s.peekNext()) {
		s.advance()

		for s.isDigit(s.peek()) {
			s.advance()
		}
	}

	num := string(s.runes[s.start:s.current])
	val, _ := strconv.ParseFloat(num, 64)
	s.addToken(number, val)
}

func (s *scanner) isDigit(r rune) bool {
	return unicode.IsDigit(r)
}

func (s *scanner) peek() rune {
	if s.current >= len(s.runes) {
		return '\000'
	}
	return s.runes[s.current]
}

func (s *scanner) peekNext() rune {
	if s.current+1 >= len(s.runes) {
		return '\000'
	}

	return s.runes[s.current+1]
}
