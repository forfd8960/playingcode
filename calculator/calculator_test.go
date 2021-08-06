package calculator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculator(t *testing.T) {
	t.Run("simple math", test1Plus2)
	t.Run("left and right parent", leftAndRightParent)
	t.Run("one parentheses", oneParentheses)
	t.Run("left parentheses not match", parenthesesNotMatch)
	t.Run("right parentheses not match", rightParenthesesNotMatch)
	t.Run("unsupport operator", unsupportOperator)
	t.Run("divisor is zero", divisorIsZero)
}

func test1Plus2(t *testing.T) {
	cal := NewCalculator("1 + 2")
	err := cal.Exec()
	assert.Nil(t, err)

	result, err := cal.Result()
	if assert.Nil(t, err) {
		assert.Equal(t, float64(3), result)
	}
}

func leftAndRightParent(t *testing.T) {
	cal := NewCalculator("(1 + 99) * (100 / 10) - 99")
	err := cal.Exec()
	assert.Nil(t, err)

	result, err := cal.Result()
	if assert.Nil(t, err) {
		assert.Equal(t, float64(901), result)
	}
}

func oneParentheses(t *testing.T) {
	cal := NewCalculator("(88 + 22 / 11)")
	err := cal.Exec()
	assert.Nil(t, err)

	result, err := cal.Result()
	if assert.Nil(t, err) {
		assert.Equal(t, float64(90), result)
	}
}

func parenthesesNotMatch(t *testing.T) {
	cal := NewCalculator("(1 + 2))")
	err := cal.Exec()
	assert.Equal(t, errInvalidExpression, err)
}

func rightParenthesesNotMatch(t *testing.T) {
	cal := NewCalculator("((1 + 2)")
	err := cal.Exec()
	assert.Equal(t, errInvalidExpression, err)
}

func unsupportOperator(t *testing.T) {
	cal := NewCalculator("(1 + 2) % (10 * 9)")
	err := cal.Exec()
	assert.Equal(t, errUnsupportToken('%'), err)
}

func divisorIsZero(t *testing.T) {
	cal := NewCalculator("10 / 0")
	err := cal.Exec()
	assert.Equal(t, errDivisorZero, err)
}
