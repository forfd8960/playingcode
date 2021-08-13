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
	s.push(&token{TkType: number, Text: "1", Value: float64(1)})
	s.push(&token{TkType: plus, Text: "+"})
	s.push(&token{TkType: number, Text: "99", Value: float64(99)})

	assert.Equal(t, s.pop(), &token{TkType: number, Text: "99", Value: float64(99)})
	assert.Equal(t, s.pop(), &token{TkType: plus, Text: "+"})
	assert.Equal(t, s.pop(), &token{TkType: number, Text: "1", Value: float64(1)})
	assert.Nil(t, s.pop())
}

func testPeek(t *testing.T) {
	s := &stack{}
	s.push(&token{TkType: number, Text: "1", Value: float64(1)})
	s.push(&token{TkType: plus, Text: "+"})
	s.push(&token{TkType: number, Text: "99", Value: float64(99)})
	assert.Equal(t, s.peek(), &token{TkType: number, Text: "99", Value: float64(99)})
	assert.Equal(t, s.length, 3)

	s.pop()
	assert.Equal(t, s.peek(), &token{TkType: plus, Text: "+"})
	assert.Equal(t, s.length, 2)
}

func testIsEmpty(t *testing.T) {
	s := &stack{}
	assert.True(t, s.isEmpty())

	s.push(&token{TkType: number, Text: "1", Value: float64(1)})
	assert.False(t, s.isEmpty())

	s.pop()
	assert.True(t, s.isEmpty())
}
