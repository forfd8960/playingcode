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
		{TkType: leftParent, Text: "(", Value: nil},
		{TkType: number, Text: "1", Value: float64(1)},
		{TkType: plus, Text: "+", Value: nil},
		{TkType: number, Text: "2", Value: float64(2)},
		{TkType: rightParent, Text: ")", Value: nil},
		{TkType: star, Text: "*", Value: nil},
		{TkType: leftParent, Text: "(", Value: nil},
		{TkType: number, Text: "3", Value: float64(3)},
		{TkType: slash, Text: "/", Value: nil},
		{TkType: number, Text: "4", Value: float64(4)},
		{TkType: rightParent, Text: ")", Value: nil},
		{TkType: plus, Text: "+", Value: nil},
		{TkType: number, Text: "11.99", Value: float64(11.99)},
		{TkType: eof, Text: "", Value: nil},
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
		{TkType: number, Text: "66.909", Value: float64(66.909)},
		{TkType: eof, Text: "", Value: nil},
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
		{TkType: number, Text: "2", Value: float64(2)},
		{TkType: doubleStar, Text: "**", Value: nil},
		{TkType: number, Text: "10", Value: float64(10)},
		{TkType: eof},
	}
	assert.Equal(t, expectTokens, sc.tokens)
}
