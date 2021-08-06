package calculator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	t.Run("pop and push", testPop)
	t.Run("peek", testPeek)
	t.Run("isEmpty", testIsEmpty)
}

func testPop(t *testing.T) {
	s := &stack{}
	s.push(&token{TkType: number, Lexeme: "1", Literal: float64(1)})
	s.push(&token{TkType: plus, Lexeme: "+"})
	s.push(&token{TkType: number, Lexeme: "99", Literal: float64(99)})

	assert.Equal(t, s.pop(), &token{TkType: number, Lexeme: "99", Literal: float64(99)})
	assert.Equal(t, s.pop(), &token{TkType: plus, Lexeme: "+"})
	assert.Equal(t, s.pop(), &token{TkType: number, Lexeme: "1", Literal: float64(1)})
	assert.Nil(t, s.pop())
}

func testPeek(t *testing.T) {
	s := &stack{}
	s.push(&token{TkType: number, Lexeme: "1", Literal: float64(1)})
	s.push(&token{TkType: plus, Lexeme: "+"})
	s.push(&token{TkType: number, Lexeme: "99", Literal: float64(99)})
	assert.Equal(t, s.peek(), &token{TkType: number, Lexeme: "99", Literal: float64(99)})
	assert.Equal(t, s.length, 3)

	s.pop()
	assert.Equal(t, s.peek(), &token{TkType: plus, Lexeme: "+"})
	assert.Equal(t, s.length, 2)
}

func testIsEmpty(t *testing.T) {
	s := &stack{}
	assert.True(t, s.isEmpty())

	s.push(&token{TkType: number, Lexeme: "1", Literal: float64(1)})
	assert.False(t, s.isEmpty())

	s.pop()
	assert.True(t, s.isEmpty())
}
