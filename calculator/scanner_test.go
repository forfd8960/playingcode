package calculator

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScanner(t *testing.T) {
	t.Run("success get tokens", testSuccessGetTokens)
	t.Run("success get numbers", testGetNumberToken)
	t.Run("unsupport token", testUnsupportToken)
	t.Run("get '**' token", getDoubleStarToken)
}

func testSuccessGetTokens(t *testing.T) {
	expression := "(1 + 2) * (3 / 4) + 11.99"
	sc := newScanner(expression)
	err := sc.scanExpression()
	assert.Nil(t, err)

	expectTokens := []*token{
		{TkType: leftParent, Lexeme: "(", Literal: nil},
		{TkType: number, Lexeme: "1", Literal: float64(1)},
		{TkType: plus, Lexeme: "+", Literal: nil},
		{TkType: number, Lexeme: "2", Literal: float64(2)},
		{TkType: rightParent, Lexeme: ")", Literal: nil},
		{TkType: star, Lexeme: "*", Literal: nil},
		{TkType: leftParent, Lexeme: "(", Literal: nil},
		{TkType: number, Lexeme: "3", Literal: float64(3)},
		{TkType: slash, Lexeme: "/", Literal: nil},
		{TkType: number, Lexeme: "4", Literal: float64(4)},
		{TkType: rightParent, Lexeme: ")", Literal: nil},
		{TkType: plus, Lexeme: "+", Literal: nil},
		{TkType: number, Lexeme: "11.99", Literal: float64(11.99)},
		{TkType: eof, Lexeme: "", Literal: nil},
	}
	assert.Equal(t, expectTokens, sc.tokens)
}

func testGetNumberToken(t *testing.T) {
	expression := "66.909"
	sc := newScanner(expression)
	err := sc.scanExpression()
	assert.Nil(t, err)

	t.Logf("tokens: %s\n", sc)
	expectTokens := []*token{
		{TkType: number, Lexeme: "66.909", Literal: float64(66.909)},
		{TkType: eof, Lexeme: "", Literal: nil},
	}
	assert.Equal(t, expectTokens, sc.tokens)
}

func testUnsupportToken(t *testing.T) {
	expression := "(1 + 2) % ( 3 * 4)"
	sc := newScanner(expression)
	err := sc.scanExpression()
	if assert.NotNil(t, err) {
		assert.EqualError(t, errors.New("unsupported token: %"), err.Error())
	}
}

func getDoubleStarToken(t *testing.T) {
	expression := "2 ** 10"
	sc := newScanner(expression)
	err := sc.scanExpression()
	assert.Nil(t, err)

	t.Logf("tokens: %s\n", sc)
	expectTokens := []*token{
		{TkType: number, Lexeme: "2", Literal: float64(2)},
		{TkType: doubleStar, Lexeme: "**", Literal: nil},
		{TkType: number, Lexeme: "10", Literal: float64(10)},
		{TkType: eof},
	}
	assert.Equal(t, expectTokens, sc.tokens)
}
