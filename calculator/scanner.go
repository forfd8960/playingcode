package calculator

import (
	"strconv"
	"unicode"
)

type tokenType int

const (
	leftParent  tokenType = iota // (
	rightParent                  // )
	minus                        // -
	plus                         // +
	slash                        // "/"
	star                         // '*'
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

// scanExpression read the expression string and split it into tokens
func (s *scanner) scanExpression() []*token {
	return nil
}

func (s *scanner) isEnd() bool {
	return s.current >= len(s.runes)
}

func (s *scanner) advance() rune {
	s.current++
	return s.runes[s.current-1]
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
