package calculator

import (
	"errors"
	"fmt"
	"math"
)

var (
	operatorPrecedence = map[tokenType]int{
		rightParent: 3,
		doubleStar:  2,
		star:        2,
		slash:       2,
		plus:        1,
		minus:       1,
		leftParent:  0,
	}

	errUnsupportOperator = func(op string) error {
		return fmt.Errorf("unsupport operator: %s", op)
	}

	errInvalidExpression = errors.New("invalid expression")
	errDivisorZero       = errors.New("divisor is zero")
)

// Calculator ...
type Calculator struct {
	exp       string
	operators *stack
	operands  *stack
}

func NewCalculator(exp string) *Calculator {
	return &Calculator{
		exp:       exp,
		operators: &stack{},
		operands:  &stack{},
	}
}

// Exec run the experssion
func (cal *Calculator) Exec() error {
	sc := newScanner(cal.exp)
	if err := sc.scanExpression(); err != nil {
		return err
	}

	return cal.putTokensToStack(sc.tokens)
}

func (cal *Calculator) putTokensToStack(tokens []*token) error {
	for tIdx := 0; tIdx < len(tokens); tIdx++ {
		t := tokens[tIdx]

		switch t.TkType {
		case eof:
		case number:
			cal.operands.push(t)
		case leftParent:
			cal.operators.push(t)
		case rightParent:
			if err := cal.popAndEvalUtilLeftParent(); err != nil {
				return err
			}
		default: // + - * /
			if err := cal.eval(t); err != nil {
				return err
			}
		}
	}

	return cal.cleanupOperator()
}

func (cal *Calculator) popAndEvalUtilLeftParent() error {
	for {
		op := cal.operators.pop()
		if op == nil {
			return errInvalidExpression
		}

		if op.TkType == leftParent {
			return nil
		}

		result, err := cal.evalBinary(op)
		if err != nil {
			return err
		}
		cal.operands.push(result)
	}
}

func (cal *Calculator) cleanupOperator() error {
	for {
		op := cal.operators.pop()
		if op == nil {
			return nil
		}

		result, err := cal.evalBinary(op)
		if err != nil {
			return err
		}
		cal.operands.push(result)
	}
}

func (cal *Calculator) eval(t *token) error {
	prec, ok := operatorPrecedence[t.TkType]
	if !ok {
		return errUnsupportOperator(t.Text)
	}

	// if there is no operator, then push current operator to ops stack
	previous := cal.operators.peek()
	if previous == nil {
		cal.operators.push(t)
		return nil
	}

	// if we have ops, and the prec > current ops prec
	// we first calculate the previous operator, and push the result into opr stack
	prevPrec := operatorPrecedence[previous.TkType]
	if prevPrec >= prec {
		result, err := cal.evalBinary(previous)
		if err != nil {
			return err
		}

		cal.operators.pop()
		cal.operands.push(result)
	}

	// and then push current operator
	cal.operators.push(t)
	return nil
}

func (cal *Calculator) evalBinary(op *token) (*token, error) {
	n2 := cal.operands.pop()
	n1 := cal.operands.pop()
	if n2 == nil || n1 == nil {
		return nil, errInvalidExpression
	}

	val1, val2 := n1.Value.(float64), n2.Value.(float64)

	result := &token{TkType: number}
	switch op.Text {
	case "+":
		result.Value = val1 + val2
	case "-":
		result.Value = val1 - val2
	case "*":
		result.Value = val1 * val2
	case "**":
		result.Value = math.Pow(val1, val2)
	case "/":
		if val2 == 0 {
			return nil, errDivisorZero
		}
		result.Value = val1 / val2
	}

	return result, nil
}

// Result get the result of the calculation
func (cal *Calculator) Result() (float64, error) {
	result := cal.operands.pop()
	if result == nil {
		return -1, errInvalidExpression
	}

	return result.Value.(float64), nil
}
